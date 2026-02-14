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
      
      <div class="content-wrapper">
        <h1 class="main-title">Gizlilik Politikası ve Hizmet Şartları</h1>
        <p class="subtitle">
          Synapic, kullanıcı gizliliğine öncelik veren ve kendi web dizinini koruyan bağımsız bir arama motoru altyapısıdır. Meta arama motoru mimarilerinin aksine, Synapic üçüncü taraf arama sağlayıcılarından sonuçları bir araya getirmez. Dizine eklenen tüm içerik, Synapic'in kendi tarama ve dizinleme altyapısı aracılığıyla toplanır, işlenir ve sunulur.        </p>
        <p class="subtitle">
          Bu belge, hizmet ekosistemimiz içinde kullanıcı verilerinin işlenmesi ve korunmasında kullanılan yöntemleri açıklamaktadır.
        </p>
        <div class="section-separator"></div>

        <h2 class="section-title">1. Gizlilik Politikası</h2>
        <p class="section-text">
          Synapic'te kullanıcı gizliliği en önemli önceliğimizdir. Kişisel kullanıcı verilerinin üçüncü taraflarla paylaşılmaması konusunda katı bir politika izliyoruz. Kullanıcı tercihleri ​​ve arama ile ilgili yapılandırmalar yalnızca yerel cihazlarda saklanır.
                </p>

                <h3 class="subsection-title">1.1 Bilgi Toplama</h3>
        <p class="section-text">
          Synapic, kullanıcı deneyimini geliştirmek amacıyla aşağıdaki veri kategorileri için yerel depolama mekanizmaları uygular:
        </p>
        <ul class="bullet-list">
          <li>
            <strong>Arama Geçmişi:</strong> Tarayıcı localStorage aracılığıyla yerel olarak tutulur ve kullanıcı tarafından herhangi bir zamanda silinebilir.
          </li>
          <li>
            <strong>Dil Tercihleri:</strong> Tutarlı ve bağlamsal olarak ilgili arama sonuçlarının sunulmasını sağlamak için yerel olarak korunur.
          </li>
          <li>
            <strong>Arayüz Tercihleri:</strong> Yalnızca yerel depolamada saklanan tema ve kullanıcı arayüzü yapılandırmaları.
          </li>
        </ul>
        <p class="section-text">
          Synapic, sunucu altyapısı içerisinde kişisel olarak tanımlanabilir kullanıcı verileri barındırmaz.
        </p>

        <h3 class="subsection-title">1.2 Dizinlenmiş Veri Mimarisi</h3>
        <p class="section-text">
          Synapic, aşağıdaki özelliklere sahip tescilli bir web tarama ve dizinleme sistemi işletir:
        </p>
        <ul class="bullet-list">
          <li>Mevcut veritabanında yaklaşık <strong>10.000.000+ dizinlenmiş URL</strong></li>
          <li>Yalnızca herkese açık web sayfalarının dizinlenmesi</li>
          <li>Gizli veya kısıtlanmış içeriğin sistematik olarak hariç tutulması</li>
          <li>Verilerin yalnızca arama sonucu sağlama amacıyla kullanılması</li>
        </ul>

        <h3 class="subsection-title">1.3 Bilgi Kullanımı</h3>
        <p class="section-text">
          Yerel olarak depolanan veriler yalnızca kullanılabilirliği artırmaya ve kullanıcı tercihlerini korumaya hizmet eder. Altyapımız analitik, reklam, kullanıcı profilleme veya pazarlama mekanizmaları olmadan çalışır.
        </p>

        <h3 class="subsection-title">1.4 Veri Güvenliği Çerçevesi</h3>
        <p class="section-text">
          Kullanıcıya özel tüm veriler yalnızca yerel depolamada bulunur. Tarayıcı verilerinin temizlenmesi kalıcı olarak silinmeyle sonuçlanır. Synapic, yerel verilerin değiştirilmesi veya alınması için hiçbir uzaktan erişim yeteneğine sahip değildir.
        </p>

        <h3 class="subsection-title">1.5 Çerezler ve Takip Politikası</h3>
        <ul class="bullet-list">
          <li>Takip çerezleri uygulaması yoktur</li>
          <li>Parmak izi (fingerprinting) teknikleri kullanılmaz</li>
          <li>Üçüncü taraf takip mekanizmalarının hiçbiri bulunmaz</li>
        </ul>

        <h3 class="subsection-title">1.6 Üçüncü Taraf Bağlantıları</h3>
        <p class="section-text">
          Synapic, arama sonuçları aracılığıyla karşılaşılan üçüncü taraf web sitesi içerikleri veya gizlilik uygulamaları konusunda sorumluluk kabul etmez.
        </p>

        <h3 class="subsection-title">1.7 Çocukların Gizliliğinin Korunması</h3>
        <p class="section-text">
          Altyapımız reşit olmayanlardan bilerek veri toplamaz. 13 yaşın altındaki kullanıcılar için ebeveyn gözetimi önerilir.
        </p>

        <h3 class="subsection-title">1.8 Veri Saklama Protokolü</h3>
        <p class="section-text">
          Yerel olarak depolanan tüm veriler, kullanıcı tarafından manuel olarak temizlenene veya uygulama kaldırılana kadar saklanır.
        </p>

        <h3 class="subsection-title">1.9 Politika Değişiklik Protokolü</h3>
        <p class="section-text">
          Bu politika önceden bildirim yapılmaksızın güncellemelere tabi olabilir.
        </p>

        <h3 class="subsection-title">1.10 İletişim Bilgileri</h3>
        <p class="section-text">
          E-posta: yigitkabak@tuta.io
        </p>

        <div class="section-separator"></div>

        <h2 class="section-title">2. Hizmet Şartları</h2>
        <p class="section-text">
          Synapic hizmetlerinin kullanımı aşağıdaki şartların kabul edildiği anlamına gelir.
        </p>

        <h3 class="subsection-title">2.1 Hizmet Kullanım Parametreleri</h3>
        <p class="section-text">
          Hizmet erişimi yalnızca kişisel ve ticari olmayan amaçlar içindir. Hizmet özelliklerinin kötüye kullanılması, manipüle edilmesi veya izinsiz otomasyonu kesinlikle yasaktır.
        </p>

        <h3 class="subsection-title">2.2 Fikri Mülkiyet Hakları</h3>
        <p class="section-text">
          Tüm Synapic içeriği ve sistem mimarileri, geçerli fikri mülkiyet yasaları kapsamında korunmaktadır. Açık kaynaklı bileşenler <strong>BSD 3-Clause Lisansı</strong> çerçevesinde çalışır.
        </p>

        <h3 class="subsection-title">2.3 Sorumluluk Sınırlamaları</h3>
        <p class="section-text">
          Hizmetler garanti olmaksızın <strong>"olduğu gibi"</strong> esasına göre sağlanır. Synapic, hizmet kullanımından kaynaklanan hasarlar için sorumluluk kabul etmez.
        </p>

        <h3 class="subsection-title">2.4 Hizmet Değişiklik Hakları</h3>
        <p class="section-text">
          Hizmetler önceden bildirim yapılmaksızın değişikliğe veya sonlandırmaya tabi olabilir.
        </p>

        <h3 class="subsection-title">2.5 Yapay Zeka Entegrasyonu</h3>
        <p class="section-text">
          Yapay zeka tarafından oluşturulan içerik yanlışlıklar içerebilir ve kritik karar verme süreçleri için tek dayanak noktası olmamalıdır.
        </p>

        <h3 class="subsection-title">2.6 Arama Sonucu Hükümleri</h3>
        <p class="section-text">
          Sonuçlar Synapic'in tescilli dizininden oluşturulur. Doğruluk garanti edilemez.
        </p>

        <h3 class="subsection-title">2.7 Kullanıcı Davranış Kuralları</h3>
        <p class="section-text">
          Kullanıcılar platform üzerinden yasa dışı faaliyetlerde bulunmaktan, sistemi kötüye kullanmaktan veya kötü amaçlı yazılım dağıtmaktan kaçınmalıdır.
        </p>

        <h3 class="subsection-title">2.8 Yargı Yetkisi</h3>
        <p class="section-text">
          Bu şartlar Türkiye Cumhuriyeti yasalarına tabidir.
        </p>

        <div class="section-separator"></div>

        <h2 class="section-title">3. Açık Kaynak Lisansı</h2>
        <p class="section-text">
          Synapic, BSD 3-Clause Lisans çerçevesi altında açık kaynak bileşenleri içerir.
        </p>
        
        <h3 class="subsection-title">3.1 BSD 3-Clause License</h3>
        
        <pre class="license-block">
BSD 3-Clause License

Copyright (c) 2023, openbytelabs

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
 list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
 this list of conditions and the following disclaimer in the documentation
 and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

THIS SOFTWARE IS PROVIDED "AS IS".
        </pre>

        <div class="section-separator"></div>

        <h2 class="section-title">4. Teknik Altyapı</h2>
        
        <h3 class="subsection-title">4.1 Teknoloji Yığını Mimarisi</h3>
        <ul class="bullet-list">
          <li><strong>Ön Uç (Frontend) Çatısı:</strong> Vue.js 3 (Composition API)</li>
          <li><strong>Arka Uç (Backend) Altyapısı:</strong> Golang</li>
          <li><strong>Web Tarayıcısı (Crawler):</strong> Özel yapım tescilli sistem</li>
          <li><strong>Dizinleme Sistemi:</strong> Tescilli dizinleme mimarisi</li>
          <li><strong>Arama Motoru:</strong> Dahili altyapı uygulaması</li>
        </ul>

        <h3 class="subsection-title">4.2 Veri Depolama Mimarisi</h3>
        <ul class="bullet-list">
          <li><strong>Sunucu Tarafı Depolama:</strong> Genel web dizini (10.000.000+ URL)</li>
          <li><strong>İstemci Tarafı Depolama:</strong> Tarayıcı localStorage uygulaması</li>
          <li>Analitik, reklam veya takip mekanizmalarının hiçbiri bulunmaz</li>
        </ul>

        <div class="section-separator"></div>

        <div class="tech-credits">
        <h3 class="subsection-title">❤️ ile şunlar kullanılarak yapıldı</h3>
        <div class="tech-logos">
          <a href="https://www.linux.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link">
            <img src="https://upload.wikimedia.org/wikipedia/commons/3/35/Tux.svg" alt="Linux" class="tech-logo">
          </a>
          <a href="https://vuejs.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link">
            <img src="https://upload.wikimedia.org/wikipedia/commons/9/95/Vue.js_Logo_2.svg" alt="Vue.js" class="tech-logo">
          </a>
          <a href="https://nodejs.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link">
            <img src="https://upload.wikimedia.org/wikipedia/commons/d/d9/Node.js_logo.svg" alt="Node.js" class="tech-logo">
          </a>
          <a href="https://golang.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link">
            <img src="https://duckduckgo.com/i/77cac52781feeb4a.png" alt="Golang" class="tech-logo">
          </a>
        </div>
      </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const menuOpen = ref(false);
const showMoreMenu = ref(false);
const isLoggedIn = ref(false);
const userEmail = ref('');
const googleAuthReady = ref(false);

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
          
          const backendResponse = await fetch('http://localhost:3000/auth/google', {
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
  text-decoration: none;
}

.nav-item.active {
  background: rgba(212, 175, 55, 0.1);
  color: #d4af37;
}

.nav-item i {
  font-size: 18px;
  width: 20px;
  text-align: center;
  flex-shrink: 0;
}

.nav-text {
  text-decoration: none;
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
  padding: 100px 32px 60px;
  position: relative;
  overflow-x: hidden;
  overflow-y: auto;
  width: 100%;
}

.blur-red {
  position: fixed;
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
  position: fixed;
  top: 15%;
  right: 20%;
  width: 300px;
  height: 300px;
  background: radial-gradient(circle, rgba(212, 175, 55, 0.06) 0%, transparent 70%);
  filter: blur(30px);
  pointer-events: none;
  z-index: 0;
}

.content-wrapper {
  width: 100%;
  max-width: 800px;
  position: relative;
  z-index: 1;
}

.main-title {
  font-family: 'Libre Baskerville', serif;
  font-size: 2.5rem;
  font-weight: 700;
  color: #d4af37;
  margin-bottom: 1.5rem;
  text-align: center;
}

.subtitle {
  color: rgba(255, 255, 255, 0.7);
  font-size: 0.95rem;
  margin-bottom: 2rem;
  text-align: center;
  line-height: 1.6;
}

.section-title {
  font-size: 1.75rem;
  font-weight: 700;
  color: #d4af37;
  margin-bottom: 1rem;
  margin-top: 2rem;
}

.subsection-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: rgba(255, 255, 255, 0.9);
  margin-bottom: 0.75rem;
  margin-top: 1.5rem;
}

.section-text {
  color: rgba(255, 255, 255, 0.7);
  font-size: 0.95rem;
  margin-bottom: 1em;
  line-height: 1.7;
}

.bullet-list {
  list-style-type: disc;
  margin-left: 1.5rem;
  margin-bottom: 1em;
  color: rgba(255, 255, 255, 0.7);
  font-size: 0.95rem;
}

.bullet-list li {
  margin-bottom: 0.75em;
  line-height: 1.7;
}

.bullet-list strong {
  color: rgba(255, 255, 255, 0.9);
}

a.tech-logo-link {
  color: #d4af37;
  text-decoration: none;
  transition: all 0.3s ease;
}

a.tech-logo-link:hover {
  color: #ffffff;
  text-decoration: none;
}

.section-separator {
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  margin-top: 3rem;
  margin-bottom: 2rem;
}

.license-block {
  background: rgba(42, 42, 42, 0.5);
  backdrop-filter: blur(10px);
  color: rgba(255, 255, 255, 0.7);
  padding: 1.5rem;
  border-radius: 12px;
  overflow: auto;
  font-size: 0.8rem;
  white-space: pre-wrap;
  line-height: 1.6;
  border: 1px solid rgba(255, 255, 255, 0.1);
  font-family: 'Courier New', monospace;
  margin: 1.5rem 0;
}

.tech-credits {
  margin-top: 3rem;
  text-align: center;
  padding: 2rem 0;
}

.tech-logos {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 3rem;
  flex-wrap: wrap;
  margin-top: 2rem;
}

.tech-logo-link {
  transition: transform 0.3s ease, opacity 0.3s ease;
  display: inline-block;
  text-decoration: none;
}

.tech-logo-link:hover {
  transform: scale(1.1);
  opacity: 0.8;
  text-decoration: none;
}

.tech-logo {
  height: 60px;
  width: auto;
  filter: brightness(0.9);
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

  .main {
    padding: 80px 24px 40px;
  }

  .main-title {
    font-size: 1.75rem;
  }

  .section-title {
    font-size: 1.5rem;
  }

  .subsection-title {
    font-size: 1.1rem;
  }

  .tech-logos {
    gap: 2rem;
  }

  .tech-logo {
    height: 50px;
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

  .main {
    padding: 70px 20px 30px;
  }

  .main-title {
    font-size: 1.5rem;
  }

  .section-title {
    font-size: 1.25rem;
  }

  .subsection-title {
    font-size: 1rem;
  }

  .section-text,
  .bullet-list {
    font-size: 0.875rem;
  }

  .license-block {
    font-size: 0.7rem;
    padding: 1rem;
  }

  .tech-logos {
    gap: 1.5rem;
  }

  .tech-logo {
    height: 40px;
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
