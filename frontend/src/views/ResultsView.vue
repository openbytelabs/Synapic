<template>
  <div class="app">
    <button class="menu-toggle-fixed" :class="{ open: menuOpen }" @click="menuOpen = !menuOpen">
      <i class="fa-solid fa-table-columns"></i>
    </button>
    <aside class="sidebar" :class="{ open: menuOpen }">
      <div class="logo">Synapic</div>
      <nav class="nav">
        <a href="/" class="nav-item" @click="menuOpen = false">
          <i class="fa-solid fa-house"></i>
          <span class="nav-text">{{ t.nav.home }}</span>
        </a>
        <a href="/api" class="nav-item" @click="menuOpen = false">
          <i class="fa-solid fa-link"></i>
          <span class="nav-text">{{ t.nav.api }}</span>
        </a>
        <a href="/terms" class="nav-item" @click="menuOpen = false">
          <i class="fa-solid fa-shield-halved"></i>
          <span class="nav-text">{{ t.nav.terms }}</span>
        </a>
        <a href="/settings" class="nav-item" @click="menuOpen = false">
          <i class="fa-solid fa-gear"></i>
          <span class="nav-text">{{ t.nav.settings }}</span>
        </a>
      </nav>
      <div class="sidebar-footer">
        <a href="/download" class="nav-item download-button" @click="menuOpen = false">
          <i class="fa-solid fa-download"></i>
          <span class="nav-text">{{ t.nav.download }}</span>
        </a>
      </div>
    </aside>

    <div class="overlay" :class="{ visible: menuOpen }" @click="closeOverlay"></div>

    <main class="main">
      <div class="blur-red"></div>
      <div class="blur-yellow"></div>

      <div class="search-header">
        <div class="search-box-top">
          <i class="fa-solid fa-magnifying-glass search-icon"></i>
          <input
            type="text"
            class="search-input-top"
            v-model="searchQuery"
            @keyup.enter="performSearch"
          />
        </div>

        <div class="search-type-selector">
          <button class="type-button" :class="{ active: searchType === 'web' }" @click="changeSearchType('web')">
            <i class="fa-solid fa-globe"></i>
            <span>{{ t.search.web }}</span>
          </button>
          <button class="type-button" :class="{ active: searchType === 'image' }" @click="changeSearchType('image')">
            <i class="fa-solid fa-image"></i>
            <span>{{ t.search.images }}</span>
          </button>
          <button class="type-button" :class="{ active: searchType === 'news' }" @click="changeSearchType('news')">
            <i class="fa-solid fa-newspaper"></i>
            <span>{{ t.search.news }}</span>
          </button>
          <button class="type-button" :class="{ active: searchType === 'maps' }" @click="changeSearchType('maps')">
            <i class="fa-solid fa-map-location-dot"></i>
            <span>{{ t.search.maps }}</span>
          </button>
        </div>
      </div>

      <div class="results-container">
        <div v-if="isLoading" class="loading-dots">
          <span></span><span></span><span></span>
        </div>
        <div v-if="!isLoading && searchType !== 'maps' && results.length === 0 && noResultsVisible" class="no-results-message">{{ t.search.noResults }}</div>

        <div v-if="aiAnswer && !isLoading && searchType === 'web'" class="ai-answer-card">
          <div class="ai-header">
            <div class="ai-logo">
              <i class="fa-solid fa-brain"></i>
              <span class="ai-title">Synapic AI</span>
            </div>
            <div class="ai-badge">{{ t.search.aiBadge }}</div>
          </div>
          <div class="ai-content" v-html="formatAIAnswer(aiAnswer)"></div>
          <div class="ai-footer">
            <i class="fa-solid fa-circle-info"></i>
            <span>{{ t.search.aiFooter }}</span>
          </div>
        </div>

        <div v-if="aiLoading && searchType === 'web'" class="ai-answer-card ai-loading">
          <div class="ai-header">
            <div class="ai-logo">
              <i class="fa-solid fa-brain"></i>
              <span class="ai-title">Synapic AI</span>
            </div>
          </div>
          <div class="ai-loading-content">
            <i class="fa-solid fa-spinner fa-spin"></i>
            <span>{{ t.search.aiLoading }}</span>
          </div>
        </div>

        <div v-if="searchType === 'maps' && searchQuery" class="maps-container">
          <iframe
            :src="mapsUrl"
            width="100%"
            height="600"
            style="border:0; border-radius: 20px;"
            allowfullscreen=""
            loading="lazy"
            referrerpolicy="no-referrer-when-downgrade"
          ></iframe>
        </div>

        <div v-if="searchType === 'image' && !isLoading" class="maintenance-container">
          <div class="maintenance-card">
            <i class="fa-solid fa-screwdriver-wrench maintenance-icon"></i>
            <h2 class="maintenance-title">{{ t.search.maintenance }}</h2>
            <p class="maintenance-desc">{{ t.search.maintenanceDesc }}</p>
          </div>
        </div>

        <div v-if="searchType === 'web' && !isLoading">
          <div class="result-item" v-for="(result, index) in paginatedResults" :key="index">
            <div v-if="isYoutubeEmbed(result.url)" class="youtube-embed-wrapper">
              <iframe
                :src="result.url"
                width="100%"
                height="315"
                frameborder="0"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                allowfullscreen
                style="border-radius: 12px;"
              ></iframe>
              <div class="result-title-below">
                <i class="fa-brands fa-youtube youtube-icon"></i>
                <span>{{ result.title || 'YouTube Video' }}</span>
              </div>
            </div>
            <template v-else>
              <div class="result-header">
                <div class="result-left">
                  <img :src="getFavicon(result.url)" class="site-favicon" @error="handleImageError" />
                  <div class="result-info">
                    <a :href="result.url" target="_blank" class="result-title-link">
                      <h3 class="result-title">{{ result.title }}</h3>
                    </a>
                    <a :href="result.url" class="result-url" target="_blank">{{ result.url }}</a>
                  </div>
                </div>
                <div class="result-menu">
                  <button class="menu-button" @click="toggleMenu(index)">
                    <i class="fa-solid fa-ellipsis-vertical"></i>
                  </button>
                  <div class="menu-dropdown" v-if="activeMenu === index">
                    <button class="menu-option" @click="copyUrl(result.url, index)">
                      <i class="fa-solid fa-copy"></i>
                      {{ t.search.copyUrl }}
                    </button>
                  </div>
                </div>
              </div>
              <p class="result-date">{{ result.date }} — {{ result.description }}</p>
            </template>
          </div>
        </div>

        <div v-if="searchType === 'news' && !isLoading">
          <div class="result-item" v-for="(result, index) in paginatedResults" :key="index">
            <div class="result-header">
              <div class="result-left">
                <img :src="getFavicon(result.url)" class="site-favicon" @error="handleImageError" />
                <div class="result-info">
                  <a :href="result.url" target="_blank" class="result-title-link">
                    <h3 class="result-title">{{ result.title }}</h3>
                  </a>
                  <a :href="result.url" class="result-url" target="_blank">{{ result.url }}</a>
                </div>
              </div>
              <div class="result-menu">
                <button class="menu-button" @click="toggleMenu(index)">
                  <i class="fa-solid fa-ellipsis-vertical"></i>
                </button>
                <div class="menu-dropdown" v-if="activeMenu === index">
                  <button class="menu-option" @click="copyUrl(result.url, index)">
                    <i class="fa-solid fa-copy"></i>
                    {{ t.search.copyUrl }}
                  </button>
                </div>
              </div>
            </div>
            <p class="result-date">{{ result.date }} — {{ result.description }}</p>
          </div>
        </div>

        <div class="pagination" v-if="totalPages > 1 && searchType !== 'maps'">
          <button class="pagination-button" :class="{ disabled: currentPage === 1 }" @click="changePage(currentPage - 1)" :disabled="currentPage === 1">
            <i class="fa-solid fa-chevron-left"></i>
          </button>
          <button
            v-for="page in displayedPages"
            :key="page"
            class="pagination-button page-number"
            :class="{ active: currentPage === page }"
            @click="changePage(page)"
          >{{ page }}</button>
          <button class="pagination-button" :class="{ disabled: currentPage === totalPages }" @click="changePage(currentPage + 1)" :disabled="currentPage === totalPages">
            <i class="fa-solid fa-chevron-right"></i>
          </button>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useLocale, detectLanguage } from '../composables/useLocale.js';

const route = useRoute();
const router = useRouter();
const { lang, t } = useLocale();
const menuOpen = ref(false);
const searchQuery = ref('');
const searchType = ref('web');
const results = ref([]);
const activeMenu = ref(null);
const isLoading = ref(false);
const aiLoading = ref(false);
const aiAnswer = ref('');
const currentPage = ref(1);
const noResultsVisible = ref(false);
let noResultsTimer = null;
const resultsPerPage = 10;


const API_BASE_URL = 'https://api.synapic.com.tr';

const mapsUrl = computed(() => {
  if (!searchQuery.value) return '';
  return `https://www.google.com/maps/embed/v1/place?key=AIzaSyBFw0Qbyq9zTFTd-tUY6dZWTgaQzuU17R8&q=${encodeURIComponent(searchQuery.value)}`;
});

const totalPages = computed(() => Math.ceil(results.value.length / resultsPerPage));

const paginatedResults = computed(() => {
  const start = (currentPage.value - 1) * resultsPerPage;
  return results.value.slice(start, start + resultsPerPage);
});

const displayedPages = computed(() => {
  const pages = [];
  const maxDisplayed = 7;
  if (totalPages.value <= maxDisplayed) {
    for (let i = 1; i <= totalPages.value; i++) pages.push(i);
  } else {
    if (currentPage.value <= 4) {
      for (let i = 1; i <= 5; i++) pages.push(i);
      pages.push('...'); pages.push(totalPages.value);
    } else if (currentPage.value >= totalPages.value - 3) {
      pages.push(1); pages.push('...');
      for (let i = totalPages.value - 4; i <= totalPages.value; i++) pages.push(i);
    } else {
      pages.push(1); pages.push('...');
      for (let i = currentPage.value - 1; i <= currentPage.value + 1; i++) pages.push(i);
      pages.push('...'); pages.push(totalPages.value);
    }
  }
  return pages;
});

const isYoutubeEmbed = (url) => url && url.includes('youtube.com/embed/');

const changeSearchType = (type) => {
  if (searchType.value === type) return;
  searchType.value = type;
  if (searchQuery.value.trim()) performSearch(type);
};

const changePage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page;
    window.scrollTo({ top: 0, behavior: 'smooth' });
  }
};

const formatAIAnswer = (text) => {
  let formatted = text
    .replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;')
    .replace(/^#{3}\s+(.+)$/gm, '<h3>$1</h3>')
    .replace(/^#{2}\s+(.+)$/gm, '<h2>$1</h2>')
    .replace(/^#{1}\s+(.+)$/gm, '<h1>$1</h1>')
    .replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>')
    .replace(/\*(.+?)\*/g, '<em>$1</em>')
    .replace(/`([^`]+)`/g, '<code>$1</code>')
    .replace(/^```[\w]*\n([\s\S]*?)```$/gm, '<pre><code>$1</code></pre>')
    .replace(/^---$/gm, '<hr>')
    .replace(/^\*\s+(.+)$/gm, '<li>$1</li>')
    .replace(/^-\s+(.+)$/gm, '<li>$1</li>')
    .replace(/^\d+\.\s+(.+)$/gm, '<li>$1</li>')
    .replace(/(<li>.*<\/li>)/gs, '<ul>$1</ul>')
    .replace(/\[([^\]]+)\]\(([^)]+)\)/g, '<a href="$2" target="_blank" rel="noopener">$1</a>')
    .replace(/\n\n/g, '</p><p>').replace(/\n/g, '<br>');
  return `<p>${formatted}</p>`;
};

const fetchAIAnswer = async (query) => {
  aiLoading.value = true;
  aiAnswer.value = '';
  try {
    const apiKey = localStorage.getItem('synapic_apikey') || '';
    const url = apiKey
      ? `${API_BASE_URL}/ai?q=${encodeURIComponent(query)}&apikey=${encodeURIComponent(apiKey)}`
      : `${API_BASE_URL}/ai?q=${encodeURIComponent(query)}`;
    const response = await fetch(url, { cache: 'no-store' });
    if (!response.ok) throw new Error(`HTTP ${response.status}`);
    const data = await response.json();
    if (data.success && data.answer) aiAnswer.value = data.answer;
  } catch {
    aiAnswer.value = '';
  } finally {
    aiLoading.value = false;
  }
};

const buildSearchUrl = (query, type) => {
  if (type === 'web') return `${API_BASE_URL}/api?q=${encodeURIComponent(query)}`;
  if (type === 'news') return `${API_BASE_URL}/news?q=${encodeURIComponent(query)}`;
  if (type === 'image') return `${API_BASE_URL}/images?q=${encodeURIComponent(query)}`;
  return `${API_BASE_URL}/api?q=${encodeURIComponent(query)}`;
};

const performSearch = async (forceType) => {
  const query = searchQuery.value.trim();
  if (!query) return;
  const currentType = typeof forceType === 'string' ? forceType : searchType.value;
  if (currentType === 'maps') return;

  if (noResultsTimer) { clearTimeout(noResultsTimer); noResultsTimer = null; }
  noResultsVisible.value = false;
  isLoading.value = true;
  aiAnswer.value = '';
  aiLoading.value = false;
  currentPage.value = 1;
  activeMenu.value = null;
  brokenImages.value = new Set();

  router.replace({ path: '/search', query: { q: query, type: currentType } });

  let fetchedResults = [];
  let shouldFetchAI = false;

  try {
    const url = buildSearchUrl(query, currentType);
    const response = await fetch(url, { cache: 'no-store' });
    if (!response.ok) throw new Error(`HTTP ${response.status}`);
    const data = await response.json();
    if (data.results && Array.isArray(data.results)) {
      fetchedResults = data.results.map(result => ({
        ...result,
        date: new Date().toLocaleDateString(lang.value === 'tr' ? 'tr-TR' : 'en-US', {
          year: 'numeric', month: 'short', day: 'numeric'
        })
      }));
      shouldFetchAI = currentType === 'web' && fetchedResults.length > 0;
    }
  } catch {
    fetchedResults = [];
  }

  results.value = fetchedResults;
  isLoading.value = false;

  if (fetchedResults.length === 0) {
    noResultsTimer = setTimeout(() => { noResultsVisible.value = true; noResultsTimer = null; }, 2000);
  }
  if (shouldFetchAI) fetchAIAnswer(query);
};

const getFavicon = (url) => {
  try {
    const domain = new URL(url).hostname;
    return `https://www.google.com/s2/favicons?domain=${domain}&sz=32`;
  } catch {
    return 'https://www.google.com/s2/favicons?domain=example.com&sz=32';
  }
};

const handleImageError = (event) => {
  event.target.src = 'https://www.google.com/s2/favicons?domain=example.com&sz=32';
};

const brokenImages = ref(new Set());
const IMAGE_EXTENSIONS = /\.(jpg|jpeg|png|gif|webp|svg|bmp|avif|tiff?)(\?.*)?$/i;

const getImageSrc = (result) => {
  if (result.description && IMAGE_EXTENSIONS.test(result.description.trim())) return result.description.trim();
  if (IMAGE_EXTENSIONS.test(result.url)) return result.url;
  try {
    const u = new URL(result.url);
    if (IMAGE_EXTENSIONS.test(u.pathname)) return result.url;
  } catch {}
  return result.url;
};

const getImageDomain = (url) => {
  try { return new URL(url).hostname.replace('www.', ''); } catch { return ''; }
};

const handleImageLoadError = (event, index) => {
  brokenImages.value = new Set([...brokenImages.value, index]);
  event.target.closest('.image-item').style.display = 'none';
};

const toggleMenu = (index) => { activeMenu.value = activeMenu.value === index ? null : index; };
const copyUrl = (url) => { navigator.clipboard.writeText(url); activeMenu.value = null; };
const closeOverlay = () => { menuOpen.value = false; };

onMounted(async () => {
  await detectLanguage();
  const queryParam = route.query.q;
  const typeParam = route.query.type || 'web';
  if (queryParam) {
    searchQuery.value = queryParam;
    searchType.value = typeParam;
    performSearch(typeParam);
  }
  document.addEventListener('click', (e) => {
    if (!e.target.closest('.result-menu')) activeMenu.value = null;
  });
});

watch(() => route.query.q, (newQuery) => {
  if (newQuery) {
    const typeParam = route.query.type || 'web';
    searchQuery.value = newQuery;
    searchType.value = typeParam;
    performSearch(typeParam);
  }
});
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap');

* { margin: 0; padding: 0; box-sizing: border-box; }

.app {
  min-height: 100vh;
  background: linear-gradient(135deg, #1a1a1d 0%, #2d1f2f 50%, #1a1a1d 100%);
  font-family: 'Inter', sans-serif; display: flex; position: relative; overflow-x: hidden;
}
.menu-toggle-fixed {
  position: fixed; top: 28px; left: 28px; background: rgba(42,42,42,0.8);
  backdrop-filter: blur(10px); border: none; color: rgba(255,255,255,0.85);
  font-size: 18px; cursor: pointer; z-index: 1002; transition: all 0.4s ease;
  border-radius: 14px; display: flex; align-items: center; justify-content: center;
  width: 46px; height: 46px; box-shadow: 0 4px 12px rgba(0,0,0,0.3);
}
.menu-toggle-fixed.open { left: 248px; }
.menu-toggle-fixed:hover { background: rgba(52,52,52,0.9); color: #d4af37; }
.sidebar {
  width: 0; height: 100vh; background: rgba(42,42,42,0.7); backdrop-filter: blur(20px);
  display: flex; flex-direction: column; padding: 0; position: fixed;
  left: -280px; top: 0; transition: all 0.4s ease; overflow: hidden;
  opacity: 0; border-radius: 0 24px 24px 0; z-index: 1000;
}
.sidebar.open { left: 0; width: 264px; padding: 28px 0; opacity: 1; box-shadow: 4px 0 20px rgba(0,0,0,0.5); overflow: visible; }
.overlay {
  position: fixed; top: 0; left: 0; width: 100%; height: 100%;
  background: rgba(0,0,0,0.5); backdrop-filter: blur(2px); z-index: 999;
  opacity: 0; pointer-events: none; transition: opacity 0.3s ease;
}
.overlay.visible { opacity: 1; pointer-events: all; }
.sidebar-footer { margin-top: auto; padding: 16px 16px 0; opacity: 0; transition: opacity 0.3s ease 0.2s; }
.sidebar.open .sidebar-footer { opacity: 1; }
.download-button { background: rgba(212,175,55,0.1); border: 1px solid rgba(212,175,55,0.3); color: #d4af37; opacity: 1 !important; }
.download-button:hover { background: rgba(212,175,55,0.2); border-color: rgba(212,175,55,0.5); color: #e8c84a; }
.download-button i { color: #d4af37; }
.logo {
  font-size: 28px; font-weight: 700;
  background: linear-gradient(135deg, #d4af37 0%, #f4e5a1 100%);
  -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text;
  margin-bottom: 40px; padding: 0 28px; text-align: center; letter-spacing: -0.5px;
  white-space: nowrap; opacity: 0; transition: opacity 0.3s ease 0.2s;
}
.sidebar.open .logo { opacity: 1; }
.nav { display: flex; flex-direction: column; gap: 8px; }
.nav-item {
  display: flex; align-items: center; gap: 16px; padding: 14px 18px;
  color: rgba(255,255,255,0.7); text-decoration: none; border-radius: 12px;
  font-size: 15px; font-weight: 500; transition: all 0.3s ease; cursor: pointer;
  white-space: nowrap; opacity: 0;
}
.sidebar.open .nav-item { opacity: 1; transition: all 0.3s ease 0.2s; }
.nav-item:hover { background: rgba(212,175,55,0.1); color: #d4af37; transform: translateX(4px); }
.nav-item i { font-size: 18px; width: 20px; text-align: center; }
.main { flex: 1; display: flex; flex-direction: column; position: relative; width: 100%; }
.blur-red, .blur-yellow { position: fixed; width: 300px; height: 300px; filter: blur(30px); pointer-events: none; z-index: 0; }
.blur-red { background: radial-gradient(circle, rgba(139,0,0,0.08) 0%, transparent 70%); top: -200px; right: -200px; }
.blur-yellow { background: radial-gradient(circle, rgba(212,175,55,0.06) 0%, transparent 70%); bottom: -200px; left: -200px; }
.search-header {
  position: sticky; top: 0; z-index: 100; background: transparent;
  padding: 32px 48px 20px; display: flex; flex-direction: column; gap: 16px;
}
.search-box-top {
  max-width: 800px; width: 100%; margin: 0 auto; position: relative;
  display: flex; align-items: center; background: rgba(255,255,255,0.03);
  border: 2px solid rgba(212,175,55,0.4); border-radius: 50px; padding: 0 24px; transition: all 0.3s ease;
}
.search-box-top:focus-within { border-color: rgba(212,175,55,0.6); background: rgba(255,255,255,0.05); }
.search-icon { color: rgba(255,255,255,0.4); font-size: 18px; margin-right: 14px; }
.search-input-top {
  flex: 1; background: transparent; border: none; outline: none; padding: 18px 0;
  color: rgba(255,255,255,0.9); font-size: 16px; font-family: 'Inter', sans-serif;
}
.search-input-top::placeholder { color: rgba(255,255,255,0.3); }
.search-type-selector {
  max-width: 800px; width: 100%; margin: 0 auto; display: flex; gap: 8px;
  background: rgba(255,255,255,0.02); border: 1px solid rgba(255,255,255,0.1);
  border-radius: 16px; padding: 6px;
}
.type-button {
  flex: 1; display: flex; align-items: center; justify-content: center; gap: 8px;
  padding: 10px 16px; background: transparent; border: none; border-radius: 8px;
  color: rgba(255,255,255,0.6); font-size: 14px; font-weight: 500;
  font-family: 'Inter', sans-serif; cursor: pointer; transition: all 0.3s ease;
}
.type-button:hover { background: rgba(255,255,255,0.05); color: rgba(255,255,255,0.8); }
.type-button.active { background: rgba(212,175,55,0.15); color: #d4af37; border: 1px solid rgba(212,175,55,0.3); }
.type-button i { font-size: 16px; }
.results-container { flex: 1; padding: 24px 48px 48px; max-width: 900px; width: 100%; margin: 0 auto; position: relative; z-index: 1; display: flex; flex-direction: column; }
.ai-answer-card { background: rgba(48,52,58,0.4); backdrop-filter: blur(10px); border: 1px solid rgba(212,175,55,0.12); border-radius: 24px; padding: 24px; margin-bottom: 24px; transition: all 0.3s ease; }
.ai-answer-card:hover { background: rgba(48,52,58,0.5); border-color: rgba(212,175,55,0.2); }
.ai-header { display: flex; align-items: center; justify-content: space-between; margin-bottom: 16px; }
.ai-logo { display: flex; align-items: center; gap: 12px; }
.ai-logo i { font-size: 28px; background: linear-gradient(135deg, #d4af37 0%, #f4e5a1 100%); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; }
.ai-title { font-size: 20px; font-weight: 600; background: linear-gradient(135deg, #d4af37 0%, #f4e5a1 100%); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; }
.ai-badge { padding: 6px 14px; background: rgba(212,175,55,0.15); border: 1px solid rgba(212,175,55,0.3); border-radius: 20px; font-size: 12px; font-weight: 500; color: #d4af37; }
.ai-content { color: rgba(255,255,255,0.85); font-size: 15px; line-height: 1.7; margin-bottom: 16px; }
.ai-content h1, .ai-content h2, .ai-content h3 { color: rgba(255,255,255,0.95); margin: 14px 0 6px; line-height: 1.3; font-weight: 600; }
.ai-content h1 { font-size: 20px; } .ai-content h2 { font-size: 17px; } .ai-content h3 { font-size: 15px; }
.ai-content ul { padding-left: 20px; margin: 8px 0; }
.ai-content li { color: rgba(255,255,255,0.8); margin: 4px 0; line-height: 1.6; }
.ai-content code { background: rgba(212,175,55,0.12); border: 1px solid rgba(212,175,55,0.2); border-radius: 5px; padding: 1px 6px; font-family: 'Courier New', monospace; font-size: 13px; color: #d4af37; }
.ai-content pre { background: rgba(0,0,0,0.35); border: 1px solid rgba(255,255,255,0.08); border-radius: 10px; padding: 14px 16px; overflow-x: auto; margin: 10px 0; }
.ai-content pre code { background: none; border: none; padding: 0; color: rgba(255,255,255,0.85); font-size: 13px; }
.ai-content a { color: #d4af37; text-decoration: underline; text-underline-offset: 3px; }
.ai-content a:hover { color: #e8c84a; }
.ai-content hr { border: none; border-top: 1px solid rgba(255,255,255,0.1); margin: 14px 0; }
.ai-content strong { color: rgba(255,255,255,0.95); font-weight: 600; }
.ai-content em { color: rgba(255,255,255,0.75); font-style: italic; }
.ai-footer { display: flex; align-items: center; gap: 8px; color: rgba(255,255,255,0.4); font-size: 13px; }
.ai-footer i { font-size: 14px; }
.ai-loading { background: rgba(48,52,58,0.5); }
.ai-loading-content { display: flex; align-items: center; gap: 12px; color: rgba(255,255,255,0.6); font-size: 15px; padding: 20px 0; }
.ai-loading-content i { font-size: 20px; }
.maps-container { width: 100%; margin-top: 20px; }
.result-item { background: rgba(48,52,58,0.4); backdrop-filter: blur(10px); border: 1px solid rgba(255,255,255,0.06); border-radius: 24px; padding: 20px 24px; margin-bottom: 16px; transition: all 0.3s ease; }
.result-item:hover { background: rgba(48,52,58,0.5); border-color: rgba(212,175,55,0.15); }
.youtube-embed-wrapper { display: flex; flex-direction: column; gap: 12px; }
.youtube-embed-wrapper iframe { width: 100%; height: 315px; border-radius: 12px; }
.result-title-below { display: flex; align-items: center; gap: 10px; color: rgba(255,255,255,0.8); font-size: 15px; font-weight: 500; }
.youtube-icon { color: #ff0000; font-size: 20px; }
.result-header { display: flex; justify-content: space-between; align-items: flex-start; margin-bottom: 8px; }
.result-left { display: flex; align-items: flex-start; gap: 12px; flex: 1; }
.site-favicon { width: 24px; height: 24px; border-radius: 4px; flex-shrink: 0; margin-top: 2px; }
.result-info { flex: 1; min-width: 0; }
.result-title-link { text-decoration: none; }
.result-title { font-size: 16px; font-weight: 500; color: #7ba4ff; margin-bottom: 4px; line-height: 1.4; }
.result-title:hover { color: #a3c4ff; text-decoration: underline; }
.result-url { font-size: 13px; color: rgba(255,255,255,0.4); text-decoration: none; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; display: block; }
.result-url:hover { color: rgba(255,255,255,0.6); }
.result-menu { position: relative; flex-shrink: 0; }
.menu-button { background: transparent; border: none; color: rgba(255,255,255,0.4); cursor: pointer; padding: 4px 8px; border-radius: 6px; transition: all 0.2s ease; }
.menu-button:hover { background: rgba(255,255,255,0.08); color: rgba(255,255,255,0.7); }
.menu-dropdown { position: absolute; right: 0; top: 100%; margin-top: 4px; background: rgba(30,30,30,0.98); backdrop-filter: blur(20px); border: 1px solid rgba(255,255,255,0.15); border-radius: 10px; padding: 6px; z-index: 50; min-width: 150px; box-shadow: 0 4px 16px rgba(0,0,0,0.4); }
.menu-option { width: 100%; display: flex; align-items: center; gap: 10px; padding: 10px 12px; background: transparent; border: none; color: rgba(255,255,255,0.8); font-size: 14px; font-family: 'Inter', sans-serif; cursor: pointer; border-radius: 6px; transition: all 0.2s ease; text-align: left; }
.menu-option:hover { background: rgba(255,255,255,0.08); }
.result-date { font-size: 13px; color: rgba(255,255,255,0.5); margin: 0; line-height: 1.5; word-wrap: break-word; overflow-wrap: break-word; }
.maintenance-container { display: flex; justify-content: center; align-items: center; padding: 60px 20px; }
.maintenance-card { display: flex; flex-direction: column; align-items: center; gap: 16px; text-align: center; }
.maintenance-icon { font-size: 48px; color: rgba(212,175,55,0.6); }
.maintenance-title { font-size: 24px; font-weight: 600; color: rgba(255,255,255,0.9); }
.maintenance-desc { font-size: 14px; color: rgba(255,255,255,0.5); line-height: 1.6; }
.image-source { font-size: 11px; color: rgba(255,255,255,0.35); margin-top: 2px; display: block; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
.image-broken { display: none !important; }
.loading-dots { display: flex; align-items: center; justify-content: center; gap: 10px; padding: 60px 20px; }
.loading-dots span { width: 12px; height: 12px; border-radius: 50%; background: #d4af37; display: inline-block; animation: bounce 1.2s ease-in-out infinite; }
.loading-dots span:nth-child(1) { animation-delay: 0s; }
.loading-dots span:nth-child(2) { animation-delay: 0.2s; }
.loading-dots span:nth-child(3) { animation-delay: 0.4s; }
@keyframes bounce { 0%, 60%, 100% { transform: translateY(0); opacity: 0.4; } 30% { transform: translateY(-16px); opacity: 1; } }
.no-results-message { color: rgba(255,255,255,0.6); font-size: 16px; text-align: center; padding: 40px 20px; display: flex; align-items: center; gap: 12px; justify-content: center; }
.pagination { display: flex; align-items: center; gap: 8px; margin-top: 32px; padding-bottom: 20px; }
.pagination-button { min-width: 40px; height: 40px; background: rgba(48,52,58,0.6); backdrop-filter: blur(10px); border: 1px solid rgba(255,255,255,0.08); border-radius: 10px; color: rgba(255,255,255,0.7); font-size: 14px; font-weight: 500; cursor: pointer; transition: all 0.3s ease; display: flex; align-items: center; justify-content: center; padding: 0 12px; font-family: 'Inter', sans-serif; }
.pagination-button:hover:not(.disabled) { background: rgba(212,175,55,0.15); border-color: rgba(212,175,55,0.3); color: rgba(255,255,255,0.9); }
.pagination-button.active { background: rgba(212,175,55,0.2); border-color: rgba(212,175,55,0.4); color: #d4af37; }
.pagination-button.disabled { opacity: 0.3; cursor: not-allowed; }
.pagination-button.page-number { min-width: 40px; }

@media (max-width: 768px) {
  .menu-toggle-fixed { top: 20px; left: 20px; }
  .menu-toggle-fixed.open { left: 20px; }
  .sidebar.open { width: 80%; max-width: 300px; border-radius: 0; }
  .search-header { padding: 30px 24px 16px; }
  .search-type-selector { flex-wrap: wrap; }
  .type-button { flex: 1 1 45%; min-width: 120px; }
  .results-container { padding: 16px 24px 40px; }
  .ai-answer-card { padding: 20px; }
  .ai-logo i { font-size: 24px; }
  .ai-title { font-size: 18px; }
  .result-item { padding: 16px 18px; }
  .result-title { font-size: 15px; }
  .result-url { font-size: 12px; }
  .result-date { font-size: 12px; }
  .youtube-embed-wrapper iframe { height: 220px; }
  .pagination { gap: 6px; margin-top: 24px; }
  .pagination-button { min-width: 36px; height: 36px; font-size: 13px; }
}
@media (max-width: 480px) {
  .menu-toggle-fixed { top: 16px; left: 16px; }
  .menu-toggle-fixed.open { left: 16px; }
  .sidebar.open { width: 85%; }
  .search-header { padding: 24px 16px 12px; }
  .type-button { flex: 1 1 100%; font-size: 13px; padding: 8px 12px; }
  .results-container { padding: 12px 16px 32px; }
  .ai-answer-card { padding: 16px; border-radius: 16px; }
  .ai-content { font-size: 14px; }
  .result-item { padding: 14px 16px; border-radius: 16px; }
  .site-favicon { width: 20px; height: 20px; }
  .result-title { font-size: 14px; }
  .result-url { font-size: 11px; }
  .result-date { font-size: 11px; }
  .youtube-embed-wrapper iframe { height: 180px; }
  .pagination { gap: 4px; margin-top: 20px; }
  .pagination-button { min-width: 32px; height: 32px; font-size: 12px; padding: 0 8px; }
}
</style>