import { parseHTML } from 'linkedom';

export interface SearchResult {
    title: string;
    link: string;
    snippet: string;
    displayUrl: string;
    source: string;
}

export interface ImageResult {
    title: string;
    image: string;
    thumbnail: string;
    link: string;
}

export interface NewsResult {
    title: string;
    snippet: string;
    link: string;
    displayUrl: string;
    source: string;
    image: string;
    date: string;
}

export interface VideoResult {
    title: string;
    url: string;
    thumbnail: string;
    source: string;
}

export interface SearchResponse {
    web: SearchResult[];
    images: ImageResult[];
    news: NewsResult[];
    videos: VideoResult[];
}

export class ConsusCrawler {
    private userAgent = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36';
    private gnewsApiKey: string;
    private trustedDomains = ['wikipedia.org', 'github.com', 'stackoverflow.com', 'medium.com', 'dev.to', 'mozilla.org', 'w3.org'];

    constructor(gnewsApiKey: string = '') {
        this.gnewsApiKey = gnewsApiKey;
    }

    private decodeDDGUrl(url: string): string {
        try {
            if (url.includes('uddg=')) {
                return decodeURIComponent(url.split('uddg=')[1].split('&')[0]);
            }
            return url.startsWith('//') ? `https:${url}` : url;
        } catch { 
            return url; 
        }
    }

    private getDisplayUrl(link: string): string {
        try {
            const url = new URL(link);
            const hostname = url.hostname;
            return hostname.startsWith('www.') ? hostname.slice(4) : hostname;
        } catch {
            return link;
        }
    }

    private calculateDomainAuthority(url: string): number {
        let score = 50;
        try {
            const urlObj = new URL(url);
            const domain = urlObj.hostname.toLowerCase();
            
            if (this.trustedDomains.some(td => domain.includes(td))) score += 30;
            if (domain.endsWith('.gov') || domain.endsWith('.edu')) score += 25;
            if (domain.endsWith('.org')) score += 15;
            if (domain.split('.').length === 2) score += 10;
            if (domain.startsWith('www.')) score += 5;
        } catch {}
        
        return Math.min(score, 100);
    }

    private calculateRelevanceScore(title: string, snippet: string, query: string): number {
        const queryWords = query.toLowerCase().split(' ').filter(w => w.length > 2);
        const titleLower = title.toLowerCase();
        const snippetLower = snippet.toLowerCase();
        
        let score = 0;
        let titleMatches = 0;
        
        queryWords.forEach(word => {
            if (titleLower.includes(word)) {
                titleMatches++;
                score += 15;
            }
            if (snippetLower.includes(word)) {
                score += 5;
            }
        });
        
        if (titleLower.includes(query.toLowerCase())) {
            score += 30;
        }
        
        if (titleMatches === queryWords.length) score += 20;
        if (title.length > 10 && title.length < 80) score += 10;
        if (snippet.length > 50 && snippet.length < 200) score += 5;
        
        return Math.min(score, 100);
    }

    private calculateFreshnessScore(position: number, totalResults: number): number {
        const positionScore = Math.max(0, 100 - (position * 5));
        const recencyBonus = position < 5 ? 20 : position < 10 ? 10 : 0;
        return Math.min(positionScore + recencyBonus, 100);
    }

    private calculateOverallScore(components: { [key: string]: number }): number {
        const weights = Object.values(components);
        const sum = weights.reduce((a, b) => a + b, 0);
        return Math.round(sum / weights.length);
    }

    public async search(query: string, localCode: string): Promise<SearchResult[]> {
        const results: any[] = [];
        try {
            const url = `https://html.duckduckgo.com/html/`;
            const body = `q=${encodeURIComponent(query)}&kl=${localCode}`;

            const controller = new AbortController();
            const timeoutId = setTimeout(() => controller.abort(), 10000);

            const response = await fetch(url, {
                method: 'POST',
                headers: {
                    'User-Agent': this.userAgent,
                    'Content-Type': 'application/x-www-form-urlencoded'
                },
                body: body,
                signal: controller.signal
            });

            clearTimeout(timeoutId);

            if (!response.ok) return [];

            const html = await response.text();
            const { document } = parseHTML(html);
            const resultElements = document.querySelectorAll('.result');
            
            let position = 0;
            
            resultElements.forEach((el: any) => {
                const titleEl = el.querySelector('.result__a');
                if (titleEl) {
                    const title = titleEl.textContent?.trim() || '';
                    const href = titleEl.getAttribute('href') || '';
                    const link = this.decodeDDGUrl(href);
                    const snippetEl = el.querySelector('.result__snippet');
                    const snippet = snippetEl?.textContent?.trim() || '';
                    const displayUrl = this.getDisplayUrl(link);
                    
                    const authorityScore = this.calculateDomainAuthority(link);
                    const relevanceScore = this.calculateRelevanceScore(title, snippet, query);
                    const freshnessScore = this.calculateFreshnessScore(position, 50);
                    const score = this.calculateOverallScore({ authorityScore, relevanceScore, freshnessScore });
                    
                    results.push({
                        title,
                        link,
                        snippet,
                        displayUrl,
                        source: "DuckDuckGo",
                        score
                    });
                    position++;
                }
            });
            
            return results
                .sort((a, b) => b.score - a.score)
                .map(({ score, ...rest }) => rest);
        } catch { 
            return []; 
        }
    }

    public async searchImages(query: string, localCode: string): Promise<ImageResult[]> {
        const results: any[] = [];
        try {
            const params = new URLSearchParams({
                q: query,
                kl: localCode,
                l: localCode,
                o: 'json',
                s: '0',
                u: 'bing',
                vqd: ''
            });

            const url = `https://duckduckgo.com/i.js?${params.toString()}`;

            const controller = new AbortController();
            const timeoutId = setTimeout(() => controller.abort(), 10000);

            const response = await fetch(url, {
                headers: {
                    'User-Agent': this.userAgent,
                    'Referer': 'https://duckduckgo.com/'
                },
                signal: controller.signal
            });

            clearTimeout(timeoutId);

            if (!response.ok) throw new Error('Image search failed');

            const data = await response.json();

            if (data && data.results) {
                data.results.forEach((item: any, index: number) => {
                    const title = item.title || '';
                    const url = item.url || '';
                    
                    let score = 100 - (index * 3);
                    if (title.toLowerCase().includes(query.toLowerCase())) score += 15;
                    if (item.width && item.height && item.width > 800 && item.height > 600) score += 10;
                    if (url) score += this.calculateDomainAuthority(url) * 0.3;
                    
                    results.push({
                        title,
                        image: item.image || '',
                        thumbnail: item.thumbnail || '',
                        link: url,
                        score: Math.min(Math.round(score), 100)
                    });
                });
            }
            
            return results
                .sort((a, b) => b.score - a.score)
                .map(({ score, ...rest }) => rest);
        } catch (error) {
            try {
                const htmlUrl = `https://duckduckgo.com/?q=${encodeURIComponent(query)}&kl=${localCode}&iax=images&ia=images`;
                
                const controller = new AbortController();
                const timeoutId = setTimeout(() => controller.abort(), 10000);

                const response = await fetch(htmlUrl, {
                    headers: { 'User-Agent': this.userAgent },
                    signal: controller.signal
                });

                clearTimeout(timeoutId);

                const htmlData = await response.text();
                const vqdMatch = htmlData.match(/vqd=['"]([^'"]+)['"]/);
                const vqd = vqdMatch ? vqdMatch[1] : null;
                
                if (vqd) {
                    const params = new URLSearchParams({ q: query, kl: localCode, vqd: vqd });
                    const apiUrl = `https://duckduckgo.com/i.js?${params.toString()}`;
                    
                    const controller2 = new AbortController();
                    const timeoutId2 = setTimeout(() => controller2.abort(), 10000);

                    const apiResponse = await fetch(apiUrl, {
                        headers: { 'User-Agent': this.userAgent },
                        signal: controller2.signal
                    });

                    clearTimeout(timeoutId2);

                    const apiData = await apiResponse.json();
                    
                    if (apiData && apiData.results) {
                        apiData.results.forEach((item: any, index: number) => {
                            let score = 100 - (index * 3);
                            results.push({
                                title: item.title || '',
                                image: item.image || '',
                                thumbnail: item.thumbnail || '',
                                link: item.url || '',
                                score: Math.round(score)
                            });
                        });
                    }
                }
                return results
                    .sort((a, b) => b.score - a.score)
                    .map(({ score, ...rest }) => rest);
            } catch { 
                return []; 
            }
        }
    }

    public async searchNews(query: string, localCode: string): Promise<NewsResult[]> {
        const results: any[] = [];
        try {
            if (!this.gnewsApiKey) return [];

            let lang = 'en';
            let country = 'us';
            
            if (localCode) {
                if (localCode.includes('-')) {
                    const parts = localCode.split('-');
                    if (parts.length === 2) {
                        lang = parts[0].toLowerCase();
                        country = parts[1].toLowerCase();
                    }
                } else {
                    lang = localCode.substring(0, 2).toLowerCase();
                    country = localCode.substring(0, 2).toLowerCase();
                }
            }

            const params = new URLSearchParams({
                q: query,
                lang: lang,
                country: country,
                apikey: this.gnewsApiKey,
                max: '10'
            });

            const url = `https://gnews.io/api/v4/search?${params.toString()}`;

            const controller = new AbortController();
            const timeoutId = setTimeout(() => controller.abort(), 10000);

            const response = await fetch(url, {
                signal: controller.signal
            });

            clearTimeout(timeoutId);

            if (!response.ok) return [];

            const data = await response.json();

            if (data && data.articles) {
                const now = Date.now();

                data.articles.forEach((item: any) => {
                    const title = item.title || '';
                    const link = item.url || '';
                    const snippet = item.description || item.content || '';
                    const source = item.source?.name || 'GNews';
                    const dateStr = item.publishedAt || '';
                    const displayUrl = this.getDisplayUrl(link);
                    const image = item.image || '';
                    
                    const authorityScore = this.calculateDomainAuthority(link);
                    const relevanceScore = this.calculateRelevanceScore(title, snippet, query);
                    
                    let freshnessScore = 50;
                    if (dateStr) {
                        const dateTs = new Date(dateStr).getTime();
                        if (!isNaN(dateTs)) {
                            const ageInHours = (now - dateTs) / (1000 * 60 * 60);
                            if (ageInHours < 24) freshnessScore = 100;
                            else if (ageInHours < 72) freshnessScore = 85;
                            else if (ageInHours < 168) freshnessScore = 70;
                            else if (ageInHours < 720) freshnessScore = 50;
                            else freshnessScore = 30;
                        }
                    }
                    
                    const score = this.calculateOverallScore({ authorityScore, relevanceScore, freshnessScore });
                    
                    results.push({
                        title,
                        snippet,
                        link,
                        displayUrl,
                        source,
                        image,
                        date: dateStr,
                        score
                    });
                });
            }
            
            return results
                .sort((a, b) => b.score - a.score)
                .map(({ score, ...rest }) => rest);
        } catch (error) {
            return [];
        }
    }

    public async searchVideos(query: string, localCode: string): Promise<VideoResult[]> {
        const results: any[] = [];
        try {
            const params = new URLSearchParams({ search_query: query });
            const url = `https://www.youtube.com/results?${params.toString()}`;

            const controller = new AbortController();
            const timeoutId = setTimeout(() => controller.abort(), 10000);

            const response = await fetch(url, {
                headers: {
                    'User-Agent': this.userAgent,
                    'Accept-Language': 'en-US,en;q=0.9'
                },
                signal: controller.signal
            });

            clearTimeout(timeoutId);

            if (!response.ok) return [];

            const data = await response.text();
            const ytInitialDataMatch = data.match(/var ytInitialData = ({.+?});/);
            
            if (ytInitialDataMatch) {
                const ytData = JSON.parse(ytInitialDataMatch[1]);
                const contents = ytData?.contents?.twoColumnSearchResultsRenderer?.primaryContents?.sectionListRenderer?.contents;

                if (contents) {
                    contents.forEach((section: any) => {
                        const items = section?.itemSectionRenderer?.contents || [];
                        items.forEach((item: any, index: number) => {
                            const videoRenderer = item?.videoRenderer;
                            if (videoRenderer) {
                                const videoId = videoRenderer.videoId;
                                const title = videoRenderer.title?.runs?.[0]?.text || '';
                                const channelName = videoRenderer.ownerText?.runs?.[0]?.text || '';
                                const thumbnails = videoRenderer.thumbnail?.thumbnails || [];
                                const thumbnail = thumbnails.length > 0 ? thumbnails[thumbnails.length - 1].url : '';

                                let score = 100 - (index * 2);
                                if (title.toLowerCase().includes(query.toLowerCase())) score += 15;
                                
                                results.push({
                                    title,
                                    url: `https://www.youtube.com/watch?v=${videoId}`,
                                    thumbnail,
                                    source: channelName,
                                    score: Math.min(Math.round(score), 100)
                                });
                            }
                        });
                    });
                }
            }
            
            return results
                .sort((a, b) => b.score - a.score)
                .slice(0, 15)
                .map(({ score, ...rest }) => rest);
        } catch (error) {
            return [];
        }
    }

    public async executeSearch(query: string, localCode: string, types: string): Promise<SearchResponse> {
        const requestedTypes = types ? types.toLowerCase().split(',').map(t => t.trim()) : ['web'];
        const response: SearchResponse = {
            web: [],
            images: [],
            news: [],
            videos: []
        };

        const promises: Promise<any>[] = [];

        if (requestedTypes.includes('web')) {
            promises.push(this.search(query, localCode).then(res => response.web = res));
        }

        if (requestedTypes.includes('image') || requestedTypes.includes('images')) {
            promises.push(this.searchImages(query, localCode).then(res => response.images = res));
        }

        if (requestedTypes.includes('news')) {
            promises.push(this.searchNews(query, localCode).then(res => response.news = res));
        }

        if (requestedTypes.includes('video') || requestedTypes.includes('videos')) {
            promises.push(this.searchVideos(query, localCode).then(res => response.videos = res));
        }

        if (promises.length === 0) {
            promises.push(this.search(query, localCode).then(res => response.web = res));
        }

        await Promise.all(promises);
        return response;
    }
}