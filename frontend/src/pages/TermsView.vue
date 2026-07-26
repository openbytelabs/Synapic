<template>
  <div class="terms-page">
    <header class="topbar">
      <div class="topbar-inner">
        <router-link to="/" class="logo-link">
          <img src="../images/synapic.png" alt="Synapic" class="topbar-logo" />
          <span class="logo-text">Synapic</span>
        </router-link>
        <h1 class="page-title">{{ t.terms.mainTitle }}</h1>
      </div>
    </header>

    <main class="content">
      <div class="content-inner">
        <h2 class="main-title">{{ t.terms.mainTitle }}</h2>
        <p class="subtitle">{{ t.terms.intro1 }}</p>
        <p class="subtitle">{{ t.terms.intro2 }}</p>
        <p class="meta-line">{{ t.terms.lastUpdated }}: 2026-06-26</p>
        <div class="card-section">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/></svg>
              {{ t.terms.privacyTitle }}
            </h2>
            <p class="card-text">{{ t.terms.privacyDesc }}</p>

            <!-- 1.1 Toplamadiklarimiz -->
            <h3 class="sub-title">{{ t.terms.noCollectionTitle }}</h3>
            <p class="card-text">{{ t.terms.noCollectionDesc }}</p>
            <ul class="bullet-list">
              <li>{{ t.terms.noCollection1 }}</li>
              <li>{{ t.terms.noCollection2 }}</li>
              <li>{{ t.terms.noCollection3 }}</li>
              <li>{{ t.terms.noCollection4 }}</li>
              <li>{{ t.terms.noCollection5 }}</li>
              <li>{{ t.terms.noCollection6 }}</li>
              <li>{{ t.terms.noCollection7 }}</li>
            </ul>

            <!-- 1.2 Lokal Veriler -->
            <h3 class="sub-title">{{ t.terms.localDataTitle }}</h3>
            <p class="card-text">{{ t.terms.localDataDesc }}</p>
            <ul class="bullet-list">
              <li>{{ t.terms.localData1 }}</li>
              <li>{{ t.terms.localData2 }}</li>
              <li>{{ t.terms.localData3 }}</li>
              <li>{{ t.terms.localData4 }}</li>
            </ul>
            <p class="card-text emphasis">{{ t.terms.localDataClearable }}</p>

            <h3 class="sub-title">{{ t.terms.serverDataTitle }}</h3>
            <p class="card-text">{{ t.terms.serverDataDesc }}</p>
            <ul class="bullet-list">
              <li>{{ t.terms.serverData1 }}</li>
              <li>{{ t.terms.serverData2 }}</li>
              <li>{{ t.terms.serverData3 }}</li>
              <li>{{ t.terms.serverData4 }}</li>
            </ul>

            <!-- 1.4 Indekslenen Veri Hacmi -->
            <h3 class="sub-title">{{ t.terms.indexDataTitle }}</h3>
            <p class="card-text">{{ t.terms.indexDataDesc }}</p>
            <ul class="bullet-list">
              <li>{{ t.terms.indexData1 }}</li>
              <li>{{ t.terms.indexData2 }}</li>
              <li>{{ t.terms.indexData3 }}</li>
            </ul>

            <!-- 1.5 Yapay Zeka Servisleri -->
            <h3 class="sub-title">{{ t.terms.aiDataTitle }}</h3>
            <p class="card-text">{{ t.terms.aiDataDesc }}</p>
            <ul class="bullet-list">
              <li>{{ t.terms.aiData1 }}</li>
              <li>{{ t.terms.aiData2 }}</li>
              <li>{{ t.terms.aiData3 }}</li>
              <li>{{ t.terms.aiData4 }}</li>
            </ul>

            <!-- 1.6 KVKK ve GDPR Uyumu -->
            <h3 class="sub-title">{{ t.terms.kvkkTitle }}</h3>
            <p class="card-text">{{ t.terms.kvkkDesc }}</p>
            <ul class="bullet-list">
              <li>{{ t.terms.kvkk1 }}</li>
              <li>{{ t.terms.kvkk2 }}</li>
              <li>{{ t.terms.kvkk3 }}</li>
              <li>{{ t.terms.kvkk4 }}</li>
              <li>{{ t.terms.kvkk5 }}</li>
            </ul>

            <!-- 1.7 Cerez ve Takip Politikasi -->
            <h3 class="sub-title">{{ t.terms.cookiesTitle }}</h3>
            <p class="card-text">{{ t.terms.cookiesDesc }}</p>
            <ul class="bullet-list">
              <li>{{ t.terms.cookies1 }}</li>
              <li>{{ t.terms.cookies2 }}</li>
              <li>{{ t.terms.cookies3 }}</li>
            </ul>

            <!-- 1.8 Veri Sorumlusu ve Iletisim -->
            <h3 class="sub-title">{{ t.terms.contactTitle }}</h3>
            <p class="card-text">{{ t.terms.contactDesc }}</p>
            <div class="contact-box">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/><polyline points="22,6 12,13 2,6"/></svg>
              <span>yigitkabak@tuta.io</span>
            </div>
          </div>
        </div>

        <!-- ============ 2. KULLANIM SARTLARI ============ -->
        <div class="card-section">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/><polyline points="10 9 9 9 8 9"/></svg>
              {{ t.terms.termsTitle }}
            </h2>
            <p class="card-text">{{ t.terms.termsDesc }}</p>

            <h3 class="sub-title">{{ t.terms.usageParams }}</h3>
            <p class="card-text">{{ t.terms.usageParamsDesc }}</p>

            <h3 class="sub-title">{{ t.terms.ipTitle }}</h3>
            <p class="card-text">{{ t.terms.ipDesc }}</p>

            <h3 class="sub-title">{{ t.terms.aiTitle }}</h3>
            <p class="card-text">{{ t.terms.aiDesc }}</p>

            <h3 class="sub-title">{{ t.terms.jurisdictionTitle }}</h3>
            <p class="card-text">{{ t.terms.jurisdictionDesc }}</p>
          </div>
        </div>

        <!-- ============ 3. KULLANICI HAKLARI (KVKK m.11) ============ -->
        <div class="card-section">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
              {{ t.terms.userRightsTitle }}
            </h2>
            <p class="card-text">{{ t.terms.userRightsDesc }}</p>
            <ul class="bullet-list">
              <li>{{ t.terms.userRights1 }}</li>
              <li>{{ t.terms.userRights2 }}</li>
              <li>{{ t.terms.userRights3 }}</li>
              <li>{{ t.terms.userRights4 }}</li>
              <li>{{ t.terms.userRights5 }}</li>
              <li>{{ t.terms.userRights6 }}</li>
            </ul>
          </div>
        </div>

        <!-- ============ 4. LISANS ============ -->
        <div class="card-section">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="16 18 22 12 16 6"/><polyline points="8 6 2 12 8 18"/></svg>
              {{ t.terms.licenseTitle }}
            </h2>
            <h3 class="sub-title">BSD 3-Clause License</h3>
            <div class="license-block">
              <pre>BSD 3-Clause License

Copyright (c) 2023, openbytelabs

Redistribution and use in source and binary forms,
with or without modification, are permitted provided
that the following conditions are met:

1. Redistributions of source code must retain the
   above copyright notice, this list of conditions
   and the following disclaimer.

2. Redistributions in binary form must reproduce the
   above copyright notice, this list of conditions
   and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the
   names of its contributors may be used to endorse
   or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS
AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED
WARRANTIES ARE DISCLAIMED.</pre>
            </div>
          </div>
        </div>

        <!-- ============ 5. TEKNIK ALTYAPI ============ -->
        <div class="card-section">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="8" rx="2" ry="2"/><rect x="2" y="14" width="20" height="8" rx="2" ry="2"/><line x1="6" y1="6" x2="6.01" y2="6"/><line x1="6" y1="18" x2="6.01" y2="18"/></svg>
              {{ t.terms.techTitle }}
            </h2>
            <ul class="bullet-list">
              <li><strong>Frontend:</strong> Vue.js 3 (Composition API)</li>
              <li><strong>Backend:</strong> Golang</li>
              <li><strong>{{ t.terms.indexingLabel }}:</strong> {{ t.terms.proprietaryArch }}</li>
              <li><strong>AI Providers:</strong> Groq, OpenRouter (yalnızca açıkça tetiklenen sorgularda)</li>
            </ul>
          </div>
        </div>

        <div class="tech-credits">
          <p class="credits-text">{{ t.terms.madeWith }}</p>
          <div class="tech-logos">
            <a href="https://www.linux.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link">
              <img src="https://upload.wikimedia.org/wikipedia/commons/3/35/Tux.svg" alt="Linux" class="tech-logo" />
            </a>
            <a href="https://vuejs.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link">
              <img src="https://upload.wikimedia.org/wikipedia/commons/9/95/Vue.js_Logo_2.svg" alt="Vue.js" class="tech-logo" />
            </a>
            <a href="https://nodejs.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link">
              <img src="https://upload.wikimedia.org/wikipedia/commons/d/d9/Node.js_logo.svg" alt="Node.js" class="tech-logo" />
            </a>
            <a href="https://golang.org/" target="_blank" rel="noopener noreferrer" class="tech-logo-link">
              <img src="https://duckduckgo.com/i/77cac52781feeb4a.png" alt="Golang" class="tech-logo" />
            </a>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import { lang, detectLanguage, tr, en } from '../composables/useLocale.js'
import { applyTheme, applyFontSize } from '../composables/settings.js'

const t = computed(() => lang.value === 'tr' ? tr : en)

onMounted(async () => {
  await detectLanguage()
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

.terms-page {
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

.content {
  flex: 1;
  max-width: 800px;
  width: 100%;
  margin: 0 auto;
  padding: 32px 28px 80px;
}

.content-inner {
  animation: fadeUp 0.3s ease both;
}

@keyframes fadeUp {
  from { opacity: 0; transform: translateY(8px); }
  to { opacity: 1; transform: translateY(0); }
}

.main-title {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 16px;
  letter-spacing: -0.02em;
}

.subtitle {
  color: var(--text-secondary);
  font-size: 0.9rem;
  margin-bottom: 12px;
  line-height: 1.7;
}

.meta-line {
  color: var(--text-muted);
  font-size: 0.78rem;
  margin-bottom: 24px;
  font-style: italic;
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
  margin-bottom: 16px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.card-title svg {
  flex-shrink: 0;
  color: var(--text-muted);
}

.card-text {
  color: var(--text-secondary);
  font-size: 0.88rem;
  line-height: 1.7;
  margin-bottom: 12px;
}

.card-text.emphasis {
  background: var(--bg-hover);
  border-left: 3px solid var(--accent-color, #4a9eff);
  padding: 10px 14px;
  border-radius: 6px;
  font-size: 0.85rem;
}

.sub-title {
  font-size: 0.92rem;
  font-weight: 600;
  color: var(--text-primary);
  margin: 20px 0 10px;
}

.sub-title:first-of-type {
  margin-top: 0;
}

.bullet-list {
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-top: 8px;
  margin-bottom: 12px;
}

.bullet-list li {
  color: var(--text-secondary);
  font-size: 0.87rem;
  line-height: 1.6;
  padding-left: 20px;
  position: relative;
}

.bullet-list li::before {
  content: "";
  position: absolute;
  left: 6px;
  top: 10px;
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: var(--text-muted);
}

.bullet-list strong {
  color: var(--text-primary);
}

.contact-box {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 16px;
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-secondary);
  font-size: 0.87rem;
  font-family: 'DM Mono', 'Courier New', monospace;
}

.contact-box svg {
  color: var(--text-muted);
  flex-shrink: 0;
}

.license-block {
  background: var(--bg-hover);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  padding: 16px;
  overflow-x: auto;
  margin-top: 12px;
}

.license-block pre {
  font-family: 'DM Mono', 'Courier New', monospace;
  font-size: 0.75rem;
  color: var(--text-secondary);
  line-height: 1.6;
  white-space: pre-wrap;
  word-wrap: break-word;
  margin: 0;
}

.tech-credits {
  margin-top: 32px;
  text-align: center;
  padding: 24px 0 8px;
}

.credits-text {
  color: var(--text-muted);
  font-size: 0.85rem;
  margin-bottom: 16px;
}

.tech-logos {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 2.5rem;
  flex-wrap: wrap;
}

.tech-logo-link {
  transition: transform 0.3s ease, opacity 0.3s ease;
  display: inline-block;
  text-decoration: none;
}

.tech-logo-link:hover {
  transform: scale(1.1);
  opacity: 0.8;
}

.tech-logo {
  height: 48px;
  width: auto;
  filter: brightness(0.85);
}

@media (max-width: 768px) {
  .content {
    padding: 24px 16px 60px;
  }

  .main-title {
    font-size: 1.5rem;
  }

  .card {
    padding: 18px;
  }

  .tech-logo {
    height: 40px;
  }

  .tech-logos {
    gap: 2rem;
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

  .main-title {
    font-size: 1.3rem;
  }

  .card {
    padding: 14px;
  }

  .card-title {
    font-size: 0.95rem;
  }

  .card-text,
  .bullet-list li {
    font-size: 0.82rem;
  }

  .license-block {
    padding: 12px;
  }

  .license-block pre {
    font-size: 0.7rem;
  }

  .tech-logos {
    gap: 1.5rem;
  }

  .tech-logo {
    height: 34px;
  }

  .sub-title {
    font-size: 0.85rem;
  }
}

@media (hover: none) and (pointer: coarse) {
  .tech-logo-link {
    -webkit-tap-highlight-color: transparent;
  }
}
</style>