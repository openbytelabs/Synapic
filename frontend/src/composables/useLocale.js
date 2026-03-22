import { ref, reactive } from 'vue'

const lang = ref('tr')
const detected = ref(false)

export const tr = {
  nav: {
    home: 'Anasayfa',
    api: 'API',
    terms: 'Güvenlik & Kurallar',
    settings: 'Ayarlar',
    more: 'Daha Fazla',
    download: 'Uygulamayı İndir',
    loginGoogle: 'Google ile Giriş Yap',
    logout: 'Çıkış Yap',
  },
  home: {
    placeholder: 'Aklından ne geçiyor...',
    hint: 'Aramak için',
    hintKey: 'Enter',
    hintSuffix: 'Tuşuna bas.',
  },
  search: {
    web: 'Web',
    images: 'Resimler',
    news: 'Haberler',
    maps: 'Haritalar',
    noResults: 'Sonuç bulunamadı. Farklı bir şey aramayı deneyin.',
    aiBadge: 'AI Destekli Cevap',
    aiFooter: 'Cevap arama sonuçlarından üretildi',
    aiLoading: 'Arama sonuçları analiz ediliyor...',
    copyUrl: 'URL Kopyala',
    maintenance: 'Bakımda',
    maintenanceDesc: 'Resim arama özelliği şu an bakım modunda. En kısa sürede geri döneceğiz.',
  },
  api: {
    loginTitle: "API'ı Kullanmak İçin Giriş Yapın",
    loginDesc: 'API özelliklerine erişmek için Google hesabınızla giriş yapmanız gerekmektedir.',
    loginBtn: 'Google ile giriş yap',
    dashTitle: 'API Kontrol Paneli',
    usedToday: 'Bugün Kullanılan',
    dailyLimit: 'Günlük Limit',
    remaining: 'Kalan',
    unlimited: 'Sınırsız',
    apiKey: 'API Anahtarı',
    copyKey: 'Kopyala',
    regenKey: 'Yenile',
    copied: 'Kopyalandı!',
  },
  settings: {
    title: 'Ayarlar',
    dataManagement: 'Bilgi Yönetimi',
    downloadData: 'Bilgimi İndir',
    downloadDesc: 'Bütün hesap bilgilerinizi JSON formatında indirebilirsiniz.',
  },
  terms: {
    title: 'Gizlilik Politikası ve Hizmet Şartları',
  },
  download: {
    title: "Synapic'i İndir",
    subtitle: "Synapic'i cihazına indirerek daha hızlı ve gizlilik odaklı arama deneyiminin keyfini çıkar.",
    macDesc: 'macOS 12 Monterey ve üzeri için',
    androidDesc: 'Android 9.0 ve üzeri için',
    apkBadge: 'APK olarak indirilebilir',
    macBtn: 'macOS için İndir',
    androidBtn: 'Android için İndir',
    privacy: 'Gizlilik öncelikli',
    fast: 'Hızlı & hafif',
    noTrack: 'Veri takibi yok',
  },
  loginSuccess: 'Giriş başarılı! API paneline yönlendiriliyorsunuz...',
  loginFail: 'Giriş başarısız. Lütfen tekrar deneyin.',
  loginFailMsg: 'Giriş başarısız: ',
  logoutSuccess: 'Başarıyla çıkış yapıldı',
  authLoading: 'Google kimlik doğrulaması yükleniyor. Lütfen birkaç saniye bekleyip tekrar deneyin.',
  authFail: 'Google kimlik doğrulaması başlatılamadı. Lütfen sayfayı yenileyip tekrar deneyin.',
}

export const en = {
  nav: {
    home: 'Home',
    api: 'API',
    terms: 'Privacy & Terms',
    settings: 'Settings',
    more: 'More',
    download: 'Download App',
    loginGoogle: 'Sign in with Google',
    logout: 'Sign Out',
  },
  home: {
    placeholder: "What's on your mind...",
    hint: 'Press',
    hintKey: 'Enter',
    hintSuffix: 'to search.',
  },
  search: {
    web: 'Web',
    images: 'Images',
    news: 'News',
    maps: 'Maps',
    noResults: 'No results found. Try searching for something else.',
    aiBadge: 'AI-Powered Answer',
    aiFooter: 'Answer generated from search results',
    aiLoading: 'Analyzing search results...',
    copyUrl: 'Copy URL',
    maintenance: 'Under Maintenance',
    maintenanceDesc: 'Image search is currently under maintenance. We will be back shortly.',
  },
  api: {
    loginTitle: 'Sign In to Use the API',
    loginDesc: 'You need to sign in with your Google account to access API features.',
    loginBtn: 'Sign in with Google',
    dashTitle: 'API Dashboard',
    usedToday: 'Used Today',
    dailyLimit: 'Daily Limit',
    remaining: 'Remaining',
    unlimited: 'Unlimited',
    apiKey: 'API Key',
    copyKey: 'Copy',
    regenKey: 'Regenerate',
    copied: 'Copied!',
  },
  settings: {
    title: 'Settings',
    dataManagement: 'Data Management',
    downloadData: 'Download My Data',
    downloadDesc: 'You can download all your account information in JSON format.',
  },
  terms: {
    title: 'Privacy Policy & Terms of Service',
  },
  download: {
    title: 'Download Synapic',
    subtitle: 'Download Synapic and enjoy a faster, privacy-focused search experience.',
    macDesc: 'For macOS 12 Monterey and later',
    androidDesc: 'For Android 9.0 and later',
    apkBadge: 'Available as APK',
    macBtn: 'Download for macOS',
    androidBtn: 'Download for Android',
    privacy: 'Privacy-first',
    fast: 'Fast & lightweight',
    noTrack: 'No data tracking',
  },
  loginSuccess: 'Login successful! Redirecting to API panel...',
  loginFail: 'Login failed. Please try again.',
  loginFailMsg: 'Login failed: ',
  logoutSuccess: 'Successfully signed out',
  authLoading: 'Google authentication is loading. Please wait a moment and try again.',
  authFail: 'Google authentication could not be initialized. Please refresh the page and try again.',
}

export async function detectLanguage() {
  if (detected.value) return lang.value
  try {
    const res = await fetch('https://api.synapic.com.tr/detect-lang', { cache: 'no-store' })
    const data = await res.json()
    lang.value = data.lang === 'tr' ? 'tr' : 'en'
  } catch {
    lang.value = 'tr'
  }
  detected.value = true
  return lang.value
}

export function useLocale() {
  const t = lang.value === 'tr' ? tr : en
  return { lang, t, detectLanguage }
}