<template>
  <div class="settings-page">
    <header class="topbar">
      <div class="topbar-inner">
        <router-link to="/" class="logo-link">
          <img src="../images/synapic.png" alt="Synapic" class="topbar-logo" />
          <span class="logo-text">Synapic</span>
        </router-link>
        <h1 class="page-title">{{ t.settings.title }}</h1>
      </div>
    </header>

    <nav class="tab-nav">
      <div class="tab-nav-inner">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          class="tab-btn"
          :class="{ active: activeTab === tab.id }"
          @click="activeTab = tab.id"
        >
          <component :is="tab.icon" />
          <span>{{ tab.label }}</span>
        </button>
      </div>
    </nav>

    <main class="content">
      <div class="content-inner">
        <section v-if="activeTab === 'appearance'" class="panel">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="13.5" cy="6.5" r="0.5"/><circle cx="17.5" cy="10.5" r="0.5"/><circle cx="8.5" cy="7.5" r="0.5"/><circle cx="6.5" cy="12.5" r="0.5"/><path d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10c.926 0 1.648-.746 1.648-1.688 0-.437-.18-.835-.437-1.125-.29-.289-.438-.652-.438-1.125a1.64 1.64 0 0 1 1.668-1.668h1.996c3.051 0 5.555-2.503 5.555-5.554C21.965 6.012 17.461 2 12 2z"/></svg>
              {{ t.settings.appearance }}
            </h2>
            <div class="setting-row">
              <label class="setting-label">{{ t.settings.fontSize }}</label>
              <div class="btn-group">
                <button
                  class="opt-btn"
                  :class="{ active: fontSize === '14px' }"
                  @click="setFontSize('14px')"
                >{{ t.settings.small }} <span class="opt-hint">14px</span></button>
                <button
                  class="opt-btn"
                  :class="{ active: fontSize === '16px' }"
                  @click="setFontSize('16px')"
                >{{ t.settings.medium }} <span class="opt-hint">16px</span></button>
                <button
                  class="opt-btn"
                  :class="{ active: fontSize === '18px' }"
                  @click="setFontSize('18px')"
                >{{ t.settings.large }} <span class="opt-hint">18px</span></button>
              </div>
            </div>
            <div class="setting-row">
              <label class="setting-label">{{ t.settings.theme }}</label>
              <div class="btn-group">
                <button
                  class="opt-btn"
                  :class="{ active: theme === 'dark' }"
                  @click="setTheme('dark')"
                >
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"/></svg>
                  {{ t.settings.dark }}
                </button>
                <button
                  class="opt-btn"
                  :class="{ active: theme === 'light' }"
                  @click="setTheme('light')"
                >
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="5"/><line x1="12" y1="1" x2="12" y2="3"/><line x1="12" y1="21" x2="12" y2="23"/><line x1="4.22" y1="4.22" x2="5.64" y2="5.64"/><line x1="18.36" y1="18.36" x2="19.78" y2="19.78"/><line x1="1" y1="12" x2="3" y2="12"/><line x1="21" y1="12" x2="23" y2="12"/><line x1="4.22" y1="19.78" x2="5.64" y2="18.36"/><line x1="18.36" y1="5.64" x2="19.78" y2="4.22"/></svg>
                  {{ t.settings.light }}
                </button>
              </div>
            </div>
          </div>
        </section>

        <section v-else-if="activeTab === 'bangs'" class="panel">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"/></svg>
              {{ t.settings.bangs }}
            </h2>
            <p class="card-desc">{{ t.settings.bangDesc }}</p>

            <div class="sub-section">
              <h3 class="sub-title">{{ t.settings.defaultBangs }}</h3>
              <div class="item-list" v-if="activeDefaultBangs.length">
                <div class="item-row" v-for="b in activeDefaultBangs" :key="b.shortcut">
                  <div class="item-info">
                    <code class="item-code">!{{ b.shortcut }}</code>
                    <span class="item-text">{{ b.site }}</span>
                  </div>
                  <button class="icon-btn danger" @click="removeBang(b.shortcut)" :title="t.settings.delete">
                    <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"/><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/></svg>
                  </button>
                </div>
              </div>
              <div class="empty-state" v-else>{{ t.settings.noBlockedSites }}</div>
            </div>

            <div class="sub-section" v-if="deletedBangs.length">
              <h3 class="sub-title sub-restore">{{ t.settings.bangActions }}</h3>
              <div class="item-list">
                <div class="item-row item-dim" v-for="s in deletedBangs" :key="'del-' + s">
                  <div class="item-info">
                    <code class="item-code">!{{ s }}</code>
                  </div>
                  <button class="icon-btn success" @click="restoreBang(s)" :title="t.settings.save">
                    <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="1 4 1 10 7 10"/><path d="M3.51 15a9 9 0 1 0 2.13-9.36L1 10"/></svg>
                  </button>
                </div>
              </div>
            </div>

            <div class="sub-section">
              <h3 class="sub-title">{{ t.settings.customBangs }}</h3>
              <div class="item-list" v-if="customBangs.length">
                <div class="item-row" v-for="b in customBangs" :key="'custom-' + b.shortcut">
                  <div class="item-info">
                    <code class="item-code">!{{ b.shortcut }}</code>
                    <span class="item-text">{{ b.site || b.url }}</span>
                  </div>
                  <button class="icon-btn danger" @click="removeBang(b.shortcut)" :title="t.settings.delete">
                    <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"/><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/></svg>
                  </button>
                </div>
              </div>
              <div class="empty-state" v-else>{{ t.settings.noBlockedSites }}</div>
            </div>

            <div class="add-form">
              <h3 class="sub-title">{{ t.settings.addBang }}</h3>
              <div class="form-row">
                <input
                  v-model="newBangShortcut"
                  type="text"
                  class="form-input"
                  :placeholder="t.settings.bangShortcut"
                  @keyup.enter="handleAddBang"
                />
                <input
                  v-model="newBangUrl"
                  type="text"
                  class="form-input wide"
                  :placeholder="'https://example.com/search?q={query}'"
                  @keyup.enter="handleAddBang"
                />
                <button class="add-btn" @click="handleAddBang">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/></svg>
                </button>
              </div>
              <p class="form-msg" v-if="bangError" :class="{ error: bangError }">{{ bangError }}</p>
            </div>
          </div>
        </section>

        <section v-else-if="activeTab === 'blocked'" class="panel">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"/></svg>
              {{ t.settings.blockedSites }}
            </h2>

            <div class="item-list" v-if="blockedSites.length">
              <div class="item-row site-row" v-for="(site, idx) in blockedSites" :key="idx">
                <div class="item-info">
                  <span class="site-url-text">{{ site.url }}</span>
                  <span class="type-badge" :class="'type-' + site.type">{{ site.type }}</span>
                </div>
                <div class="site-actions">
                  <button
                    v-if="site.type !== 'block'"
                    class="icon-btn danger"
                    @click="changeSiteType(idx, 'block')"
                    title="Block"
                  >
                    <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"/></svg>
                  </button>
                  <button
                    v-if="site.type !== 'restrict'"
                    class="icon-btn warn"
                    @click="changeSiteType(idx, 'restrict')"
                    title="Restrict"
                  >
                    <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"/></svg>
                  </button>
                  <button
                    class="icon-btn success"
                    @click="promoteSite(site.url)"
                    :title="t.settings.promoteSite"
                  >
                    <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="19" x2="12" y2="5"/><polyline points="5 12 12 5 19 12"/></svg>
                  </button>
                  <button
                    class="icon-btn danger"
                    @click="removeBlockedSite(site.url)"
                    :title="t.settings.delete"
                  >
                    <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"/><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/></svg>
                  </button>
                </div>
              </div>
            </div>
            <div class="empty-state" v-else>{{ t.settings.noBlockedSites }}</div>

            <div class="add-form">
              <h3 class="sub-title">{{ t.settings.addSite }}</h3>
              <div class="form-row">
                <input
                  v-model="newSiteUrl"
                  type="text"
                  class="form-input"
                  :placeholder="t.settings.siteUrl"
                  @keyup.enter="handleAddSite('block')"
                />
                <button class="type-btn block-type" @click="handleAddSite('block')">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"/></svg>
                  {{ t.settings.blockSite }}
                </button>
                <button class="type-btn restrict-type" @click="handleAddSite('restrict')">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"/></svg>
                  {{ t.settings.restrictSite }}
                </button>
                <button class="type-btn promote-type" @click="handleAddSite('promote')">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="19" x2="12" y2="5"/><polyline points="5 12 12 5 19 12"/></svg>
                  {{ t.settings.promoteSite }}
                </button>
              </div>
            </div>
          </div>
        </section>

        <section v-else-if="activeTab === 'language'" class="panel">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="2" y1="12" x2="22" y2="12"/><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"/></svg>
              {{ t.settings.language }}
            </h2>
            <div class="btn-group lang-group">
              <button
                class="opt-btn lang-opt"
                :class="{ active: lang === 'tr' }"
                @click="setLanguageSetting('tr')"
              >{{ t.settings.turkish }}</button>
              <button
                class="opt-btn lang-opt"
                :class="{ active: lang === 'en' }"
                @click="setLanguageSetting('en')"
              >{{ t.settings.english }}</button>
            </div>
          </div>
        </section>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, h, onMounted } from 'vue'
import { lang, detectLanguage, setLanguage, tr, en } from '../composables/useLocale.js'
import { useBangs } from '../composables/useBangs.js'
import { useBlockedSites } from '../composables/useBlockedSites.js'
import { applyTheme, applyFontSize } from '../composables/settings.js'

const { bangs, customBangs, deletedBangs, addBang, removeBang, restoreBang } = useBangs()
const { blockedSites, addBlockedSite, removeBlockedSite, promoteSite } = useBlockedSites()

const activeTab = ref('appearance')
const fontSize = ref('16px')
const theme = ref('dark')

const newBangShortcut = ref('')
const newBangUrl = ref('')
const bangError = ref('')

const newSiteUrl = ref('')

const t = computed(() => lang.value === 'tr' ? tr : en)

const DEFAULT_BANGS_LIST = [
  { shortcut: 'g', url: 'https://www.google.com/search?q={query}', site: 'Google' },
  { shortcut: 'yt', url: 'https://www.youtube.com/results?search_query={query}', site: 'YouTube' },
  { shortcut: 'w', url: 'https://en.wikipedia.org/wiki/Special:Search?search={query}', site: 'Wikipedia' },
  { shortcut: 'tw', url: 'https://twitter.com/search?q={query}', site: 'Twitter/X' },
  { shortcut: 'r', url: 'https://www.reddit.com/search/?q={query}', site: 'Reddit' },
  { shortcut: 'gh', url: 'https://github.com/search?q={query}', site: 'GitHub' },
  { shortcut: 'mdn', url: 'https://developer.mozilla.org/en-US/search?q={query}', site: 'MDN' },
  { shortcut: 'so', url: 'https://stackoverflow.com/search?q={query}', site: 'Stack Overflow' },
  { shortcut: 'ddg', url: 'https://duckduckgo.com/?q={query}', site: 'DuckDuckGo' },
  { shortcut: 'maps', url: 'https://www.google.com/maps/search/{query}', site: 'Google Maps' }
]

const activeDefaultBangs = computed(() => {
  return DEFAULT_BANGS_LIST.filter(b => !deletedBangs.value.includes(b.shortcut))
})

const AppearanceIcon = {
  render() {
    return h('svg', { width: 15, height: 15, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2, 'stroke-linecap': 'round', 'stroke-linejoin': 'round' }, [
      h('circle', { cx: 13.5, cy: 6.5, r: 0.5 }),
      h('circle', { cx: 17.5, cy: 10.5, r: 0.5 }),
      h('circle', { cx: 8.5, cy: 7.5, r: 0.5 }),
      h('circle', { cx: 6.5, cy: 12.5, r: 0.5 }),
      h('path', { d: 'M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10c.926 0 1.648-.746 1.648-1.688 0-.437-.18-.835-.437-1.125-.29-.289-.438-.652-.438-1.125a1.64 1.64 0 0 1 1.668-1.668h1.996c3.051 0 5.555-2.503 5.555-5.554C21.965 6.012 17.461 2 12 2z' })
    ])
  }
}

const BangsIcon = {
  render() {
    return h('svg', { width: 15, height: 15, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2, 'stroke-linecap': 'round', 'stroke-linejoin': 'round' }, [
      h('polygon', { points: '13 2 3 14 12 14 11 22 21 10 12 10 13 2' })
    ])
  }
}

const BlockedIcon = {
  render() {
    return h('svg', { width: 15, height: 15, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2, 'stroke-linecap': 'round', 'stroke-linejoin': 'round' }, [
      h('circle', { cx: 12, cy: 12, r: 10 }),
      h('line', { x1: 4.93, y1: 4.93, x2: 19.07, y2: 19.07 })
    ])
  }
}

const LanguageIcon = {
  render() {
    return h('svg', { width: 15, height: 15, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2, 'stroke-linecap': 'round', 'stroke-linejoin': 'round' }, [
      h('circle', { cx: 12, cy: 12, r: 10 }),
      h('line', { x1: 2, y1: 12, x2: 22, y2: 12 }),
      h('path', { d: 'M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z' })
    ])
  }
}

const tabs = computed(() => [
  { id: 'appearance', label: t.value.settings.appearance, icon: AppearanceIcon },
  { id: 'bangs', label: t.value.settings.bangs, icon: BangsIcon },
  { id: 'blocked', label: t.value.settings.blockedSites, icon: BlockedIcon },
  { id: 'language', label: t.value.settings.language, icon: LanguageIcon }
])

const setFontSize = (size) => {
  fontSize.value = size
  document.documentElement.style.fontSize = size
  localStorage.setItem('synapic_font_size', size)
}

const setTheme = (newTheme) => {
  theme.value = newTheme
  localStorage.setItem('synapic_theme', newTheme)
  applyTheme()
}

const handleAddBang = () => {
  bangError.value = ''
  const shortcut = newBangShortcut.value.trim()
  const url = newBangUrl.value.trim()
  if (!shortcut || !url) {
    bangError.value = lang.value === 'tr' ? 'Kısayol ve URL gerekli.' : 'Shortcut and URL are required.'
    return
  }
  if (!url.includes('{query}')) {
    bangError.value = lang.value === 'tr' ? 'URL {query} içermeli.' : 'URL must contain {query}.'
    return
  }
  const existing = bangs.value.find(b => b.shortcut === shortcut)
  if (existing) {
    bangError.value = lang.value === 'tr' ? 'Bu kısayol zaten mevcut.' : 'This shortcut already exists.'
    return
  }
  try {
    const siteName = new URL(url.replace('{query}', 'example')).hostname.replace('www.', '')
    addBang({ shortcut, url, site: siteName })
  } catch {
    addBang({ shortcut, url, site: url })
  }
  newBangShortcut.value = ''
  newBangUrl.value = ''
}

const changeSiteType = (index, newType) => {
  const site = blockedSites.value[index]
  if (site) {
    addBlockedSite(site.url, newType)
  }
}

const handleAddSite = (type) => {
  const url = newSiteUrl.value.trim()
  if (!url) return
  addBlockedSite(url, type)
  newSiteUrl.value = ''
}

const setLanguageSetting = (newLang) => {
  setLanguage(newLang)
}

onMounted(async () => {
  await detectLanguage()

  const savedFontSize = localStorage.getItem('synapic_font_size')
  if (savedFontSize) {
    fontSize.value = savedFontSize
    document.documentElement.style.fontSize = savedFontSize
  }

  const savedTheme = localStorage.getItem('synapic_theme')
  if (savedTheme) {
    theme.value = savedTheme
  }
  applyTheme()
  applyFontSize()
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=DM+Sans:opsz,wght@9..40,300;9..40,400;9..40,500;9..40,600&display=swap');

*, *::before, *::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

.settings-page {
  min-height: 100vh;
  background: var(--bg-primary);
  font-family: 'DM Sans', system-ui, -apple-system, sans-serif;
  color: var(--text-primary);
  -webkit-font-smoothing: antialiased;
  display: flex;
  flex-direction: column;
  transition: background-color 0.3s ease, color 0.3s ease;
}

.topbar {
  position: sticky;
  top: 0;
  z-index: 100;
  background: var(--bg-primary);
  border-bottom: 1px solid var(--border-color);
  padding: 12px 20px;
  transition: background-color 0.3s ease, border-color 0.3s ease;
}

.topbar-inner {
  max-width: 800px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  gap: 20px;
}

.logo-link {
  display: flex;
  align-items: center;
  gap: 8px;
  text-decoration: none;
  flex-shrink: 0;
  color: var(--text-primary);
}

.topbar-logo {
  width: 28px;
  height: 28px;
  object-fit: contain;
  border-radius: 6px;
  flex-shrink: 0;
}

.logo-text {
  font-size: 1.1rem;
  font-weight: 600;
  color: var(--text-primary);
  letter-spacing: -0.01em;
}

.page-title {
  font-size: 1.1rem;
  font-weight: 500;
  color: var(--text-secondary);
}

.tab-nav {
  position: sticky;
  top: 49px;
  z-index: 99;
  background: var(--bg-primary);
  border-bottom: 1px solid var(--border-color);
  transition: background-color 0.3s ease, border-color 0.3s ease;
}

.tab-nav-inner {
  max-width: 800px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  padding: 0 20px;
}

.tab-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 12px 16px;
  background: none;
  border: none;
  border-bottom: 2px solid transparent;
  color: var(--text-muted);
  font-family: 'DM Sans', system-ui, sans-serif;
  font-size: 0.85rem;
  font-weight: 500;
  cursor: pointer;
  transition: color 0.2s, border-color 0.2s;
  white-space: nowrap;
}

.tab-btn:hover {
  color: var(--text-secondary);
}

.tab-btn.active {
  color: var(--text-primary);
  border-bottom-color: var(--text-primary);
}

.tab-btn svg {
  flex-shrink: 0;
}

.content {
  flex: 1;
  max-width: 800px;
  width: 100%;
  margin: 0 auto;
  padding: 24px 28px 80px;
}

.panel {
  animation: fadeUp 0.3s ease both;
}

@keyframes fadeUp {
  from {
    opacity: 0;
    transform: translateY(8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.card {
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  padding: 24px;
  transition: background-color 0.3s ease, border-color 0.3s ease;
}

.card-title {
  font-size: 1.05rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 20px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.card-title svg {
  flex-shrink: 0;
  color: var(--text-muted);
}

.card-desc {
  color: var(--text-muted);
  font-size: 0.87rem;
  line-height: 1.6;
  margin-bottom: 20px;
}

.setting-row {
  margin-bottom: 20px;
}

.setting-row:last-child {
  margin-bottom: 0;
}

.setting-label {
  display: block;
  font-size: 0.87rem;
  font-weight: 500;
  color: var(--text-secondary);
  margin-bottom: 10px;
}

.btn-group {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.opt-btn {
  padding: 8px 18px;
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-muted);
  font-size: 0.85rem;
  font-weight: 500;
  font-family: 'DM Sans', system-ui, sans-serif;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  gap: 6px;
}

.opt-btn:hover {
  background: var(--bg-hover);
  border-color: var(--border-light);
}

.opt-btn.active {
  background: var(--accent);
  border-color: var(--accent);
  color: #ffffff;
}

[data-theme="light"] .opt-btn.active {
  color: #ffffff;
}

.opt-hint {
  font-size: 0.75rem;
  opacity: 0.5;
}

.lang-group {
  gap: 10px;
}

.lang-opt {
  min-width: 110px;
  justify-content: center;
}

.sub-section {
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid var(--border-color);
}

.sub-section:last-of-type {
  border-bottom: none;
  margin-bottom: 0;
  padding-bottom: 0;
}

.sub-title {
  font-size: 0.9rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 12px;
}

.sub-restore {
  color: var(--text-secondary);
}

.item-list {
  display: flex;
  flex-direction: column;
  gap: 4px;
  max-height: 320px;
  overflow-y: auto;
}

.item-list::-webkit-scrollbar {
  width: 4px;
}

.item-list::-webkit-scrollbar-track {
  background: transparent;
}

.item-list::-webkit-scrollbar-thumb {
  background: var(--scrollbar-thumb);
  border-radius: 10px;
}

.item-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 12px;
  background: var(--bg-hover);
  border-radius: 8px;
  border: 1px solid transparent;
  transition: all 0.15s ease;
}

.item-row:hover {
  background: var(--border-color);
  border-color: var(--border-color);
}

.item-dim {
  opacity: 0.5;
}

.item-info {
  display: flex;
  align-items: center;
  gap: 10px;
  flex: 1;
  min-width: 0;
}

.item-code {
  font-family: 'DM Mono', 'Courier New', monospace;
  font-size: 0.8rem;
  font-weight: 600;
  color: var(--text-primary);
  background: var(--bg-hover);
  padding: 3px 8px;
  border-radius: 4px;
  white-space: nowrap;
  flex-shrink: 0;
}

.item-text {
  font-size: 0.85rem;
  color: var(--text-muted);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.site-url-text {
  font-size: 0.82rem;
  color: var(--text-muted);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: 'DM Mono', 'Courier New', monospace;
}

.type-badge {
  font-size: 0.68rem;
  font-weight: 600;
  text-transform: uppercase;
  padding: 2px 8px;
  border-radius: 20px;
  letter-spacing: 0.5px;
  flex-shrink: 0;
}

.type-block {
  background: rgba(239, 68, 68, 0.12);
  color: var(--danger);
  border: 1px solid rgba(239, 68, 68, 0.25);
}

.type-restrict {
  background: rgba(251, 191, 36, 0.12);
  color: var(--warning);
  border: 1px solid rgba(251, 191, 36, 0.25);
}

.type-promote {
  background: rgba(34, 197, 94, 0.12);
  color: var(--success);
  border: 1px solid rgba(34, 197, 94, 0.25);
}

.site-actions {
  display: flex;
  gap: 4px;
  flex-shrink: 0;
  margin-left: 8px;
}

.icon-btn {
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 6px;
  color: var(--text-muted);
  cursor: pointer;
  transition: all 0.15s ease;
  flex-shrink: 0;
}

.icon-btn:hover {
  background: var(--border-color);
  color: var(--text-primary);
}

.icon-btn.danger:hover {
  background: rgba(239, 68, 68, 0.12);
  border-color: rgba(239, 68, 68, 0.25);
  color: var(--danger);
}

.icon-btn.success:hover {
  background: rgba(34, 197, 94, 0.12);
  border-color: rgba(34, 197, 94, 0.25);
  color: var(--success);
}

.icon-btn.warn:hover {
  background: rgba(251, 191, 36, 0.12);
  border-color: rgba(251, 191, 36, 0.25);
  color: var(--warning);
}

.empty-state {
  text-align: center;
  color: var(--text-muted);
  font-size: 0.85rem;
  padding: 20px;
}

.add-form {
  margin-top: 18px;
  padding-top: 18px;
  border-top: 1px solid var(--border-color);
}

.form-row {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.form-input {
  flex: 1;
  min-width: 120px;
  padding: 9px 14px;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-primary);
  font-size: 0.85rem;
  font-family: 'DM Sans', system-ui, sans-serif;
  transition: all 0.2s ease;
}

.form-input:focus {
  outline: none;
  border-color: var(--border-light);
  background: var(--border-color);
}

.form-input::placeholder {
  color: var(--text-muted);
}

.form-input.wide {
  flex: 2;
}

.add-btn {
  width: 38px;
  height: 38px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-primary);
  cursor: pointer;
  transition: all 0.2s ease;
  flex-shrink: 0;
}

.add-btn:hover {
  background: var(--border-color);
  border-color: var(--border-light);
}

.form-msg {
  font-size: 0.8rem;
  color: var(--text-muted);
  margin-top: 8px;
}

.form-msg.error {
  color: var(--danger);
}

.type-btn {
  padding: 9px 14px;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-muted);
  font-size: 0.8rem;
  font-weight: 500;
  font-family: 'DM Sans', system-ui, sans-serif;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  gap: 5px;
  white-space: nowrap;
  flex-shrink: 0;
}

.type-btn:hover {
  background: var(--border-color);
  border-color: var(--border-light);
}

.block-type:hover {
  background: rgba(239, 68, 68, 0.12);
  border-color: rgba(239, 68, 68, 0.25);
  color: var(--danger);
}

.restrict-type:hover {
  background: rgba(251, 191, 36, 0.12);
  border-color: rgba(251, 191, 36, 0.25);
  color: var(--warning);
}

.promote-type:hover {
  background: rgba(34, 197, 94, 0.12);
  border-color: rgba(34, 197, 94, 0.25);
  color: var(--success);
}

.site-row {
  flex-direction: column;
  align-items: flex-start;
  gap: 8px;
}

.site-row .site-actions {
  margin-left: 0;
  width: 100%;
  justify-content: flex-end;
}

@media (max-width: 640px) {
  .tab-btn span {
    display: none;
  }

  .tab-btn {
    padding: 12px 14px;
  }

  .content {
    padding: 20px 16px 60px;
  }

  .card {
    padding: 18px;
  }

  .form-row {
    flex-direction: column;
  }

  .form-input,
  .form-input.wide {
    min-width: 100%;
  }

  .type-btn {
    width: 100%;
    justify-content: center;
  }

  .add-btn {
    width: 100%;
    height: auto;
    padding: 9px;
  }

  .site-row .site-actions {
    justify-content: flex-end;
  }
}

@media (max-width: 480px) {
  .page-title {
    font-size: 1rem;
  }

  .card-title {
    font-size: 0.95rem;
  }

  .opt-btn {
    padding: 7px 14px;
    font-size: 0.82rem;
  }

  .btn-group {
    gap: 6px;
  }
}
</style>
