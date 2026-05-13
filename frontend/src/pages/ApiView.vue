<template>
  <div class="api-page">
    <header class="topbar">
      <div class="topbar-inner">
        <router-link to="/" class="logo-link">
          <img src="../images/synapic.png" alt="Synapic" class="topbar-logo" />
          <span class="logo-text">Synapic</span>
        </router-link>
        <h1 class="page-title">{{ t.nav.api }}</h1>
        <div v-if="isLoggedIn" class="topbar-user">
          <span class="topbar-email">{{ userEmail }}</span>
          <button class="icon-btn-sm" @click="handleSignOut" :title="t.nav.logout">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/><polyline points="16 17 21 12 16 7"/><line x1="21" y1="12" x2="9" y2="12"/></svg>
          </button>
        </div>
      </div>
    </header>

    <main class="content">
      <div class="content-inner" v-if="!isLoggedIn">
        <div class="login-card">
          <div class="login-icon-wrap">
            <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></svg>
          </div>
          <h2 class="login-title">{{ t.api.loginTitle }}</h2>
          <p class="login-desc">{{ t.api.loginDesc }}</p>
          <button class="google-login-btn" @click="handleGoogleLogin">
            <svg width="18" height="18" viewBox="0 0 24 24"><path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92a5.06 5.06 0 0 1-2.2 3.32v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.1z" fill="#4285F4"/><path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/><path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/><path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/></svg>
            <span>{{ t.api.loginBtn }}</span>
          </button>
        </div>
      </div>

      <div class="content-inner" v-else>
        <div class="stats-grid">
          <div class="stat-card">
            <div class="stat-icon">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>
            </div>
            <div class="stat-info">
              <span class="stat-label">{{ t.api.usedToday }}</span>
              <span class="stat-value">{{ stats.searchesToday }}</span>
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-icon">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><polyline points="12 6 12 12 16 14"/></svg>
            </div>
            <div class="stat-info">
              <span class="stat-label">{{ t.api.dailyLimit }}</span>
              <span class="stat-value">{{ stats.dailyLimit === -1 ? t.api.unlimited : stats.dailyLimit }}</span>
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-icon">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M13 2L3 14h9l-1 8 10-12h-9l1-8z"/></svg>
            </div>
            <div class="stat-info">
              <span class="stat-label">{{ t.api.remaining }}</span>
              <span class="stat-value">{{ stats.remaining === -1 ? t.api.unlimited : stats.remaining }}</span>
            </div>
          </div>
        </div>

        <div class="card-section">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 2l-2 2m-7.61 7.61a5.5 5.5 0 1 1-7.778 7.778 5.5 5.5 0 0 1 7.777-7.777zm0 0L15.5 7.5m0 0l3 3L22 7l-3-3m-3.5 3.5L19 4"/></svg>
              {{ t.api.apiKey }}
            </h2>
            <div class="key-display">
              <code class="key-text">{{ maskedApiKey }}</code>
              <div class="key-actions">
                <button class="icon-btn" @click="toggleApiKeyVisibility" :title="showFullKey ? t.api.hide : t.api.show">
                  <svg v-if="showFullKey" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"/><line x1="1" y1="1" x2="23" y2="23"/></svg>
                  <svg v-else width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>
                </button>
                <button class="icon-btn" @click="copyApiKey" :title="t.api.copyKey">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"/><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"/></svg>
                </button>
              </div>
            </div>
            <button class="regen-btn" @click="regenerateKey">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="23 4 23 10 17 10"/><polyline points="1 20 1 14 7 14"/><path d="M3.51 9a9 9 0 0 1 14.85-3.36L23 10M1 14l4.64 4.36A9 9 0 0 0 20.49 15"/></svg>
              {{ t.api.regenKey }}
            </button>
          </div>
        </div>

        <div class="card-section">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="22 12 18 12 15 21 9 3 6 12 2 12"/></svg>
              {{ t.api.usageTimeline }}
            </h2>
            <div class="timeline-header">
              <span>{{ t.api.last24Hours }}</span>
              <span class="timeline-date">{{ currentDate }}</span>
            </div>
            <div class="timeline-chart">
              <div v-for="(hour, index) in usageData" :key="index" class="timeline-bar-wrap">
                <div class="bar-tooltip">
                  <span class="bar-tooltip-time">{{ hour.hour }}:00</span>
                  <span class="bar-tooltip-count">{{ hour.count }} {{ t.api.searches }}</span>
                </div>
                <div class="timeline-bar" :style="{ height: getBarHeight(hour.count) + '%' }"></div>
                <span class="timeline-label">{{ hour.hour }}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="card-section">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"/></svg>
              {{ t.api.apiTest }}
            </h2>
            <div class="test-input-group">
              <label class="test-label">{{ t.api.searchQuery }}</label>
              <input
                v-model="testQuery"
                type="text"
                class="form-input"
                :placeholder="t.api.searchPlaceholder"
                @keyup.enter="testApiSearch"
              />
            </div>
            <button class="test-btn" @click="testApiSearch" :disabled="testLoading || !testQuery">
              <svg v-if="testLoading" class="spin-icon" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12a9 9 0 1 1-6.219-8.56"/></svg>
              <svg v-else width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="5 3 19 12 5 21 5 3"/></svg>
              <span>{{ testLoading ? t.api.searching : t.api.testSearch }}</span>
            </button>

            <div v-if="testResponse" class="test-response">
              <h3 class="test-response-title">{{ t.api.result }}</h3>
              <div class="test-stats">
                <span class="test-stat-item"><strong>{{ t.api.totalResults }}</strong> {{ testResponse.total }}</span>
                <span class="test-stat-item" v-if="testResponse.total_used !== undefined"><strong>{{ t.api.used }}</strong> {{ testResponse.total_used }}</span>
                <span class="test-stat-item" v-if="testResponse.remaining !== undefined"><strong>{{ t.api.remaining }}</strong> {{ testResponse.remaining }}</span>
              </div>
              <div class="test-results">
                <div v-for="(result, idx) in testResponse.results?.slice(0, 3)" :key="idx" class="test-result-item">
                  <a :href="result.url" target="_blank" class="test-result-link">{{ result.title || result.url }}</a>
                  <p class="test-result-desc">{{ result.description }}</p>
                  <span class="test-result-score">{{ t.api.score }} {{ result.score.toFixed(2) }}</span>
                </div>
              </div>
            </div>

            <div v-if="testError" class="test-error">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"/><line x1="12" y1="9" x2="12" y2="13"/><line x1="12" y1="17" x2="12.01" y2="17"/></svg>
              <span>{{ testError }}</span>
            </div>
          </div>
        </div>

        <div class="card-section">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/><polyline points="10 9 9 9 8 9"/></svg>
              {{ t.api.apiDocumentation }}
            </h2>

            <h3 class="doc-subtitle">Endpoint</h3>
            <div class="code-block">
              <code>https://api.synapicsearch.com/api/search?q=YOUR_QUERY&apikey=YOUR_API_KEY</code>
            </div>

            <h3 class="doc-subtitle">{{ t.api.authentication }}</h3>
            <p class="doc-text">{{ t.api.authDesc }}</p>
            <h4 class="doc-subsubtitle">{{ t.api.option1 }}</h4>
            <div class="code-block"><code>?apikey={{ apiKey }}</code></div>
            <h4 class="doc-subsubtitle">{{ t.api.option2 }}</h4>
            <div class="code-block"><code>X-API-Key: {{ apiKey }}</code></div>

            <h3 class="doc-subtitle">{{ t.api.exampleCurl }}</h3>
            <div class="code-block">
              <pre>curl -X GET "https://api.synapicsearch.com/api/search?q=github&apikey={{ apiKey }}"</pre>
            </div>

            <h3 class="doc-subtitle">{{ t.api.exampleJs }}</h3>
            <div class="code-block">
              <pre>fetch('https://api.synapicsearch.com/api/search?q=github&apikey={{ apiKey }}')
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error('Error:', error));</pre>
            </div>

            <h3 class="doc-subtitle">{{ t.api.responseFormat }}</h3>
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
  "total_used": 15,
  "remaining": 85,
  "daily_limit": 100
}</pre>
            </div>

            <h3 class="doc-subtitle">{{ t.api.errorCodes }}</h3>
            <div class="error-codes">
              <div class="error-code-item">
                <strong>API key required</strong>
                <span>{{ t.api.errorMissing }}</span>
              </div>
              <div class="error-code-item">
                <strong>Invalid API key</strong>
                <span>{{ t.api.errorInvalid }}</span>
              </div>
              <div class="error-code-item">
                <strong>Daily limit exceeded</strong>
                <span>{{ t.api.errorExceeded }}</span>
              </div>
              <div class="error-code-item">
                <strong>Query required</strong>
                <span>{{ t.api.errorQuery }}</span>
              </div>
            </div>

            <h3 class="doc-subtitle">{{ t.api.limits }}</h3>
            <ul class="doc-list">
              <li>{{ t.api.limit1 }}</li>
              <li>{{ t.api.limit2 }}</li>
              <li>{{ t.api.limit3 }}</li>
              <li>{{ t.api.limit4 }}</li>
            </ul>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { lang, detectLanguage, tr, en } from '../composables/useLocale.js'
import { applyTheme, applyFontSize } from '../composables/settings.js'

const router = useRouter()
const showMoreMenu = ref(false)
const apiKey = ref('')
const showFullKey = ref(false)
const isLoggedIn = ref(false)
const userEmail = ref('')
const googleAuthReady = ref(false)
const stats = ref({ searchesToday: 0, dailyLimit: 100, remaining: 100 })
const testQuery = ref('')
const testLoading = ref(false)
const testResponse = ref(null)
const testError = ref('')
const usageData = ref([])
const currentTheme = ref('dark')
let statsInterval = null

const t = computed(() => lang.value === 'tr' ? tr : en)

const API_BASE = 'https://api.synapicsearch.com'

const maskedApiKey = computed(() => {
  if (!apiKey.value) return '••••••••••••••••••••••••••••••••'
  if (showFullKey.value) return apiKey.value
  return apiKey.value.substring(0, 15) + '••••••••••••••••••••'
})

const currentDate = computed(() => {
  const now = new Date()
  return now.toLocaleDateString(lang.value === 'tr' ? 'tr-TR' : 'en-US', { day: 'numeric', month: 'long', year: 'numeric' })
})

const getBarHeight = (count) => {
  const maxCount = Math.max(...usageData.value.map(d => d.count), 1)
  return maxCount === 0 ? 0 : (count / maxCount) * 100
}

const initializeUsageData = () => {
  const data = []
  const now = new Date()
  const currentHour = now.getHours()
  for (let i = 0; i < 24; i++) {
    const hour = (currentHour - 23 + i + 24) % 24
    data.push({ hour: hour.toString().padStart(2, '0'), count: 0 })
  }
  usageData.value = data
}

const updateUsageData = () => {
  if (stats.value.searchesToday > 0) {
    const currentHourStr = new Date().getHours().toString().padStart(2, '0')
    const hourIndex = usageData.value.findIndex(d => d.hour === currentHourStr)
    if (hourIndex !== -1) usageData.value[hourIndex].count = stats.value.searchesToday
  }
}

const fetchStats = async () => {
  try {
    const key = localStorage.getItem('apiKey')
    if (!key) return
    const response = await fetch(`${API_BASE}/user/stats`, {
      method: 'GET', mode: 'cors', credentials: 'include',
      headers: { 'X-API-Key': key, 'Content-Type': 'application/json' }
    })
    if (response.ok) {
      const data = await response.json()
      stats.value = { searchesToday: data.searches_today || 0, dailyLimit: data.daily_limit || 100, remaining: data.remaining || 0 }
      updateUsageData()
    }
  } catch (error) { console.error('Stats error:', error) }
}

const testApiSearch = async () => {
  if (!testQuery.value.trim()) return
  testLoading.value = true
  testError.value = ''
  testResponse.value = null
  try {
    const response = await fetch(`${API_BASE}/api/search?q=${encodeURIComponent(testQuery.value)}&apikey=${apiKey.value}`, { method: 'GET', mode: 'cors' })
    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.error || t.value.api.searchFailed)
    }
    testResponse.value = await response.json()
    await fetchStats()
  } catch (error) { testError.value = error.message }
  finally { testLoading.value = false }
}

const toggleApiKeyVisibility = () => { showFullKey.value = !showFullKey.value }

const copyApiKey = async () => {
  try {
    await navigator.clipboard.writeText(apiKey.value)
    alert(t.value.api.copySuccess)
  } catch { console.error('Copy failed') }
}

const regenerateKey = async () => {
  if (!confirm(t.value.api.regenConfirm)) return
  try {
    const response = await fetch(`${API_BASE}/auth/regenerate`, {
      method: 'POST', mode: 'cors', credentials: 'include',
      headers: { 'Content-Type': 'application/json', 'X-API-Key': apiKey.value }
    })
    const data = await response.json()
    if (data.success) {
      apiKey.value = data.apikey
      localStorage.setItem('apiKey', data.apikey)
      alert(t.value.api.regenSuccess)
    } else {
      alert(t.value.api.regenFail + (data.message || ''))
    }
  } catch { alert(t.value.api.regenFailRetry) }
}

const loadGoogleScript = () => {
  return new Promise((resolve, reject) => {
    if (window.google?.accounts) { resolve(); return }
    const script = document.createElement('script')
    script.src = 'https://accounts.google.com/gsi/client'
    script.async = true; script.defer = true
    script.onload = resolve; script.onerror = reject
    document.head.appendChild(script)
  })
}

const initGoogleAuth = async () => {
  try { await loadGoogleScript(); googleAuthReady.value = true }
  catch (error) { console.error('Google script error:', error) }
}

const handleGoogleLogin = async () => {
  if (!googleAuthReady.value) {
    alert(t.value.api.authLoading)
    await initGoogleAuth()
    if (!googleAuthReady.value) { alert(t.value.api.authFail); return }
  }
  try {
    const client = window.google.accounts.oauth2.initTokenClient({
      client_id: '449095352954-sltui6f9t112ld2abk52s8fgfhle9f0s.apps.googleusercontent.com',
      scope: 'email profile',
      callback: async (response) => {
        if (response.error) { alert(t.value.api.loginFail); return }
        try {
          const userInfoResponse = await fetch('https://www.googleapis.com/oauth2/v3/userinfo', {
            headers: { 'Authorization': `Bearer ${response.access_token}` }
          })
          const userInfo = await userInfoResponse.json()
          const backendResponse = await fetch(`${API_BASE}/auth/google`, {
            method: 'POST', mode: 'cors', credentials: 'include',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ google_id: userInfo.sub, email: userInfo.email }),
          })
          const data = await backendResponse.json()
          if (data.success) {
            localStorage.setItem('apiKey', data.apikey)
            localStorage.setItem('userEmail', userInfo.email)
            apiKey.value = data.apikey
            isLoggedIn.value = true
            userEmail.value = userInfo.email
            showMoreMenu.value = false
            initializeUsageData()
            fetchStats()
            statsInterval = setInterval(fetchStats, 5000)
            alert(t.value.api.loginSuccess)
          } else { alert(t.value.api.loginFailMsg + (data.message || '')) }
        } catch { alert(t.value.api.loginFail) }
      },
    })
    client.requestAccessToken()
  } catch { alert(t.value.api.loginFail) }
}

const handleSignOut = async () => {
  if (window.google?.accounts?.id) { window.google.accounts.id.disableAutoSelect() }
  if (statsInterval) clearInterval(statsInterval)
  localStorage.removeItem('apiKey')
  localStorage.removeItem('userEmail')
  isLoggedIn.value = false
  userEmail.value = ''
  apiKey.value = ''
  showMoreMenu.value = false
  testQuery.value = ''
  testResponse.value = null
  testError.value = ''
  alert(t.value.api.logoutSuccess)
  router.push('/')
}

onMounted(async () => {
  await detectLanguage()
  applyTheme()
  applyFontSize()
  currentTheme.value = localStorage.getItem('synapic_theme') || 'dark'
  const key = localStorage.getItem('apiKey')
  const email = localStorage.getItem('userEmail')
  if (key && email) {
    apiKey.value = key
    isLoggedIn.value = true
    userEmail.value = email
    initializeUsageData()
    fetchStats()
    statsInterval = setInterval(fetchStats, 5000)
  }
  initGoogleAuth()
})

onUnmounted(() => { if (statsInterval) clearInterval(statsInterval) })
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=DM+Sans:opsz,wght@9..40,300;9..40,400;9..40,500;9..40,600&display=swap');

*, *::before, *::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

.api-page {
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
  max-width: 860px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  gap: 16px;
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

.topbar-user {
  margin-left: auto;
  display: flex;
  align-items: center;
  gap: 10px;
}

.topbar-email {
  font-size: 0.82rem;
  color: var(--text-muted);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  max-width: 180px;
}

.icon-btn-sm {
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

.icon-btn-sm:hover {
  background: rgba(239, 68, 68, 0.12);
  border-color: rgba(239, 68, 68, 0.25);
  color: var(--danger);
}

.content {
  flex: 1;
  max-width: 860px;
  width: 100%;
  margin: 0 auto;
  padding: 24px 28px 80px;
}

.content-inner {
  animation: fadeUp 0.3s ease both;
}

@keyframes fadeUp {
  from { opacity: 0; transform: translateY(8px); }
  to { opacity: 1; transform: translateY(0); }
}

.login-card {
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  padding: 48px 32px;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  transition: background-color 0.3s ease, border-color 0.3s ease;
}

.login-icon-wrap {
  width: 80px;
  height: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 20px;
  color: var(--text-muted);
  margin-bottom: 24px;
}

.login-title {
  font-size: 1.3rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 12px;
}

.login-desc {
  color: var(--text-muted);
  font-size: 0.87rem;
  line-height: 1.6;
  max-width: 420px;
  margin-bottom: 28px;
}

.google-login-btn {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 24px;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-primary);
  font-size: 0.87rem;
  font-weight: 500;
  font-family: 'DM Sans', system-ui, sans-serif;
  cursor: pointer;
  transition: all 0.2s ease;
}

.google-login-btn:hover {
  background: var(--border-color);
  border-color: var(--border-light);
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
  margin-bottom: 16px;
}

.stat-card {
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  padding: 20px;
  display: flex;
  align-items: center;
  gap: 14px;
  transition: background-color 0.3s ease, border-color 0.3s ease;
}

.stat-icon {
  width: 42px;
  height: 42px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 10px;
  color: var(--text-muted);
  flex-shrink: 0;
}

.stat-info {
  display: flex;
  flex-direction: column;
  min-width: 0;
}

.stat-label {
  font-size: 0.78rem;
  font-weight: 500;
  color: var(--text-muted);
  margin-bottom: 4px;
}

.stat-value {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--text-primary);
  line-height: 1.2;
}

.card-section {
  margin-bottom: 16px;
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

.key-display {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 14px;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  margin-bottom: 12px;
}

.key-text {
  flex: 1;
  font-family: 'DM Mono', 'Courier New', monospace;
  font-size: 0.82rem;
  color: var(--text-primary);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.key-actions {
  display: flex;
  gap: 4px;
  flex-shrink: 0;
}

.icon-btn {
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-card);
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

.regen-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  padding: 10px 16px;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-secondary);
  font-size: 0.85rem;
  font-weight: 500;
  font-family: 'DM Sans', system-ui, sans-serif;
  cursor: pointer;
  transition: all 0.2s ease;
}

.regen-btn:hover {
  background: rgba(239, 68, 68, 0.1);
  border-color: rgba(239, 68, 68, 0.25);
  color: var(--danger);
}

.timeline-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  color: var(--text-secondary);
  font-size: 0.85rem;
}

.timeline-date {
  color: var(--text-muted);
}

.timeline-chart {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  height: 180px;
  gap: 3px;
}

.timeline-bar-wrap {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100%;
  justify-content: flex-end;
  position: relative;
}

.timeline-bar-wrap:hover .bar-tooltip {
  opacity: 1;
  visibility: visible;
  transform: translateX(-50%) translateY(0);
}

.timeline-bar {
  width: 100%;
  background: var(--accent);
  border-radius: 3px 3px 0 0;
  transition: all 0.2s ease;
  cursor: pointer;
  min-height: 2px;
  opacity: 0.7;
}

.timeline-bar:hover {
  opacity: 1;
}

.bar-tooltip {
  position: absolute;
  bottom: 100%;
  left: 50%;
  transform: translateX(-50%) translateY(4px);
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  color: var(--text-primary);
  padding: 8px 12px;
  border-radius: 6px;
  font-size: 0.75rem;
  white-space: nowrap;
  opacity: 0;
  visibility: hidden;
  pointer-events: none;
  transition: all 0.2s ease;
  margin-bottom: 8px;
  z-index: 10;
  box-shadow: 0 4px 12px rgba(0,0,0,0.15);
}

.bar-tooltip::after {
  content: "";
  position: absolute;
  top: 100%;
  left: 50%;
  transform: translateX(-50%);
  border: 5px solid transparent;
  border-top-color: var(--border-color);
}

.bar-tooltip-time {
  font-weight: 600;
  display: block;
  margin-bottom: 2px;
  color: var(--text-primary);
}

.bar-tooltip-count {
  color: var(--text-muted);
  font-size: 0.7rem;
}

.timeline-label {
  font-size: 0.6rem;
  color: var(--text-muted);
  margin-top: 6px;
}

.test-input-group {
  margin-bottom: 14px;
}

.test-label {
  display: block;
  color: var(--text-secondary);
  font-size: 0.85rem;
  font-weight: 500;
  margin-bottom: 8px;
}

.form-input {
  width: 100%;
  padding: 10px 14px;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-primary);
  font-size: 0.87rem;
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

.test-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  padding: 11px 16px;
  background: var(--accent);
  border: 1px solid var(--accent);
  border-radius: 8px;
  color: #ffffff;
  font-size: 0.87rem;
  font-weight: 500;
  font-family: 'DM Sans', system-ui, sans-serif;
  cursor: pointer;
  transition: all 0.2s ease;
}

.test-btn:hover:not(:disabled) {
  opacity: 0.9;
}

.test-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.spin-icon {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.test-response {
  margin-top: 18px;
  padding: 18px;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 8px;
}

.test-response-title {
  font-size: 0.92rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 14px;
}

.test-stats {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
  margin-bottom: 14px;
  padding-bottom: 14px;
  border-bottom: 1px solid var(--border-color);
}

.test-stat-item {
  font-size: 0.82rem;
  color: var(--text-secondary);
}

.test-stat-item strong {
  color: var(--text-primary);
}

.test-results {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.test-result-item {
  padding: 12px;
  background: var(--bg-card);
  border-radius: 8px;
  border: 1px solid var(--border-color);
}

.test-result-link {
  display: block;
  color: var(--accent);
  font-weight: 500;
  font-size: 0.85rem;
  text-decoration: none;
  margin-bottom: 4px;
  transition: opacity 0.2s ease;
}

.test-result-link:hover {
  opacity: 0.8;
}

.test-result-desc {
  font-size: 0.8rem;
  color: var(--text-muted);
  margin-bottom: 4px;
  line-height: 1.5;
}

.test-result-score {
  font-size: 0.72rem;
  color: var(--text-muted);
}

.test-error {
  margin-top: 14px;
  padding: 12px 14px;
  background: rgba(239, 68, 68, 0.08);
  border: 1px solid rgba(239, 68, 68, 0.2);
  border-radius: 8px;
  color: var(--danger);
  font-size: 0.85rem;
  display: flex;
  align-items: center;
  gap: 8px;
}

.doc-subtitle {
  font-size: 0.95rem;
  font-weight: 600;
  color: var(--text-primary);
  margin: 24px 0 10px;
}

.doc-subtitle:first-of-type {
  margin-top: 0;
}

.doc-subsubtitle {
  font-size: 0.85rem;
  font-weight: 500;
  color: var(--text-secondary);
  margin: 14px 0 6px;
}

.doc-text {
  color: var(--text-secondary);
  font-size: 0.85rem;
  margin-bottom: 10px;
  line-height: 1.6;
}

.doc-list {
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 6px;
  margin-top: 8px;
}

.doc-list li {
  color: var(--text-secondary);
  font-size: 0.85rem;
  line-height: 1.6;
  padding-left: 18px;
  position: relative;
}

.doc-list li::before {
  content: "";
  position: absolute;
  left: 4px;
  top: 9px;
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: var(--text-muted);
}

.code-block {
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  padding: 14px;
  overflow-x: auto;
  margin-bottom: 14px;
}

.code-block code,
.code-block pre {
  font-family: 'DM Mono', 'Courier New', monospace;
  font-size: 0.78rem;
  color: var(--text-primary);
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
  gap: 8px;
  margin-top: 10px;
}

.error-code-item {
  display: flex;
  flex-direction: column;
  gap: 3px;
  padding: 12px;
  background: var(--bg-hover);
  border-radius: 8px;
  border: 1px solid var(--border-color);
}

.error-code-item strong {
  color: var(--danger);
  font-size: 0.82rem;
  font-family: 'DM Mono', 'Courier New', monospace;
}

.error-code-item span {
  color: var(--text-muted);
  font-size: 0.8rem;
}

@media (max-width: 768px) {
  .topbar-email {
    max-width: 120px;
  }

  .content {
    padding: 20px 16px 60px;
  }

  .stats-grid {
    grid-template-columns: 1fr;
  }

  .stat-card {
    padding: 16px;
  }

  .login-card {
    padding: 36px 24px;
  }

  .card {
    padding: 18px;
  }

  .timeline-chart {
    height: 140px;
  }

  .timeline-label {
    font-size: 0.55rem;
  }
}

@media (max-width: 480px) {
  .topbar-email {
    display: none;
  }

  .page-title {
    font-size: 1rem;
  }

  .content {
    padding: 16px 12px 48px;
  }

  .login-card {
    padding: 28px 16px;
  }

  .login-title {
    font-size: 1.1rem;
  }

  .login-icon-wrap {
    width: 64px;
    height: 64px;
  }

  .login-icon-wrap svg {
    width: 36px;
    height: 36px;
  }

  .stat-value {
    font-size: 1.3rem;
  }

  .key-display {
    flex-direction: column;
    align-items: stretch;
  }

  .key-actions {
    justify-content: flex-end;
  }

  .card {
    padding: 14px;
  }

  .card-title {
    font-size: 0.95rem;
  }

  .code-block {
    padding: 10px;
  }

  .code-block code,
  .code-block pre {
    font-size: 0.7rem;
  }

  .timeline-chart {
    height: 120px;
  }

  .timeline-label {
    font-size: 0.5rem;
  }

  .test-stats {
    flex-direction: column;
    gap: 6px;
  }
}

@media (hover: none) and (pointer: coarse) {
  .icon-btn,
  .icon-btn-sm,
  .regen-btn,
  .test-btn,
  .google-login-btn {
    -webkit-tap-highlight-color: transparent;
  }
}
</style>
