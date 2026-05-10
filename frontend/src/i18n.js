const translations = {
  tr: {
    nav: {
      home: 'Ana Sayfa',
      settings: 'Ayarlar',
      logo: 'Arama'
    },
    home: {
      title: 'Arama',
      placeholder: 'Ara...',
      searchBtn: 'Ara',
      resultsFound: 'sonuç bulundu',
      blockedHidden: 'engelli domain gizlendi',
      noResults: 'Sonuç bulunamadı',
      addFav: 'Favorilere Ekle',
      removeFav: 'Favorilerden Çıkar',
      blockDomain: 'Domaini Engelle',
      addedFav: 'Favorilere eklendi',
      removedFav: 'Favorilerden çıkarıldı',
      blocked: 'engellendi',
      bangPrefix: 'Bang araması:',
      bangHint: '!g arama → Google, !yt arama → YouTube, !w arama → Wikipedia, !r arama → Reddit'
    },
    settings: {
      title: 'Ayarlar',
      subtitle: 'Tercihlerini özelleştir',
      appearance: 'Görünüm',
      blockedSites: 'Engellenen Domainler',
      favorites: 'Favoriler',
      theme: 'Tema',
      themeDesc: 'Karanlık veya aydınlık tema seç',
      dark: 'Karanlık',
      light: 'Aydınlık',
      fontSize: 'Yazı Boyutu',
      fontSizeDesc: 'Yazı boyutunu büyük veya küçük yap',
      normal: 'Normal',
      large: 'Büyük',
      language: 'Dil',
      languageDesc: 'Arayüz dilini değiştir',
      turkish: 'Türkçe',
      english: 'English',
      autoDetect: 'Otomatik Algıla',
      blockedDomains: 'Engellenen Domainler',
      blockedDesc: 'Engellemek istediğin domainleri ekle. Bu domainleri içeren sonuçlar gizlenir.',
      blockPlaceholder: 'ornekdomain.com',
      blockBtn: 'Engelle',
      blockedList: 'Engellenen Listesi',
      noBlocked: 'Henüz engellenen domain yok',
      favSites: 'Favori Domainler',
      favDesc: 'Arama sonuçlarından eklediğin favoriler burada listelenir.',
      noFavs: 'Henüz favori eklenmedi'
    }
  },
  en: {
    nav: {
      home: 'Home',
      settings: 'Settings',
      logo: 'Search'
    },
    home: {
      title: 'Search',
      placeholder: 'Search...',
      searchBtn: 'Search',
      resultsFound: 'results found',
      blockedHidden: 'blocked domains hidden',
      noResults: 'No results found',
      addFav: 'Add to Favorites',
      removeFav: 'Remove from Favorites',
      blockDomain: 'Block Domain',
      addedFav: 'Added to favorites',
      removedFav: 'Removed from favorites',
      blocked: 'blocked',
      bangPrefix: 'Bang search:',
      bangHint: '!g query → Google, !yt query → YouTube, !w query → Wikipedia, !r query → Reddit'
    },
    settings: {
      title: 'Settings',
      subtitle: 'Customize your preferences',
      appearance: 'Appearance',
      blockedSites: 'Blocked Domains',
      favorites: 'Favorites',
      theme: 'Theme',
      themeDesc: 'Choose dark or light theme',
      dark: 'Dark',
      light: 'Light',
      fontSize: 'Font Size',
      fontSizeDesc: 'Make text larger or smaller',
      normal: 'Normal',
      large: 'Large',
      language: 'Language',
      languageDesc: 'Change interface language',
      turkish: 'Turkish',
      english: 'English',
      autoDetect: 'Auto Detect',
      blockedDomains: 'Blocked Domains',
      blockedDesc: 'Add domains to block. Results containing these domains will be hidden.',
      blockPlaceholder: 'example.com',
      blockBtn: 'Block',
      blockedList: 'Blocked List',
      noBlocked: 'No blocked domains yet',
      favSites: 'Favorite Domains',
      favDesc: 'Your favorites from search results are listed here.',
      noFavs: 'No favorites added yet'
    }
  }
}

const bangs = {
  g: 'https://www.google.com/search?q=',
  google: 'https://www.google.com/search?q=',
  yt: 'https://www.youtube.com/results?search_query=',
  youtube: 'https://www.youtube.com/results?search_query=',
  w: 'https://en.wikipedia.org/wiki/Special:Search?search=',
  wiki: 'https://en.wikipedia.org/wiki/Special:Search?search=',
  wikipedia: 'https://en.wikipedia.org/wiki/Special:Search?search=',
  r: 'https://www.reddit.com/search/?q=',
  reddit: 'https://www.reddit.com/search/?q=',
  gh: 'https://github.com/search?q=',
  github: 'https://github.com/search?q=',
  so: 'https://stackoverflow.com/search?q=',
  stackoverflow: 'https://stackoverflow.com/search?q=',
  ddg: 'https://duckduckgo.com/?q=',
  duckduckgo: 'https://duckduckgo.com/?q=',
  tw: 'https://twitter.com/search?q=',
  twitter: 'https://twitter.com/search?q=',
  x: 'https://twitter.com/search?q=',
  mdn: 'https://developer.mozilla.org/en-US/search?q=',
  npm: 'https://www.npmjs.com/search?q=',
  py: 'https://pypi.org/search/?q=',
  pypi: 'https://pypi.org/search/?q='
}

export function getTranslation(lang, section, key) {
  const l = translations[lang] || translations.tr
  if (l[section] && l[section][key] !== undefined) {
    return l[section][key]
  }
  return key
}

export function t(section, key) {
  const lang = localStorage.getItem('lang') || 'tr'
  return getTranslation(lang, section, key)
}

export function resolveBang(query) {
  const trimmed = query.trim()
  if (!trimmed.startsWith('!')) return null
  const parts = trimmed.slice(1).split(/\s+(.+)/)
  if (parts.length < 2 || !parts[0] || !parts[1]) return null
  const prefix = parts[0].toLowerCase()
  const searchQuery = parts[1]
  const url = bangs[prefix]
  if (url) {
    return url + encodeURIComponent(searchQuery)
  }
  return null
}

export function detectBrowserLang() {
  const nav = navigator.language || navigator.userLanguage || 'tr'
  return nav.toLowerCase().startsWith('en') ? 'en' : 'tr'
}

export default translations
