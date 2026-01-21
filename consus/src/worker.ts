import { ConsusCrawler } from './crawler';
import { fetchOpenRouterResponse } from './ai';

interface Env {
    OPENROUTER_API_KEY: string;
    GNEWS_API_KEY: string;
}

const localeMap: { [key: string]: string } = {
    tr: 'tr-tr', de: 'de-de', us: 'us-en', fr: 'fr-fr', ru: 'ru-ru',
    jp: 'jp-jp', es: 'es-es', it: 'it-it', cn: 'cn-zh', gb: 'uk-en',
    br: 'br-pt', ar: 'xa-ar', nl: 'nl-nl', pl: 'pl-pl', kr: 'kr-ko',
    in: 'in-en', ca: 'ca-en', au: 'au-en', sa: 'sa-ar', se: 'se-sv',
    no: 'no-no', dk: 'dk-da', fi: 'fi-fi', gr: 'gr-el', il: 'il-he',
    mx: 'mx-es', id: 'id-id', th: 'th-th', vn: 'vn-vi', za: 'za-en'
};

export default {
    async fetch(request: Request, env: Env): Promise<Response> {
        const url = new URL(request.url);
        
        if (url.pathname !== '/api/search') {
            return new Response('Not Found', { status: 404 });
        }

        const query = url.searchParams.get('q');
        const userLocals = url.searchParams.get('kl');
        const types = url.searchParams.get('type') || 'web';

        if (!query) {
            return Response.json({ error: "You must enter a query." }, { status: 400 });
        }

        const crawler = new ConsusCrawler(env.GNEWS_API_KEY);

        if (types.toLowerCase() === 'ai') {
            try {
                const aiResponse = await fetchOpenRouterResponse(query, env.OPENROUTER_API_KEY);
                return Response.json({
                    query: query,
                    type: "ai",
                    searchSource: "Synapic AI",
                    wiki: null,
                    response: aiResponse
                });
            } catch (error) {
                return Response.json({ error: "AI error" }, { status: 500 });
            }
        }

        const locals = userLocals 
            ? Array.from(new Set(userLocals.split(',').map(l => l.trim().toLowerCase())))
            : ['tr'];

        const searchTasks = locals.map(async (lang) => {
            const code = localeMap[lang] || 'wt-wt';
            return await crawler.executeSearch(query, code, types);
        });

        const allData = await Promise.all(searchTasks);
        const requestedTypes = types.toLowerCase().split(',').map(t => t.trim());

        if (requestedTypes.includes('image') || requestedTypes.includes('images')) {
            const allImages: any[] = [];
            allData.forEach(searchResult => {
                if (searchResult.images) allImages.push(...searchResult.images);
            });

            const uniqueImages = Array.from(
                new Map(allImages.map(item => [item.image, item])).values()
            );

            return Response.json({
                query,
                type: "image",
                searchSource: "Images",
                wiki: null,
                images: uniqueImages
            });
        }

        if (requestedTypes.includes('news')) {
            const allNews: any[] = [];
            allData.forEach(searchResult => {
                if (searchResult.news) allNews.push(...searchResult.news);
            });

            const uniqueNews = Array.from(
                new Map(allNews.map(item => [item.link, item])).values()
            );

            return Response.json({
                query,
                type: "news",
                searchSource: "News",
                wiki: null,
                newsResults: uniqueNews
            });
        }

        if (requestedTypes.includes('video') || requestedTypes.includes('videos')) {
            const allVideos: any[] = [];
            allData.forEach(searchResult => {
                if (searchResult.videos) allVideos.push(...searchResult.videos);
            });

            const uniqueVideos = Array.from(
                new Map(allVideos.map(item => [item.url, item])).values()
            );

            return Response.json({
                query,
                type: "video",
                searchSource: "YouTube Videos",
                wiki: null,
                videos: uniqueVideos
            });
        }

        const allResults: any[] = [];
        allData.forEach(searchResult => {
            if (searchResult.web) allResults.push(...searchResult.web);
        });

        const uniqueResults = Array.from(
            new Map(allResults.map(item => [item.link, item])).values()
        );

        return Response.json({
            query,
            type: types,
            searchSource: "Web",
            wiki: null,
            results: uniqueResults
        });
    }
};