<template>
  <div class="home">
    <nav class="topnav">
      <div class="nav-right">
        <router-link to="/api" class="nav-item nav-api">
          <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="22 12 18 12 15 21 9 3 6 12 2 12"/>
          </svg>
          <span>{{ t.nav.api }}</span>
        </router-link>
        <router-link to="/settings" class="nav-item nav-settings">
          <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="12" r="3"/>
            <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83-2.83l.06-.06A1.65 1.65 0 0 0 4.68 15a1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 2.83-2.83l.06.06A1.65 1.65 0 0 0 9 4.68a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 2.83l-.06.06A1.65 1.65 0 0 0 19.4 9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"/>
          </svg>
          <span>{{ t.home.settings }}</span>
        </router-link>
      </div>
    </nav>

    <main class="center">
      <div class="brand">
        <img src="../images/synapic.png" alt="Synapic" class="brand-logo" />
        <h1 class="wordmark">Synapic</h1>
      </div>

      <form class="search-form" @submit.prevent="handleSearch">
        <div class="search-wrap" :class="{ focused: isFocused }">
          <input
            ref="inputRef"
            v-model="query"
            type="text"
            class="search-input"
            :placeholder="t.home.placeholder"
            autocomplete="off"
            spellcheck="false"
            @focus="isFocused = true"
            @blur="isFocused = false"
          />
          <button type="submit" class="search-btn" aria-label="Search">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="var(--text-secondary)" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="11" cy="11" r="8"/>
              <line x1="21" y1="21" x2="16.65" y2="16.65"/>
            </svg>
          </button>
        </div>
      </form>
    </main>

    <footer class="footer">
      <div class="footer-links">
        <router-link to="/terms" class="footer-link">
          <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/></svg>
          <span>{{ t.nav.terms }}</span>
        </router-link>
      </div>
      <span class="footer-copy">&copy; Synapic. Search honestly.</span>
    </footer>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { lang, detectLanguage, tr, en } from '../composables/useLocale.js'
import { initSettings } from '../composables/settings.js'
import { useBangs } from '../composables/useBangs.js'

const router = useRouter()
const { checkBang } = useBangs()
const query = ref('')
const isFocused = ref(false)
const inputRef = ref(null)

const t = computed(() => lang.value === 'tr' ? tr : en)

onMounted(async () => {
  await detectLanguage()
  initSettings()
  inputRef.value?.focus()
})

function handleSearch() {
  const q = query.value.trim()
  if (!q) return
  if (q.startsWith('!')) {
    const result = checkBang(q)
    if (result) {
      const url = result.bang.url.replace('{query}', encodeURIComponent(result.query))
      window.open(url, '_blank')
      return
    }
  }
  router.push({ path: '/search', query: { q } })
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=DM+Sans:opsz,wght@9..40,300;9..40,400;9..40,500&display=swap');

*, *::before, *::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

.home {
  min-height: 100vh;
  background: var(--bg-primary);
  display: flex;
  flex-direction: column;
  font-family: var(--font-family, 'DM Sans', system-ui, sans-serif);
  color: var(--text-primary);
  -webkit-font-smoothing: antialiased;
  transition: background-color 0.3s ease, color 0.3s ease;
}

.topnav {
  display: flex;
  justify-content: flex-end;
  padding: 16px 28px;
}

.nav-right {
  display: flex;
  gap: 22px;
  align-items: center;
}

.nav-item {
  font-size: 13.5px;
  color: var(--text-muted);
  cursor: pointer;
  transition: color 0.15s;
}

.nav-item:hover {
  color: var(--text-secondary);
}

.nav-settings {
  display: flex;
  align-items: center;
  gap: 6px;
  text-decoration: none;
  color: var(--text-muted);
}

.nav-settings:hover {
  color: var(--text-secondary);
}

.nav-settings svg {
  flex-shrink: 0;
}

.nav-api {
  display: flex;
  align-items: center;
  gap: 6px;
  text-decoration: none;
  color: var(--text-muted);
}

.nav-api:hover {
  color: var(--text-secondary);
}

.nav-api svg {
  flex-shrink: 0;
}

.center {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 0 24px 80px;
  gap: 30px;
}

.brand {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 14px;
}

.brand-logo {
  width: 80px;
  height: 80px;
  object-fit: contain;
  border-radius: 16px;
}

.wordmark {
  font-size: 44px;
  font-weight: 300;
  color: var(--text-primary);
  letter-spacing: -1.5px;
  line-height: 1;
  transition: color 0.3s ease;
}

.search-form {
  width: 100%;
  max-width: 630px;
}

.search-wrap {
  display: flex;
  align-items: center;
  background: var(--bg-input);
  border-radius: 100px;
  height: 52px;
  padding: 0 7px 0 22px;
  transition: background 0.2s ease;
  border: 1px solid var(--border-color);
}

.search-wrap.focused {
  background: var(--bg-hover);
  border-color: var(--border-light);
}

.search-input {
  flex: 1;
  background: none;
  border: none;
  outline: none;
  font-family: var(--font-family, 'DM Sans', system-ui, sans-serif);
  font-size: 16px;
  font-weight: 300;
  color: var(--text-primary);
  caret-color: var(--text-primary);
  letter-spacing: 0.01em;
  min-width: 0;
  transition: color 0.3s ease;
}

.search-input::placeholder {
  color: var(--text-muted);
}

.search-btn {
  background: var(--bg-secondary);
  border: none;
  cursor: pointer;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  transition: background 0.15s ease;
}

.search-btn:hover {
  background: var(--bg-hover);
}

.search-btn svg {
  flex-shrink: 0;
}

.search-btn:hover svg {
  stroke: var(--text-primary);
}

.footer {
  padding: 14px 24px 26px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
}

.footer-links {
  display: flex;
  gap: 20px;
  align-items: center;
}

.footer-link {
  display: flex;
  align-items: center;
  gap: 5px;
  font-size: 12px;
  color: var(--text-muted);
  text-decoration: none;
  cursor: pointer;
  transition: color 0.15s;
}

.footer-link:hover {
  color: var(--text-secondary);
}

.footer-link svg {
  flex-shrink: 0;
  opacity: 0.7;
}

.footer-copy {
  font-size: 12px;
  color: var(--text-muted);
  transition: color 0.3s ease;
}

@media (max-width: 480px) {
  .wordmark {
    font-size: 34px;
  }
  .search-wrap {
    height: 48px;
  }
  .search-input {
    font-size: 15px;
  }
  .nav-settings span {
    display: none;
  }
  .nav-api span {
    display: none;
  }
}
</style>