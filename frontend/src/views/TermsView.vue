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
        <div class="more-container">
          <a href="#" class="nav-item" @click="handleMoreClick">
            <i class="fa-solid fa-ellipsis"></i>
            <span class="nav-text">{{ t.nav.more }}</span>
          </a>
          <div class="more-menu" v-if="showMoreMenu">
            <button class="more-menu-item google-button" @click="handleGoogleLogin" v-if="!isLoggedIn">
              <i class="fa-brands fa-google"></i>
              <span>{{ t.nav.loginGoogle }}</span>
            </button>
            <div v-if="isLoggedIn" class="user-info">
              <div class="user-email">{{ userEmail }}</div>
              <button class="more-menu-item" @click="handleSignOut">
                <i class="fa-solid fa-arrow-right-from-bracket"></i>
                <span>{{ t.nav.logout }}</span>
              </button>
            </div>
          </div>
        </div>
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

      <!-- TURKISH CONTENT -->
      <div class="content-wrapper" v-if="lang === 'tr'">
        <h1 class="main-title">Gizlilik Politikası ve Hizmet Şartları</h1>
        <p class="subtitle">Synapic, kullanıcı gizliliğine öncelik veren ve kendi web dizinini koruyan bağımsız bir arama motoru altyapısıdır. Meta arama motoru mimarilerinin aksine, Synapic üçüncü taraf arama sağlayıcılarından sonuçları bir araya getirmez.</p>
        <p class="subtitle">Bu belge, hizmet ekosistemimiz içinde kullanıcı verilerinin işlenmesi ve korunmasında kullanılan yöntemleri açıklamaktadır.</p>
        <div class="section-separator"></div>
        <h2 class="section-title">1. Gizlilik Politikası</h2>
        <p class="section-text">Synapic'te kullanıcı gizliliği en önemli önceliğimizdir. Kişisel kullanıcı verilerinin üçüncü taraflarla paylaşılmaması konusunda katı bir politika izliyoruz.</p>
        <h3 class="subsection-title">1.1 Bilgi Toplama</h3>
        <ul class="bullet-list">
          <li><strong>Arama Geçmişi:</strong> Tarayıcı localStorage aracılığıyla yerel olarak tutulur.</li>
          <li><strong>Dil Tercihleri:</strong> Tutarlı arama sonuçları için yerel olarak korunur.</li>
          <li><strong>Arayüz Tercihleri:</strong> Yalnızca yerel depolamada saklanan tema yapılandırmaları.</li>
        </ul>
        <h3 class="subsection-title">1.2 Dizinlenmiş Veri Mimarisi</h3>
        <ul class="bullet-list">
          <li>Mevcut veritabanında yaklaşık <strong>20.000.000+ dizinlenmiş URL</strong></li>
          <li>Yalnızca herkese açık web sayfalarının dizinlenmesi</li>
          <li>Verilerin yalnızca arama sonucu sağlama amacıyla kullanılması</li>
        </ul>
        <h3 class="subsection-title">1.3 Çerezler ve Takip Politikası</h3>
        <ul class="bullet-list">
          <li>Takip çerezleri uygulaması yoktur</li>
          <li>Parmak izi (fingerprinting) teknikleri kullanılmaz</li>
          <li>Üçüncü taraf takip mekanizmalarının hiçbiri bulunmaz</li>
        </ul>
        <h3 class="subsection-title">1.4 İletişim Bilgileri</h3>
        <p class="section-text">E-posta: yigitkabak@tuta.io</p>
        <div class="section-separator"></div>
        <h2 class="section-title">2. Hizmet Şartları</h2>
        <p class="section-text">Synapic hizmetlerinin kullanımı aşağıdaki şartların kabul edildiği anlamına gelir.</p>
        <h3 class="subsection-title">2.1 Hizmet Kullanım Parametreleri</h3>
        <p class="section-text">Hizmet erişimi yalnızca kişisel ve ticari olmayan amaçlar içindir. Kötüye kullanım veya izinsiz otomasyon kesinlikle yasaktır.</p>
        <h3 class="subsection-title">2.2 Fikri Mülkiyet Hakları</h3>
        <p class="section-text">Tüm Synapic içeriği geçerli fikri mülkiyet yasaları kapsamında korunmaktadır. Açık kaynaklı bileşenler <strong>BSD 3-Clause Lisansı</strong> çerçevesinde çalışır.</p>
        <h3 class="subsection-title">2.3 Yapay Zeka Entegrasyonu</h3>
        <p class="section-text">Yapay zeka tarafından oluşturulan içerik yanlışlıklar içerebilir ve kritik karar verme için tek dayanak noktası olmamalıdır.</p>
        <h3 class="subsection-title">2.4 Yargı Yetkisi</h3>
        <p class="section-text">Bu şartlar Türkiye Cumhuriyeti yasalarına tabidir.</p>
        <div class="section-separator"></div>
        <h2 class="section-title">3. Açık Kaynak Lisansı</h2>
        <h3 class="subsection-title">3.1 BSD 3-Clause License</h3>
        <pre class="license-block">BSD 3-Clause License

Copyright (c) 2023, openbytelabs

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES ARE DISCLAIMED.</pre>
        <div class="section-separator"></div>
        <h2 class="section-title">4. Teknik Altyapı</h2>
        <ul class="bullet-list">
          <li><strong>Frontend:</strong> Vue.js 3 (Composition API)</li>
          <li><strong>Backend:</strong> Golang</li>
          <li><strong>Dizinleme:</strong> Tescilli mimari</li>
        </ul>
        <div class="section-separator"></div>
        <div class="tech-credits">
          <h3 class="subsection-title">❤️ ile şunlar kullanılarak yapıldı</h3>
          <div class="tech-logos">
            <a href="https://www.linux.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link"><img src="https://upload.wikimedia.org/wikipedia/commons/3/35/Tux.svg" alt="Linux" class="tech-logo"></a>
            <a href="https://vuejs.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link"><img src="https://upload.wikimedia.org/wikipedia/commons/9/95/Vue.js_Logo_2.svg" alt="Vue.js" class="tech-logo"></a>
            <a href="https://nodejs.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link"><img src="https://upload.wikimedia.org/wikipedia/commons/d/d9/Node.js_logo.svg" alt="Node.js" class="tech-logo"></a>
            <a href="https://golang.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link"><img src="https://duckduckgo.com/i/77cac52781feeb4a.png" alt="Golang" class="tech-logo"></a>
          </div>
        </div>
      </div>

      <!-- ENGLISH CONTENT -->
      <div class="content-wrapper" v-else>
        <h1 class="main-title">Privacy Policy & Terms of Service</h1>
        <p class="subtitle">Synapic is an independent search engine infrastructure that prioritizes user privacy and maintains its own web index. Unlike meta search engines, Synapic does not aggregate results from third-party providers.</p>
        <p class="subtitle">This document outlines the methods used to handle and protect user data within our service ecosystem.</p>
        <div class="section-separator"></div>
        <h2 class="section-title">1. Privacy Policy</h2>
        <p class="section-text">User privacy is our top priority at Synapic. We follow a strict policy of not sharing personal user data with third parties.</p>
        <h3 class="subsection-title">1.1 Information Collection</h3>
        <ul class="bullet-list">
          <li><strong>Search History:</strong> Stored locally via browser localStorage and can be deleted by the user at any time.</li>
          <li><strong>Language Preferences:</strong> Kept locally to provide consistent search results.</li>
          <li><strong>Interface Preferences:</strong> Theme and UI configurations stored only in local storage.</li>
        </ul>
        <h3 class="subsection-title">1.2 Indexed Data Architecture</h3>
        <ul class="bullet-list">
          <li>Approximately <strong>20,000,000+ indexed URLs</strong> in the current database</li>
          <li>Only publicly accessible web pages are indexed</li>
          <li>Data is used solely for providing search results</li>
        </ul>
        <h3 class="subsection-title">1.3 Cookies & Tracking Policy</h3>
        <ul class="bullet-list">
          <li>No tracking cookies are used</li>
          <li>No fingerprinting techniques are employed</li>
          <li>No third-party tracking mechanisms exist</li>
        </ul>
        <h3 class="subsection-title">1.4 Contact Information</h3>
        <p class="section-text">Email: yigitkabak@tuta.io</p>
        <div class="section-separator"></div>
        <h2 class="section-title">2. Terms of Service</h2>
        <p class="section-text">Use of Synapic services implies acceptance of the following terms.</p>
        <h3 class="subsection-title">2.1 Service Usage Parameters</h3>
        <p class="section-text">Service access is for personal and non-commercial purposes only. Misuse or unauthorized automation is strictly prohibited.</p>
        <h3 class="subsection-title">2.2 Intellectual Property Rights</h3>
        <p class="section-text">All Synapic content is protected under applicable intellectual property laws. Open-source components operate under the <strong>BSD 3-Clause License</strong>.</p>
        <h3 class="subsection-title">2.3 AI Integration</h3>
        <p class="section-text">AI-generated content may contain inaccuracies and should not be used as the sole basis for critical decision-making.</p>
        <h3 class="subsection-title">2.4 Jurisdiction</h3>
        <p class="section-text">These terms are governed by the laws of the Republic of Turkey.</p>
        <div class="section-separator"></div>
        <h2 class="section-title">3. Open Source License</h2>
        <h3 class="subsection-title">3.1 BSD 3-Clause License</h3>
        <pre class="license-block">BSD 3-Clause License

Copyright (c) 2023, openbytelabs

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES ARE DISCLAIMED.</pre>
        <div class="section-separator"></div>
        <h2 class="section-title">4. Technical Infrastructure</h2>
        <ul class="bullet-list">
          <li><strong>Frontend:</strong> Vue.js 3 (Composition API)</li>
          <li><strong>Backend:</strong> Golang</li>
          <li><strong>Indexing:</strong> Proprietary architecture</li>
        </ul>
        <div class="section-separator"></div>
        <div class="tech-credits">
          <h3 class="subsection-title">Made with ❤️ using</h3>
          <div class="tech-logos">
            <a href="https://www.linux.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link"><img src="https://upload.wikimedia.org/wikipedia/commons/3/35/Tux.svg" alt="Linux" class="tech-logo"></a>
            <a href="https://vuejs.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link"><img src="https://upload.wikimedia.org/wikipedia/commons/9/95/Vue.js_Logo_2.svg" alt="Vue.js" class="tech-logo"></a>
            <a href="https://nodejs.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link"><img src="https://upload.wikimedia.org/wikipedia/commons/d/d9/Node.js_logo.svg" alt="Node.js" class="tech-logo"></a>
            <a href="https://golang.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link"><img src="https://duckduckgo.com/i/77cac52781feeb4a.png" alt="Golang" class="tech-logo"></a>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { lang, detectLanguage, tr, en } from '../composables/useLocale.js';

const router = useRouter();
const menuOpen = ref(false);
const showMoreMenu = ref(false);
const isLoggedIn = ref(false);
const userEmail = ref('');
const googleAuthReady = ref(false);

const t = computed(() => lang.value === 'tr' ? tr : en);

const handleMoreClick = (e) => { e.preventDefault(); e.stopPropagation(); showMoreMenu.value = !showMoreMenu.value; };
const closeOverlay = () => { menuOpen.value = false; showMoreMenu.value = false; };

const loadGoogleScript = () => {
  return new Promise((resolve, reject) => {
    if (window.google?.accounts) { resolve(); return; }
    const script = document.createElement('script');
    script.src = 'https://accounts.google.com/gsi/client';
    script.async = true; script.defer = true;
    script.onload = resolve; script.onerror = reject;
    document.head.appendChild(script);
  });
};

const initGoogleAuth = async () => {
  try { await loadGoogleScript(); googleAuthReady.value = true; }
  catch (error) { console.error('Google script error:', error); }
};

const handleGoogleLogin = async () => {
  if (!googleAuthReady.value) {
    alert(t.value.authLoading);
    await initGoogleAuth();
    if (!googleAuthReady.value) { alert(t.value.authFail); return; }
  }
  try {
    const client = window.google.accounts.oauth2.initTokenClient({
      client_id: '449095352954-sltui6f9t112ld2abk52s8fgfhle9f0s.apps.googleusercontent.com',
      scope: 'email profile',
      callback: async (response) => {
        if (response.error) { alert(t.value.loginFail); return; }
        try {
          const userInfoResponse = await fetch('https://www.googleapis.com/oauth2/v3/userinfo', {
            headers: { 'Authorization': `Bearer ${response.access_token}` }
          });
          const userInfo = await userInfoResponse.json();
          const backendResponse = await fetch('https://api.synapic.com.tr/auth/google', {
            method: 'POST', headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ google_id: userInfo.sub, email: userInfo.email }),
          });
          const data = await backendResponse.json();
          if (data.success) {
            localStorage.setItem('apiKey', data.apikey);
            localStorage.setItem('userEmail', userInfo.email);
            isLoggedIn.value = true;
            userEmail.value = userInfo.email;
            showMoreMenu.value = false;
            alert(t.value.loginSuccess);
            setTimeout(() => { router.push('/api'); }, 500);
          } else { alert(t.value.loginFailMsg + (data.message || '')); }
        } catch { alert(t.value.loginFail); }
      },
    });
    client.requestAccessToken();
  } catch { alert(t.value.loginFail); }
};

const handleSignOut = async () => {
  try { if (window.google?.accounts?.id) { window.google.accounts.id.disableAutoSelect(); } } catch {}
  localStorage.removeItem('apiKey');
  localStorage.removeItem('userEmail');
  isLoggedIn.value = false;
  userEmail.value = '';
  showMoreMenu.value = false;
  alert(t.value.logoutSuccess);
};

onMounted(async () => {
  await detectLanguage();
  const apiKey = localStorage.getItem('apiKey');
  const email = localStorage.getItem('userEmail');
  if (apiKey && email) { isLoggedIn.value = true; userEmail.value = email; }
  initGoogleAuth();
});
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Libre+Baskerville:wght@400;700&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600&display=swap');
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.1/css/all.min.css');

* { margin: 0; padding: 0; box-sizing: border-box; }
.app { width: 100vw; height: 100vh; display: flex; background: linear-gradient(135deg, #1a1a1d 0%, #2d1f2f 50%, #1a1a1d 100%); font-family: 'Inter', sans-serif; overflow: hidden; position: relative; }
.menu-toggle-fixed { position: fixed; top: 28px; left: 28px; background: rgba(42,42,42,0.8); backdrop-filter: blur(10px); border: none; color: #ffffff; font-size: 16px; padding: 10px 12px; cursor: pointer; transition: all 0.4s ease; border-radius: 10px; z-index: 1000; display: flex; align-items: center; justify-content: center; width: 40px; height: 40px; }
.menu-toggle-fixed.open { left: 248px; }
.menu-toggle-fixed:hover { background: rgba(255,255,255,0.15); }
.overlay { position: fixed; top: 0; left: 0; width: 100vw; height: 100vh; background: rgba(0,0,0,0.5); backdrop-filter: blur(2px); opacity: 0; pointer-events: none; transition: opacity 0.3s ease; z-index: 998; }
.overlay.visible { opacity: 1; pointer-events: all; }
.sidebar { width: 0; height: 100vh; background: rgba(42,42,42,0.7); backdrop-filter: blur(20px); display: flex; flex-direction: column; padding: 0; position: fixed; left: -264px; top: 0; transition: all 0.4s ease; overflow: hidden; opacity: 0; border-radius: 0 24px 24px 0; z-index: 999; }
.sidebar.open { left: 0; width: 264px; padding: 28px 0; opacity: 1; overflow: visible; }
.sidebar-footer { margin-top: auto; padding: 16px 16px 0; opacity: 0; transition: opacity 0.3s ease 0.2s; }
.sidebar.open .sidebar-footer { opacity: 1; }
.download-button { background: rgba(212,175,55,0.1); border: 1px solid rgba(212,175,55,0.3); color: #d4af37; opacity: 1 !important; }
.download-button:hover { background: rgba(212,175,55,0.2); border-color: rgba(212,175,55,0.5); color: #e8c84a; }
.download-button i { color: #d4af37; }
.logo { font-family: 'Libre Baskerville', serif; font-size: 28px; font-weight: 700; color: #ffffff; padding: 0 28px; margin-bottom: 32px; white-space: nowrap; opacity: 0; transition: opacity 0.3s ease 0.2s; }
.sidebar.open .logo { opacity: 1; }
.nav { display: flex; flex-direction: column; gap: 4px; padding: 0 16px; position: relative; overflow: visible; }
.nav-item { display: flex; align-items: center; gap: 16px; padding: 14px 12px; color: #e0e0e0; text-decoration: none; font-size: 16px; font-weight: 400; transition: all 0.3s ease; border-radius: 12px; white-space: nowrap; opacity: 0; cursor: pointer; }
.sidebar.open .nav-item { opacity: 1; transition: all 0.3s ease 0.2s; }
.nav-item:hover { background: rgba(255,255,255,0.05); text-decoration: none; }
.nav-item.active { background: rgba(212,175,55,0.1); color: #d4af37; }
.nav-item i { font-size: 18px; width: 20px; text-align: center; flex-shrink: 0; }
.nav-text { text-decoration: none; }
.more-container { position: relative; }
.more-menu { position: absolute; bottom: 100%; left: 0; right: 0; margin-bottom: 8px; background: rgba(30,30,30,0.98); backdrop-filter: blur(20px); border: 1px solid rgba(255,255,255,0.2); border-radius: 12px; padding: 8px; z-index: 100; min-width: 220px; box-shadow: 0 4px 20px rgba(0,0,0,0.5); animation: slideUp 0.2s ease; }
@keyframes slideUp { from { opacity: 0; transform: translateY(10px); } to { opacity: 1; transform: translateY(0); } }
.user-info { display: flex; flex-direction: column; gap: 8px; }
.user-email { padding: 12px; color: rgba(255,255,255,0.7); font-size: 14px; border-bottom: 1px solid rgba(255,255,255,0.1); word-break: break-word; }
.more-menu-item { width: 100%; display: flex; align-items: center; gap: 12px; padding: 12px; background: transparent; border: none; color: rgba(255,255,255,0.85); font-size: 15px; font-family: 'Inter', sans-serif; cursor: pointer; border-radius: 8px; transition: all 0.3s ease; text-align: left; }
.more-menu-item:hover { background: rgba(255,255,255,0.05); }
.more-menu-item.google-button { background: rgba(66,133,244,0.15); border: 1px solid rgba(66,133,244,0.3); }
.more-menu-item.google-button:hover { background: rgba(66,133,244,0.25); border-color: rgba(66,133,244,0.4); }
.more-menu-item i { font-size: 16px; width: 20px; text-align: center; }
.main { flex: 1; display: flex; flex-direction: column; align-items: center; padding: 100px 32px 60px; position: relative; overflow-x: hidden; overflow-y: auto; width: 100%; }
.blur-red { position: fixed; top: 15%; left: 15%; width: 300px; height: 300px; background: radial-gradient(circle, rgba(139,0,0,0.08) 0%, transparent 70%); filter: blur(30px); pointer-events: none; z-index: 0; }
.blur-yellow { position: fixed; top: 15%; right: 20%; width: 300px; height: 300px; background: radial-gradient(circle, rgba(212,175,55,0.06) 0%, transparent 70%); filter: blur(30px); pointer-events: none; z-index: 0; }
.content-wrapper { width: 100%; max-width: 800px; position: relative; z-index: 1; }
.main-title { font-family: 'Libre Baskerville', serif; font-size: 2.5rem; font-weight: 700; color: #d4af37; margin-bottom: 1.5rem; text-align: center; }
.subtitle { color: rgba(255,255,255,0.7); font-size: 0.95rem; margin-bottom: 2rem; text-align: center; line-height: 1.6; }
.section-title { font-size: 1.75rem; font-weight: 700; color: #d4af37; margin-bottom: 1rem; margin-top: 2rem; }
.subsection-title { font-size: 1.25rem; font-weight: 600; color: rgba(255,255,255,0.9); margin-bottom: 0.75rem; margin-top: 1.5rem; }
.section-text { color: rgba(255,255,255,0.7); font-size: 0.95rem; margin-bottom: 1em; line-height: 1.7; }
.bullet-list { list-style-type: disc; margin-left: 1.5rem; margin-bottom: 1em; color: rgba(255,255,255,0.7); font-size: 0.95rem; }
.bullet-list li { margin-bottom: 0.75em; line-height: 1.7; }
.bullet-list strong { color: rgba(255,255,255,0.9); }
.section-separator { border-top: 1px solid rgba(255,255,255,0.1); margin-top: 3rem; margin-bottom: 2rem; }
.license-block { background: rgba(42,42,42,0.5); backdrop-filter: blur(10px); color: rgba(255,255,255,0.7); padding: 1.5rem; border-radius: 12px; overflow: auto; font-size: 0.8rem; white-space: pre-wrap; line-height: 1.6; border: 1px solid rgba(255,255,255,0.1); font-family: 'Courier New', monospace; margin: 1.5rem 0; }
.tech-credits { margin-top: 3rem; text-align: center; padding: 2rem 0; }
.tech-logos { display: flex; justify-content: center; align-items: center; gap: 3rem; flex-wrap: wrap; margin-top: 2rem; }
.tech-logo-link { transition: transform 0.3s ease, opacity 0.3s ease; display: inline-block; text-decoration: none; }
.tech-logo-link:hover { transform: scale(1.1); opacity: 0.8; }
.tech-logo { height: 60px; width: auto; filter: brightness(0.9); }

@media (max-width: 768px) {
  .menu-toggle-fixed { top: 20px; left: 20px; width: 44px; height: 44px; font-size: 18px; }
  .menu-toggle-fixed.open { left: 20px; }
  .sidebar { border-radius: 0; }
  .sidebar.open { width: 80%; max-width: 300px; }
  .main { padding: 80px 24px 40px; }
  .main-title { font-size: 1.75rem; }
  .section-title { font-size: 1.5rem; }
  .subsection-title { font-size: 1.1rem; }
  .tech-logos { gap: 2rem; }
  .tech-logo { height: 50px; }
  .blur-red, .blur-yellow { width: 200px; height: 200px; }
}
@media (max-width: 480px) {
  .menu-toggle-fixed { top: 16px; left: 16px; }
  .main { padding: 70px 20px 30px; }
  .main-title { font-size: 1.5rem; }
  .section-title { font-size: 1.25rem; }
  .subsection-title { font-size: 1rem; }
  .section-text, .bullet-list { font-size: 0.875rem; }
  .license-block { font-size: 0.7rem; padding: 1rem; }
  .tech-logos { gap: 1.5rem; }
  .tech-logo { height: 40px; }
  .sidebar.open { width: 85%; }
  .logo { font-size: 24px; padding: 0 20px; }
  .nav { padding: 0 12px; }
  .nav-item { padding: 12px 10px; font-size: 15px; }
  .blur-red, .blur-yellow { width: 150px; height: 150px; filter: blur(20px); }
}
@media (hover: none) and (pointer: coarse) {
  .menu-toggle-fixed, .nav-item { -webkit-tap-highlight-color: transparent; }
  .nav-item:active { background: rgba(255,255,255,0.1); }
  .menu-toggle-fixed:active { background: rgba(255,255,255,0.2); }
}
</style>