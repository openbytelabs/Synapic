<script setup>
import { ref, watch, onMounted, computed, nextTick, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { lang, detectLanguage, tr, en } from '../composables/useLocale.js'
import { useBangs } from '../composables/useBangs.js'
import { useBlockedSites } from '../composables/useBlockedSites.js'
import { initSettings } from '../composables/settings.js'
import { usePrisms } from '../composables/usePrisms.js'

const { checkBang } = useBangs()
const { getFilteredResults, addBlockedSite } = useBlockedSites()
const { allPrisms, activePrismId, setActivePrism, filterByPrism } = usePrisms()

const API_BASE = 'https://api.synapicsearch.com'

const route = useRoute()
const router = useRouter()

const t = computed(() => lang.value === 'tr' ? tr : en)

const activeTab = ref('all')
const searchQuery = ref('')
const inputRef = ref(null)
const results = ref([])
const imageResults = ref([])
const newsResults = ref([])
const loading = ref(false)
const wikiData = ref(null)
const relatedSearches = ref([])
const previewImage = ref(null)
const previewVisible = ref(false)
const resultsCount = ref(0)
const searchTime = ref(0)
const wikiLoading = ref(false)
const hasSearched = ref(false)
const lastSearchedQuery = ref('')

const weatherData = ref(null)
const weatherLoading = ref(false)
const showCalculator = ref(false)
const calcInput = ref('')
const calcResult = ref('')
const calcHistory = ref([])
const activeMenuIdx = ref(null)

const aiAnswer = ref(null)
const aiLoading = ref(false)
const showAiAnswer = ref(true)
const aiMode = ref('auto')
const aiExpanded = ref(false)

const expandedTexts = ref({})

watch(aiAnswer, () => { aiExpanded.value = false })

const toggleText = (idx) => {
  expandedTexts.value[idx] = !expandedTexts.value[idx]
}

const interleavedResults = computed(() => {
  if (results.value.length === 0) return []
  const items = []
  let relatedInserted = false
  let imagesInserted = false
  for (let i = 0; i < results.value.length; i++) {
    items.push({ type: 'result', data: results.value[i], index: i })
    if (i === 2 && results.value.length > 4 && !isMobile.value) {
      if (relatedSearches.value.length > 0 && !relatedInserted) {
        items.push({ type: 'related', data: relatedSearches.value })
        relatedInserted = true
      }
      if (imageResults.value.length > 0 && !imagesInserted) {
        const validImages = imageResults.value.slice(0, 8).filter(img => isValidImageUrl(img.thumbnail || img.url || img.src || img.image))
        if (validImages.length > 0) {
          items.push({ type: 'images', data: validImages })
          imagesInserted = true
        }
      }
    }
  }
  return items
})
const handleResize = () => { isMobile.value = window.innerWidth <= 768 }
const RESULTS_PER_PAGE = 15
const currentPage = ref(1)

function isValidImageUrl(url) {
  if (!url || typeof url !== 'string') return false
  return url.match(/\.(jpg|jpeg|png|gif|webp|svg|bmp|ico)(\?.*)?$/i) !== null || url.includes('googleusercontent') || url.includes('gstatic') || url.includes('favicon')
}

const totalResultItems = computed(() => {
  return interleavedResults.value.filter(item => item.type === 'result').length
})

const totalPages = computed(() => {
  return Math.max(1, Math.ceil(totalResultItems.value / RESULTS_PER_PAGE))
})

const paginatedResults = computed(() => {
  const all = interleavedResults.value
  if (all.length === 0) return []
  const start = (currentPage.value - 1) * RESULTS_PER_PAGE
  const end = start + RESULTS_PER_PAGE
  let resultCount = 0
  const pageItems = []
  for (const item of all) {
    if (item.type === 'result') {
      if (resultCount >= start && resultCount < end) {
        pageItems.push(item)
      }
      resultCount++
    } else if (item.type === 'related' || item.type === 'images') {
      const prevResult = resultCount - 1
      if (prevResult >= start && prevResult < end) {
        pageItems.push(item)
      }
    }
  }
  return pageItems
})

const paginationRange = computed(() => {
  const total = totalPages.value
  const current = currentPage.value
  const pages = []
  if (total <= 7) {
    for (let i = 1; i <= total; i++) pages.push(i)
  } else {
    pages.push(1)
    if (current > 3) pages.push('...')
    const startP = Math.max(2, current - 1)
    const endP = Math.min(total - 1, current + 1)
    for (let i = startP; i <= endP; i++) pages.push(i)
    if (current < total - 2) pages.push('...')
    pages.push(total)
  }
  return pages
})

function goToPage(page) {
  if (page < 1 || page > totalPages.value) return
  currentPage.value = page
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

function nextPage() {
  goToPage(currentPage.value + 1)
}

function prevPage() {
  goToPage(currentPage.value - 1)
}

const mapResults = ref([])
const mapLoading = ref(false)
const mapInstance = ref(null)
const mapMarkers = ref([])
const mapContainer = ref(null)
const leafletLoaded = ref(false)

const tabs = computed(() => [
  { id: 'all', label: t.value.results.all, icon: 'web' },
  { id: 'images', label: t.value.results.images, icon: 'image' },
  { id: 'news', label: t.value.results.news, icon: 'news' },
  { id: 'maps', label: t.value.results.maps, icon: 'maps' }
])

const weatherCodes = {
  0: { tr: 'Açık', en: 'Clear sky', icon: '☀️' },
  1: { tr: 'Çoğunlukla Açık', en: 'Mainly clear', icon: '🌤️' },
  2: { tr: 'Parçalı Bulutlu', en: 'Partly cloudy', icon: '⛅' },
  3: { tr: 'Kapalı', en: 'Overcast', icon: '☁️' },
  45: { tr: 'Sisli', en: 'Foggy', icon: '🌫️' },
  48: { tr: 'Kırağılı Sis', en: 'Freezing fog', icon: '🌫️' },
  51: { tr: 'Hafif Çisenti', en: 'Light drizzle', icon: '🌦️' },
  53: { tr: 'Orta Çisenti', en: 'Moderate drizzle', icon: '🌦️' },
  55: { tr: 'Yoğun Çisenti', en: 'Dense drizzle', icon: '🌧️' },
  61: { tr: 'Hafif Yağmur', en: 'Slight rain', icon: '🌧️' },
  63: { tr: 'Orta Yağmur', en: 'Moderate rain', icon: '🌧️' },
  65: { tr: 'Şiddetli Yağmur', en: 'Heavy rain', icon: '🌧️' },
  71: { tr: 'Hafif Kar', en: 'Slight snow', icon: '🌨️' },
  73: { tr: 'Orta Kar', en: 'Moderate snow', icon: '🌨️' },
  75: { tr: 'Yoğun Kar', en: 'Heavy snow', icon: '❄️' },
  80: { tr: 'Sağanak', en: 'Rain showers', icon: '🌦️' },
  81: { tr: 'Orta Sağanak', en: 'Moderate showers', icon: '🌧️' },
  82: { tr: 'Şiddetli Sağanak', en: 'Violent showers', icon: '⛈️' },
  85: { tr: 'Kar Sağanağı', en: 'Snow showers', icon: '🌨️' },
  86: { tr: 'Yoğun Kar Sağanağı', en: 'Heavy snow showers', icon: '❄️' },
  95: { tr: 'Gök Gürültülü Fırtına', en: 'Thunderstorm', icon: '⛈️' },
  96: { tr: 'Dolu ile Fırtına', en: 'Thunderstorm with hail', icon: '⛈️' },
  99: { tr: 'Şiddetli Dolu Fırtınası', en: 'Thunderstorm with heavy hail', icon: '⛈️' }
}

function getWeatherInfo(code) {
  return weatherCodes[code] || { tr: 'Bilinmiyor', en: 'Unknown', icon: '🌡️' }
}

function detectWidgets(query) {
  const q = query.toLowerCase().trim()
  weatherData.value = null
  showCalculator.value = false
  const weatherPatternTr = /^hava\s+durumu\s+(.+)$/
  const weatherPatternEn = /^weather\s+(?:in\s+|of\s+)?(.+)$/
  const calcPatternTr = /^(hesap\s*makinesi|hesaplayıcı)$/
  const calcPatternEn = /^(calculator|calc)$/
  let weatherMatch = q.match(weatherPatternTr) || q.match(weatherPatternEn)
  if (weatherMatch) {
    fetchWeather(weatherMatch[1])
    return
  }
  if (calcPatternTr.test(q) || calcPatternEn.test(q)) {
    showCalculator.value = true
    return
  }
}

async function fetchWeather(city) {
  weatherLoading.value = true
  weatherData.value = null
  try {
    const geoRes = await fetch(
      `https://nominatim.openstreetmap.org/search?format=json&q=${encodeURIComponent(city)}&limit=1`
    )
    const geoData = await geoRes.json()
    if (!geoData.length) { weatherLoading.value = false; return }
    const lat = geoData[0].lat
    const lon = geoData[0].lon
    const weatherRes = await fetch(
      `https://api.open-meteo.com/v1/forecast?latitude=${lat}&longitude=${lon}&current=temperature_2m,relative_humidity_2m,apparent_temperature,wind_speed_10m,weather_code`
    )
    const weather = await weatherRes.json()
    weatherData.value = {
      city: geoData[0].display_name.split(',')[0],
      country: geoData[0].display_name.split(',').slice(1, 2).join(',').trim(),
      temp: weather.current?.temperature_2m,
      feelsLike: weather.current?.apparent_temperature,
      humidity: weather.current?.relative_humidity_2m,
      windSpeed: weather.current?.wind_speed_10m,
      weatherCode: weather.current?.weather_code,
      lat,
      lon
    }
  } catch {
    weatherData.value = null
  }
  weatherLoading.value = false
}

function handleCalc() {
  if (!calcInput.value.trim()) return
  try {
    const sanitized = calcInput.value.replace(/[^0-9+\-*/().%\s]/g, '')
    const result = Function('"use strict"; return (' + sanitized + ')')()
    if (result !== undefined && result !== null && !isNaN(result)) {
      calcResult.value = String(result)
      calcHistory.value.unshift({ expression: calcInput.value, result: calcResult.value })
      if (calcHistory.value.length > 10) calcHistory.value.pop()
    } else {
      calcResult.value = 'Error'
    }
  } catch {
    calcResult.value = 'Error'
  }
}

function getDomain(url) {
  try {
    if (!url || typeof url !== 'string') return ''
    const urlWithProtocol = url.match(/^https?:\/\//) ? url : 'https://' + url
    const u = new URL(urlWithProtocol)
    return u.hostname.replace(/^www\./, '')
  } catch {
    return ''
  }
}

function getPath(url) {
  try {
    if (!url || typeof url !== 'string') return ''
    const urlWithProtocol = url.match(/^https?:\/\//) ? url : 'https://' + url
    const u = new URL(urlWithProtocol)
    const p = u.pathname + u.search
    if (p.length > 50) return p.substring(0, 50) + '...'
    return p || '/'
  } catch {
    return ''
  }
}

function handleImageError(e) {
  const img = e.target
  if (!img) return
  img.style.display = 'none'
  const parent = img.parentNode
  if (!parent) return
  const existing = parent.querySelector('.image-placeholder')
  if (existing) return
  const placeholder = document.createElement('div')
  placeholder.className = 'image-placeholder'
  placeholder.innerHTML = '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"><rect x="3" y="3" width="18" height="18" rx="2"/><circle cx="8.5" cy="8.5" r="1.5"/><path d="M21 15l-5-5L5 21"/></svg>'
  parent.appendChild(placeholder)
}

function generateRelatedSearches(query) {
  const words = query.toLowerCase().split(/\s+/).filter(w => w.length > 2)
  if (words.length === 0) return []
  const suffixes = [
    'definition', 'meaning', 'examples', 'tutorial', 'explained',
    'vs', 'alternatives', 'how to use', 'benefits', 'review',
    'guide', 'best practices', 'free', 'online', 'near me',
    '2025', 'tips', 'for beginners', 'advanced', 'tools',
    'türkçe', 'nedir', 'nasıl yapılır', 'örnekleri'
  ]
  const suggestions = new Set()
  if (words.length === 1) {
    const w = words[0]
    for (let i = 0; i < suffixes.length && suggestions.size < 8; i++) {
      suggestions.add(`${w} ${suffixes[i]}`)
    }
  } else {
    const full = query
    for (let i = 0; i < suffixes.length && suggestions.size < 4; i++) {
      suggestions.add(`${full} ${suffixes[i]}`)
    }
    for (let i = 0; i < words.length && suggestions.size < 8; i++) {
      suggestions.add(`${words[i]} explained`)
    }
  }
  return Array.from(suggestions).slice(0, 8)
}

async function fetchWithRetry(url, options = {}, retries = 4, delay = 800) {
  for (let i = 0; i <= retries; i++) {
    try {
      const res = await fetch(url, { ...options, signal: AbortSignal.timeout(12000) })
      if (res.ok) {
        return res
      }
      if (i < retries && res.status >= 500) {
        await new Promise(resolve => setTimeout(resolve, delay * (i + 1)))
        continue
      }
      return res
    } catch (err) {
      if (i < retries) {
        await new Promise(resolve => setTimeout(resolve, delay * (i + 1)))
        continue
      }
      throw err
    }
  }
  throw new Error('Max retries exceeded')
}

async function fetchWebResults(query) {
  try {
    const res = await fetchWithRetry(`${API_BASE}/api?q=${encodeURIComponent(query)}`, {}, 4, 800)
    if (!res.ok) return { results: [], count: 0, time: 0 }
    const data = await res.json()
    if (data.error) return { results: [], count: 0, time: 0 }
    let rawResults = data.results || data.web?.results || data.organic || []
    const results = rawResults.map(r => ({
      url: r.url || r.link || r.href || r.source_url || r.webpage || r.canonical_url || r.uri || r.address || '',
      title: r.title || r.name || r.heading || r.headline || 'Başlık yok',
      description: r.description || r.snippet || r.content || r.summary || r.body || r.text || '',
      date: r.date || r.published || r.publishedDate || r.age || '',
      thumbnail: r.thumbnail || r.image || r.img || r.icon || ''
    }))
    return {
      results,
      count: data.total || data.total_results || data.count || results.length || 0,
      time: data.time || 0
    }
  } catch (e) {
    console.error('fetchWebResults error:', e)
    return { results: [], count: 0, time: 0 }
  }
}

async function fetchImageResults(query) {
  try {
    const res = await fetchWithRetry(`${API_BASE}/images?q=${encodeURIComponent(query)}`, {}, 1, 200)
    if (!res.ok) return []
    const data = await res.json()
    if (data.error) return []
    let rawImages = data.results || data.images || data.data || []
    return rawImages.map(img => ({
      url: img.url || img.image || img.src || img.link || img.source_url || '',
      thumbnail: img.thumbnail || img.thumb || img.preview || img.url || img.image || img.src || '',
      title: img.title || img.alt || img.name || img.caption || '',
      source_url: img.source_url || img.source || img.webpage || img.page_url || img.url || ''
    }))
  } catch (e) {
    console.error('fetchImageResults error:', e)
    return []
  }
}

async function fetchNewsResults(query) {
  try {
    const res = await fetchWithRetry(`${API_BASE}/news?q=${encodeURIComponent(query)}`, {}, 1, 200)
    if (!res.ok) return []
    const data = await res.json()
    if (data.error) return []
    let rawNews = data.results || data.news || data.data || []
    return rawNews.map(n => ({
      url: n.url || n.link || n.href || n.source_url || n.webpage || n.uri || '',
      title: n.title || n.headline || n.name || 'Başlık yok',
      description: n.description || n.snippet || n.content || n.summary || n.body || '',
      date: n.date || n.published || n.publishedDate || n.age || n.time || ''
    }))
  } catch (e) {
    console.error('fetchNewsResults error:', e)
    return []
  }
}

async function fetchWiki(query) {
  wikiLoading.value = true
  wikiData.value = null
  try {
    const langParam = lang.value === 'tr' ? 'tr' : 'en'
    const res = await fetch(
      `https://${langParam}.wikipedia.org/api/rest_v1/page/summary/${encodeURIComponent(query)}`
    )
    if (!res.ok) {
      const res2 = await fetch(
        `https://${langParam}.wikipedia.org/w/api.php?action=query&list=search&srsearch=${encodeURIComponent(query)}&format=json&origin=*`
      )
      if (!res2.ok) { wikiLoading.value = false; return }
      const data2 = await res2.json()
      const first = data2.query?.search?.[0]
      if (!first) { wikiLoading.value = false; return }
      const res3 = await fetch(
        `https://${langParam}.wikipedia.org/api/rest_v1/page/summary/${encodeURIComponent(first.title)}`
      )
      if (!res3.ok) { wikiLoading.value = false; return }
      const data3 = await res3.json()
      if (data3.type === 'disambiguation' || !data3.extract) { wikiLoading.value = false; return }
      const isWikiSource = getDomain(data3.content_urls?.desktop?.page || '') === `${langParam}.wikipedia.org`
      wikiData.value = {
        title: data3.title,
        extract: data3.extract,
        thumbnail: data3.thumbnail?.source || data3.originalimage?.source,
        url: data3.content_urls?.desktop?.page,
        isValid: isWikiSource
      }
      wikiLoading.value = false
      return
    }
    const data = await res.json()
    if (data.type === 'disambiguation' || !data.extract) { wikiLoading.value = false; return }
    const isWikiSource = getDomain(data.content_urls?.desktop?.page || '') === `${lang.value === 'tr' ? 'tr' : 'en'}.wikipedia.org`
    wikiData.value = {
      title: data.title,
      extract: data.extract,
      thumbnail: data.thumbnail?.source || data.originalimage?.source,
      url: data.content_urls?.desktop?.page,
      isValid: isWikiSource
    }
  } catch {
    wikiData.value = null
  }
  wikiLoading.value = false
}

function loadLeaflet() {
  return new Promise((resolve, reject) => {
    if (window.L) { resolve(); return }
    const link = document.createElement('link')
    link.rel = 'stylesheet'
    link.href = 'https://unpkg.com/leaflet@1.9.4/dist/leaflet.css'
    document.head.appendChild(link)
    const script = document.createElement('script')
    script.src = 'https://unpkg.com/leaflet@1.9.4/dist/leaflet.js'
    script.onload = resolve
    script.onerror = reject
    document.head.appendChild(script)
  })
}

async function initMap() {
  if (mapInstance.value) return
  if (!mapContainer.value) return
  await loadLeaflet()
  leafletLoaded.value = true
  mapInstance.value = window.L.map(mapContainer.value).setView([41.0082, 28.9784], 6)
  window.L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; OpenStreetMap contributors'
  }).addTo(mapInstance.value)
}

async function searchMapLocations(query) {
  mapLoading.value = true
  mapResults.value = []
  if (!leafletLoaded.value) await loadLeaflet()
  if (!mapInstance.value && mapContainer.value) {
    await initMap()
    await nextTick()
  }
  try {
    const res = await fetch(
      `https://nominatim.openstreetmap.org/search?format=json&q=${encodeURIComponent(query)}&limit=10&addressdetails=1`
    )
    const data = await res.json()
    mapResults.value = data
    if (mapInstance.value && data.length > 0) {
      mapMarkers.value.forEach(m => m.remove())
      mapMarkers.value = []
      const bounds = window.L.latLngBounds(data.map(item => [parseFloat(item.lat), parseFloat(item.lon)]))
      data.forEach(item => {
        const marker = window.L.marker([parseFloat(item.lat), parseFloat(item.lon)]).addTo(mapInstance.value)
        marker.bindPopup(`<strong>${item.display_name}</strong>`)
        mapMarkers.value.push(marker)
      })
      mapInstance.value.fitBounds(bounds, { padding: [30, 30] })
    } else if (mapInstance.value && data.length === 0) {
      mapMarkers.value.forEach(m => m.remove())
      mapMarkers.value = []
    }
  } catch {
    mapResults.value = []
  }
  mapLoading.value = false
}

function determineAutoMode(query) {
  const lower = query.toLowerCase()
  const complexKeywords = ['compare', 'vs ', 'versus', 'impact of', 'relationship between', 'causes of', 'benefits of', 'history of', 'how does', 'why does', 'difference between', 'analyze', 'explain']
  if (complexKeywords.some(k => lower.includes(k))) return 'graphrag'
  return 'full_text'
}

async function fetchAIAnswer(query, searchType) {
  aiLoading.value = true
  aiAnswer.value = null
  showAiAnswer.value = true

  const mode = aiMode.value === 'auto' ? determineAutoMode(query) : aiMode.value

  const readSSE = (url) => new Promise((resolve) => {
    const es = new EventSource(url)
    es.addEventListener('ai_answer', (e) => {
      try {
        const data = JSON.parse(e.data)
        if (data.answer) resolve(data.answer)
        else resolve(null)
      } catch { resolve(null) }
      es.close()
    })
    es.addEventListener('ai_error', () => { resolve(null); es.close() })
    es.addEventListener('no_results', () => { resolve(null); es.close() })
    es.addEventListener('error', () => { resolve(null); es.close() })
    setTimeout(() => { resolve(null); es.close() }, 30000)
  })

  if (mode === 'full_text') {
    const answer = await readSSE(`${API_BASE}/ai?q=${encodeURIComponent(query)}&type=${searchType}`)
    if (answer) aiAnswer.value = answer
  } else if (mode === 'graphrag') {
    const answer = await readSSE(`${API_BASE}/graphrag?q=${encodeURIComponent(query)}&type=${searchType}`)
    if (answer) aiAnswer.value = answer
  } else if (mode === 'hybrid') {
    const answer = await readSSE(`${API_BASE}/graphrag?q=${encodeURIComponent(query)}&type=${searchType}`)
    if (answer) {
      aiAnswer.value = answer
    } else {
      const answer2 = await readSSE(`${API_BASE}/ai?q=${encodeURIComponent(query)}&type=${searchType}`)
      if (answer2) aiAnswer.value = answer2
    }
  }

  aiLoading.value = false
}

async function performSearch(query) {
  if (!query.trim()) return
  lastSearchedQuery.value = query
  loading.value = true
  hasSearched.value = true
  currentPage.value = 1
  results.value = []
  imageResults.value = []
  newsResults.value = []
  relatedSearches.value = generateRelatedSearches(query)
  wikiData.value = null
  mapResults.value = []
  aiAnswer.value = null
  aiLoading.value = false
  detectWidgets(query)
  const [webData, imgData, newsData] = await Promise.all([
    fetchWebResults(query),
    fetchImageResults(query),
    fetchNewsResults(query)
  ])
  const webFiltered = getFilteredResults(webData.results)
  const prismFiltered = filterByPrism(webFiltered)
  if (prismFiltered.length === 0 && webFiltered.length === 0) {
    let retryCount = 0
    while (retryCount < 3) {
      retryCount++
      await new Promise(resolve => setTimeout(resolve, 1200 * retryCount))
      const retryWeb = await fetchWebResults(query)
      const retryFiltered = getFilteredResults(retryWeb.results)
      const retryPrism = filterByPrism(retryFiltered)
      if (retryPrism.length > 0 || retryFiltered.length > 0) {
        results.value = retryPrism.length > 0 ? retryPrism : retryFiltered
        resultsCount.value = retryWeb.count
        searchTime.value = retryWeb.time
        fetchWiki(query)
        if (activeTab.value === 'all') fetchAIAnswer(query, 'web')
        loading.value = false
        return
      }
    }
  }
  results.value = prismFiltered
  resultsCount.value = webData.count
  searchTime.value = webData.time
  imageResults.value = imgData
  newsResults.value = filterByPrism(getFilteredResults(newsData))
  if (results.value.length > 0) {
    fetchWiki(query)
    if (activeTab.value === 'all') fetchAIAnswer(query, 'web')
  }
  loading.value = false
}

function handlePrismSwitch(id) {
  if (activePrismId.value === id) {
    setActivePrism('default')
  } else {
    setActivePrism(id)
  }
  if (searchQuery.value) {
    performSearch(searchQuery.value)
  }
}

function handleSearchSubmit() {
  const q = searchQuery.value.trim()
  if (!q) return
  if (q.startsWith('!')) {
    const result = checkBang(q)
    if (result) {
      const url = result.bang.url.replace('{query}', encodeURIComponent(result.query))
      window.open(url, '_blank')
      return
    }
  }
  activeTab.value = 'all'
  lastSearchedQuery.value = q
  performSearch(q)
  router.push({ path: '/search', query: { q } }).catch(() => {})
}

function switchTab(tabId) {
  activeTab.value = tabId
  if (tabId === 'maps' && hasSearched.value && searchQuery.value) {
    nextTick(async () => {
      if (!mapInstance.value) {
        await initMap()
        await nextTick()
        setTimeout(() => {
          if (mapInstance.value) mapInstance.value.invalidateSize()
        }, 200)
      }
      searchMapLocations(searchQuery.value)
    })
  }
}

function clearInput() {
  searchQuery.value = ''
  inputRef.value?.focus()
}

function openPreview(result) {
  previewImage.value = result
  previewVisible.value = true
}

function closePreview() {
  previewVisible.value = false
  setTimeout(() => { previewImage.value = null }, 300)
}

function handleRelatedClick(query) {
  searchQuery.value = query
  activeTab.value = 'all'
  performSearch(query)
  router.push({ path: '/search', query: { q: query } })
}

function openMapLink(lat, lon) {
  window.open(`https://www.openstreetmap.org/?mlat=${lat}&mlon=${lon}#map=15/${lat}/${lon}`, '_blank')
}

function focusMapLocation(loc) {
  if (!mapInstance.value) return
  const lat = parseFloat(loc.lat)
  const lon = parseFloat(loc.lon)
  mapInstance.value.setView([lat, lon], 15)
  mapMarkers.value.forEach(m => m.closePopup())
  const marker = mapMarkers.value.find((_, i) => mapResults.value[i] === loc)
  if (marker) marker.openPopup()
}

function toggleResultMenu(idx, event) {
  event.stopPropagation()
  if (activeMenuIdx.value === idx) {
    activeMenuIdx.value = null
  } else {
    activeMenuIdx.value = idx
  }
}

function blockDomainFromResult(url) {
  const domain = getDomain(url)
  addBlockedSite(domain, 'block')
  results.value = results.value.filter(r => getDomain(r.url) !== domain)
  activeMenuIdx.value = null
}

watch(() => route.query.q, (newQ) => {
  if (newQ && newQ !== lastSearchedQuery.value) {
    searchQuery.value = newQ
    activeTab.value = 'all'
    lastSearchedQuery.value = newQ
    performSearch(newQ)
  }
})

watch(searchQuery, (newVal) => {
  const currentQ = route.query.q
  if (newVal && newVal !== currentQ && newVal !== lastSearchedQuery.value) {
    lastSearchedQuery.value = ''
  }
})

watch(previewVisible, (val) => {
  if (val) {
    document.body.style.overflow = 'hidden'
  } else {
    document.body.style.overflow = ''
  }
})

function handleDocumentClick() {
  if (activeMenuIdx.value !== null) {
    activeMenuIdx.value = null
  }
}

const isMobile = ref(window.innerWidth <= 768)

function renderMarkdown(text) {
  if (!text) return ''
  let html = text
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')

  html = html.replace(/```(\w*)\n([\s\S]*?)```/g, '<pre><code class="language-$1">$2</code></pre>')
  html = html.replace(/`([^`\n]+)`/g, '<code>$1</code>')
  html = html.replace(/^#### (.*$)/gm, '<h4>$1</h4>')
  html = html.replace(/^### (.*$)/gm, '<h3>$1</h3>')
  html = html.replace(/^## (.*$)/gm, '<h2>$1</h2>')
  html = html.replace(/^# (.*$)/gm, '<h1>$1</h1>')
  html = html.replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
  html = html.replace(/\*(.*?)\*/g, '<em>$1</em>')
  html = html.replace(/^> (.*$)/gm, '<blockquote>$1</blockquote>')
  html = html.replace(/^- (.*$)/gm, '<li>$1</li>')
  html = html.replace(/^(\d+)\. (.*$)/gm, '<li class="ordered">$1. $2</li>')

  let lines = html.split('\n')
  let result = []
  let inList = false
  let listType = ''

  for (let i = 0; i < lines.length; i++) {
    let line = lines[i].trim()
    if (!line) {
      if (inList) {
        result.push(`</${listType}>`)
        inList = false
      }
      continue
    }
    if (line.startsWith('<li')) {
      if (!inList) {
        listType = 'ul'
        inList = true
        result.push('<ul>')
      }
      result.push(line)
      continue
    }
    if (line.startsWith('<h') || line.startsWith('<pre') || line.startsWith('<blockquote>')) {
      if (inList) {
        result.push(`</${listType}>`)
        inList = false
      }
      result.push(line)
      continue
    }
    if (inList) {
      result.push(line)
    } else {
      result.push(`<p>${line}</p>`)
    }
  }
  if (inList) {
    result.push(`</${listType}>`)
  }

  return result.join('\n')
}

onMounted(async () => {
  await detectLanguage()
  initSettings()
  aiMode.value = localStorage.getItem('synapic_ai_mode') || 'auto'
  document.addEventListener('click', handleDocumentClick)
  window.addEventListener('resize', handleResize)
  const q = route.query.q
  if (q) {
    searchQuery.value = q
    performSearch(q)
  }
})

onUnmounted(() => {
  document.removeEventListener('click', handleDocumentClick)
  window.removeEventListener('resize', handleResize)
  if (mapInstance.value) {
    mapInstance.value.remove()
    mapInstance.value = null
  }
})
</script>

<template>
  <div class="results-page">
    <header class="search-header">
      <div class="header-inner">
        <router-link to="/" class="logo-link">
          <span class="logo-text">Synapic</span>
        </router-link>
        <form class="header-search-form" @submit.prevent="handleSearchSubmit">
          <div class="header-search-bar">
            <svg class="search-icon-svg" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="11" cy="11" r="8"/>
              <line x1="21" y1="21" x2="16.65" y2="16.65"/>
            </svg>
            <input
              ref="inputRef"
              v-model="searchQuery"
              type="text"
              class="header-search-input"
              :placeholder="t.home.placeholder"
              autocomplete="off"
              spellcheck="false"
            />
            <button v-if="searchQuery" type="button" class="clear-btn" @click="clearInput">
              <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/>
              </svg>
            </button>
            <button type="submit" class="search-submit-btn" aria-label="Search">
              <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="11" cy="11" r="8"/>
                <line x1="21" y1="21" x2="16.65" y2="16.65"/>
              </svg>
            </button>
          </div>
        </form>
      </div>
    </header>

    <nav class="tab-bar">
      <div class="tab-bar-inner">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          class="tab-btn"
          :class="{ active: activeTab === tab.id }"
          @click="switchTab(tab.id)"
        >
          <svg v-if="tab.icon === 'web'" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="12" r="10"/><line x1="2" y1="12" x2="22" y2="12"/><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"/>
          </svg>
          <svg v-else-if="tab.icon === 'image'" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/>
          </svg>
          <svg v-else-if="tab.icon === 'news'" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M4 22h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v16a2 2 0 0 1-2 2Zm0 0a2 2 0 0 1-2-2v-9c0-1.1.9-2 2-2h2"/><line x1="10" y1="6" x2="18" y2="6"/><line x1="10" y1="10" x2="18" y2="10"/><line x1="10" y1="14" x2="14" y2="14"/>
          </svg>
          <svg v-else-if="tab.icon === 'maps'" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <polygon points="1 6 1 22 8 18 16 22 23 18 23 2 16 6 8 2 1 6"/>
          </svg>
          <span>{{ tab.label }}</span>
        </button>
      </div>
    </nav>

    <div class="prism-bar" v-if="allPrisms.length > 1">
      <div class="prism-bar-inner">
        <div class="prism-select-wrapper">
          <span class="prism-select-label">{{ lang === 'tr' ? 'Prizma' : 'Prism' }}</span>
          <select
            class="prism-select"
            :value="activePrismId"
            @change="handlePrismSwitch($event.target.value)"
            :aria-label="t.results.prismSelectLabel"
          >
            <option value="default">{{ t.results.prismSelectDefault }}</option>
            <option
              v-for="prism in allPrisms.filter(p => p.id !== 'default')"
              :key="prism.id"
              :value="prism.id"
            >{{ lang.value === 'tr' ? prism.nametr : prism.name }}</option>
          </select>
          <svg class="prism-select-arrow" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="6 9 12 15 18 9"/>
          </svg>
        </div>
      </div>
    </div>

    <main class="content-area">
      <div v-if="loading" class="loading-state">
        <div class="dots-loader">
          <span class="dot"></span>
          <span class="dot"></span>
          <span class="dot"></span>
        </div>
        <p class="loading-text">{{ t.results.searching }}</p>
      </div>

      <template v-else-if="hasSearched">
        <div v-if="activeTab === 'all'" class="all-content">
          <div v-if="weatherLoading" class="widget-loading">
            <div class="widget-spinner"></div>
            <span>{{ t.results.searchLocation }}</span>
          </div>

          <div v-if="weatherData" class="weather-widget">
            <div class="weather-main">
              <div class="weather-icon-large">{{ getWeatherInfo(weatherData.weatherCode).icon }}</div>
              <div class="weather-temp-group">
                <span class="weather-temp">{{ Math.round(weatherData.temp) }}°C</span>
                <span class="weather-desc">{{ getWeatherInfo(weatherData.weatherCode)[lang.value] }}</span>
                <span class="weather-location">{{ weatherData.city }}, {{ weatherData.country }}</span>
              </div>
            </div>
            <div class="weather-details">
              <div class="weather-detail">
                <span class="weather-detail-label">{{ t.results.feelsLike }}</span>
                <span class="weather-detail-value">{{ Math.round(weatherData.feelsLike) }}°C</span>
              </div>
              <div class="weather-detail">
                <span class="weather-detail-label">{{ t.results.wind }}</span>
                <span class="weather-detail-value">{{ weatherData.windSpeed }} km/h</span>
              </div>
              <div class="weather-detail">
                <span class="weather-detail-label">{{ t.results.humidity }}</span>
                <span class="weather-detail-value">%{{ weatherData.humidity }}</span>
              </div>
            </div>
          </div>

          <div v-if="showCalculator" class="calculator-widget">
            <div class="calc-header">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="4" y="2" width="16" height="20" rx="2"/><line x1="8" y1="6" x2="16" y2="6"/><line x1="8" y1="10" x2="8" y2="10.01"/><line x1="12" y1="10" x2="12" y2="10.01"/><line x1="8" y1="14" x2="8" y2="14.01"/><line x1="8" y1="18" x2="8" y2="18.01"/><line x1="12" y1="18" x2="12" y2="18.01"/><line x1="16" y1="18" x2="16" y2="18.01"/>
              </svg>
              <span>{{ t.results.calculator }}</span>
            </div>
            <div class="calc-body">
              <form class="calc-form" @submit.prevent="handleCalc">
                <input v-model="calcInput" type="text" class="calc-input" placeholder="2 + 2 * 3" autocomplete="off" />
                <button type="submit" class="calc-btn">{{ t.results.calculate }}</button>
              </form>
              <div v-if="calcResult" class="calc-result">
                <span class="calc-result-label">{{ t.results.calculatorResult }}</span>
                <span class="calc-result-value">{{ calcResult }}</span>
              </div>
              <div v-if="calcHistory.length > 0" class="calc-history">
                <div v-for="(item, idx) in calcHistory" :key="idx" class="calc-history-item" @click="calcInput = item.expression; handleCalc()">
                  <span class="calc-history-expr">{{ item.expression }}</span>
                  <span class="calc-history-result">= {{ item.result }}</span>
                </div>
              </div>
            </div>
          </div>

          <div v-if="results.length === 0 && !weatherData && !showCalculator && !aiLoading && !aiAnswer" class="empty-state">
            <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/><line x1="8" y1="11" x2="14" y2="11"/>
            </svg>
            <p class="empty-query">{{ t.results.noResults }} "{{ searchQuery }}"</p>
            <p class="empty-hint">{{ t.results.tryDifferent }}</p>
            <button class="try-again-btn" @click="handleSearchSubmit">{{ t.results.tryAgain }}</button>
          </div>

          <div v-if="results.length > 0" class="results-layout">
            <section class="results-list-section">
              <div v-if="aiLoading || aiAnswer" class="ai-answer-section">
                <div class="ai-answer-header" @click="showAiAnswer = !showAiAnswer">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2a10 10 0 1 0 10 10H12V2z"/><path d="M20 12a8 8 0 0 0-8-8v8h8z" opacity="0.5"/></svg>
                  <span>{{ t.results.aiAnswer }}</span>
                  <svg class="ai-chevron" :class="{ collapsed: !showAiAnswer }" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 12 15 18 9"/></svg>
                </div>
                <div v-if="showAiAnswer" class="ai-answer-body">
                  <div v-if="aiLoading" class="widget-loading" style="padding: 20px;">
                    <div class="widget-spinner"></div>
                    <span>{{ t.results.aiGenerating }}</span>
                  </div>
                  <div v-else-if="aiAnswer" class="ai-answer-wrapper">
                    <div class="ai-answer-text" :class="{ collapsed: !aiExpanded }" v-html="renderMarkdown(aiAnswer)"></div>
                    <button v-if="!aiExpanded" class="ai-read-more" @click="aiExpanded = true">{{ t.results.readMore }}</button>
                  </div>
                  <div v-else class="ai-answer-failed">
                    <span>{{ t.results.aiFailed }}</span>
                  </div>
                </div>
              </div>

              <p v-if="resultsCount > 0" class="results-meta">
                {{ resultsCount.toLocaleString() }} {{ t.results.for }} "{{ searchQuery }}"
              </p>

              <div class="results-list">
                <template v-for="(item, idx) in paginatedResults" :key="item.type === 'result' ? 'r-' + item.index : item.type + '-' + idx">
                  <article v-if="item.type === 'result'" class="result-card">
                    <div class="result-header">
                      <img
                        v-if="getDomain(item.data.url)"
                        :src="'https://www.google.com/s2/favicons?domain=' + getDomain(item.data.url) + '&sz=32'"
                        :alt="getDomain(item.data.url)"
                        class="result-favicon"
                      />
                      <div class="result-url-info">
                        <span v-if="getDomain(item.data.url)" class="result-domain">{{ getDomain(item.data.url) }}</span>
                        <span v-if="getPath(item.data.url)" class="result-path">{{ getPath(item.data.url) }}</span>
                      </div>
                      <div class="result-actions">
                        <button class="result-menu-btn" @click="toggleResultMenu(idx, $event)" :title="t.results.siteOptions">
                          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                            <circle cx="12" cy="5" r="1"/><circle cx="12" cy="12" r="1"/><circle cx="12" cy="19" r="1"/>
                          </svg>
                        </button>
                        <div v-if="activeMenuIdx === idx" class="result-dropdown">
                          <button class="result-dropdown-item block" @click.stop="blockDomainFromResult(item.data.url)">
                            <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"/></svg>
                            <span>{{ t.results.blockDomain }}</span>
                          </button>
                        </div>
                      </div>
                    </div>
                    <a :href="item.data.url" target="_blank" rel="noopener noreferrer" class="result-title">{{ item.data.title ? item.data.title.split(' ').slice(0, 35).join(' ') + (item.data.title.split(' ').length > 35 ? '...' : '') : '' }}</a>
                    <p v-if="item.data.description" class="result-desc">{{ item.data.description }}</p>
                    <div v-if="item.data.text" class="result-text-wrapper">
                      <p class="result-text" :class="{ expanded: expandedTexts[idx] }">{{ expandedTexts[idx] ? item.data.text : (item.data.text.length > 220 ? item.data.text.slice(0, 220) + '...' : item.data.text) }}</p>
                      <button v-if="item.data.text.length > 220" @click.prevent="toggleText(idx)" class="text-toggle-btn">
                        {{ expandedTexts[idx] ? (lang === 'tr' ? 'Daha az göster' : 'Show less') : (lang === 'tr' ? 'Metni göster' : 'Show text') }}
                      </button>
                    </div>
                  </article>

                  <section v-else-if="item.type === 'related'" class="inline-related-section">
                    <div class="inline-related-title">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <polyline points="22 12 18 12 15 21 9 3 6 12 2 12"/>
                      </svg>
                      <span>{{ t.results.relatedSearches }}</span>
                    </div>
                    <div class="inline-related-grid">
                      <button
                        v-for="related in item.data"
                        :key="related"
                        class="inline-related-chip"
                        @click="handleRelatedClick(related)"
                      >
                        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                          <circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/>
                        </svg>
                        <span>{{ related }}</span>
                      </button>
                    </div>
                  </section>

                  <section v-else-if="item.type === 'images'" class="inline-images-section">
                    <div class="inline-images-title">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/>
                      </svg>
                      <span>{{ t.results.inlineImages }}</span>
                    </div>
                    <div class="inline-images-scroll">
                      <div
                        v-for="(img, imgIdx) in item.data"
                        :key="img.url || imgIdx"
                        class="inline-image-card"
                        @click="openPreview(img)"
                      >
                        <div class="inline-image-wrapper">
                          <img
                            :src="img.thumbnail || img.url || img.src || img.image"
                            :alt="img.title || 'Image'"
                            loading="lazy"
                            @error="handleImageError"
                          />
                        </div>
                        <span class="inline-image-title">{{ img.title || '' }}</span>
                      </div>
                    </div>
                  </section>
                </template>
              </div>

              <nav v-if="totalPages > 1" class="pagination">
                <button class="page-btn page-nav" :disabled="currentPage === 1" @click="goToPage(1)">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                    <polyline points="11 17 6 12 11 7"/><polyline points="18 17 13 12 18 7"/>
                  </svg>
                </button>
                <button class="page-btn page-nav" :disabled="currentPage === 1" @click="prevPage">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                    <polyline points="15 18 9 12 15 6"/>
                  </svg>
                </button>
                <template v-for="p in paginationRange" :key="p">
                  <span v-if="p === '...'" class="page-ellipsis">...</span>
                  <button v-else class="page-btn" :class="{ active: currentPage === p }" @click="goToPage(p)">{{ p }}</button>
                </template>
                <button class="page-btn page-nav" :disabled="currentPage === totalPages" @click="nextPage">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                    <polyline points="9 18 15 12 9 6"/>
                  </svg>
                </button>
                <button class="page-btn page-nav" :disabled="currentPage === totalPages" @click="goToPage(totalPages)">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                    <polyline points="13 17 18 12 13 7"/><polyline points="6 17 11 12 6 7"/>
                  </svg>
                </button>
              </nav>
            </section>

            <aside v-if="wikiData && wikiData.isValid && !isMobile" class="wiki-panel">
              <div class="wiki-card">
                <div class="wiki-header">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <circle cx="12" cy="12" r="10"/><line x1="12" y1="16" x2="12" y2="12"/><line x1="12" y1="8" x2="12.01" y2="8"/>
                  </svg>
                  <span class="wiki-label">{{ t.results.about }} {{ wikiData.title }}</span>
                </div>
                <img v-if="wikiData.thumbnail" :src="wikiData.thumbnail" :alt="wikiData.title" class="wiki-thumb" loading="lazy" />
                <p class="wiki-extract">{{ wikiData.extract }}</p>
                <a :href="wikiData.url" target="_blank" rel="noopener noreferrer" class="wiki-read-more">{{ t.results.readMore }} &rarr;</a>
              </div>
            </aside>
          </div>
        </div>

        <div v-else-if="activeTab === 'images'" class="images-content">
          <div v-if="imageResults.length === 0" class="empty-state">
            <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
              <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/>
            </svg>
            <p class="empty-query">{{ t.results.noResults }} "{{ searchQuery }}"</p>
            <p class="empty-hint">{{ t.results.tryDifferent }}</p>
          </div>
          <template v-else>
            <p class="results-meta">{{ imageResults.length.toLocaleString() }} {{ t.results.images }} {{ t.results.for }} "{{ searchQuery }}"</p>
            <div class="images-grid">
              <div v-for="(img, index) in imageResults" :key="img.url || index" class="image-card" @click="openPreview(img)">
                <div class="image-wrapper">
                  <img :src="img.thumbnail || img.url || img.src || img.image" :alt="img.title || img.alt || 'Image'" class="image-thumb" loading="lazy" @error="handleImageError" />
                </div>
                <div class="image-meta">
                  <span class="image-title">{{ img.title || '' }}</span>
                  <span class="image-domain">{{ getDomain(img.source_url || img.url || img.src || '') }}</span>
                </div>
              </div>
            </div>
          </template>
        </div>

        <div v-else-if="activeTab === 'news'" class="news-content">
          <div v-if="newsResults.length === 0" class="empty-state">
            <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
              <path d="M4 22h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v16a2 2 0 0 1-2 2Zm0 0a2 2 0 0 1-2-2v-9c0-1.1.9-2 2-2h2"/><line x1="10" y1="6" x2="18" y2="6"/><line x1="10" y1="10" x2="18" y2="10"/><line x1="10" y1="14" x2="14" y2="14"/>
            </svg>
            <p class="empty-query">{{ t.results.noResults }} "{{ searchQuery }}"</p>
            <p class="empty-hint">{{ t.results.tryDifferent }}</p>
          </div>
          <template v-else>
            <p class="results-meta">{{ newsResults.length.toLocaleString() }} {{ t.results.news }} {{ t.results.for }} "{{ searchQuery }}"</p>
            <div class="news-list">
              <article v-for="(news, index) in newsResults" :key="news.url || index" class="news-card">
                <div class="news-header">
                  <img v-if="getDomain(news.url)" :src="'https://www.google.com/s2/favicons?domain=' + getDomain(news.url) + '&sz=32'" :alt="getDomain(news.url)" class="news-favicon" loading="lazy" />
                  <div class="news-source-info">
                    <span class="news-source">{{ getDomain(news.url) }}</span>
                    <span class="news-time">{{ news.date || news.published || news.age || t.results.publishedDate }}</span>
                  </div>
                </div>
                <a :href="news.url" target="_blank" rel="noopener noreferrer" class="news-title">{{ news.title }}</a>
                <p v-if="news.description" class="news-desc">{{ news.description }}</p>
              </article>
            </div>
          </template>
        </div>

        <div v-else-if="activeTab === 'maps'" class="maps-content">
          <div v-if="mapLoading" class="widget-loading">
            <div class="widget-spinner"></div>
            <span>{{ t.results.searchLocation }}</span>
          </div>
          <div class="maps-layout">
            <div class="map-container" ref="mapContainer"></div>
            <div class="map-sidebar">
              <h3 class="map-sidebar-title">{{ t.results.locationResults }}</h3>
              <div v-if="mapResults.length === 0 && !mapLoading" class="map-no-results">{{ t.results.noLocations }} "{{ searchQuery }}"</div>
              <div class="map-results-list">
                <div v-for="(loc, idx) in mapResults" :key="idx" class="map-result-item" @click="focusMapLocation(loc)">
                  <div class="map-result-icon">
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M21 10c0 7-9 13-13a9 9 0 0 1 18 0z"/><circle cx="12" cy="10" r="3"/>
                    </svg>
                  </div>
                  <div class="map-result-info">
                    <span class="map-result-name">{{ loc.display_name.split(',')[0] }}</span>
                    <span class="map-result-address">{{ loc.display_name }}</span>
                  </div>
                  <button class="map-open-btn" @click.stop="openMapLink(loc.lat, loc.lon)" :title="t.results.openInMap">
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"/><polyline points="15 3 21 3 21 9"/><line x1="10" y1="14" x2="21" y2="3"/>
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </template>

      <div v-if="!hasSearched" class="welcome-state">
        <img src="../images/synapic.png" alt="Synapic" class="welcome-logo" />
        <p class="welcome-text">{{ t.home.placeholder }}</p>
      </div>
    </main>

    <Transition name="preview-fade">
      <div v-if="previewVisible && previewImage" class="preview-overlay" @click.self="closePreview">
        <div class="preview-panel">
          <button class="preview-close" @click="closePreview">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/>
            </svg>
          </button>
          <div class="preview-image-wrapper">
            <img :src="previewImage.url || previewImage.src || previewImage.image" :alt="previewImage.title || 'Preview'" class="preview-image" />
          </div>
          <div class="preview-info">
            <h3 class="preview-title">{{ previewImage.title || 'Untitled' }}</h3>
            <span class="preview-domain">{{ getDomain(previewImage.source_url || previewImage.url || previewImage.src || '') }}</span>
            <div class="preview-actions">
              <a :href="previewImage.source_url || previewImage.url || previewImage.src" target="_blank" rel="noopener noreferrer" class="preview-source-btn">
                <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"/><polyline points="15 3 21 3 21 9"/><line x1="10" y1="14" x2="21" y2="3"/>
                </svg>
                {{ t.results.viewSource }}
              </a>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=DM+Sans:wght@400;500;600;700&family=DM+Mono:wght@400;500&display=swap');

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.results-page {
  min-height: 100vh;
  background: var(--bg-primary);
  font-family: 'DM Sans', sans-serif;
  color: var(--text-primary);
  transition: background-color 0.3s ease, color 0.3s ease;
  overflow-x: hidden;
}

.search-header {
  position: sticky;
  top: 0;
  z-index: 100;
  background: var(--bg-primary);
  border-bottom: 1px solid var(--border-color);
  padding: 10px 20px;
  transition: background-color 0.3s ease, border-color 0.3s ease;
}

.header-inner {
  max-width: 960px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  gap: 16px;
}

.logo-link {
  display: flex;
  align-items: center;
  gap: 7px;
  text-decoration: none;
  flex-shrink: 0;
  color: var(--text-primary);
  transition: color 0.3s ease;
}

.logo-text {
  font-size: 1.05rem;
  font-weight: 700;
  color: var(--text-primary);
  letter-spacing: -0.01em;
}

.header-search-form {
  flex: 1;
  min-width: 0;
  max-width: 560px;
}

.header-search-bar {
  display: flex;
  align-items: center;
  gap: 8px;
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: 20px;
  padding: 8px 12px 8px 14px;
  transition: border-color 0.2s ease, box-shadow 0.2s ease, background-color 0.3s ease;
}

.header-search-bar:focus-within {
  border-color: var(--text-muted);
  box-shadow: 0 0 0 2px var(--result-hover);
}

.search-icon-svg {
  flex-shrink: 0;
  color: var(--text-muted);
}

.header-search-input {
  flex: 1;
  background: none;
  border: none;
  outline: none;
  font-family: 'DM Sans', sans-serif;
  font-size: 0.9rem;
  color: var(--text-primary);
  min-width: 0;
  transition: color 0.3s ease;
}

.header-search-input::placeholder {
  color: var(--text-muted);
}

.clear-btn {
  background: none;
  border: none;
  color: var(--text-muted);
  cursor: pointer;
  padding: 2px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  transition: color 0.15s, background 0.15s;
}

.clear-btn:hover {
  color: var(--text-primary);
  background: var(--result-hover);
}

.search-submit-btn {
  background: var(--bg-hover);
  border: none;
  cursor: pointer;
  width: 30px;
  height: 30px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  color: var(--text-muted);
  transition: background 0.15s ease, color 0.15s ease;
}

.search-submit-btn:hover {
  background: var(--border-light);
  color: var(--text-primary);
}

.tab-bar {
  position: sticky;
  top: 49px;
  z-index: 99;
  background: var(--bg-primary);
  border-bottom: 1px solid var(--border-color);
  transition: background-color 0.3s ease, border-color 0.3s ease;
}

.tab-bar-inner {
  max-width: 960px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  gap: 0;
  padding: 0 20px;
}

.tab-btn {
  display: flex;
  align-items: center;
  gap: 5px;
  padding: 10px 14px;
  background: none;
  border: none;
  border-bottom: 2px solid transparent;
  color: var(--text-muted);
  font-family: 'DM Sans', sans-serif;
  font-size: 0.82rem;
  font-weight: 500;
  cursor: pointer;
  transition: color 0.2s, border-color 0.2s;
  white-space: nowrap;
}

.tab-btn:hover {
  color: var(--text-secondary);
}

.tab-btn.active {
  color: var(--accent);
  border-bottom-color: var(--accent);
}

.content-area {
  max-width: 960px;
  margin: 0 auto;
  padding: 16px 20px 80px;
  overflow-x: hidden;
}

.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 20px;
  gap: 16px;
}

.dots-loader {
  display: flex;
  gap: 6px;
}

.prism-bar {
  background: var(--bg-primary);
  border-bottom: 1px solid var(--border-color);
  position: sticky;
  top: 89px;
  z-index: 97;
}

.prism-bar-inner {
  display: flex;
  align-items: center;
  padding: 7px 20px;
  max-width: 960px;
  margin: 0 auto;
}

.prism-select-wrapper {
  display: flex;
  align-items: center;
  gap: 7px;
  background: var(--bg-secondary);
  border: 1.5px solid var(--border-color);
  border-radius: 10px;
  padding: 5px 10px 5px 10px;
  cursor: pointer;
  transition: border-color 0.2s, box-shadow 0.2s;
  min-width: 140px;
  max-width: 240px;
}

.prism-select-wrapper:focus-within {
  border-color: var(--accent);
  box-shadow: 0 0 0 2px var(--result-hover);
}

.prism-select-label {
  font-size: 0.75rem;
  font-weight: 600;
  color: var(--accent);
  white-space: nowrap;
  flex-shrink: 0;
  letter-spacing: 0.01em;
}

.prism-select-arrow {
  flex-shrink: 0;
  color: var(--text-muted);
  pointer-events: none;
}

.prism-select {
  flex: 1;
  background: none;
  border: none;
  outline: none;
  font-family: 'DM Sans', sans-serif;
  font-size: 0.82rem;
  font-weight: 500;
  color: var(--text-primary);
  cursor: pointer;
  appearance: none;
  -webkit-appearance: none;
  min-width: 0;
}

.prism-select option {
  background: var(--bg-secondary);
  color: var(--text-primary);
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: var(--dot-color);
  animation: dotPulse 1.4s ease-in-out infinite both;
}

.dot:nth-child(2) { animation-delay: 0.16s; }
.dot:nth-child(3) { animation-delay: 0.32s; }

@keyframes dotPulse {
  0%, 80%, 100% { transform: scale(0.4); opacity: 0.3; }
  40% { transform: scale(1); opacity: 1; }
}

.loading-text {
  font-size: 0.88rem;
  color: var(--text-muted);
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 20px;
  gap: 10px;
  color: var(--text-muted);
}

.empty-query {
  font-size: 1rem;
  color: var(--text-primary);
}

.empty-hint {
  font-size: 0.85rem;
  color: var(--text-muted);
  margin-bottom: 6px;
}

.try-again-btn {
  padding: 8px 20px;
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: 18px;
  color: var(--accent);
  font-family: 'DM Sans', sans-serif;
  font-size: 0.85rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s, border-color 0.2s;
}

.try-again-btn:hover {
  background: var(--bg-hover);
  border-color: var(--border-light);
}

.results-layout {
  display: flex;
  gap: 24px;
  align-items: flex-start;
}

.results-list-section {
  flex: 1;
  min-width: 0;
}

.pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
  margin-top: 24px;
  padding: 12px 0;
}

.page-btn {
  min-width: 34px;
  height: 34px;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 6px;
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-secondary);
  font-family: 'DM Sans', sans-serif;
  font-size: 0.82rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.15s, border-color 0.15s, color 0.15s;
}

.page-btn:hover:not(:disabled):not(.active) {
  background: var(--bg-hover);
  border-color: var(--border-light);
  color: var(--text-primary);
}

.page-btn.active {
  background: var(--accent);
  border-color: var(--accent);
  color: var(--bg-primary);
  font-weight: 600;
}

.page-btn:disabled {
  opacity: 0.3;
  cursor: not-allowed;
}

.page-nav {
  padding: 0 8px;
}

.page-ellipsis {
  width: 28px;
  text-align: center;
  color: var(--text-muted);
  font-size: 0.82rem;
  user-select: none;
}

.results-meta {
  font-size: 0.78rem;
  color: var(--text-muted);
  margin-bottom: 14px;
  font-family: 'DM Mono', monospace;
}

.results-list {
  display: flex;
  flex-direction: column;
  gap: 2px;
  max-width: 620px;
  overflow: hidden;
}

.result-card {
  background: transparent;
  border-radius: 8px;
  padding: 10px 12px;
  transition: background 0.15s;
}

.result-card:hover {
  background: var(--result-hover);
}

.result-header {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 2px;
  position: relative;
  padding-right: 36px;
  overflow: visible;
}

.result-favicon {
  width: 16px;
  height: 16px;
  border-radius: 2px;
  flex-shrink: 0;
  display: block;
}

.result-url-info {
  display: flex;
  align-items: center;
  gap: 4px;
  min-width: 0;
  flex: 1;
  overflow: hidden;
}

.result-domain {
  font-size: 0.76rem;
  color: var(--text-secondary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  min-width: 0;
}

.result-path {
  font-size: 0.72rem;
  color: var(--text-muted);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: 'DM Mono', monospace;
}

.result-actions {
  flex-shrink: 0;
  position: absolute;
  right: 0;
  top: 50%;
  transform: translateY(-50%);
}

.result-menu-btn {
  width: 26px;
  height: 26px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: none;
  border: none;
  border-radius: 6px;
  color: var(--text-muted);
  cursor: pointer;
  transition: background 0.15s, color 0.15s;
  opacity: 0;
}

.result-card:hover .result-menu-btn {
  opacity: 1;
}

.result-menu-btn:hover {
  background: var(--bg-hover);
  color: var(--text-primary);
}

.result-dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 4px;
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: 10px;
  padding: 4px;
  min-width: 180px;
  z-index: 200;
  box-shadow: 0 4px 20px rgba(0,0,0,0.18);
  animation: dropIn 0.15s ease both;
}

@keyframes dropIn {
  from { opacity: 0; transform: translateY(-4px) scale(0.97); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}

.result-dropdown-item {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
  padding: 8px 12px;
  background: none;
  border: none;
  border-radius: 6px;
  color: var(--text-secondary);
  font-family: 'DM Sans', sans-serif;
  font-size: 0.82rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.12s, color 0.12s;
  white-space: nowrap;
}

.result-dropdown-item svg {
  flex-shrink: 0;
}

.result-dropdown-item.block:hover {
  background: rgba(239, 68, 68, 0.1);
  color: var(--danger);
}

.result-title {
  display: block;
  font-size: 0.95rem;
  font-weight: 500;
  color: var(--accent);
  text-decoration: none;
  line-height: 1.35;
  margin-bottom: 2px;
  transition: color 0.15s;
  word-break: break-word;
  white-space: normal;
  overflow-wrap: break-word;
}

.result-title:hover {
  text-decoration: underline;
  color: var(--accent-hover);
}

.result-desc {
  font-size: 0.82rem;
  color: var(--text-secondary);
  line-height: 1.5;
  word-break: break-word;
}

.inline-related-section {
  margin: 16px 0;
  padding: 14px 16px;
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: 10px;
  overflow: hidden;
  min-width: 0;
}

.inline-related-title {
  display: flex;
  align-items: center;
  gap: 7px;
  margin-bottom: 10px;
  font-size: 0.85rem;
  font-weight: 600;
  color: var(--text-secondary);
}

.inline-related-title svg {
  flex-shrink: 0;
  color: var(--accent);
}

.inline-related-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  min-width: 0;
}

.inline-related-chip {
  display: flex;
  align-items: center;
  gap: 7px;
  padding: 8px 14px;
  background: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: 20px;
  color: var(--text-primary);
  font-family: 'DM Sans', sans-serif;
  font-size: 0.82rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.15s, border-color 0.15s, color 0.15s;
  white-space: nowrap;
}

.inline-related-chip:hover {
  background: var(--bg-hover);
  border-color: var(--accent);
  color: var(--accent);
}

.inline-related-chip svg {
  flex-shrink: 0;
  color: var(--text-muted);
  transition: color 0.15s;
}

.inline-related-chip:hover svg {
  color: var(--accent);
}

.inline-images-section {
  margin: 16px 0;
  overflow: hidden;
  min-width: 0;
}

.inline-images-title {
  display: flex;
  align-items: center;
  gap: 7px;
  margin-bottom: 10px;
  font-size: 0.85rem;
  font-weight: 600;
  color: var(--text-secondary);
}

.inline-images-title svg {
  flex-shrink: 0;
  color: var(--accent);
}

.inline-images-scroll {
  display: flex;
  gap: 8px;
  overflow-x: auto;
  padding-bottom: 6px;
  scroll-snap-type: x mandatory;
  width: 100%;
  min-width: 0;
}

.inline-images-scroll::-webkit-scrollbar {
  height: 4px;
}

.inline-images-scroll::-webkit-scrollbar-track {
  background: transparent;
}

.inline-images-scroll::-webkit-scrollbar-thumb {
  background: var(--scrollbar-thumb);
  border-radius: 2px;
}

.inline-image-card {
  flex-shrink: 0;
  width: 120px;
  cursor: pointer;
  scroll-snap-align: start;
  border-radius: 8px;
  overflow: hidden;
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  transition: border-color 0.2s, transform 0.2s;
}

.inline-image-card:hover {
  border-color: var(--border-light);
  transform: translateY(-2px);
}

.inline-image-wrapper {
  width: 100%;
  aspect-ratio: 1;
  background: var(--bg-primary);
  overflow: hidden;
  position: relative;
}

.inline-image-wrapper img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s;
}

.inline-image-card:hover .inline-image-wrapper img {
  transform: scale(1.05);
}

.inline-image-title {
  display: block;
  padding: 5px 7px;
  font-size: 0.68rem;
  color: var(--text-secondary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  line-height: 1.3;
}

.inline-image-wrapper .image-placeholder {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-secondary);
  color: var(--text-muted);
  width: 100%;
  height: 100%;
}

.inline-image-wrapper .image-placeholder svg {
  width: 24px;
  height: 24px;
}

.wiki-panel {
  width: 280px;
  flex-shrink: 0;
  position: sticky;
  top: 100px;
}

.wiki-card {
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: 10px;
  padding: 16px;
}

.wiki-header {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 10px;
  color: var(--accent);
}

.wiki-label {
  font-size: 0.78rem;
  color: var(--accent);
  font-weight: 500;
}

.wiki-thumb {
  width: 100%;
  border-radius: 6px;
  margin-bottom: 10px;
  aspect-ratio: 16 / 10;
  object-fit: cover;
  background: var(--bg-primary);
}

.wiki-extract {
  font-size: 0.8rem;
  color: var(--text-secondary);
  line-height: 1.55;
  margin-bottom: 10px;
  max-height: 180px;
  overflow-y: auto;
}

.wiki-extract::-webkit-scrollbar {
  width: 3px;
}

.wiki-extract::-webkit-scrollbar-track {
  background: transparent;
}

.wiki-extract::-webkit-scrollbar-thumb {
  background: var(--scrollbar-thumb);
  border-radius: 2px;
}

.wiki-read-more {
  font-size: 0.8rem;
  color: var(--accent);
  text-decoration: none;
  font-weight: 500;
  transition: color 0.15s;
}

.wiki-read-more:hover {
  color: var(--accent-hover);
  text-decoration: underline;
}

.images-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  gap: 8px;
  margin-top: 10px;
}

.image-card {
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  overflow: hidden;
  cursor: pointer;
  transition: border-color 0.2s, transform 0.2s;
}

.image-card:hover {
  border-color: var(--border-light);
  transform: translateY(-1px);
}

.image-wrapper {
  position: relative;
  width: 100%;
  aspect-ratio: 1;
  background: var(--bg-primary);
  overflow: hidden;
}

.image-thumb {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s;
}

.image-card:hover .image-thumb {
  transform: scale(1.03);
}

.image-placeholder {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-secondary);
  color: var(--text-muted);
  width: 100%;
  height: 100%;
}

.image-placeholder svg {
  width: 32px;
  height: 32px;
}

.image-meta {
  padding: 6px 8px;
}

.image-title {
  display: block;
  font-size: 0.74rem;
  font-weight: 500;
  color: var(--text-primary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  line-height: 1.3;
}

.image-domain {
  display: block;
  font-size: 0.68rem;
  color: var(--text-muted);
  margin-top: 1px;
  font-family: 'DM Mono', monospace;
}

.news-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  max-width: 620px;
  margin-top: 10px;
}

.news-card {
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  padding: 12px 14px;
  transition: border-color 0.15s, background 0.15s;
}

.news-card:hover {
  border-color: var(--border-light);
  background: var(--result-hover);
}

.news-header {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 4px;
}

.news-favicon {
  width: 14px;
  height: 14px;
  border-radius: 2px;
  flex-shrink: 0;
}

.news-source-info {
  display: flex;
  align-items: center;
  gap: 6px;
}

.news-source {
  font-size: 0.76rem;
  font-weight: 500;
  color: var(--text-secondary);
}

.news-time {
  font-size: 0.7rem;
  color: var(--text-muted);
  font-family: 'DM Mono', monospace;
}

.news-title {
  display: block;
  font-size: 0.92rem;
  font-weight: 600;
  color: var(--accent);
  text-decoration: none;
  line-height: 1.35;
  margin-bottom: 4px;
  transition: color 0.15s;
}

.news-title:hover {
  color: var(--accent-hover);
  text-decoration: underline;
}

.news-desc {
  font-size: 0.8rem;
  color: var(--text-secondary);
  line-height: 1.5;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.weather-widget {
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 20px;
}

.weather-main {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 16px;
}

.weather-icon-large {
  font-size: 3rem;
  line-height: 1;
}

.weather-temp-group {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.weather-temp {
  font-size: 2.2rem;
  font-weight: 700;
  color: var(--text-primary);
  line-height: 1.1;
}

.weather-desc {
  font-size: 0.88rem;
  color: var(--text-secondary);
  font-weight: 500;
}

.weather-location {
  font-size: 0.8rem;
  color: var(--text-muted);
}

.weather-details {
  display: flex;
  gap: 20px;
  padding-top: 14px;
  border-top: 1px solid var(--border-color);
}

.weather-detail {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.weather-detail-label {
  font-size: 0.72rem;
  color: var(--text-muted);
  text-transform: uppercase;
  letter-spacing: 0.3px;
}

.weather-detail-value {
  font-size: 0.92rem;
  font-weight: 600;
  color: var(--text-primary);
}

.calculator-widget {
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  overflow: hidden;
  margin-bottom: 20px;
}

.calc-header {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 14px 18px;
  background: var(--bg-hover);
  border-bottom: 1px solid var(--border-color);
  font-size: 0.85rem;
  font-weight: 600;
  color: var(--text-primary);
}

.calc-header svg {
  color: var(--accent);
}

.calc-body {
  padding: 16px 18px;
}

.calc-form {
  display: flex;
  gap: 8px;
  margin-bottom: 12px;
}

.calc-input {
  flex: 1;
  padding: 10px 14px;
  background: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-primary);
  font-family: 'DM Mono', monospace;
  font-size: 0.95rem;
  outline: none;
  transition: border-color 0.2s;
}

.calc-input:focus {
  border-color: var(--accent);
}

.calc-input::placeholder {
  color: var(--text-muted);
}

.calc-btn {
  padding: 10px 18px;
  background: var(--accent);
  color: var(--bg-primary);
  border: none;
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 0.85rem;
  font-weight: 600;
  cursor: pointer;
  transition: opacity 0.15s;
}

.calc-btn:hover {
  opacity: 0.85;
}

.calc-result {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 14px;
  background: var(--bg-primary);
  border-radius: 8px;
  margin-bottom: 12px;
}

.calc-result-label {
  font-size: 0.75rem;
  color: var(--text-muted);
  text-transform: uppercase;
}

.calc-result-value {
  font-size: 1.3rem;
  font-weight: 700;
  color: var(--accent);
  font-family: 'DM Mono', monospace;
}

.calc-history {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.calc-history-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 6px 10px;
  border-radius: 6px;
  cursor: pointer;
  transition: background 0.15s;
}

.calc-history-item:hover {
  background: var(--bg-hover);
}

.calc-history-expr {
  font-size: 0.8rem;
  color: var(--text-muted);
  font-family: 'DM Mono', monospace;
}

.calc-history-result {
  font-size: 0.85rem;
  font-weight: 600;
  color: var(--text-primary);
  font-family: 'DM Mono', monospace;
}

.widget-loading {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  padding: 30px 20px;
  color: var(--text-muted);
  font-size: 0.85rem;
  margin-bottom: 16px;
}

.widget-spinner {
  width: 18px;
  height: 18px;
  border: 2px solid var(--border-color);
  border-top-color: var(--text-muted);
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.maps-layout {
  display: flex;
  gap: 0;
  height: calc(100vh - 160px);
  min-height: 400px;
  border: 1px solid var(--border-color);
  border-radius: 10px;
  overflow: hidden;
}

.map-container {
  flex: 1;
  min-width: 0;
  z-index: 1;
}

.map-sidebar {
  width: 280px;
  flex-shrink: 0;
  background: var(--bg-secondary);
  border-left: 1px solid var(--border-color);
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.map-sidebar-title {
  font-size: 0.82rem;
  font-weight: 600;
  color: var(--text-primary);
  padding: 12px 14px;
  border-bottom: 1px solid var(--border-color);
  flex-shrink: 0;
}

.map-no-results {
  padding: 20px 14px;
  text-align: center;
  color: var(--text-muted);
  font-size: 0.82rem;
}

.map-results-list {
  flex: 1;
  overflow-y: auto;
  padding: 6px;
}

.map-results-list::-webkit-scrollbar {
  width: 4px;
}

.map-results-list::-webkit-scrollbar-track {
  background: transparent;
}

.map-results-list::-webkit-scrollbar-thumb {
  background: var(--scrollbar-thumb);
  border-radius: 2px;
}

.map-result-item {
  display: flex;
  align-items: flex-start;
  gap: 8px;
  padding: 10px 8px;
  border-radius: 8px;
  cursor: pointer;
  transition: background 0.15s;
}

.map-result-item:hover {
  background: var(--bg-hover);
}

.map-result-icon {
  flex-shrink: 0;
  padding-top: 2px;
  color: var(--accent);
}

.map-result-info {
  flex: 1;
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.map-result-name {
  font-size: 0.82rem;
  font-weight: 600;
  color: var(--text-primary);
}

.map-result-address {
  font-size: 0.72rem;
  color: var(--text-muted);
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.map-open-btn {
  flex-shrink: 0;
  width: 26px;
  height: 26px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 5px;
  color: var(--text-muted);
  cursor: pointer;
  transition: all 0.15s ease;
  margin-top: 1px;
}

.map-open-btn:hover {
  background: var(--border-color);
  color: var(--accent);
  border-color: var(--border-light);
}

.preview-overlay {
  position: fixed;
  inset: 0;
  z-index: 1000;
  background: var(--overlay-bg);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  backdrop-filter: blur(4px);
}

.preview-panel {
  position: relative;
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  max-width: 680px;
  width: 100%;
  max-height: 90vh;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.preview-close {
  position: absolute;
  top: 10px;
  right: 10px;
  z-index: 10;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(0,0,0,0.6);
  border: 1px solid var(--border-light);
  border-radius: 50%;
  color: var(--text-primary);
  cursor: pointer;
  transition: background 0.15s, color 0.15s;
}

.preview-close:hover {
  background: rgba(255,255,255,0.1);
}

.preview-image-wrapper {
  width: 100%;
  max-height: 55vh;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-primary);
}

.preview-image {
  max-width: 100%;
  max-height: 55vh;
  object-fit: contain;
}

.preview-info {
  padding: 14px 18px;
}

.preview-title {
  font-size: 0.95rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 4px;
}

.preview-domain {
  font-size: 0.78rem;
  color: var(--text-muted);
  display: block;
  margin-bottom: 10px;
}

.preview-actions {
  display: flex;
  gap: 8px;
}

.preview-source-btn {
  display: flex;
  align-items: center;
  gap: 5px;
  padding: 6px 12px;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 6px;
  color: var(--text-secondary);
  text-decoration: none;
  font-size: 0.78rem;
  font-weight: 500;
  transition: all 0.15s ease;
}

.preview-source-btn:hover {
  background: var(--border-color);
  color: var(--accent);
}

.preview-fade-enter-active,
.preview-fade-leave-active {
  transition: opacity 0.25s ease;
}

.preview-fade-enter-from,
.preview-fade-leave-to {
  opacity: 0;
}

.welcome-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 50vh;
  gap: 16px;
  color: var(--text-muted);
}

.welcome-logo {
  width: 56px;
  height: 56px;
  object-fit: contain;
  border-radius: 12px;
  opacity: 0.5;
}

.welcome-text {
  font-size: 0.95rem;
  color: var(--text-muted);
}

.ai-answer-section {
  margin-bottom: 16px;
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-left: 4px solid var(--accent);
  border-radius: 10px;
  overflow: hidden;
  transition: background-color 0.3s ease, border-color 0.3s ease;
}

.ai-answer-header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 16px;
  cursor: pointer;
  color: var(--text-primary);
  transition: background 0.2s ease;
  user-select: none;
}

.ai-answer-header:hover {
  background: var(--bg-hover);
}

.ai-answer-header svg:first-child {
  color: var(--accent);
  flex-shrink: 0;
}

.ai-answer-header span {
  font-size: 0.88rem;
  font-weight: 600;
  flex: 1;
}

.ai-chevron {
  color: var(--text-muted);
  transition: transform 0.2s ease;
  flex-shrink: 0;
}

.ai-chevron.collapsed {
  transform: rotate(-90deg);
}

.ai-answer-body {
  border-top: 1px solid var(--border-color);
  padding: 16px;
}

.ai-answer-text {
  font-size: 0.88rem;
  line-height: 1.65;
  color: var(--text-secondary);
  word-break: break-word;
  overflow-wrap: break-word;
}

.ai-answer-text.collapsed {
  display: -webkit-box;
  -webkit-line-clamp: 4;
  -webkit-box-orient: vertical;
  overflow: hidden;
  max-height: 7.2em;
  position: relative;
}

.ai-answer-wrapper {
  position: relative;
}

.ai-read-more {
  margin-top: 10px;
  background: none;
  border: none;
  color: var(--accent);
  font-size: 0.84rem;
  font-weight: 600;
  cursor: pointer;
  padding: 4px 0;
  transition: opacity 0.2s;
}

.ai-read-more:hover {
  opacity: 0.8;
  text-decoration: underline;
}

.ai-answer-text :deep(p) {
  margin-bottom: 10px;
}

.ai-answer-text :deep(p:last-child) {
  margin-bottom: 0;
}

.ai-answer-text :deep(ul), .ai-answer-text :deep(ol) {
  margin-left: 20px;
  margin-bottom: 10px;
}

.ai-answer-text :deep(li) {
  margin-bottom: 4px;
}

.ai-answer-text :deep(strong) {
  color: var(--text-primary);
  font-weight: 600;
}

.ai-answer-text :deep(h1), .ai-answer-text :deep(h2), .ai-answer-text :deep(h3) {
  color: var(--text-primary);
  margin-top: 12px;
  margin-bottom: 8px;
}

.ai-answer-failed {
  padding: 8px 0;
  font-size: 0.82rem;
  color: var(--text-muted);
  text-align: center;
}

@media (max-width: 768px) {
  .search-header { padding: 8px 12px; }
  .header-inner { gap: 10px; }
  .header-search-form { max-width: none; }
  .header-search-bar { padding: 8px 10px 8px 12px; border-radius: 24px; }
  .header-search-input { font-size: 16px; }
  .clear-btn { width: 36px; height: 36px; min-width: 36px; }
  .search-submit-btn { width: 36px; height: 36px; min-width: 36px; }
  .tab-bar { top: 46px; }
  .tab-bar-inner { padding: 0 8px; overflow-x: auto; -webkit-overflow-scrolling: touch; scrollbar-width: none; }
  .tab-bar-inner::-webkit-scrollbar { display: none; }
  .tab-btn { padding: 10px 14px; font-size: 0.78rem; min-height: 44px; flex-shrink: 0; }
  .content-area { max-width: 100%; padding: 12px 12px 60px; }
  .results-layout { flex-direction: column; gap: 16px; }
  .results-list { max-width: 100%; gap: 8px; overflow: visible; }
  .results-list-section { min-width: 0; overflow: visible; }
  .result-card { background: var(--bg-secondary); border: 1px solid var(--border-color); border-radius: 8px; padding: 10px 12px; min-width: 0; overflow: visible; }
  .result-card:hover { border-color: var(--border-light); background: var(--result-hover); }
  .result-header { gap: 6px; margin-bottom: 2px; min-width: 0; justify-content: flex-start; width: 100%; position: relative; padding-right: 36px; overflow: visible; }
  .result-favicon { width: 16px; height: 16px; }
  .result-title { font-size: 0.92rem; font-weight: 600; line-height: 1.35; margin-bottom: 2px; display: block; word-break: break-word; white-space: normal; }
  .result-desc { font-size: 0.8rem; color: var(--text-secondary); line-height: 1.5; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; word-break: break-word; }
  .result-path { max-width: 100px; flex-shrink: 1; min-width: 0; }
  .result-domain { max-width: 180px; }
  .result-url-info { gap: 5px; min-width: 0; flex: 1; overflow: hidden; justify-content: flex-start; }
  .result-menu-btn { opacity: 1; width: 36px; height: 36px; border-radius: 8px; }
  .result-dropdown { min-width: 180px; right: -12px; }
  .result-dropdown-item { min-height: 44px; padding: 10px 12px; }
  .wiki-panel { width: 100%; position: static; order: 2; }
  .wiki-card { padding: 14px; }
  .wiki-header { margin-bottom: 8px; }
  .wiki-thumb { margin-bottom: 8px; }
  .wiki-extract { font-size: 0.82rem; line-height: 1.55; max-height: 160px; }
  .wiki-read-more { display: inline-block; padding-top: 4px; }
  .maps-layout { flex-direction: column; height: auto; min-height: 350px; }
  .map-container { height: 280px; min-height: 250px; }
  .map-sidebar { width: 100%; border-left: none; border-top: 1px solid var(--border-color); max-height: 280px; }
  .map-sidebar-title { padding: 10px 12px; font-size: 0.8rem; min-height: 44px; display: flex; align-items: center; }
  .map-result-item { padding: 10px 10px; min-height: 48px; gap: 10px; }
  .map-result-name { font-size: 0.82rem; }
  .map-result-address { font-size: 0.72rem; }
  .map-open-btn { width: 36px; height: 36px; min-width: 36px; border-radius: 8px; }
  .images-grid { grid-template-columns: repeat(auto-fill, minmax(110px, 1fr)); gap: 6px; }
  .image-card { border-radius: 8px; }
  .image-meta { padding: 6px; }
  .news-list { max-width: 100%; }
  .news-card { padding: 12px 10px; }
  .news-header { min-height: 28px; }
  .news-title { font-size: 0.9rem; line-height: 1.35; min-height: 44px; }
  .news-desc { font-size: 0.82rem; -webkit-line-clamp: 3; }
  .weather-widget { padding: 14px; border-radius: 10px; }
  .weather-main { gap: 12px; margin-bottom: 12px; }
  .weather-icon-large { font-size: 2.5rem; }
  .weather-temp { font-size: 2rem; }
  .weather-details { flex-wrap: wrap; gap: 12px; padding-top: 12px; }
  .weather-detail { min-width: 80px; }
  .calculator-widget { margin-bottom: 16px; }
  .calc-header { padding: 12px 14px; font-size: 0.82rem; min-height: 44px; }
  .calc-body { padding: 14px; }
  .calc-form { gap: 8px; margin-bottom: 10px; }
  .calc-input { padding: 10px 12px; font-size: 16px; min-height: 44px; border-radius: 8px; }
  .calc-btn { padding: 10px 16px; font-size: 0.85rem; min-height: 44px; min-width: 44px; border-radius: 8px; }
  .calc-result { padding: 10px 12px; min-height: 44px; border-radius: 8px; }
  .calc-history-item { padding: 8px 10px; min-height: 40px; }
  .inline-related-section { padding: 12px; margin: 12px 0; border-radius: 8px; overflow: hidden; min-width: 0; }
  .inline-related-title { margin-bottom: 8px; font-size: 0.82rem; }
  .inline-related-grid { display: flex; flex-wrap: nowrap; overflow-x: auto; gap: 8px; padding-bottom: 4px; -webkit-overflow-scrolling: touch; scroll-snap-type: x mandatory; scrollbar-width: none; width: 100%; min-width: 0; }
  .inline-related-grid::-webkit-scrollbar { display: none; }
  .inline-related-chip { padding: 8px 12px; font-size: 0.78rem; min-height: 40px; flex-shrink: 0; scroll-snap-align: start; border-radius: 20px; }
  .inline-images-section { margin: 12px 0; overflow: hidden; min-width: 0; }
  .inline-images-scroll { gap: 8px; padding-bottom: 4px; overflow-x: auto; -webkit-overflow-scrolling: touch; scrollbar-width: none; width: 100%; min-width: 0; }
  .inline-images-scroll::-webkit-scrollbar { display: none; }
  .inline-image-card { width: 110px; min-height: 130px; }
  .inline-image-title { font-size: 0.66rem; padding: 4px 6px; }
  .pagination { gap: 4px; margin-top: 20px; padding: 8px 0; overflow-x: auto; -webkit-overflow-scrolling: touch; scrollbar-width: none; justify-content: flex-start; }
  .pagination::-webkit-scrollbar { display: none; }
  .page-btn { min-width: 44px; height: 44px; font-size: 0.82rem; border-radius: 10px; flex-shrink: 0; }
  .page-nav { padding: 0 8px; min-width: 44px; }
  .page-ellipsis { width: 44px; height: 44px; display: flex; align-items: center; justify-content: center; flex-shrink: 0; }
  .preview-overlay { padding: 12px; padding-top: 48px; align-items: flex-start; backdrop-filter: blur(6px); }
  .preview-panel { max-width: 100%; max-height: calc(100vh - 60px); border-radius: 12px; }
  .preview-close { width: 40px; height: 40px; top: -52px; right: 0; border-radius: 50%; background: rgba(0,0,0,0.7); }
  .preview-image-wrapper { max-height: 50vh; }
  .preview-image { max-height: 50vh; }
  .preview-info { padding: 12px 14px; margin-top: 0; }
  .preview-title { font-size: 0.9rem; line-height: 1.35; }
  .preview-domain { font-size: 0.76rem; margin-bottom: 8px; }
  .preview-source-btn { padding: 10px 14px; font-size: 0.82rem; min-height: 44px; border-radius: 8px; }
  .empty-state { padding: 40px 16px; }
  .try-again-btn { padding: 12px 24px; min-height: 44px; font-size: 0.88rem; border-radius: 22px; }
  .widget-loading { padding: 24px 16px; min-height: 44px; }
  .welcome-state { min-height: 40vh; }
  .welcome-logo { width: 48px; height: 48px; }
  .results-meta { margin-bottom: 10px; }
  .map-no-results { padding: 16px 12px; min-height: 44px; display: flex; align-items: center; justify-content: center; }
  .loading-state { padding: 50px 20px; }
  .loading-text { font-size: 0.85rem; }
  .prism-bar { top: 80px; }
  .prism-bar-inner { padding: 6px 12px; gap: 5px; }
  .ai-answer-section { border-radius: 10px; margin-bottom: 12px; }
  .ai-answer-header { padding: 12px; min-height: 44px; }
  .ai-answer-body { padding: 14px 12px; }
  .ai-answer-text { font-size: 0.84rem; }
}

@media (max-width: 480px) {
  .search-header { padding: 6px 8px; }
  .header-inner { gap: 6px; overflow: hidden; }
  .logo-text { display: none; }
  .logo-link { min-width: 32px; min-height: 32px; display: flex; align-items: center; justify-content: center; }
  .header-search-form { max-width: none; }
  .header-search-bar { padding: 8px 8px 8px 10px; border-radius: 22px; gap: 6px; }
  .header-search-input { font-size: 16px; padding: 0; }
  .clear-btn { width: 34px; height: 34px; min-width: 34px; }
  .search-submit-btn { width: 34px; height: 34px; min-width: 34px; }
  .tab-bar { top: 42px; }
  .tab-btn { padding: 10px 12px; font-size: 0.76rem; gap: 4px; min-height: 42px; }
  .tab-btn span { font-size: 0.72rem; }
  .content-area { padding: 10px 8px 50px; }
  .results-layout { gap: 12px; }
  .results-list-section { padding: 0; overflow: hidden; }
  .results-list { overflow: hidden; }
  .result-card { padding: 8px 8px; border-radius: 8px; min-width: 0; overflow: hidden; }
  .result-header { gap: 5px; margin-bottom: 1px; min-width: 0; justify-content: flex-start; width: 100%; padding-right: 32px; overflow: visible; }
  .result-favicon { width: 16px; height: 16px; }
  .result-domain { font-size: 0.72rem; flex-shrink: 1; min-width: 0; max-width: none; }
  .result-path { display: none; }
  .result-title { font-size: 0.86rem; font-weight: 600; line-height: 1.3; margin-bottom: 2px; display: block; word-break: break-word; white-space: normal; }
  .result-desc { font-size: 0.78rem; color: var(--text-secondary); line-height: 1.45; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; word-break: break-word; }
  .result-menu-btn { width: 34px; height: 34px; min-width: 34px; }
  .result-menu-btn svg { width: 12px; height: 12px; }
  .result-dropdown { min-width: 170px; right: -4px; border-radius: 10px; }
  .result-dropdown-item { padding: 8px 10px; font-size: 0.78rem; min-height: 40px; }
  .results-meta { font-size: 0.7rem; margin-bottom: 8px; }
  .inline-related-section { padding: 10px; margin: 10px 0; overflow: hidden; min-width: 0; }
  .inline-related-title { font-size: 0.78rem; margin-bottom: 6px; }
  .inline-related-grid { gap: 6px; width: 100%; min-width: 0; }
  .inline-related-chip { padding: 6px 10px; font-size: 0.72rem; min-height: 36px; border-radius: 18px; }
  .inline-related-chip svg { display: none; }
  .inline-images-section { margin: 10px 0; overflow: hidden; min-width: 0; }
  .inline-images-scroll { gap: 6px; width: 100%; min-width: 0; }
  .inline-image-card { width: 90px; min-height: 110px; }
  .inline-image-title { font-size: 0.62rem; padding: 3px 5px; }
  .images-grid { grid-template-columns: repeat(auto-fill, minmax(85px, 1fr)); gap: 5px; }
  .image-meta { padding: 4px 5px; }
  .image-title { font-size: 0.68rem; }
  .image-domain { font-size: 0.62rem; }
  .news-list { max-width: 100%; gap: 6px; }
  .news-card { padding: 10px 8px; border-radius: 8px; }
  .news-header { margin-bottom: 3px; gap: 5px; }
  .news-source { font-size: 0.72rem; }
  .news-time { font-size: 0.66rem; }
  .news-title { font-size: 0.86rem; line-height: 1.3; }
  .news-desc { font-size: 0.78rem; -webkit-line-clamp: 2; line-height: 1.45; }
  .weather-widget { padding: 12px; margin-bottom: 12px; }
  .weather-main { gap: 10px; margin-bottom: 10px; }
  .weather-icon-large { font-size: 2rem; }
  .weather-temp { font-size: 1.6rem; }
  .weather-details { gap: 10px; padding-top: 10px; }
  .weather-detail { min-width: 70px; }
  .weather-detail-label { font-size: 0.68rem; }
  .weather-detail-value { font-size: 0.88rem; }
  .calculator-widget { margin-bottom: 12px; border-radius: 10px; }
  .calc-header { padding: 10px 12px; font-size: 0.8rem; }
  .calc-body { padding: 12px; }
  .calc-form { gap: 6px; margin-bottom: 8px; }
  .calc-input { padding: 10px 10px; font-size: 16px; border-radius: 8px; }
  .calc-btn { padding: 10px 14px; font-size: 0.82rem; border-radius: 8px; }
  .calc-result { padding: 8px 10px; flex-direction: column; align-items: flex-start; gap: 2px; border-radius: 6px; }
  .calc-result-value { font-size: 1.1rem; }
  .calc-history-item { padding: 6px 8px; min-height: 36px; border-radius: 6px; }
  .wiki-card { padding: 12px; }
  .wiki-header { gap: 5px; margin-bottom: 8px; }
  .wiki-label { font-size: 0.74rem; }
  .wiki-extract { font-size: 0.8rem; line-height: 1.5; max-height: 140px; margin-bottom: 8px; }
  .wiki-read-more { font-size: 0.78rem; }
  .pagination { gap: 3px; margin-top: 16px; padding: 6px 0; }
  .page-btn { min-width: 40px; height: 40px; font-size: 0.76rem; border-radius: 8px; }
  .page-nav { min-width: 40px; padding: 0 6px; }
  .page-ellipsis { width: 40px; height: 40px; font-size: 0.76rem; }
  .preview-overlay { padding: 8px; padding-top: 44px; }
  .preview-panel { max-height: calc(100vh - 52px); border-radius: 10px; }
  .preview-close { width: 36px; height: 36px; top: -48px; right: 0; }
  .preview-close svg { width: 16px; height: 16px; }
  .preview-image-wrapper { max-height: 45vh; }
  .preview-image { max-height: 45vh; }
  .preview-info { padding: 10px 12px; }
  .preview-title { font-size: 0.86rem; }
  .preview-domain { font-size: 0.72rem; margin-bottom: 6px; }
  .preview-source-btn { padding: 8px 12px; font-size: 0.76rem; min-height: 40px; border-radius: 8px; }
  .maps-layout { min-height: 300px; }
  .map-container { height: 250px; min-height: 220px; }
  .map-sidebar { max-height: 240px; }
  .map-result-item { padding: 8px 8px; min-height: 44px; gap: 8px; }
  .map-result-icon { padding-top: 3px; }
  .map-result-icon svg { width: 12px; height: 12px; }
  .map-result-name { font-size: 0.78rem; }
  .map-result-address { font-size: 0.68rem; }
  .map-open-btn { width: 32px; height: 32px; min-width: 32px; border-radius: 6px; }
  .map-open-btn svg { width: 11px; height: 11px; }
  .empty-state { padding: 30px 12px; gap: 8px; }
  .empty-state svg { width: 32px; height: 32px; }
  .empty-query { font-size: 0.88rem; padding: 0 8px; text-align: center; }
  .empty-hint { font-size: 0.78rem; }
  .try-again-btn { padding: 10px 20px; min-height: 44px; font-size: 0.84rem; border-radius: 22px; }
  .loading-state { padding: 36px 12px; }
  .welcome-state { min-height: 35vh; }
  .welcome-logo { width: 40px; height: 40px; }
  .welcome-text { font-size: 0.88rem; }
  .widget-loading { padding: 20px 12px; font-size: 0.8rem; }
  .ai-answer-section { border-radius: 8px; margin-bottom: 10px; }
}

@media (max-width: 359px) {
  .tab-btn span { display: none; }
  .tab-btn { padding: 10px; min-width: 44px; justify-content: center; }
  .inline-related-chip svg { display: none; }
  .weather-icon-large { font-size: 1.8rem; }
  .weather-temp { font-size: 1.4rem; }
  .weather-details { flex-direction: column; gap: 8px; }
  .weather-detail { flex-direction: row; justify-content: space-between; align-items: center; gap: 8px; }
  .content-area { padding: 8px 6px 40px; }
  .result-card { padding: 8px 6px; }
  .inline-image-card { width: 80px; min-height: 100px; }
  .images-grid { grid-template-columns: repeat(auto-fill, minmax(75px, 1fr)); gap: 4px; }
  .preview-overlay { padding: 4px; padding-top: 40px; }
  .preview-close { width: 34px; height: 34px; top: -44px; }
}

@media (hover: none) and (pointer: coarse) {
  .result-card { -webkit-tap-highlight-color: transparent; }
  .result-card:active { background: var(--result-hover); }
  .news-card { -webkit-tap-highlight-color: transparent; }
  .news-card:active { background: var(--result-hover); border-color: var(--border-light); }
  .image-card { -webkit-tap-highlight-color: transparent; }
  .image-card:active { border-color: var(--border-light); opacity: 0.85; }
  .inline-image-card { -webkit-tap-highlight-color: transparent; }
  .inline-image-card:active { border-color: var(--border-light); opacity: 0.85; }
  .inline-related-chip { -webkit-tap-highlight-color: transparent; }
  .inline-related-chip:active { background: var(--bg-hover); border-color: var(--accent); color: var(--accent); }
  .tab-btn { -webkit-tap-highlight-color: transparent; }
  .page-btn { -webkit-tap-highlight-color: transparent; }
  .page-btn:active:not(:disabled) { background: var(--bg-hover); border-color: var(--border-light); transform: scale(0.95); }
  .map-result-item { -webkit-tap-highlight-color: transparent; }
  .map-result-item:active { background: var(--bg-hover); }
  .try-again-btn { -webkit-tap-highlight-color: transparent; }
  .try-again-btn:active { background: var(--bg-hover); transform: scale(0.97); }
  .preview-source-btn { -webkit-tap-highlight-color: transparent; }
  .preview-source-btn:active { background: var(--border-color); color: var(--accent); }
  .map-open-btn { -webkit-tap-highlight-color: transparent; }
  .map-open-btn:active { background: var(--border-color); color: var(--accent); }
  .calc-btn { -webkit-tap-highlight-color: transparent; }
  .calc-btn:active { opacity: 0.75; }
  .wiki-read-more { -webkit-tap-highlight-color: transparent; }
  .ai-answer-header { -webkit-tap-highlight-color: transparent; }
  .ai-answer-header:active { background: var(--bg-hover); }
}

.result-text-wrapper { margin-top: 6px; }
.result-text { font-size: 0.84rem; line-height: 1.5; color: var(--text-secondary); margin: 0; white-space: pre-wrap; word-break: break-word; }
.result-text.expanded { color: var(--text-primary); }
.text-toggle-btn { background: none; border: none; color: var(--accent); font-size: 0.78rem; cursor: pointer; padding: 4px 0; margin-top: 2px; font-weight: 500; }
.text-toggle-btn:hover { text-decoration: underline; }

</style>