<template>
  <div class="app-container">
    <header class="app-header">
      <button aria-label="Grid view" class="icon-btn" @click="toggleGridView(true)">
        <i class="fas fa-th-large"></i>
      </button>
      <button aria-label="Menu" class="icon-btn" @click="toggleControlCenter(true)">
        <i class="fas fa-bars"></i>
      </button>
    </header>

    <main class="main-content" :class="{ 'content-shifted': isControlCenterOpen }">
      <div class="logo-wrapper">
        <h1 class="logo-text">
          Synapic
        </h1>
      </div>

      <div class="search-section">
        <form aria-label="Search form" class="search-form" action="/search" method="GET" @submit="handleSearchSubmit">
          <label class="sr-only" for="search">Ara</label>
          <div class="search-input-container">
            <input 
              class="search-input" 
              id="search" 
              name="query" 
              placeholder="Aklında ne var?" 
              type="search"
              v-model="searchQuery"
              @focus="handleSearchFocus"
              @blur="handleSearchBlur"
              @input="handleSearchInput"
            />
            <button 
              type="button" 
              class="clear-btn" 
              :class="{ 'hidden': !searchQuery }"
              @click="clearSearch"
            >
              <i class="fas fa-times"></i>
            </button>
            <div class="divider"></div>
            <button aria-label="Search" class="search-submit-btn" type="submit">
              <i class="fas fa-magnifying-glass"></i>
            </button>
          </div>
          <input type="hidden" name="kl" :value="currentLang">
        </form>

        <div v-if="showHistory && recentSearches.length > 0" class="history-container">
          <div class="history-content">
            <h3 class="history-title">Recent Searches</h3>
            <div class="history-list">
              <a 
                v-for="(item, index) in recentSearches.slice(0, 5)" 
                :key="index"
                href="#" 
                class="history-item"
                @click.prevent="applyHistorySearch(item.title)"
              >
                <i class="fas fa-magnifying-glass"></i>
                <span>{{ cleanHistoryTitle(item.title) }}</span>
              </a>
            </div>
          </div>
        </div>
      </div>

      <div class="weather-container">
        <span class="weather-info">
          <i class="fas fa-sun weather-icon"></i>
          <span>{{ weatherText }}</span>
        </span>
      </div>
    </main>

    <footer class="app-footer">
      <div class="footer-content">
        <a href="https://github.com/synapic" target="_blank" rel="noopener noreferrer" class="footer-link">
          <i class="fab fa-github"></i>
          <span>GitHub</span>
        </a>
        
        <div class="system-status">
          <span :class="['status-indicator', systemStatusClass]"></span>
          <span class="status-text">{{ systemStatusText }}</span>
        </div>
        
        <a href="https://uptime.synapic.com.tr" target="_blank" rel="noopener noreferrer" class="footer-link">
          <i class="fas fa-chart-line"></i>
          <span>Monitor</span>
        </a>
      </div>
    </footer>

    <div class="overlay control-center-overlay" :class="{ 'is-active': isControlCenterOpen }" @click.self="toggleControlCenter(false)">
      <button class="overlay-close" aria-label="Close" @click="toggleControlCenter(false)">
        <i class="fas fa-xmark"></i>
      </button>
      <div class="overlay-content right-slide">
        <h2 class="overlay-title">Kontrol Merkezi</h2>
        <div class="menu-container">
          <a href="#" class="menu-item" @click.prevent="openSearchOptions">
            <div class="menu-icon-wrapper">
              <i class="fas fa-gear"></i>
            </div>
            <span class="menu-text">Arama Seçenekleri</span>
          </a>
          <a href="/terms" class="menu-item">
            <div class="menu-icon-wrapper">
              <i class="fa-solid fa-lock"></i>
            </div>
            <span class="menu-text">Gizlilik & Kurallar</span>
          </a>
          
          <div class="settings-section">
            <label class="settings-label">Dil Seçimi</label>
            <select class="custom-select" v-model="currentLang">
              <option value="tr">Turkish</option>
<option value="en">English</option>
<option value="de">Deutsch</option>
<option value="fr">Français</option>
<option value="es">Español</option>
<option value="it">Italiano</option>
<option value="pt">Português</option>
<option value="ru">Русский</option>
<option value="ar">العربية</option>
<option value="zh">中文</option>
<option value="ja">日本語</option>
<option value="ko">한국어</option>
<option value="nl">Nederlands</option>
<option value="sv">Svenska</option>
<option value="pl">Polski</option>
<option value="cs">Čeština</option>
<option value="da">Dansk</option>
<option value="fi">Suomi</option>
<option value="no">Norsk</option>
<option value="el">Ελληνικά</option>
<option value="hu">Magyar</option>
<option value="ro">Română</option>
<option value="bg">Български</option>
<option value="uk">Українська</option>
<option value="he">עברית</option>
<option value="fa">فارسی</option>
<option value="hi">हिन्दी</option>
<option value="ur">اردو</option>
<option value="bn">বাংলা</option>
<option value="ta">தமிழ்</option>
<option value="te">తెలుగు</option>
<option value="ml">മലയാളം</option>
<option value="th">ไทย</option>
<option value="vi">Tiếng Việt</option>
<option value="id">Bahasa Indonesia</option>
<option value="ms">Bahasa Melayu</option>
<option value="tl">Filipino</option>
<option value="sw">Kiswahili</option>
<option value="af">Afrikaans</option>
<option value="sq">Shqip</option>
<option value="hr">Hrvatski</option>
<option value="sr">Српски</option>
<option value="sk">Slovenčina</option>
<option value="sl">Slovenščina</option>
<option value="et">Eesti</option>
<option value="lv">Latviešu</option>
<option value="lt">Lietuvių</option>
<option value="is">Íslenska</option>
<option value="ga">Gaeilge</option>
<option value="cy">Cymraeg</option>
<option value="eu">Euskara</option>
<option value="ca">Català</option>
<option value="gl">Galego</option>
<option value="mt">Malti</option>
<option value="lb">Lëtzebuergesch</option>
<option value="mk">Македонски</option>
<option value="ka">ქართული</option>
<option value="hy">Հայերեն</option>
<option value="az">Azərbaycan</option>
<option value="kk">Қазақша</option>
<option value="uz">Oʻzbek</option>
<option value="ky">Кыргызча</option>
<option value="mn">Монгол</option>
<option value="ne">नेपाली</option>
<option value="si">සිංහල</option>
<option value="my">မြန်မာ</option>
<option value="km">ខ្មែរ</option>
<option value="lo">ລາວ</option>
<option value="am">አማርኛ</option>
<option value="zu">isiZulu</option>
<option value="xh">isiXhosa</option>
<option value="st">Sesotho</option>
<option value="tn">Setswana</option>
<option value="ts">Xitsonga</option>
<option value="ss">SiSwati</option>
<option value="ve">Tshivenda</option>
<option value="nr">isiNdebele</option>
<option value="so">Soomaali</option>
<option value="ha">Hausa</option>
<option value="ig">Igbo</option>
<option value="yo">Yorùbá</option>
<option value="ee">Eʋegbe</option>
<option value="rw">Kinyarwanda</option>
<option value="rn">Kirundi</option>
<option value="ny">Chichewa</option>
<option value="mg">Malagasy</option>
<option value="mi">Māori</option>
<option value="sm">Gagana Samoa</option>
<option value="to">Lea Faka-Tonga</option>
<option value="fj">Vosa Vakaviti</option>
<option value="pa">ਪੰਜਾਬੀ</option>
<option value="gu">ગુજરાતી</option>
<option value="mr">मराठी</option>
<option value="kn">ಕನ್ನಡ</option>
<option value="or">ଓଡ଼ିଆ</option>
<option value="as">অসমীয়া</option>
<option value="sa">संस्कृतम्</option>
<option value="bo">བོད་ཡིག</option>
<option value="ug">ئۇيغۇرچە</option>
<option value="tt">Татарча</option>
            </select>

            <button type="button" class="action-btn" @click="runPageTranslation()">
              <i class="fas fa-language"></i>
              <span>Çeviriyi Uygula</span>
            </button>
          </div>

          <div id="google_translate_element" style="display:none !important;"></div>
        </div>
      </div>
    </div>

    <div class="overlay search-options-overlay" :class="{ 'is-active': isSearchOptionsOpen }" @click.self="toggleSearchOptions(false)">
      <button class="overlay-close" aria-label="Close" @click="toggleSearchOptions(false)">
        <i class="fas fa-xmark"></i>
      </button>
      <div class="overlay-content right-slide">
        <h2 class="overlay-title">Arama Seçenekleri</h2>
        <div class="options-list">
          <div class="option-card">
            <div class="option-header">
              <div class="option-info">
                <i class="fas fa-location-dot option-icon"></i>
                <label for="locationBased" class="option-label">Lokal-Tabanlı Arama</label>
              </div>
              <label class="toggle-switch">
                <input type="checkbox" id="locationBased" v-model="isLocationBased">
                <span class="slider"></span>
              </label>
            </div>
          </div>
          
          <div class="option-card">
            <label for="languageSelect" class="settings-label">Dil Seçimi</label>
            <select id="languageSelect" class="custom-select" v-model="currentLang">
              <option value="tr">Turkish</option>
<option value="de">Deutsch</option>
<option value="us">English (US)</option>
<option value="fr">Français</option>
<option value="ru">Русский</option>
<option value="jp">日本語</option>
<option value="es">Español</option>
<option value="it">Italiano</option>
<option value="cn">简体中文</option>
<option value="gb">English (UK)</option>
<option value="br">Português (BR)</option>
<option value="ar">العربية</option>
<option value="nl">Nederlands</option>
<option value="pl">Polski</option>
<option value="kr">한국어</option>
<option value="in">English (IN)</option>
<option value="ca">English (CA)</option>
<option value="au">English (AU)</option>
<option value="sa">العربية (SA)</option>
<option value="se">Svenska</option>
<option value="no">Norsk</option>
<option value="dk">Dansk</option>
<option value="fi">Suomi</option>
<option value="gr">Ελληνικά</option>
<option value="il">עברית</option>
<option value="mx">Español (MX)</option>
<option value="id">Bahasa Indonesia</option>
<option value="th">ไทย</option>
<option value="vn">Tiếng Việt</option>
<option value="za">English (ZA)</option>
            </select>
          </div>
          
          <button class="action-btn primary" @click="saveSettings">
            <i class="fas fa-check"></i>
            <span>Ayarları Kaydet</span>
          </button>
        </div>
      </div>
    </div>

    <div class="overlay grid-view-overlay" :class="{ 'is-active': isGridViewOpen }" @click.self="toggleGridView(false)">
      <button class="overlay-close" aria-label="Close" @click="toggleGridView(false)">
        <i class="fas fa-xmark"></i>
      </button>
      <div class="modal-content">
        <h2 class="modal-title">Hızlı Arama</h2>
        <form aria-label="Quick search form" class="quick-search-form" action="/search" method="GET">
          <div class="quick-search-wrapper">
            <input 
              class="quick-search-input"
              name="query" 
              type="search"
              v-model="quickSearchQuery"
              :placeholder="quickSearchPlaceholder"
              ref="quickSearchInputRef"
            />
            <button aria-label="Quick Search Submit" class="quick-search-btn" type="submit">
              <i class="fas fa-magnifying-glass"></i>
            </button>
          </div>
          <input type="hidden" name="type" :value="quickSearchType">
          <input type="hidden" name="kl" :value="currentLang">
        </form>

        <nav class="quick-links">
          <a href="#" @click.prevent="setQuickSearchType('web')" class="quick-link">
            <i class="fas fa-globe active-icon"></i>
            <span class="link-text">Webde ara</span>
          </a>
          <a href="#" @click.prevent="setQuickSearchType('image')" class="quick-link">
            <i class="fas fa-image active-icon"></i>
            <span class="link-text">Resimlerde ara</span>
          </a>
          <a href="#" @click.prevent="setQuickSearchType('news')" class="quick-link">
            <i class="fas fa-newspaper active-icon"></i>
            <span class="link-text">Haberlerde ara</span>
          </a>
          <a href="#" @click.prevent="setQuickSearchType('maps')" class="quick-link">
            <i class="fa-sharp fa-solid fa-map-location active-icon"></i>
            <span class="link-text">Haritalarda ara</span>
          </a>
          <a href="#" @click.prevent="setQuickSearchType('video')" class="quick-link">
            <i class="fas fa-video active-icon"></i>
            <span class="link-text">Videolarda ara</span>
          </a>
        </nav>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, nextTick } from 'vue';

const systemStatusText = ref('Checking...');
const systemStatusClass = ref('checking');

const checkSystemStatus = async () => {
  try {
    const response = await fetch('http://uptimeapi.synapic.com.tr/api/stats?apiKey=synapicsearch');
    if (!response.ok) throw new Error('API error');
    
    const data = await response.json();
    
    const allOnline = data.every(service => service.status === 'online');
    const someOffline = data.some(service => service.status === 'offline');
    const someIssues = data.some(service => service.status !== 'online' && service.status !== 'offline');
    
    if (allOnline) {
      systemStatusText.value = 'Tüm sistemler çalışır durumda';
      systemStatusClass.value = 'online';
    } else if (someOffline) {
      systemStatusText.value = 'Sistem kesintisi';
      systemStatusClass.value = 'offline';
    } else if (someIssues) {
      systemStatusText.value = 'Kısmi kesinti';
      systemStatusClass.value = 'warning';
    } else {
      systemStatusText.value = 'Durum bilinmiyor';
      systemStatusClass.value = 'checking';
    }
  } catch (error) {
    console.error('Status check failed:', error);
    systemStatusText.value = 'Status unavailable';
    systemStatusClass.value = 'offline';
  }
};


const loadGoogleTranslate = () => {
  if (!window.google || !window.google.translate) {
    const script = document.createElement('script');
    script.src = 'https://translate.google.com/translate_a/element.js?cb=googleTranslateElementInit';
    document.head.appendChild(script);
  }
};

window.googleTranslateElementInit = () => {
  new window.google.translate.TranslateElement({
    pageLanguage: 'tr',
    autoDisplay: false
  }, 'google_translate_element');
};

const runPageTranslation = (langCode) => {
  const lang = langCode || currentLang.value;
  localStorage.setItem('synapicSearchLang', lang);
  const cookieValue = `/tr/${lang}`;
  const expires = "; expires=" + new Date(Date.now() + 86400000).toUTCString();

  document.cookie = `googtrans=${cookieValue}${expires}; path=/`;
  document.cookie = `googtrans=${cookieValue}${expires}; path=/; domain=${window.location.hostname}`;
  
  window.location.reload();
};

onMounted(() => {

  checkSystemStatus();
  setInterval(checkSystemStatus, 30000);

  const savedLang = localStorage.getItem('synapicSearchLang');
  
  if (savedLang) {
    currentLang.value = savedLang;
    const cookieValue = `/tr/${savedLang}`;
    const decodeCookie = decodeURIComponent(document.cookie);
    
    if (!decodeCookie.includes(`googtrans=${cookieValue}`)) {
      const expires = "; expires=" + new Date(Date.now() + 86400000).toUTCString();
      document.cookie = `googtrans=${cookieValue}${expires}; path=/`;
      document.cookie = `googtrans=${cookieValue}${expires}; path=/; domain=${window.location.hostname}`;
      window.location.reload();
    }
  }

  loadGoogleTranslate();
  
  const observer = new MutationObserver(() => {
    const selectors = ['.goog-te-banner-frame', '.skiptranslate', '#goog-gt-tt', '.goog-te-balloon-frame'];
    selectors.forEach(s => {
      const el = document.querySelector(s);
      if (el) el.style.setProperty('display', 'none', 'important');
    });
    document.body.style.top = '0px';
  });

  observer.observe(document.documentElement, { childList: true, subtree: true });
});

const searchQuery = ref('');
const showHistory = ref(false);
const recentSearches = ref([]);
const weatherText = ref('Loading...');
const statusText = ref('All services active');
const statusClass = ref('green');

const isControlCenterOpen = ref(false);
const isSearchOptionsOpen = ref(false);
const isGridViewOpen = ref(false);

const currentLang = ref('en');
const isLocationBased = ref(false);

const quickSearchQuery = ref('');
const quickSearchType = ref('web');
const quickSearchPlaceholder = ref('Webde Ara...');
const quickSearchInputRef = ref(null);

const WEATHERAPI_KEY = '74fad8405f164079a0093741252507';
const IPINFO_API_TOKEN = 'b372d7ccabf9aa';

onMounted(() => {
  const savedLang = localStorage.getItem('synapicSearchLang');
  if (savedLang) currentLang.value = savedLang;


  const savedLoc = localStorage.getItem('synapicLocationBased');
  isLocationBased.value = savedLoc === 'true';

  fetchWeatherAndLocation();
  updateGlobalStatusDisplay();
  setInterval(updateGlobalStatusDisplay, 10000);
});

watch(isControlCenterOpen, (val) => {
  if (val) {
    document.body.style.overflow = 'hidden';
  } else {
    document.body.style.overflow = '';
  }
});

const getDeviceId = () => {
  let deviceId = localStorage.getItem('synapicDeviceId');
  if (!deviceId) {
    const randomDigits = Math.floor(100000 + Math.random() * 900000);
    deviceId = `apm_${randomDigits}`;
    localStorage.setItem('synapicDeviceId', deviceId);
  }
  return deviceId;
};

const fetchWeatherAndLocation = async () => {
  try {
    const ipResponse = await fetch(`https://ipinfo.io/json?token=${IPINFO_API_TOKEN}`);
    if (!ipResponse.ok) throw new Error('IP Info error');
    const ipData = await ipResponse.json();
    const [lat, lon] = ipData.loc.split(',');
    const city = ipData.city || 'Unknown';

    const weatherResponse = await fetch(`https://api.weatherapi.com/v1/current.json?key=${WEATHERAPI_KEY}&q=${lat},${lon}&aqi=no`);
    if (!weatherResponse.ok) throw new Error('Weather API error');
    const weatherData = await weatherResponse.json();
    
    weatherText.value = `${weatherData.location.name || city} • ${Math.round(weatherData.current.temp_c)}°C`;
  } catch (error) {
    console.error(error);
    weatherText.value = 'Location/Weather not available';
  }
};

const updateGlobalStatusDisplay = async () => {
  try {
    const response = await fetch('/api/latest-status');
    let data;
    if (response.ok) {
      data = await response.json();
    } else {
      data = { overallStatus: 'Status Information Unavailable' };
    }
    
    statusText.value = data.overallStatus;
    if (data.overallStatus === 'All services active') statusClass.value = 'green';
    else if (data.overallStatus === 'Some services interrupted') statusClass.value = 'yellow';
    else statusClass.value = 'red';

  } catch (error) {
    statusText.value = 'Status Information Unavailable';
    statusClass.value = 'red';
  }
};

const fetchRecentSearches = async () => {
  const deviceId = getDeviceId();
  try {
    const response = await fetch('/api/history', {
      headers: { 'X-Device-ID': deviceId }
    });
    if (response.ok) {
      recentSearches.value = await response.json();
    }
  } catch (error) {
    console.error(error);
  }
};

const saveSearchToHistory = async () => {
  if (!searchQuery.value.trim()) return;
  const deviceId = getDeviceId();
  const title = searchQuery.value;
  const url = `/search?query=${encodeURIComponent(title)}&type=web&kl=${currentLang.value}`;
  
  try {
    await fetch('/api/save-history', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ title, url, deviceId }),
      keepalive: true
    });
  } catch (error) {
    console.error(error);
  }
};

const handleSearchSubmit = () => {
  saveSearchToHistory();
};

const handleSearchFocus = () => {
  fetchRecentSearches();
  showHistory.value = true;
};

const handleSearchBlur = () => {
  setTimeout(() => {
    showHistory.value = false;
  }, 200);
};

const handleSearchInput = () => {
  if (searchQuery.value.length === 0) showHistory.value = false;
};

const clearSearch = () => {
  searchQuery.value = '';
  showHistory.value = false;
};

const cleanHistoryTitle = (title) => {
  return title.replace(/Search: (.*?) \((.*?)\ssearch\)/, '$1');
};

const applyHistorySearch = (title) => {
  const cleanTitle = cleanHistoryTitle(title);
  searchQuery.value = cleanTitle;
  saveSearchToHistory(); 
  window.location.href = `/search?query=${encodeURIComponent(cleanTitle)}&type=web&kl=${currentLang.value}`;
};

const toggleControlCenter = (show) => {
  isControlCenterOpen.value = show;
};

const openSearchOptions = () => {
  toggleControlCenter(false);
  setTimeout(() => {
    isSearchOptionsOpen.value = true;
  }, 350);
};

const toggleSearchOptions = (show) => {
  isSearchOptionsOpen.value = show;
};

const saveSettings = () => {
  localStorage.setItem('synapicSearchLang', currentLang.value);
  localStorage.setItem('synapicLocationBased', isLocationBased.value);
  toggleSearchOptions(false);
};

const toggleGridView = (show) => {
  isGridViewOpen.value = show;
  if (show) {
    quickSearchQuery.value = searchQuery.value;
    nextTick(() => {
      quickSearchInputRef.value?.focus();
    });
  }
};

const setQuickSearchType = (type) => {
  quickSearchType.value = type;
  switch (type) {
    case 'web': quickSearchPlaceholder.value = "Web'de Ara..."; break;
    case 'image': quickSearchPlaceholder.value = "Resimlerde Ara..."; break;
    case 'news': quickSearchPlaceholder.value = "Haberlerde Ara..."; break;
    case 'video': quickSearchPlaceholder.value = "Videolarda Ara..."; break;
    default: quickSearchPlaceholder.value = "Ara...";
  }
  nextTick(() => {
    quickSearchInputRef.value?.focus();
  });
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&display=swap');
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.7.2/css/all.min.css');

* {
  -webkit-tap-highlight-color: transparent;
  -webkit-touch-callout: none;
}

.app-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #000000;
  color: #ffffff;
  font-family: 'Inter', sans-serif;
  overflow-x: hidden;
  position: relative;
}

.app-header {
  display: flex;
  justify-content: flex-end;
  padding: 1rem;
  gap: 1.5rem;
}

@media (max-width: 768px) {
  .app-header {
    padding: 0.75rem;
    gap: 1rem;
  }
}

.icon-btn {
  color: #ffffff;
  font-size: 1.25rem;
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.5rem;
  min-width: 44px;
  min-height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
}

@media (max-width: 768px) {
  .icon-btn {
    font-size: 1.1rem;
  }
}

.main-content {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 0 1rem;
  padding-bottom: 5rem;
  transition: transform 0.35s cubic-bezier(0.4, 0, 0.2, 1);
}

@media (max-width: 768px) {
  .main-content {
    padding: 0 1rem;
    padding-bottom: 6rem;
  }
  
  .main-content.content-shifted {
    transform: none;
  }
}

.logo-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  max-width: 600px;
  margin-top: 3rem;
  margin-bottom: 2rem;
}

@media (max-width: 768px) {
  .logo-wrapper {
    margin-top: 2rem;
    margin-bottom: 1.5rem;
  }
}

.logo-text {
  color: #ffffff;
  font-weight: 800;
  font-size: 96px;
  line-height: 128px;
  letter-spacing: -0.025em;
  margin: 0;
}

@media (max-width: 768px) {
  .logo-text {
    font-size: 56px;
    line-height: 72px;
  }
}

@media (max-width: 480px) {
  .logo-text {
    font-size: 48px;
    line-height: 64px;
  }
}

.search-section {
  width: 100%;
  max-width: 600px;
  position: relative;
}

@media (max-width: 768px) {
  .search-section {
    max-width: 100%;
  }
}

.search-form {
  width: 100%;
}

.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border-width: 0;
}

.search-input-container {
  position: relative;
  display: flex;
  align-items: center;
  background-color: #1C1C1E;
  border-radius: 9999px;
  padding: 0.5rem 1rem;
  border: 1px solid #1C1C1E;
  transition: box-shadow 0.2s;
  min-height: 48px;
}

@media (max-width: 768px) {
  .search-input-container {
    padding: 0.6rem 1rem;
    min-height: 52px;
  }
}

.search-input-container:focus-within {
  box-shadow: 0 0 0 1px #5A5A5F;
}

.search-input {
  flex-grow: 1;
  background-color: transparent;
  color: #E0E0E0;
  font-size: 1rem;
  border: none;
  outline: none;
}

@media (max-width: 768px) {
  .search-input {
    font-size: 16px;
  }
}

.search-input::placeholder {
  color: #8E8E93;
}

.clear-btn {
  color: #8E8E93;
  margin-left: 0.5rem;
  background: none;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  padding: 0.5rem;
  min-width: 44px;
  min-height: 44px;
}

.clear-btn:hover {
  color: #ffffff;
}

.clear-btn.hidden {
  display: none;
}

.divider {
  height: 1.5rem;
  width: 1px;
  background-color: #5A5A5F;
  margin: 0 0.5rem;
}

.search-submit-btn {
  color: #8E8E93;
  background: none;
  border: none;
  cursor: pointer;
  margin-left: 0.5rem;
  display: flex;
  align-items: center;
  padding: 0.5rem;
  min-width: 44px;
  min-height: 44px;
}

.search-submit-btn:hover {
  color: #ffffff;
}

.history-container {
  position: absolute;
  top: 100%;
  left: 0;
  width: 100%;
  background-color: #2C2C2E;
  border-radius: 1rem;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
  z-index: 10;
  margin-top: 0.5rem;
  max-height: 70vh;
  overflow-y: auto;
}

@media (max-width: 768px) {
  .history-container {
    max-height: 60vh;
  }
}

.history-content {
  padding: 1rem;
}

.history-title {
  font-size: 0.875rem;
  font-weight: 700;
  color: #8E8E93;
  margin-bottom: 0.5rem;
  margin-top: 0;
}

.history-list {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.history-item {
  display: flex;
  align-items: center;
  color: #E0E0E0;
  font-size: 0.875rem;
  text-decoration: none;
  padding: 0.8rem 1.2rem;
  border-radius: 0.6rem;
  transition: background-color 0.2s, color 0.2s;
  min-height: 44px;
}

@media (max-width: 768px) {
  .history-item {
    padding: 1rem 1.2rem;
    min-height: 48px;
  }
}

.history-item:hover {
  background-color: #2a2a2a;
  color: #ffffff;
}

.history-item i {
  margin-right: 0.85rem;
  color: #888;
  font-size: 1.1rem;
  transition: color 0.2s;
}

.history-item:hover i {
  color: #ffffff;
}

.weather-container {
  margin-top: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.875rem;
  color: #8E8E93;
  padding: 0.5rem;
}

.weather-info {
  display: flex;
  align-items: center;
  gap: 0.25rem;
}

.weather-icon {
  color: #facc15;
}

.app-footer {
  width: 100%;
  padding: 1.25rem 2rem;
  position: fixed;
  bottom: 0;
  left: 0;
  background: rgba(10, 10, 10, 0.95);
  backdrop-filter: blur(20px);
  border-top: 1px solid rgba(255, 255, 255, 0.08);
  z-index: 100;
}

@media (max-width: 768px) {
  .app-footer {
    padding: 1rem;
  }
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1.5rem;
}

@media (max-width: 768px) {
  .footer-content {
    gap: 0.5rem;
    padding: 0;
  }
}

@media (max-width: 480px) {
  .footer-content {
    gap: 0.4rem;
  }
}

.footer-link {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #9ca3af;
  text-decoration: none;
  font-size: 0.875rem;
  font-weight: 500;
  padding: 0.5rem 0.75rem;
  border-radius: 0.5rem;
  transition: all 0.2s ease;
  min-height: 44px;
}

@media (max-width: 768px) {
  .footer-link {
    flex-direction: column;
    gap: 0.2rem;
    padding: 0.3rem 0.4rem;
    font-size: 0.65rem;
    min-height: auto;
  }
  
  .footer-link i {
    font-size: 0.95rem;
  }
  
  .footer-link span {
    display: none;
  }
}

@media (max-width: 480px) {
  .footer-link {
    padding: 0.25rem 0.3rem;
  }
  
  .footer-link i {
    font-size: 0.9rem;
  }
}

@media (min-width: 769px) {
  .footer-link:hover {
    color: #10b981;
    background: rgba(16, 185, 129, 0.1);
  }
}

.system-status {
  display: flex;
  align-items: center;
  gap: 0.6rem;
  padding: 0.5rem 1rem;
  background: rgba(28, 28, 30, 0.8);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 2rem;
  transition: all 0.3s ease;
}

@media (max-width: 768px) {
  .system-status {
    padding: 0.35rem 0.7rem;
    gap: 0.35rem;
  }
}

@media (max-width: 480px) {
  .system-status {
    padding: 0.3rem 0.6rem;
    gap: 0.3rem;
  }
}

.system-status:hover {
  background: rgba(28, 28, 30, 0.95);
  border-color: rgba(255, 255, 255, 0.15);
}

.status-indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  position: relative;
  flex-shrink: 0;
}

@media (max-width: 768px) {
  .status-indicator {
    width: 7px;
    height: 7px;
  }
}

.status-indicator::before {
  content: '';
  position: absolute;
  inset: -4px;
  border-radius: 50%;
  animation: ripple 2s ease-out infinite;
}

.status-indicator.online {
  background-color: #10b981;
  box-shadow: 0 0 8px rgba(16, 185, 129, 0.5);
}

.status-indicator.online::before {
  background: radial-gradient(circle, rgba(16, 185, 129, 0.4) 0%, transparent 70%);
}

.status-indicator.warning {
  background-color: #f59e0b;
  box-shadow: 0 0 8px rgba(245, 158, 11, 0.5);
}

.status-indicator.warning::before {
  background: radial-gradient(circle, rgba(245, 158, 11, 0.4) 0%, transparent 70%);
}

.status-indicator.offline {
  background-color: #ef4444;
  box-shadow: 0 0 8px rgba(239, 68, 68, 0.5);
}

.status-indicator.offline::before {
  background: radial-gradient(circle, rgba(239, 68, 68, 0.4) 0%, transparent 70%);
}

.status-indicator.checking {
  background-color: #6b7280;
}

.status-indicator.checking::before {
  background: radial-gradient(circle, rgba(107, 114, 128, 0.4) 0%, transparent 70%);
}

@keyframes ripple {
  0% {
    transform: scale(1);
    opacity: 1;
  }
  100% {
    transform: scale(2.5);
    opacity: 0;
  }
}

.status-text {
  font-size: 0.8rem;
  color: #e5e7eb;
  font-weight: 500;
  letter-spacing: 0.01em;
}

@media (max-width: 768px) {
  .status-text {
    font-size: 0.65rem;
  }
}

@media (max-width: 480px) {
  .status-text {
    font-size: 0.6rem;
  }
}

.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.7);
  z-index: 1000;
  display: flex;
  opacity: 0;
  visibility: hidden;
  transition: opacity 0.3s ease-in-out, visibility 0.3s ease-in-out;
  overscroll-behavior: contain;
}

.overlay.is-active {
  opacity: 1;
  visibility: visible;
}

.control-center-overlay {
  align-items: flex-start;
  justify-content: flex-end;
}

.overlay-content {
  background: #1a1a1a;
  color: white;
  text-align: left;
}

.overlay-content.right-slide {
  position: absolute;
  top: 0;
  right: 0;
  height: 100%;
  width: 360px;
  max-width: 85vw;
  background: linear-gradient(180deg, #1a1a1a 0%, #0d0d0d 100%);
  box-shadow: -8px 0 30px rgba(0, 0, 0, 0.7);
  transform: translateX(100%);
  transition: transform 0.35s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  flex-direction: column;
  padding: 1.75rem;
  overflow-y: auto;
  border-left: 1px solid #2a2a2a;
  -webkit-overflow-scrolling: touch;
}

@media (max-width: 768px) {
  .overlay-content.right-slide {
    width: 90vw;
    max-width: 340px;
    padding: 1.5rem;
  }
}

.overlay.is-active .overlay-content.right-slide {
  transform: translateX(0);
}

.overlay-close {
  position: absolute;
  top: 1rem;
  right: 1.5rem;
  color: #999;
  cursor: pointer;
  z-index: 1001;
  background: none;
  border: none;
  padding: 0.5rem;
  transition: color 0.2s ease-in-out;
  min-width: 44px;
  min-height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
}

@media (max-width: 768px) {
  .overlay-close {
    top: 0.75rem;
    right: 1rem;
  }
}

.overlay-close:hover {
  color: #007aff;
}

.overlay-close i {
  font-size: 1.75rem;
}

.overlay-title {
  font-size: 1.75rem;
  font-weight: 700;
  margin-bottom: 1.75rem;
  color: white;
  letter-spacing: -0.02em;
}

@media (max-width: 768px) {
  .overlay-title {
    font-size: 1.5rem;
    margin-bottom: 1.5rem;
  }
}

.menu-container {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  width: 100%;
}

.menu-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem 1.25rem;
  background: rgba(42, 42, 42, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 1rem;
  color: #e0e0e0;
  text-decoration: none;
  transition: all 0.2s ease;
  min-height: 56px;
}

@media (max-width: 768px) {
  .menu-item {
    padding: 1rem;
    min-height: 60px;
  }
}

.menu-item:hover {
  background: rgba(42, 42, 42, 0.9);
  border-color: rgba(255, 255, 255, 0.1);
  transform: translateY(-1px);
}

.menu-icon-wrapper {
  width: 40px;
  height: 40px;
  border-radius: 0.75rem;
  background: linear-gradient(135deg, #007aff 0%, #0051d5 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.menu-icon-wrapper i {
  font-size: 1.25rem;
  color: white;
}

.menu-text {
  font-size: 0.95rem;
  font-weight: 500;
  letter-spacing: -0.01em;
}

.settings-section {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  margin-top: 1rem;
  padding-top: 1rem;
  border-top: 1px solid rgba(255, 255, 255, 0.06);
}

.settings-label {
  color: #9ca3af;
  font-size: 0.8rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 0.25rem;
}

.custom-select {
  background: rgba(28, 28, 30, 0.8);
  color: white;
  padding: 0.85rem 1rem;
  border-radius: 0.75rem;
  border: 1px solid rgba(255, 255, 255, 0.08);
  font-size: 0.95rem;
  min-height: 48px;
  width: 100%;
  font-family: 'Inter', sans-serif;
  transition: all 0.2s ease;
}

.custom-select:focus {
  outline: none;
  border-color: #007aff;
  background: rgba(28, 28, 30, 0.95);
}

@media (max-width: 768px) {
  .custom-select {
    font-size: 16px;
    padding: 0.95rem 1rem;
    min-height: 52px;
  }
}

.action-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  padding: 0.85rem 1.5rem;
  background: linear-gradient(135deg, #007aff 0%, #0051d5 100%);
  color: white;
  border-radius: 0.75rem;
  border: none;
  font-weight: 600;
  font-size: 0.95rem;
  cursor: pointer;
  transition: all 0.2s ease;
  min-height: 48px;
  box-shadow: 0 4px 12px rgba(0, 122, 255, 0.3);
}

.action-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(0, 122, 255, 0.4);
}

.action-btn.primary {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.action-btn.primary:hover {
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
}

@media (max-width: 768px) {
  .action-btn {
    min-height: 52px;
    font-size: 1rem;
  }
}

.options-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.option-card {
  background: rgba(42, 42, 42, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 1rem;
  padding: 1.25rem;
  transition: all 0.2s ease;
}

.option-card:hover {
  background: rgba(42, 42, 42, 0.8);
  border-color: rgba(255, 255, 255, 0.1);
}

.option-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
}

.option-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  flex: 1;
}

.option-icon {
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #007aff 0%, #0051d5 100%);
  border-radius: 0.5rem;
  color: white;
  font-size: 1rem;
}

.option-label {
  color: white;
  font-size: 0.95rem;
  font-weight: 500;
}

@media (max-width: 768px) {
  .option-label {
    font-size: 0.9rem;
  }
}

.toggle-switch {
  position: relative;
  display: inline-block;
  width: 50px;
  height: 28px;
  flex-shrink: 0;
}

.toggle-switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(58, 58, 58, 0.8);
  transition: .3s;
  border-radius: 28px;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.slider:before {
  position: absolute;
  content: "";
  height: 22px;
  width: 22px;
  left: 3px;
  bottom: 2px;
  background-color: white;
  transition: .3s;
  border-radius: 50%;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

input:checked + .slider {
  background: linear-gradient(135deg, #007aff 0%, #0051d5 100%);
  border-color: transparent;
}

input:checked + .slider:before {
  transform: translateX(22px);
}

:deep(body) {
  top: 0px !important;
  position: static !important;
}

:deep(.goog-te-banner-frame),
:deep(.goog-te-banner),
:deep(iframe.skiptranslate),
:deep(#goog-gt-tt),
:deep(.goog-tooltip),
:deep(.goog-te-balloon-frame),
:deep(.skiptranslate),
:deep(.goog-te-gadget-icon) {
  display: none !important;
  visibility: hidden !important;
  opacity: 0 !important;
  height: 0 !important;
  width: 0 !important;
}

:deep(.goog-text-highlight) {
  background-color: transparent !important;
  box-shadow: none !important;
}

#google_translate_element {
  display: none !important;
}

.grid-view-overlay {
  justify-content: center;
  align-items: center;
  padding: 1rem;
}

@media (max-width: 768px) {
  .grid-view-overlay {
    padding: 0.5rem;
  }
}

.modal-content {
  background-color: #1a1a1a;
  padding: 2.5rem;
  border-radius: 1.25rem;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.6);
  width: 90%;
  max-width: 450px;
  position: relative;
  transform: translateY(-30px);
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1), opacity 0.3s ease-out;
  opacity: 0;
  border: 1px solid rgba(255, 255, 255, 0.08);
  max-height: 90vh;
  overflow-y: auto;
  -webkit-overflow-scrolling: touch;
}

@media (max-width: 768px) {
  .modal-content {
    padding: 1.75rem;
    max-height: 85vh;
    width: 95%;
  }
}

@media (max-width: 480px) {
  .modal-content {
    padding: 1.5rem;
    max-height: 90vh;
  }
}

.grid-view-overlay.is-active .modal-content {
  transform: translateY(0);
  opacity: 1;
}

.modal-title {
  font-size: 1.875rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
  color: white;
  text-align: center;
  letter-spacing: -0.02em;
}

@media (max-width: 768px) {
  .modal-title {
    font-size: 1.5rem;
    margin-bottom: 1.25rem;
  }
}

@media (max-width: 480px) {
  .modal-title {
    font-size: 1.25rem;
  }
}

.quick-search-form {
  width: 100%;
  margin-bottom: 1.5rem;
}

@media (max-width: 768px) {
  .quick-search-form {
    margin-bottom: 1.25rem;
  }
}

.quick-search-wrapper {
  position: relative;
  margin-bottom: 1rem;
}

.quick-search-input {
  width: 100%;
  border-radius: 9999px;
  background-color: #18181B;
  color: #e0e0e0;
  padding: 0.75rem 3rem 0.75rem 1.5rem;
  font-size: 1rem;
  border: 1px solid transparent;
  outline: none;
  box-sizing: border-box;
  min-height: 48px;
}

@media (max-width: 768px) {
  .quick-search-input {
    font-size: 16px;
    padding: 0.85rem 3.5rem 0.85rem 1.5rem;
    min-height: 52px;
  }
}

.quick-search-input:focus {
  box-shadow: 0 0 0 1px #007aff;
  background-color: #2a2a2a;
}

.quick-search-input::placeholder {
  color: #8e8e93;
}

.quick-search-btn {
  position: absolute;
  right: 1rem;
  top: 50%;
  transform: translateY(-50%);
  color: #999;
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.5rem;
  min-width: 44px;
  min-height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.quick-search-btn:hover {
  color: white;
}

.quick-links {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  font-size: 1.125rem;
}

@media (max-width: 768px) {
  .quick-links {
    gap: 0.5rem;
  }
}

.quick-link {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem;
  border-radius: 0.5rem;
  text-decoration: none;
  transition: background-color 0.2s;
  min-height: 48px;
}

@media (max-width: 768px) {
  .quick-link {
    padding: 1rem;
    min-height: 52px;
  }
}

.quick-link:hover {
  background-color: #2a2a2a;
}

.active-icon {
  color: #007aff;
  width: 1.5rem;
  text-align: center;
  font-size: 1.25rem;
}

.link-text {
  color: white;
  font-size: 0.95rem;
}

@media (max-width: 768px) {
  .link-text {
    font-size: 1rem;
  }
}

@media (hover: none) and (pointer: coarse) {
  .icon-btn:active,
  .search-submit-btn:active,
  .clear-btn:active,
  .menu-item:active,
  .quick-link:active,
  .action-btn:active {
    opacity: 0.7;
  }
}
</style>