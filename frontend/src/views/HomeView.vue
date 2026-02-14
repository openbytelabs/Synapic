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
              <span>Google ile giriş yap</span>
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
      <h1 class="hero-title">Synapic</h1>
      <div class="search-container">
        <div class="search-box">
          <i class="fa-solid fa-magnifying-glass search-icon"></i>
          <input 
            type="text" 
            class="search-input" 
            placeholder="Aklından ne geçiyor..."
            v-model="searchQuery"
            @keyup.enter="performSearch"
          />
        </div>
        <p class="search-hint">Aramak için <span class="key">Enter</span> Tuşuna bas.</p>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const menuOpen = ref(false);
const searchQuery = ref('');
const showMoreMenu = ref(false);
const isLoggedIn = ref(false);
const userEmail = ref('');
const googleAuthReady = ref(false);

const performSearch = () => {
  if (searchQuery.value.trim()) {
    router.push(`/search?q=${encodeURIComponent(searchQuery.value)}`);
  }
};

const handleMoreClick = (e) => {
  e.preventDefault();
  e.stopPropagation();
  console.log('More clicked, current state:', showMoreMenu.value);
  showMoreMenu.value = !showMoreMenu.value;
  console.log('New state:', showMoreMenu.value);
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
    console.log('Google Auth başarıyla yüklendi');
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
            localStorage.setItem('apiKey', data.api_key);
            localStorage.setItem('userEmail', userInfo.email);
            isLoggedIn.value = true;
            userEmail.value = userInfo.email;
            showMoreMenu.value = false;
            
            alert('Giriş başarılı! API paneline yönlendiriliyorsunuz...');
            setTimeout(() => {
              router.push('/api');
            }, 500);
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
  
  localStorage.removeItem('apiKey');
  localStorage.removeItem('userEmail');
  isLoggedIn.value = false;
  userEmail.value = '';
  showMoreMenu.value = false;
  alert('Başarıyla çıkış yapıldı');
};

onMounted(async () => {
  const apiKey = localStorage.getItem('apiKey');
  const email = localStorage.getItem('userEmail');
  if (apiKey && email) {
    isLoggedIn.value = true;
    userEmail.value = email;
  }
  
  initGoogleAuth();
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
  height: 100vh;
  display: flex;
  background: linear-gradient(135deg, #1a1a1d 0%, #2d1f2f 50%, #1a1a1d 100%);
  font-family: 'Inter', sans-serif;
  overflow: hidden;
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

.more-menu-item.google-button {
  background: rgba(66, 133, 244, 0.15);
  border: 1px solid rgba(66, 133, 244, 0.3);
}

.more-menu-item.google-button:hover {
  background: rgba(66, 133, 244, 0.25);
  border-color: rgba(66, 133, 244, 0.4);
}

.more-menu-item i {
  font-size: 16px;
  width: 20px;
  text-align: center;
}

.main {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow: hidden;
  width: 100%;
}

.blur-red {
  position: absolute;
  top: 15%;
  left: 15%;
  width: 300px;
  height: 300px;
  background: radial-gradient(circle, rgba(139, 0, 0, 0.08) 0%, transparent 70%);
  filter: blur(30px);
  pointer-events: none;
  z-index: 0;
}

.blur-yellow {
  position: absolute;
  top: 15%;
  right: 20%;
  width: 300px;
  height: 300px;
  background: radial-gradient(circle, rgba(212, 175, 55, 0.06) 0%, transparent 70%);
  filter: blur(30px);
  pointer-events: none;
  z-index: 0;
}

.hero-title {
  font-family: 'Libre Baskerville', serif;
  font-size: 140px;
  font-weight: 700;
  color: #d4af37;
  margin-bottom: 60px;
  position: relative;
  z-index: 1;
  user-select: none;
}

.search-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
  width: 100%;
  max-width: 680px;
  padding: 0 32px;
  position: relative;
  z-index: 1;
}

.search-box {
  width: 100%;
  display: flex;
  align-items: center;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.15);
  border-radius: 32px;
  padding: 16px 28px;
  transition: all 0.3s ease;
}

.search-box:focus-within {
  border-color: rgba(212, 175, 55, 0.4);
  background: rgba(255, 255, 255, 0.05);
}

.search-icon {
  color: rgba(255, 255, 255, 0.4);
  font-size: 20px;
  margin-right: 16px;
}

.search-input {
  flex: 1;
  background: transparent;
  border: none;
  outline: none;
  color: rgba(255, 255, 255, 0.5);
  font-size: 16px;
  font-family: 'Inter', sans-serif;
  font-weight: 400;
}

.search-input::placeholder {
  color: rgba(255, 255, 255, 0.4);
}

.search-hint {
  color: rgba(255, 255, 255, 0.35);
  font-size: 14px;
  font-weight: 400;
  display: flex;
  align-items: center;
  gap: 6px;
}

.key {
  background: rgba(255, 255, 255, 0.08);
  padding: 3px 8px;
  border-radius: 6px;
  font-size: 13px;
  border: 1px solid rgba(255, 255, 255, 0.1);
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

  .hero-title {
    font-size: 80px;
    margin-bottom: 40px;
    text-align: center;
  }

  .search-container {
    padding: 0 24px;
  }

  .search-box {
    padding: 14px 20px;
  }

  .search-icon {
    font-size: 18px;
    margin-right: 12px;
  }

  .search-input {
    font-size: 15px;
  }

  .search-hint {
    font-size: 13px;
  }

  .blur-red,
  .blur-yellow {
    width: 200px;
    height: 200px;
  }
}

@media (max-width: 480px) {
  .menu-toggle-fixed {
    top: 16px;
    left: 16px;
  }

  .hero-title {
    font-size: 60px;
    margin-bottom: 30px;
  }

  .search-container {
    padding: 0 20px;
  }

  .search-box {
    padding: 12px 18px;
    border-radius: 24px;
  }

  .search-input {
    font-size: 14px;
  }

  .search-hint {
    font-size: 12px;
  }

  .key {
    padding: 2px 6px;
    font-size: 12px;
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

  .blur-red,
  .blur-yellow {
    width: 150px;
    height: 150px;
    filter: blur(20px);
  }
}

@media (max-width: 768px) and (orientation: landscape) {
  .hero-title {
    font-size: 50px;
    margin-bottom: 20px;
  }

  .search-container {
    gap: 12px;
  }

  .search-box {
    padding: 10px 16px;
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