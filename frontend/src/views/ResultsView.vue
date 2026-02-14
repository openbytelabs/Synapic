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
          <span class="nav-text">Home</span>
        </a>
        <a href="/api" class="nav-item" @click="menuOpen = false">
          <i class="fa-solid fa-link"></i>
          <span class="nav-text">API</span>
        </a>
        <a href="/terms" class="nav-item" @click="menuOpen = false">
          <i class="fa-solid fa-shield-halved"></i>
          <span class="nav-text">Privacy & Terms</span>
        </a>
        <a href="/settings" class="nav-item" @click="menuOpen = false">
          <i class="fa-solid fa-gear"></i>
          <span class="nav-text">Settings</span>
        </a>
      </nav>
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
          <button
            class="type-button"
            :class="{ active: searchType === 'web' }"
            @click="changeSearchType('web')"
          >
            <i class="fa-solid fa-globe"></i>
            <span>Web</span>
          </button>
          <button
            class="type-button"
            :class="{ active: searchType === 'image' }"
            @click="changeSearchType('image')"
          >
            <i class="fa-solid fa-image"></i>
            <span>Images</span>
          </button>
          <button
            class="type-button"
            :class="{ active: searchType === 'news' }"
            @click="changeSearchType('news')"
          >
            <i class="fa-solid fa-newspaper"></i>
            <span>News</span>
          </button>
          <button
            class="type-button"
            :class="{ active: searchType === 'maps' }"
            @click="changeSearchType('maps')"
          >
            <i class="fa-solid fa-map-location-dot"></i>
            <span>Maps</span>
          </button>
        </div>
      </div>

      <div class="results-container">
        <div v-if="isLoading" class="loading-dots">
          <span></span>
          <span></span>
          <span></span>
        </div>
        <div v-if="!isLoading && searchType !== 'maps' && results.length === 0 && noResultsVisible" class="no-results-message">No results found. Try searching for something.</div>

        <div v-if="aiAnswer && !isLoading && searchType === 'web'" class="ai-answer-card">
          <div class="ai-header">
            <div class="ai-logo">
              <i class="fa-solid fa-brain"></i>
              <span class="ai-title">Synapic AI</span>
            </div>
            <div class="ai-badge">AI-Powered Answer</div>
          </div>
          <div class="ai-content" v-html="formatAIAnswer(aiAnswer)"></div>
          <div class="ai-footer">
            <i class="fa-solid fa-circle-info"></i>
            <span>Answer generated from search results</span>
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
            <span>Analyzing search results...</span>
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

        <div v-if="searchType === 'image' && !isLoading" class="image-grid">
          <div
            v-for="(result, index) in paginatedResults"
            :key="index"
            class="image-item"
          >
            <a :href="result.url" target="_blank">
              <img :src="result.url" :alt="result.title" @error="handleImageLoadError" />
            </a>
            <div class="image-info">
              <a :href="result.url" target="_blank" class="image-title">{{ result.title }}</a>
            </div>
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
                      Copy URL
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
                    Copy URL
                  </button>
                </div>
              </div>
            </div>
            <p class="result-date">{{ result.date }} — {{ result.description }}</p>
          </div>
        </div>

        <div class="pagination" v-if="totalPages > 1 && searchType !== 'maps'">
          <button
            class="pagination-button"
            :class="{ disabled: currentPage === 1 }"
            @click="changePage(currentPage - 1)"
            :disabled="currentPage === 1"
          >
            <i class="fa-solid fa-chevron-left"></i>
          </button>

          <button
            v-for="page in displayedPages"
            :key="page"
            class="pagination-button page-number"
            :class="{ active: currentPage === page }"
            @click="changePage(page)"
          >
            {{ page }}
          </button>

          <button
            class="pagination-button"
            :class="{ disabled: currentPage === totalPages }"
            @click="changePage(currentPage + 1)"
            :disabled="currentPage === totalPages"
          >
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

const route = useRoute();
const router = useRouter();
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

const mapsUrl = computed(() => {
  if (!searchQuery.value) return '';
  return `https://www.google.com/maps/embed/v1/place?key=AIzaSyBFw0Qbyq9zTFTd-tUY6dZWTgaQzuU17R8&q=${encodeURIComponent(searchQuery.value)}`;
});

const totalPages = computed(() => Math.ceil(results.value.length / resultsPerPage));

const paginatedResults = computed(() => {
  const start = (currentPage.value - 1) * resultsPerPage;
  const end = start + resultsPerPage;
  return results.value.slice(start, end);
});

const displayedPages = computed(() => {
  const pages = [];
  const maxDisplayed = 7;

  if (totalPages.value <= maxDisplayed) {
    for (let i = 1; i <= totalPages.value; i++) {
      pages.push(i);
    }
  } else {
    if (currentPage.value <= 4) {
      for (let i = 1; i <= 5; i++) {
        pages.push(i);
      }
      pages.push('...');
      pages.push(totalPages.value);
    } else if (currentPage.value >= totalPages.value - 3) {
      pages.push(1);
      pages.push('...');
      for (let i = totalPages.value - 4; i <= totalPages.value; i++) {
        pages.push(i);
      }
    } else {
      pages.push(1);
      pages.push('...');
      for (let i = currentPage.value - 1; i <= currentPage.value + 1; i++) {
        pages.push(i);
      }
      pages.push('...');
      pages.push(totalPages.value);
    }
  }

  return pages;
});

const isYoutubeEmbed = (url) => {
  return url && url.includes('youtube.com/embed/');
};

const changeSearchType = (type) => {
  if (searchType.value === type) return;
  searchType.value = type;
  if (searchQuery.value.trim()) {
    performSearch(type);
  }
};

const changePage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page;
    window.scrollTo({ top: 0, behavior: 'smooth' });
  }
};

const formatAIAnswer = (text) => {
  let formatted = text
    .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
    .replace(/\*(.*?)\*/g, '<em>$1</em>')
    .replace(/\n\n/g, '</p><p>')
    .replace(/\n/g, '<br>');

  return `<p>${formatted}</p>`;
};

const fetchAIAnswer = async (query) => {
  aiLoading.value = true;
  aiAnswer.value = '';

  try {
    const response = await fetch(`https://api.synapic.com.tr/ai?q=${encodeURIComponent(query)}`);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();

    if (data.answer) {
      aiAnswer.value = data.answer;
    } else {
      aiAnswer.value = 'Unable to generate AI summary at this time.';
    }
  } catch (error) {
    console.error('AI Error:', error);
    aiAnswer.value = 'Error generating AI summary.';
  } finally {
    aiLoading.value = false;
  }
};

const buildSearchUrl = (query, type) => {
  if (type === 'web') {
    return `https://api.synapic.com.tr/api?q=${encodeURIComponent(query)}`;
  } else if (type === 'news') {
    return `https://api.synapic.com.tr/news?q=${encodeURIComponent(query)}`;
  } else if (type === 'image') {
    return `https://api.synapic.com.tr/images?q=${encodeURIComponent(query)}`;
  }
  return `https://api.synapic.com.tr/api?q=${encodeURIComponent(query)}`;
};

const performSearch = async (forceType) => {
  if (!searchQuery.value.trim()) return;

  const currentType = forceType || searchType.value;

  if (currentType === 'maps') {
    return;
  }

  if (noResultsTimer) {
    clearTimeout(noResultsTimer);
    noResultsTimer = null;
  }
  noResultsVisible.value = false;

  isLoading.value = true;
  aiAnswer.value = '';
  aiLoading.value = false;
  currentPage.value = 1;
  activeMenu.value = null;

  let fetchedResults = [];
  let shouldFetchAI = false;

  try {
    const url = buildSearchUrl(searchQuery.value, currentType);
    const response = await fetch(url);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();

    if (data.results && Array.isArray(data.results)) {
      fetchedResults = data.results.map(result => ({
        ...result,
        date: new Date().toLocaleDateString('en-US', {
          year: 'numeric',
          month: 'short',
          day: 'numeric'
        })
      }));
      shouldFetchAI = currentType === 'web' && fetchedResults.length > 0;
    }
  } catch (error) {
    console.error('Search error:', error);
    fetchedResults = [];
  }

  results.value = fetchedResults;
  isLoading.value = false;

  if (fetchedResults.length === 0 && currentType !== 'maps') {
    noResultsTimer = setTimeout(() => {
      noResultsVisible.value = true;
      noResultsTimer = null;
    }, 60000);
  }

  if (shouldFetchAI) {
    fetchAIAnswer(searchQuery.value);
  }
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

const handleImageLoadError = (event) => {
  event.target.src = 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="200" height="200"%3E%3Crect fill="%23333" width="200" height="200"/%3E%3Ctext fill="%23999" x="50%25" y="50%25" text-anchor="middle" dy=".3em"%3EImage Error%3C/text%3E%3C/svg%3E';
  event.target.style.objectFit = 'contain';
  event.target.style.backgroundColor = 'rgba(48, 52, 58, 0.8)';
};

const toggleMenu = (index) => {
  activeMenu.value = activeMenu.value === index ? null : index;
};

const copyUrl = (url, index) => {
  navigator.clipboard.writeText(url);
  activeMenu.value = null;
};

const closeOverlay = () => {
  menuOpen.value = false;
};

onMounted(() => {
  const queryParam = route.query.q;
  const typeParam = route.query.type || 'web';

  if (queryParam) {
    searchQuery.value = queryParam;
    searchType.value = typeParam;
    performSearch(typeParam);
  }

  document.addEventListener('click', (e) => {
    if (!e.target.closest('.result-menu')) {
      activeMenu.value = null;
    }
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

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.app {
  min-height: 100vh;
  background: linear-gradient(135deg, #1a1a1d 0%, #2d1f2f 50%, #1a1a1d 100%);
  font-family: 'Inter', sans-serif;
  display: flex;
  position: relative;
  overflow-x: hidden;
}

.menu-toggle-fixed {
  position: fixed;
  top: 28px;
  left: 28px;
  background: rgba(42, 42, 42, 0.8);
  backdrop-filter: blur(10px);
  border: none;
  color: rgba(255, 255, 255, 0.85);
  font-size: 18px;
  cursor: pointer;
  z-index: 1002;
  transition: all 0.4s ease;
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 46px;
  height: 46px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.menu-toggle-fixed.open {
  left: 248px;
}

.menu-toggle-fixed:hover {
  background: rgba(52, 52, 52, 0.9);
  color: #d4af37;
}

.sidebar {
  width: 0;
  height: 100vh;
  background: rgba(42, 42, 42, 0.7);
  backdrop-filter: blur(20px);
  display: flex;
  flex-direction: column;
  padding: 0;
  position: fixed;
  left: -280px;
  top: 0;
  transition: all 0.4s ease;
  overflow: hidden;
  opacity: 0;
  border-radius: 0 24px 24px 0;
  z-index: 1000;
}

.sidebar.open {
  left: 0;
  width: 264px;
  padding: 28px 0;
  opacity: 1;
  box-shadow: 4px 0 20px rgba(0, 0, 0, 0.5);
  overflow: visible;
}

.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(2px);
  z-index: 999;
  opacity: 0;
  pointer-events: none;
  transition: opacity 0.3s ease;
}

.overlay.visible {
  opacity: 1;
  pointer-events: all;
}

.logo {
  font-size: 28px;
  font-weight: 700;
  background: linear-gradient(135deg, #d4af37 0%, #f4e5a1 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 40px;
  padding: 0 28px;
  text-align: center;
  letter-spacing: -0.5px;
  white-space: nowrap;
  opacity: 0;
  transition: opacity 0.3s ease 0.2s;
}

.sidebar.open .logo {
  opacity: 1;
}

.nav {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 14px 18px;
  color: rgba(255, 255, 255, 0.7);
  text-decoration: none;
  border-radius: 12px;
  font-size: 15px;
  font-weight: 500;
  transition: all 0.3s ease;
  cursor: pointer;
  white-space: nowrap;
  opacity: 0;
}

.sidebar.open .nav-item {
  opacity: 1;
  transition: all 0.3s ease 0.2s;
}

.nav-item:hover {
  background: rgba(212, 175, 55, 0.1);
  color: #d4af37;
  transform: translateX(4px);
}

.nav-item i {
  font-size: 18px;
  width: 20px;
  text-align: center;
}

.main {
  flex: 1;
  display: flex;
  flex-direction: column;
  position: relative;
  width: 100%;
}

.blur-red,
.blur-yellow {
  position: fixed;
  width: 300px;
  height: 300px;
  filter: blur(30px);
  pointer-events: none;
  z-index: 0;
}

.blur-red {
  background: radial-gradient(circle, rgba(139, 0, 0, 0.08) 0%, transparent 70%);
  top: -200px;
  right: -200px;
}

.blur-yellow {
  background: radial-gradient(circle, rgba(212, 175, 55, 0.06) 0%, transparent 70%);
  bottom: -200px;
  left: -200px;
}

.search-header {
  position: sticky;
  top: 0;
  z-index: 100;
  background: transparent;
  padding: 32px 48px 20px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.search-box-top {
  max-width: 800px;
  width: 100%;
  margin: 0 auto;
  position: relative;
  display: flex;
  align-items: center;
  background: rgba(255, 255, 255, 0.03);
  border: 2px solid rgba(212, 175, 55, 0.4);
  border-radius: 50px;
  padding: 0 24px;
  transition: all 0.3s ease;
}

.search-box-top:focus-within {
  border-color: rgba(212, 175, 55, 0.6);
  background: rgba(255, 255, 255, 0.05);
}

.search-icon {
  color: rgba(255, 255, 255, 0.4);
  font-size: 18px;
  margin-right: 14px;
}

.search-input-top {
  flex: 1;
  background: transparent;
  border: none;
  outline: none;
  padding: 18px 0;
  color: rgba(255, 255, 255, 0.9);
  font-size: 16px;
  font-family: 'Inter', sans-serif;
}

.search-input-top::placeholder {
  color: rgba(255, 255, 255, 0.3);
}

.search-type-selector {
  max-width: 800px;
  width: 100%;
  margin: 0 auto;
  display: flex;
  gap: 8px;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 6px;
}

.type-button {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 10px 16px;
  background: transparent;
  border: none;
  border-radius: 8px;
  color: rgba(255, 255, 255, 0.6);
  font-size: 14px;
  font-weight: 500;
  font-family: 'Inter', sans-serif;
  cursor: pointer;
  transition: all 0.3s ease;
}

.type-button:hover {
  background: rgba(255, 255, 255, 0.05);
  color: rgba(255, 255, 255, 0.8);
}

.type-button.active {
  background: rgba(212, 175, 55, 0.15);
  color: #d4af37;
  border: 1px solid rgba(212, 175, 55, 0.3);
}

.type-button i {
  font-size: 16px;
}

.results-container {
  flex: 1;
  padding: 24px 48px 48px;
  max-width: 900px;
  width: 100%;
  margin: 0 auto;
  position: relative;
  z-index: 1;
  display: flex;
  flex-direction: column;
}

.ai-answer-card {
  background: rgba(48, 52, 58, 0.4);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(212, 175, 55, 0.12);
  border-radius: 24px;
  padding: 24px;
  margin-bottom: 24px;
  transition: all 0.3s ease;
}

.ai-answer-card:hover {
  background: rgba(48, 52, 58, 0.5);
  border-color: rgba(212, 175, 55, 0.2);
}

.ai-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 16px;
}

.ai-logo {
  display: flex;
  align-items: center;
  gap: 12px;
}

.ai-logo i {
  font-size: 28px;
  background: linear-gradient(135deg, #d4af37 0%, #f4e5a1 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.ai-title {
  font-size: 20px;
  font-weight: 600;
  background: linear-gradient(135deg, #d4af37 0%, #f4e5a1 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.ai-badge {
  padding: 6px 14px;
  background: rgba(212, 175, 55, 0.15);
  border: 1px solid rgba(212, 175, 55, 0.3);
  border-radius: 20px;
  font-size: 12px;
  font-weight: 500;
  color: #d4af37;
}

.ai-content {
  color: rgba(255, 255, 255, 0.85);
  font-size: 15px;
  line-height: 1.7;
  margin-bottom: 16px;
}

.ai-footer {
  display: flex;
  align-items: center;
  gap: 8px;
  color: rgba(255, 255, 255, 0.4);
  font-size: 13px;
}

.ai-footer i {
  font-size: 14px;
}

.ai-loading {
  background: rgba(48, 52, 58, 0.5);
}

.ai-loading-content {
  display: flex;
  align-items: center;
  gap: 12px;
  color: rgba(255, 255, 255, 0.6);
  font-size: 15px;
  padding: 20px 0;
}

.ai-loading-content i {
  font-size: 20px;
}

.maps-container {
  width: 100%;
  margin-top: 20px;
}

.image-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 16px;
  margin-top: 20px;
}

.image-item {
  position: relative;
  background: rgba(48, 52, 58, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  overflow: hidden;
  transition: all 0.3s ease;
}

.image-item:hover {
  background: rgba(48, 52, 58, 0.7);
  border-color: rgba(212, 175, 55, 0.2);
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.3);
}

.image-item a {
  display: block;
}

.image-item img {
  width: 100%;
  height: 200px;
  object-fit: cover;
  display: block;
}

.image-item img[src$=".svg"] {
  object-fit: contain;
  background: rgba(48, 52, 58, 0.8);
  padding: 10px;
}

.image-info {
  padding: 12px;
}

.image-title {
  color: #7ba4ff;
  font-size: 14px;
  text-decoration: none;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.image-title:hover {
  color: #a3c4ff;
  text-decoration: underline;
}

.result-item {
  background: rgba(48, 52, 58, 0.4);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 24px;
  padding: 20px 24px;
  margin-bottom: 16px;
  transition: all 0.3s ease;
}

.result-item:hover {
  background: rgba(48, 52, 58, 0.5);
  border-color: rgba(212, 175, 55, 0.15);
}

.youtube-embed-wrapper {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.youtube-embed-wrapper iframe {
  width: 100%;
  height: 315px;
  border-radius: 12px;
}

.result-title-below {
  display: flex;
  align-items: center;
  gap: 10px;
  color: rgba(255, 255, 255, 0.8);
  font-size: 15px;
  font-weight: 500;
}

.youtube-icon {
  color: #ff0000;
  font-size: 20px;
}

.result-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 8px;
}

.result-left {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  flex: 1;
}

.site-favicon {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  margin-top: 2px;
  flex-shrink: 0;
}

.result-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
  flex: 1;
  min-width: 0;
}

.result-title-link {
  text-decoration: none;
  cursor: pointer;
}

.result-title {
  font-size: 16px;
  font-weight: 500;
  color: #7ba4ff;
  margin: 0;
  cursor: pointer;
  word-wrap: break-word;
  overflow-wrap: break-word;
  transition: color 0.2s ease;
}

.result-title-link:hover .result-title {
  color: #a3c4ff;
  text-decoration: underline;
}

.result-url {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.5);
  text-decoration: none;
  word-wrap: break-word;
  overflow-wrap: break-word;
}

.result-url:hover {
  color: rgba(255, 255, 255, 0.7);
}

.result-menu {
  position: relative;
  margin-left: 16px;
}

.menu-button {
  background: transparent;
  border: none;
  color: rgba(255, 255, 255, 0.4);
  font-size: 16px;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 6px;
  transition: all 0.2s ease;
}

.menu-button:hover {
  background: rgba(255, 255, 255, 0.05);
  color: rgba(255, 255, 255, 0.7);
}

.menu-dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 4px;
  background: rgba(42, 42, 42, 0.95);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 12px;
  padding: 6px;
  min-width: 140px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.4);
  z-index: 10;
}

.menu-option {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 12px;
  background: transparent;
  border: none;
  color: rgba(255, 255, 255, 0.85);
  font-size: 14px;
  font-family: 'Inter', sans-serif;
  cursor: pointer;
  border-radius: 8px;
  transition: all 0.2s ease;
  text-align: left;
}

.menu-option:hover {
  background: rgba(255, 255, 255, 0.08);
}

.menu-option i {
  font-size: 13px;
  width: 14px;
}

.result-date {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.5);
  margin: 0;
  line-height: 1.5;
  word-wrap: break-word;
  overflow-wrap: break-word;
}

.loading-dots {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  padding: 60px 20px;
}

.loading-dots span {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: #d4af37;
  display: inline-block;
  animation: bounce 1.2s ease-in-out infinite;
}

.loading-dots span:nth-child(1) {
  animation-delay: 0s;
}

.loading-dots span:nth-child(2) {
  animation-delay: 0.2s;
}

.loading-dots span:nth-child(3) {
  animation-delay: 0.4s;
}

@keyframes bounce {
  0%, 60%, 100% {
    transform: translateY(0);
    opacity: 0.4;
  }
  30% {
    transform: translateY(-16px);
    opacity: 1;
  }
}

.no-results-message {
  color: rgba(255, 255, 255, 0.6);
  font-size: 16px;
  text-align: center;
  padding: 40px 20px;
  display: flex;
  align-items: center;
  gap: 12px;
  justify-content: center;
}

.pagination {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 32px;
  padding-bottom: 20px;
}

.pagination-button {
  min-width: 40px;
  height: 40px;
  background: rgba(48, 52, 58, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 10px;
  color: rgba(255, 255, 255, 0.7);
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 12px;
  font-family: 'Inter', sans-serif;
}

.pagination-button:hover:not(.disabled) {
  background: rgba(212, 175, 55, 0.15);
  border-color: rgba(212, 175, 55, 0.3);
  color: rgba(255, 255, 255, 0.9);
}

.pagination-button.active {
  background: rgba(212, 175, 55, 0.2);
  border-color: rgba(212, 175, 55, 0.4);
  color: #d4af37;
}

.pagination-button.disabled {
  opacity: 0.3;
  cursor: not-allowed;
}

.pagination-button.page-number {
  min-width: 40px;
}

@media (max-width: 768px) {
  .menu-toggle-fixed {
    top: 20px;
    left: 20px;
  }

  .menu-toggle-fixed.open {
    left: 20px;
  }

  .sidebar.open {
    width: 80%;
    max-width: 300px;
    border-radius: 0;
  }

  .search-header {
    padding: 30px 24px 16px;
  }

  .search-type-selector {
    flex-wrap: wrap;
  }

  .type-button {
    flex: 1 1 45%;
    min-width: 120px;
  }

  .results-container {
    padding: 16px 24px 40px;
  }

  .image-grid {
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 12px;
  }

  .image-item img {
    height: 150px;
  }

  .ai-answer-card {
    padding: 20px;
  }

  .ai-logo i {
    font-size: 24px;
  }

  .ai-title {
    font-size: 18px;
  }

  .result-item {
    padding: 16px 18px;
  }

  .result-title {
    font-size: 15px;
  }

  .result-url {
    font-size: 12px;
  }

  .result-date {
    font-size: 12px;
  }

  .youtube-embed-wrapper iframe {
    height: 220px;
  }

  .pagination {
    gap: 6px;
    margin-top: 24px;
  }

  .pagination-button {
    min-width: 36px;
    height: 36px;
    font-size: 13px;
  }
}

@media (max-width: 480px) {
  .menu-toggle-fixed {
    top: 16px;
    left: 16px;
  }

  .menu-toggle-fixed.open {
    left: 16px;
  }

  .sidebar.open {
    width: 85%;
  }

  .search-header {
    padding: 24px 16px 12px;
  }

  .type-button {
    flex: 1 1 100%;
    font-size: 13px;
    padding: 8px 12px;
  }

  .type-button span {
    display: inline;
  }

  .results-container {
    padding: 12px 16px 32px;
  }

  .image-grid {
    grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  }

  .image-item img {
    height: 120px;
  }

  .ai-answer-card {
    padding: 16px;
    border-radius: 16px;
  }

  .ai-content {
    font-size: 14px;
  }

  .result-item {
    padding: 14px 16px;
    border-radius: 16px;
  }

  .site-favicon {
    width: 20px;
    height: 20px;
  }

  .result-title {
    font-size: 14px;
  }

  .result-url {
    font-size: 11px;
  }

  .result-date {
    font-size: 11px;
  }

  .youtube-embed-wrapper iframe {
    height: 180px;
  }

  .pagination {
    gap: 4px;
    margin-top: 20px;
  }

  .pagination-button {
    min-width: 32px;
    height: 32px;
    font-size: 12px;
    padding: 0 8px;
  }
}
</style>
