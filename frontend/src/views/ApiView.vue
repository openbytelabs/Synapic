<template>
  <div class="app">
    <button class="menu-toggle-fixed" :class="{ open: menuOpen }" @click="menuOpen = !menuOpen">
      <i class="fa-solid fa-table-columns"></i>
    </button>

    <aside class="sidebar" :class="{ open: menuOpen }">
      <div class="logo">Synapic</div>

      <nav class="nav">
        <a href="/" class="nav-item active" @click="menuOpen = false">
          <i class="fa-solid fa-house"></i>
          <span class="nav-text">Anasayfa</span>
        </a>
        <a href="/api" class="nav-item" @click="menuOpen = false">
          <i class="fa-solid fa-link"></i>
          <span class="nav-text">API</span>
        </a>
        <a href="/terms" class="nav-item" @click="menuOpen = false">
            <i class="fa-solid fa-shield-halved"></i>
            <span class="nav-text">Güvenlik & Kurallar</span>
          </a>
          <a href="/settings" class="nav-item" @click="menuOpen = false">
          <i class="fa-solid fa-gear"></i>
          <span class="nav-text">Ayarlar</span>
        </a>
        <div class="more-container">
          <a href="#" class="nav-item" @click="handleMoreClick">
            <i class="fa-solid fa-ellipsis"></i>
            <span class="nav-text">Daha Fazla</span>
          </a>
          <div class="more-menu" v-if="showMoreMenu">
            <button class="more-menu-item google-button" @click="handleGoogleLogin" v-if="!isLoggedIn">
              <i class="fa-brands fa-google"></i>
              <span>Google ile Giriş Yap</span>
            </button>
            <div v-if="isLoggedIn" class="user-info">
              <div class="user-email">{{ userEmail }}</div>
              <button class="more-menu-item" @click="handleSignOut">
                <i class="fa-solid fa-arrow-right-from-bracket"></i>
                <span>Çıkış Yap</span>
              </button>
            </div>
          </div>
        </div>
      </nav>
    </aside>

    <div class="overlay" :class="{ visible: menuOpen }" @click="closeOverlay"></div>

    <main class="main">
      <div class="blur-red"></div>
      <div class="blur-yellow"></div>

      <div class="api-container" v-if="!isLoggedIn">
        <div class="login-prompt">
          <i class="fa-solid fa-lock login-icon"></i>
          <h2 class="login-title">API'ı Kullanmak İçin Giriş Yapın</h2>
          <p class="login-description">API özelliklerine erişmek için Google hesabınızla giriş yapmanız gerekmektedir.</p>
          <button class="login-button" @click="handleGoogleLogin">
            <i class="fa-brands fa-google"></i>
            <span>Google ile giriş yap</span>
          </button>
        </div>
      </div>

      <div class="api-container" v-else>
        <h1 class="api-title">API Kontrol Paneli</h1>

        <div class="stats-grid">
          <div class="stat-card">
            <div class="stat-icon">
              <i class="fa-solid fa-search"></i>
            </div>
            <div class="stat-info">
              <div class="stat-label">Bugün Kullanılan</div>
              <div class="stat-value">{{ stats.searchesToday }}</div>
            </div>
          </div>

          <div class="stat-card">
            <div class="stat-icon">
              <i class="fa-solid fa-clock"></i>
            </div>
            <div class="stat-info">
              <div class="stat-label">Günlük Limit</div>
              <div class="stat-value">{{ stats.dailyLimit === -1 ? 'Sınırsız' : stats.dailyLimit }}</div>
            </div>
          </div>

          <div class="stat-card">
            <div class="stat-icon">
              <i class="fa-solid fa-battery-three-quarters"></i>
            </div>
            <div class="stat-info">
              <div class="stat-label">Kalan Hak</div>
              <div class="stat-value">{{ stats.remaining === -1 ? 'Sınırsız' : stats.remaining }}</div>
            </div>
          </div>
        </div>

        <div class="api-key-section">
          <h2 class="section-title">API Anahtarınız</h2>
          <div class="api-key-card">
            <div class="api-key-display">
              <code class="api-key-text">{{ maskedApiKey }}</code>
              <div class="api-key-actions">
                <button class="action-button" @click="toggleApiKeyVisibility" :title="showFullKey ? 'Gizle' : 'Göster'">
                  <i :class="showFullKey ? 'fa-solid fa-eye-slash' : 'fa-solid fa-eye'"></i>
                </button>
                <button class="action-button" @click="copyApiKey" title="Kopyala">
                  <i class="fa-solid fa-copy"></i>
                </button>
              </div>
            </div>
            <button class="regenerate-button" @click="regenerateKey">
              <i class="fa-solid fa-rotate"></i>
              Yeni Anahtar Oluştur
            </button>
          </div>
        </div>

        <div class="usage-section">
          <h2 class="section-title">Kullanım Zaman Çizelgesi</h2>
          <div class="timeline-card">
            <div class="timeline-header">
              <span>Son 24 Saat</span>
              <span class="timeline-date">{{ currentDate }}</span>
            </div>
            <div class="timeline-chart">
              <div 
                v-for="(hour, index) in usageData" 
                :key="index" 
                class="timeline-bar-container"
              >
                <div class="tooltip">
                  <span class="tooltip-time">{{ hour.hour }}:00</span>
                  <span class="tooltip-count">{{ hour.count }} arama</span>
                </div>
                <div 
                  class="timeline-bar" 
                  :style="{ height: getBarHeight(hour.count) + '%' }"
                ></div>
                <div class="timeline-label">{{ hour.hour }}</div>
              </div>
            </div>
          </div>
        </div>

        <div class="test-section">
          <h2 class="section-title">API Testi</h2>
          <div class="test-card">
            <div class="test-input-group">
              <label class="test-label">Arama Sorgusu</label>
              <input 
                v-model="testQuery" 
                type="text" 
                class="test-input" 
                placeholder="Örnek: github"
                @keyup.enter="testApiSearch"
              />
            </div>
            <button class="test-button" @click="testApiSearch" :disabled="testLoading || !testQuery">
              <i :class="testLoading ? 'fa-solid fa-spinner fa-spin' : 'fa-solid fa-play'"></i>
              {{ testLoading ? 'Aranıyor...' : 'Aramayı Test Et' }}
            </button>
            
            <div v-if="testResponse" class="test-response">
              <h3 class="test-response-title">Sonuç:</h3>
              <div class="test-stats">
                <span class="test-stat-item"><strong>Toplam Sonuç:</strong> {{ testResponse.total }}</span>
                <span class="test-stat-item" v-if="testResponse.total_used !== undefined"><strong>Kullanılan:</strong> {{ testResponse.total_used }}</span>
                <span class="test-stat-item" v-if="testResponse.remaining !== undefined"><strong>Kalan:</strong> {{ testResponse.remaining }}</span>
              </div>
              <div class="test-results">
                <div v-for="(result, idx) in testResponse.results?.slice(0, 3)" :key="idx" class="test-result-item">
                  <a :href="result.url" target="_blank" class="test-result-link">{{ result.title || result.url }}</a>
                  <p class="test-result-desc">{{ result.description }}</p>
                  <span class="test-result-score">Skor: {{ result.score.toFixed(2) }}</span>
                </div>
              </div>
            </div>

            <div v-if="testError" class="test-error">
              <i class="fa-solid fa-exclamation-triangle"></i>
              <span>{{ testError }}</span>
            </div>
          </div>
        </div>

        <div class="documentation-section">
          <h2 class="section-title">API Dokümantasyonu</h2>
          <div class="doc-card">
            <h3 class="doc-subtitle">Endpoint</h3>
            <div class="code-block">
              <code>https://api.synapic.com.tr/api/search?q=YOUR_QUERY&apikey=YOUR_API_KEY</code>
            </div>
            
            <h3 class="doc-subtitle">Kimlik Doğrulama</h3>
            <p class="doc-text">API anahtarınızı query parameter olarak veya header ile gönderin:</p>
            
            <h4 class="doc-subsubtitle">Seçenek 1: Query Parameter (Önerilen)</h4>
            <div class="code-block">
              <code>?apikey={{ apiKey }}</code>
            </div>

            <h4 class="doc-subsubtitle">Seçenek 2: Header</h4>
            <div class="code-block">
              <code>X-API-Key: {{ apiKey }}</code>
            </div>

            <h3 class="doc-subtitle">Örnek İstek (cURL)</h3>
            <div class="code-block">
              <pre>curl -X GET "https://api.synapic.com.tr/api/search?q=github&apikey={{ apiKey }}"</pre>
            </div>

            <h3 class="doc-subtitle">Örnek İstek (JavaScript)</h3>
            <div class="code-block">
              <pre>fetch('https://api.synapic.com.tr/api/search?q=github&apikey={{ apiKey }}')
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error('Error:', error));</pre>
            </div>

            <h3 class="doc-subtitle">Yanıt Formatı</h3>
            <div class="code-block">
              <pre>{
  "results": [
    {
      "url": "https://github.com",
      "title": "GitHub",
      "description": "Where the world builds software",
      "score": 1234.56
    }
  ],
  "total": 1,
  "total_used": 15,      // Bugün kullanılan sorgu sayısı
  "remaining": 85,       // Kalan sorgu hakkı
  "daily_limit": 100     // Günlük limit
}</pre>
            </div>

            <h3 class="doc-subtitle">Hata Kodları</h3>
            <div class="error-codes">
              <div class="error-code-item">
                <strong>API key required</strong>
                <span>API anahtarı eksik</span>
              </div>
              <div class="error-code-item">
                <strong>Invalid or expired API key</strong>
                <span>Geçersiz veya süresi dolmuş API anahtarı</span>
              </div>
              <div class="error-code-item">
                <strong>Daily search limit exceeded</strong>
                <span>Günlük arama limiti aşıldı (100 arama/gün)</span>
              </div>
              <div class="error-code-item">
                <strong>Query parameter 'q' is required</strong>
                <span>Sorgu parametresi eksik</span>
              </div>
            </div>

            <h3 class="doc-subtitle">Limitler</h3>
            <ul class="doc-list">
              <li>Günlük 100 arama hakkı (ücretsiz kullanıcılar)</li>
              <li>Her gece otomatik olarak sıfırlanır</li>
              <li>Master key sınırsız erişim sağlar</li>
              <li>Rate limit: Dakikada maksimum 60 istek</li>
            </ul>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const menuOpen = ref(false);
const showMoreMenu = ref(false);
const apiKey = ref('');
const showFullKey = ref(false);
const isLoggedIn = ref(false);
const userEmail = ref('');
const googleAuthReady = ref(false);
const stats = ref({
  searchesToday: 0,
  dailyLimit: 100,
  remaining: 100
});

const testQuery = ref('');
const testLoading = ref(false);
const testResponse = ref(null);
const testError = ref('');

const usageData = ref([]);
let statsInterval = null;

const maskedApiKey = computed(() => {
  if (!apiKey.value) return '••••••••••••••••••••••••••••••••';
  if (showFullKey.value) return apiKey.value;
  
  const prefix = apiKey.value.substring(0, 15);
  return prefix + '••••••••••••••••••••';
});

const currentDate = computed(() => {
  const now = new Date();
  return now.toLocaleDateString('tr-TR', { day: 'numeric', month: 'long', year: 'numeric' });
});

const getBarHeight = (count) => {
  const maxCount = Math.max(...usageData.value.map(d => d.count), 1);
  if (maxCount === 0) return 0;
  return (count / maxCount) * 100;
};

const initializeUsageData = () => {
  const data = [];
  const now = new Date();
  const currentHour = now.getHours();
  
  for (let i = 0; i < 24; i++) {
    const hour = (currentHour - 23 + i + 24) % 24;
    data.push({
      hour: hour.toString().padStart(2, '0'),
      count: 0
    });
  }
  
  usageData.value = data;
};

const updateUsageData = () => {
  if (stats.value.searchesToday > 0) {
    const now = new Date();
    const currentHour = now.getHours();
    const currentHourStr = currentHour.toString().padStart(2, '0');
    
    const hourIndex = usageData.value.findIndex(d => d.hour === currentHourStr);
    if (hourIndex !== -1) {
      usageData.value[hourIndex].count = stats.value.searchesToday;
    }
  }
};

const fetchStats = async () => {
  try {
    const key = localStorage.getItem('apiKey');
    if (!key) return;

    const response = await fetch('https://api.synapic.com.tr/user/stats', {
      headers: {
        'X-API-Key': key
      }
    });

    if (response.ok) {
      const data = await response.json();
      stats.value = {
        searchesToday: data.searches_today || 0,
        dailyLimit: data.daily_limit || 100,
        remaining: data.remaining || 0
      };
      
      updateUsageData();
    }
  } catch (error) {
    console.error('İstatistik alınamadı:', error);
  }
};

const testApiSearch = async () => {
  if (!testQuery.value.trim()) return;
  
  testLoading.value = true;
  testError.value = '';
  testResponse.value = null;

  try {
    const response = await fetch(`https://api.synapic.com.tr/api/search?q=${encodeURIComponent(testQuery.value)}&apikey=${apiKey.value}`);
    
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || errorData.message || 'Arama başarısız');
    }

    const data = await response.json();
    testResponse.value = data;
    
    // İstatistikleri güncelle
    await fetchStats();
  } catch (error) {
    testError.value = error.message;
  } finally {
    testLoading.value = false;
  }
};

const toggleApiKeyVisibility = () => {
  showFullKey.value = !showFullKey.value;
};

const copyApiKey = async () => {
  try {
    await navigator.clipboard.writeText(apiKey.value);
    alert('API anahtarı panoya kopyalandı!');
  } catch (error) {
    console.error('Kopyalama başarısız:', error);
  }
};

const regenerateKey = async () => {
  if (!confirm('API anahtarınızı yenilemek istediğinizden emin misiniz? Eski anahtarınız artık çalışmayacak.')) {
    return;
  }

  try {
    const response = await fetch('https://api.synapic.com.tr//auth/regenerate', {
      method: 'POST',
      headers: {
        'X-API-Key': apiKey.value
      }
    });

    const data = await response.json();
    
    if (data.success) {
      apiKey.value = data.apikey;
      localStorage.setItem('apiKey', data.apikey);
      alert('API anahtarı başarıyla yenilendi!');
    } else {
      alert('Anahtar yenileme başarısız: ' + (data.message || 'Bilinmeyen hata'));
    }
  } catch (error) {
    console.error('Anahtar yenileme hatası:', error);
    alert('Anahtar yenileme başarısız. Lütfen tekrar deneyin.');
  }
};

const handleMoreClick = (e) => {
  e.preventDefault();
  e.stopPropagation();
  showMoreMenu.value = !showMoreMenu.value;
};

const closeOverlay = () => {
  menuOpen.value = false;
  showMoreMenu.value = false;
};

const loadGoogleScript = () => {
  return new Promise((resolve, reject) => {
    if (window.google?.accounts) {
      resolve();
      return;
    }

    const script = document.createElement('script');
    script.src = 'https://accounts.google.com/gsi/client';
    script.async = true;
    script.defer = true;
    script.onload = resolve;
    script.onerror = reject;
    document.head.appendChild(script);
  });
};

const initGoogleAuth = async () => {
  try {
    await loadGoogleScript();
    googleAuthReady.value = true;
  } catch (error) {
    console.error('Google script yüklenirken hata:', error);
  }
};

const handleGoogleLogin = async () => {
  if (!googleAuthReady.value) {
    alert('Google kimlik doğrulaması yükleniyor. Lütfen birkaç saniye bekleyip tekrar deneyin.');
    await initGoogleAuth();
    if (!googleAuthReady.value) {
      alert('Google kimlik doğrulaması başlatılamadı. Lütfen sayfayı yenileyip tekrar deneyin.');
      return;
    }
  }

  try {
    const client = window.google.accounts.oauth2.initTokenClient({
      client_id: '449095352954-sltui6f9t112ld2abk52s8fgfhle9f0s.apps.googleusercontent.com',
      scope: 'email profile',
      callback: async (response) => {
        if (response.error) {
          console.error('Google login error:', response);
          alert('Giriş başarısız. Lütfen tekrar deneyin.');
          return;
        }

        try {
          const userInfoResponse = await fetch('https://www.googleapis.com/oauth2/v3/userinfo', {
            headers: {
              'Authorization': `Bearer ${response.access_token}`
            }
          });
          
          const userInfo = await userInfoResponse.json();
          
          const backendResponse = await fetch('https://api.synapic.com.tr/auth/google', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              google_id: userInfo.sub,
              email: userInfo.email,
            }),
          });

          const data = await backendResponse.json();
          
          if (data.success) {
            localStorage.setItem('apiKey', data.apikey);
            localStorage.setItem('userEmail', userInfo.email);
            apiKey.value = data.apikey;
            isLoggedIn.value = true;
            userEmail.value = userInfo.email;
            showMoreMenu.value = false;
            
            initializeUsageData();
            fetchStats();
            statsInterval = setInterval(fetchStats, 5000);
            
            alert('Giriş başarılı!');
          } else {
            alert('Giriş başarısız: ' + (data.message || 'Bilinmeyen hata'));
          }
        } catch (error) {
          console.error('Backend error:', error);
          alert('Giriş başarısız. Lütfen tekrar deneyin.');
        }
      },
    });

    client.requestAccessToken();
  } catch (error) {
    console.error('Giriş başarısız:', error);
    alert('Giriş başarısız. Lütfen tekrar deneyin.');
  }
};

const handleSignOut = async () => {
  try {
    if (window.google?.accounts?.id) {
      window.google.accounts.id.disableAutoSelect();
    }
  } catch (error) {
    console.error('Google sign out error:', error);
  }
  
  if (statsInterval) {
    clearInterval(statsInterval);
  }
  
  localStorage.removeItem('apiKey');
  localStorage.removeItem('userEmail');
  isLoggedIn.value = false;
  userEmail.value = '';
  apiKey.value = '';
  showMoreMenu.value = false;
  testQuery.value = '';
  testResponse.value = null;
  testError.value = '';
  alert('Başarıyla çıkış yapıldı');
  router.push('/');
};

onMounted(async () => {
  const key = localStorage.getItem('apiKey');
  const email = localStorage.getItem('userEmail');
  
  if (key && email) {
    apiKey.value = key;
    isLoggedIn.value = true;
    userEmail.value = email;
    
    initializeUsageData();
    fetchStats();
    statsInterval = setInterval(fetchStats, 5000);
  }
  
  initGoogleAuth();
});

onUnmounted(() => {
  if (statsInterval) {
    clearInterval(statsInterval);
  }
});
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Libre+Baskerville:wght@400;700&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600&display=swap');
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.1/css/all.min.css');

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.app {
  width: 100vw;
  min-height: 100vh;
  display: flex;
  background: linear-gradient(135deg, #1a1a1d 0%, #2d1f2f 50%, #1a1a1d 100%);
  font-family: 'Inter', sans-serif;
  overflow-x: hidden;
  position: relative;
}

.menu-toggle-fixed {
  position: fixed;
  top: 28px;
  left: 28px;
  background: rgba(42, 42, 42, 0.8);
  backdrop-filter: blur(10px);
  border: none;
  color: #ffffff;
  font-size: 16px;
  padding: 10px 12px;
  cursor: pointer;
  transition: all 0.4s ease;
  border-radius: 10px;
  z-index: 1000;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
}

.menu-toggle-fixed.open {
  left: 248px;
}

.menu-toggle-fixed:hover {
  background: rgba(255, 255, 255, 0.15);
}

.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(2px);
  opacity: 0;
  pointer-events: none;
  transition: opacity 0.3s ease;
  z-index: 998;
}

.overlay.visible {
  opacity: 1;
  pointer-events: all;
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
  left: -264px;
  top: 0;
  transition: all 0.4s ease;
  overflow: hidden;
  opacity: 0;
  border-radius: 0 24px 24px 0;
  z-index: 999;
}

.sidebar.open {
  left: 0;
  width: 264px;
  padding: 28px 0;
  opacity: 1;
  overflow: visible;
}

.logo {
  font-family: 'Libre Baskerville', serif;
  font-size: 28px;
  font-weight: 700;
  color: #ffffff;
  padding: 0 28px;
  margin-bottom: 32px;
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
  gap: 4px;
  padding: 0 16px;
  position: relative;
  overflow: visible;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 14px 12px;
  color: #e0e0e0;
  text-decoration: none;
  font-size: 16px;
  font-weight: 400;
  transition: all 0.3s ease;
  border-radius: 12px;
  white-space: nowrap;
  opacity: 0;
  cursor: pointer;
}

.sidebar.open .nav-item {
  opacity: 1;
  transition: all 0.3s ease 0.2s;
}

.nav-item:hover {
  background: rgba(255, 255, 255, 0.05);
}

.nav-item i {
  font-size: 18px;
  width: 20px;
  text-align: center;
  flex-shrink: 0;
}

.more-container {
  position: relative;
}

.more-menu {
  position: absolute;
  bottom: 100%;
  left: 0;
  right: 0;
  margin-bottom: 8px;
  background: rgba(30, 30, 30, 0.98);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  padding: 8px;
  z-index: 100;
  min-width: 220px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.5);
  animation: slideUp 0.2s ease;
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.more-menu-item {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background: transparent;
  border: none;
  color: rgba(255, 255, 255, 0.85);
  font-size: 15px;
  font-family: 'Inter', sans-serif;
  cursor: pointer;
  border-radius: 8px;
  transition: all 0.3s ease;
  text-align: left;
}

.more-menu-item:hover {
  background: rgba(255, 255, 255, 0.05);
}

.more-menu-item i {
  font-size: 16px;
  width: 20px;
  text-align: center;
}

.google-button {
  background: rgba(66, 133, 244, 0.15);
  border: 1px solid rgba(66, 133, 244, 0.3);
}

.google-button:hover {
  background: rgba(66, 133, 244, 0.25);
  border-color: rgba(66, 133, 244, 0.4);
}

.user-info {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.user-email {
  padding: 12px;
  color: rgba(255, 255, 255, 0.7);
  font-size: 14px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  word-break: break-word;
}

.main {
  flex: 1;
  min-height: 100vh;
  overflow-y: auto;
  position: relative;
  z-index: 1;
}

.blur-red,
.blur-yellow {
  position: fixed;
  width: 500px;
  height: 500px;
  border-radius: 50%;
  filter: blur(120px);
  opacity: 0.15;
  pointer-events: none;
  z-index: 0;
}

.blur-red {
  background: radial-gradient(circle, rgba(220, 38, 38, 0.4) 0%, transparent 70%);
  top: 10%;
  right: 10%;
}

.blur-yellow {
  background: radial-gradient(circle, rgba(212, 175, 55, 0.3) 0%, transparent 70%);
  bottom: 15%;
  left: 10%;
}

.api-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 80px 40px 60px;
  position: relative;
  z-index: 1;
}

.login-prompt {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  min-height: 60vh;
  padding: 60px 20px;
}

.login-icon {
  font-size: 80px;
  color: #d4af37;
  margin-bottom: 32px;
  opacity: 0.9;
}

.login-title {
  font-size: 36px;
  font-weight: 700;
  color: #ffffff;
  margin-bottom: 16px;
  line-height: 1.3;
}

.login-description {
  font-size: 16px;
  color: rgba(255, 255, 255, 0.7);
  margin-bottom: 40px;
  max-width: 500px;
  line-height: 1.6;
}

.login-button {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px 32px;
  background: linear-gradient(135deg, #4285f4 0%, #357ae8 100%);
  border: none;
  border-radius: 12px;
  color: white;
  font-size: 16px;
  font-weight: 600;
  font-family: 'Inter', sans-serif;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 20px rgba(66, 133, 244, 0.3);
}

.login-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 30px rgba(66, 133, 244, 0.4);
}

.api-title {
  font-size: 48px;
  font-weight: 700;
  color: #ffffff;
  margin-bottom: 40px;
  text-align: center;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 24px;
  margin-bottom: 48px;
}

.stat-card {
  background: rgba(48, 52, 58, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  padding: 24px;
  display: flex;
  align-items: center;
  gap: 20px;
  transition: all 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.2);
  border-color: rgba(212, 175, 55, 0.3);
}

.stat-icon {
  width: 60px;
  height: 60px;
  background: linear-gradient(135deg, rgba(212, 175, 55, 0.2) 0%, rgba(212, 175, 55, 0.1) 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  color: #d4af37;
}

.stat-info {
  flex: 1;
}

.stat-label {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.6);
  margin-bottom: 8px;
  font-weight: 500;
}

.stat-value {
  font-size: 32px;
  font-weight: 700;
  color: #ffffff;
}

.api-key-section,
.usage-section,
.test-section,
.documentation-section {
  margin-bottom: 48px;
}

.section-title {
  font-size: 24px;
  font-weight: 700;
  color: rgba(255, 255, 255, 0.9);
  margin-bottom: 20px;
}

.api-key-card {
  background: rgba(48, 52, 58, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.api-key-display {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 16px;
  background: rgba(0, 0, 0, 0.3);
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.api-key-text {
  flex: 1;
  font-family: 'Courier New', monospace;
  font-size: 14px;
  color: #d4af37;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.api-key-actions {
  display: flex;
  gap: 8px;
}

.action-button {
  width: 36px;
  height: 36px;
  background: rgba(212, 175, 55, 0.1);
  border: 1px solid rgba(212, 175, 55, 0.2);
  border-radius: 8px;
  color: rgba(212, 175, 55, 0.8);
  font-size: 14px;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.action-button:hover {
  background: rgba(212, 175, 55, 0.15);
  border-color: rgba(212, 175, 55, 0.3);
  color: #d4af37;
}

.regenerate-button {
  width: 100%;
  padding: 12px 20px;
  background: rgba(212, 175, 55, 0.15);
  border: 1px solid rgba(212, 175, 55, 0.3);
  border-radius: 10px;
  color: #d4af37;
  font-size: 15px;
  font-weight: 500;
  font-family: 'Inter', sans-serif;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.regenerate-button:hover {
  background: rgba(212, 175, 55, 0.25);
  border-color: rgba(212, 175, 55, 0.4);
}

.timeline-card {
  background: rgba(48, 52, 58, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  padding: 24px;
}

.timeline-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  color: rgba(255, 255, 255, 0.7);
  font-size: 14px;
}

.timeline-date {
  color: rgba(255, 255, 255, 0.5);
}

.timeline-chart {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  height: 200px;
  gap: 4px;
}

.timeline-bar-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100%;
  justify-content: flex-end;
  position: relative;
}

.timeline-bar-container:hover .tooltip {
  opacity: 1;
  visibility: visible;
  transform: translateX(-50%) translateY(0);
}

.timeline-bar {
  width: 100%;
  background: linear-gradient(180deg, #d4af37 0%, rgba(212, 175, 55, 0.4) 100%);
  border-radius: 4px 4px 0 0;
  transition: all 0.3s ease;
  cursor: pointer;
  min-height: 2px;
}

.timeline-bar:hover {
  background: linear-gradient(180deg, #f0c857 0%, rgba(240, 200, 87, 0.6) 100%);
  transform: scaleY(1.05);
}

.tooltip {
  position: absolute;
  bottom: 100%;
  left: 50%;
  transform: translateX(-50%) translateY(-5px);
  background: rgba(0, 0, 0, 0.95);
  color: #fff;
  padding: 10px 14px;
  border-radius: 8px;
  font-size: 13px;
  white-space: nowrap;
  opacity: 0;
  visibility: hidden;
  pointer-events: none;
  transition: all 0.3s ease;
  margin-bottom: 12px;
  z-index: 1000;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.5);
  border: 1px solid rgba(212, 175, 55, 0.3);
}

.tooltip::after {
  content: "";
  position: absolute;
  top: 100%;
  left: 50%;
  transform: translateX(-50%);
  border: 6px solid transparent;
  border-top-color: rgba(0, 0, 0, 0.95);
}

.tooltip-time {
  color: #d4af37;
  font-weight: 600;
  display: block;
  margin-bottom: 4px;
  font-size: 13px;
}

.tooltip-count {
  color: rgba(255, 255, 255, 0.9);
  font-size: 12px;
}

.timeline-label {
  font-size: 10px;
  color: rgba(255, 255, 255, 0.4);
  margin-top: 8px;
}

.test-card {
  background: rgba(48, 52, 58, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  padding: 24px;
}

.test-input-group {
  margin-bottom: 16px;
}

.test-label {
  display: block;
  color: rgba(255, 255, 255, 0.7);
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 8px;
}

.test-input {
  width: 100%;
  padding: 12px 16px;
  background: rgba(0, 0, 0, 0.3);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 10px;
  color: #ffffff;
  font-size: 15px;
  font-family: 'Inter', sans-serif;
  transition: all 0.3s ease;
}

.test-input:focus {
  outline: none;
  border-color: rgba(212, 175, 55, 0.5);
  background: rgba(0, 0, 0, 0.4);
}

.test-button {
  width: 100%;
  padding: 14px 20px;
  background: linear-gradient(135deg, rgba(212, 175, 55, 0.2) 0%, rgba(212, 175, 55, 0.15) 100%);
  border: 1px solid rgba(212, 175, 55, 0.3);
  border-radius: 10px;
  color: #d4af37;
  font-size: 15px;
  font-weight: 600;
  font-family: 'Inter', sans-serif;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
}

.test-button:hover:not(:disabled) {
  background: linear-gradient(135deg, rgba(212, 175, 55, 0.3) 0%, rgba(212, 175, 55, 0.25) 100%);
  border-color: rgba(212, 175, 55, 0.5);
}

.test-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.test-response {
  margin-top: 20px;
  padding: 20px;
  background: rgba(0, 0, 0, 0.3);
  border: 1px solid rgba(212, 175, 55, 0.2);
  border-radius: 10px;
}

.test-response-title {
  font-size: 16px;
  font-weight: 600;
  color: #d4af37;
  margin-bottom: 16px;
}

.test-stats {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
  margin-bottom: 16px;
  padding-bottom: 16px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.test-stat-item {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.7);
}

.test-stat-item strong {
  color: rgba(255, 255, 255, 0.9);
}

.test-results {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.test-result-item {
  padding: 12px;
  background: rgba(255, 255, 255, 0.03);
  border-radius: 8px;
  border: 1px solid rgba(255, 255, 255, 0.05);
}

.test-result-link {
  display: block;
  color: #d4af37;
  font-weight: 500;
  font-size: 14px;
  text-decoration: none;
  margin-bottom: 6px;
  transition: color 0.3s ease;
}

.test-result-link:hover {
  color: #f0c857;
  text-decoration: underline;
}

.test-result-desc {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.6);
  margin-bottom: 6px;
  line-height: 1.5;
}

.test-result-score {
  font-size: 11px;
  color: rgba(212, 175, 55, 0.7);
}

.test-error {
  margin-top: 16px;
  padding: 12px 16px;
  background: rgba(220, 38, 38, 0.1);
  border: 1px solid rgba(220, 38, 38, 0.3);
  border-radius: 8px;
  color: #ef4444;
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.doc-card {
  background: rgba(48, 52, 58, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  padding: 24px;
}

.doc-subtitle {
  font-size: 18px;
  font-weight: 600;
  color: rgba(255, 255, 255, 0.85);
  margin: 24px 0 12px;
}

.doc-subtitle:first-child {
  margin-top: 0;
}

.doc-subsubtitle {
  font-size: 16px;
  font-weight: 500;
  color: rgba(255, 255, 255, 0.75);
  margin: 16px 0 8px;
}

.doc-text {
  color: rgba(255, 255, 255, 0.7);
  font-size: 14px;
  margin-bottom: 12px;
  line-height: 1.6;
}

.doc-list {
  list-style-position: inside;
  color: rgba(255, 255, 255, 0.7);
  font-size: 14px;
  line-height: 1.8;
  padding-left: 4px;
}

.doc-list li {
  margin-bottom: 8px;
}

.code-block {
  background: rgba(0, 0, 0, 0.4);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  padding: 16px;
  overflow-x: auto;
  margin-bottom: 16px;
}

.code-block code,
.code-block pre {
  font-family: 'Courier New', monospace;
  font-size: 13px;
  color: #d4af37;
  line-height: 1.6;
}

.code-block pre {
  margin: 0;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.error-codes {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-top: 12px;
}

.error-code-item {
  display: flex;
  flex-direction: column;
  gap: 4px;
  padding: 12px;
  background: rgba(0, 0, 0, 0.3);
  border-radius: 8px;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.error-code-item strong {
  color: #ef4444;
  font-size: 14px;
  font-family: 'Courier New', monospace;
}

.error-code-item span {
  color: rgba(255, 255, 255, 0.6);
  font-size: 13px;
}

@media (max-width: 768px) {
  .menu-toggle-fixed {
    top: 20px;
    left: 20px;
    width: 44px;
    height: 44px;
    font-size: 18px;
  }

  .menu-toggle-fixed.open {
    left: 20px;
  }

  .sidebar {
    border-radius: 0;
  }

  .sidebar.open {
    width: 80%;
    max-width: 300px;
  }

  .api-container {
    padding: 60px 24px 40px;
  }

  .api-title {
    font-size: 36px;
  }

  .login-title {
    font-size: 28px;
  }

  .login-icon {
    font-size: 60px;
  }

  .stats-grid {
    grid-template-columns: 1fr;
  }

  .stat-card {
    padding: 20px;
  }

  .stat-icon {
    width: 50px;
    height: 50px;
    font-size: 20px;
  }

  .stat-value {
    font-size: 28px;
  }

  .timeline-chart {
    height: 150px;
  }
}

@media (max-width: 480px) {
  .menu-toggle-fixed {
    top: 16px;
    left: 16px;
  }

  .api-container {
    padding: 50px 16px 32px;
  }

  .api-title {
    font-size: 28px;
  }

  .login-title {
    font-size: 24px;
  }

  .login-description {
    font-size: 14px;
  }

  .section-title {
    font-size: 20px;
  }

  .api-key-display {
    flex-direction: column;
    align-items: stretch;
  }

  .api-key-actions {
    justify-content: center;
  }

  .code-block {
    padding: 12px;
  }

  .code-block code,
  .code-block pre {
    font-size: 11px;
  }

  .sidebar.open {
    width: 85%;
  }

  .logo {
    font-size: 24px;
    padding: 0 20px;
  }

  .nav {
    padding: 0 12px;
  }

  .nav-item {
    padding: 12px 10px;
    font-size: 15px;
  }
}

@media (hover: none) and (pointer: coarse) {
  .menu-toggle-fixed,
  .nav-item {
    -webkit-tap-highlight-color: transparent;
  }

  .nav-item:active {
    background: rgba(255, 255, 255, 0.1);
  }

  .menu-toggle-fixed:active {
    background: rgba(255, 255, 255, 0.2);
  }
}
</style>