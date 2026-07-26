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
        <button v-for="tab in tabs" :key="tab.id" class="tab-btn" :class="{ active: activeTab === tab.id }" @click="activeTab = tab.id">
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
              <label class="setting-label">{{ t.settings.theme }}</label>
              <select v-model="theme" @change="setTheme($event.target.value)" class="s-select">
                <option value="dark">{{ t.settings.dark }}</option>
                <option value="light">{{ t.settings.light }}</option>
                <option value="auto">{{ t.settings.auto }}</option>
              </select>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.accentColor }}</label>
              <div class="color-swatches">
                <button v-for="c in accentPresets" :key="c.value" class="color-swatch" :class="{ active: accentColor === c.value }" :style="{ background: c.value }" @click="setAccentColor(c.value)" :title="c.name">
                  <svg v-if="accentColor === c.value" width="14" height="14" viewBox="0 0 24 24" fill="none" :stroke="getContrastText(c.value)" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"/></svg>
                </button>
                <label class="color-swatch custom-swatch" :class="{ active: isCustomAccent }" title="Custom">
                  <input type="color" :value="accentColor" @input="e => setAccentColor(e.target.value)" class="custom-color-input" />
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="#fff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/></svg>
                </label>
              </div>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.fontSize }}</label>
              <select v-model="fontSize" @change="setFontSize($event.target.value)" class="s-select">
                <option value="13px">{{ t.settings.small }} (13px)</option>
                <option value="15px">{{ t.settings.medium }} (15px)</option>
                <option value="17px">{{ t.settings.large }} (17px)</option>
                <option value="19px">XL (19px)</option>
              </select>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.fontFamily }}</label>
              <select v-model="fontFamily" @change="setFontFamily($event.target.value)" class="s-select">
                <optgroup :label="t.settings.sansSerif">
                  <option value="sans">System UI</option>
                  <option value="inter">Inter</option>
                  <option value="geist">Geist</option>
                  <option value="dm">DM Sans</option>
                </optgroup>
                <optgroup :label="t.settings.serif">
                  <option value="serif">Georgia</option>
                  <option value="playfair">Playfair Display</option>
                </optgroup>
                <optgroup :label="t.settings.monospace">
                  <option value="mono">Courier New</option>
                  <option value="fira">Fira Code</option>
                </optgroup>
                <optgroup v-if="customFonts.length" :label="t.settings.customFontsLabel">
                  <option v-for="f in customFonts" :key="f.value" :value="f.value">{{ f.name }}</option>
                </optgroup>
              </select>
              <div class="font-upload-area">
                <input ref="fontFileInput" type="file" accept=".ttf,.otf,.woff,.woff2" class="hidden-input" @change="handleFontUpload" />
                <button class="upload-font-btn" @click="$refs.fontFileInput.click()">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="17 8 12 3 7 8"/><line x1="12" y1="3" x2="12" y2="15"/></svg>
                  {{ t.settings.uploadFont }}
                </button>
                <span class="upload-hint">{{ t.settings.uploadFontHint }}</span>
              </div>
              <div v-if="customFonts.length" class="custom-font-list">
                <div v-for="f in customFonts" :key="f.value" class="custom-font-item">
                  <span class="custom-font-name">{{ f.name }}</span>
                  <button class="icon-btn danger" @click="removeCustomFont(f.value)" :title="t.settings.delete">
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
                  </button>
                </div>
              </div>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.borderRadius }}</label>
              <select v-model="borderRadius" @change="setBorderRadius($event.target.value)" class="s-select">
                <option value="sharp">{{ t.settings.sharp }}</option>
                <option value="rounded">{{ t.settings.rounded }}</option>
                <option value="pill">{{ t.settings.pill }}</option>
              </select>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.searchBarStyle }}</label>
              <select v-model="searchBarStyle" @change="setSearchBarStyle($event.target.value)" class="s-select">
                <option value="default">{{ t.settings.defaultStyle }}</option>
                <option value="outlined">{{ t.settings.outlined }}</option>
                <option value="filled">{{ t.settings.filled }}</option>
                <option value="glass">{{ t.settings.glass }}</option>
              </select>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.homepageLayout }}</label>
              <select v-model="homepageLayout" @change="setHomepageLayout($event.target.value)" class="s-select">
                <option value="centered">{{ t.settings.centered }}</option>
                <option value="minimal">{{ t.settings.minimal }}</option>
                <option value="compact">{{ t.settings.compact }}</option>
              </select>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.cardShadow }}</label>
              <select v-model="cardShadow" @change="setCardShadow($event.target.value)" class="s-select">
                <option value="none">{{ t.settings.none }}</option>
                <option value="light">{{ t.settings.lightShadow }}</option>
                <option value="medium">{{ t.settings.mediumShadow }}</option>
                <option value="heavy">{{ t.settings.heavyShadow }}</option>
              </select>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.animationSpeed }}</label>
              <select v-model="animationSpeed" @change="setAnimationSpeed($event.target.value)" class="s-select">
                <option value="none">{{ t.settings.animNone }}</option>
                <option value="fast">{{ t.settings.fast }}</option>
                <option value="normal">{{ t.settings.normalSpeed }}</option>
                <option value="slow">{{ t.settings.slow }}</option>
              </select>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.backgroundStyle }}</label>
              <select v-model="backgroundStyle" @change="setBackgroundStyle($event.target.value)" class="s-select">
                <option value="plain">{{ t.settings.bgPlain }}</option>
                <option value="dots">{{ t.settings.bgDots }}</option>
              </select>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.density }}</label>
              <select v-model="density" @change="setDensity($event.target.value)" class="s-select">
                <option value="compact">{{ t.settings.compact }}</option>
                <option value="normal">{{ t.settings.normalSpeed }}</option>
                <option value="relaxed">{{ t.settings.relaxed }}</option>
              </select>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.logoSize }}</label>
              <select v-model="logoSize" @change="setLogoSize($event.target.value)" class="s-select">
                <option value="small">{{ t.settings.smallSize }}</option>
                <option value="medium">{{ t.settings.mediumSize }}</option>
                <option value="large">{{ t.settings.largeSize }}</option>
              </select>
            </div>

            <div class="divider"></div>

            <div class="toggle-row" @click="setBlur(!blurEnabled)">
              <div class="toggle-info">
                <span class="toggle-title">{{ t.settings.blur }}</span>
                <span class="toggle-desc">{{ t.settings.blurDesc }}</span>
              </div>
              <div class="toggle-switch" :class="{ active: blurEnabled }"><div class="toggle-knob"></div></div>
            </div>

            <div class="toggle-row" @click="setReducedMotion(!reducedMotion)">
              <div class="toggle-info">
                <span class="toggle-title">{{ t.settings.reducedMotion }}</span>
                <span class="toggle-desc">{{ t.settings.reducedMotionDesc }}</span>
              </div>
              <div class="toggle-switch" :class="{ active: reducedMotion }"><div class="toggle-knob"></div></div>
            </div>

            <div class="toggle-row" @click="setHighlight(!highlightEnabled)">
              <div class="toggle-info">
                <span class="toggle-title">{{ t.settings.highlight }}</span>
                <span class="toggle-desc">{{ t.settings.highlightDesc }}</span>
              </div>
              <div class="toggle-switch" :class="{ active: highlightEnabled }"><div class="toggle-knob"></div></div>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.resultDensity }}</label>
              <p class="setting-desc">{{ t.settings.resultDensityDesc }}</p>
              <select v-model="resultDensity" @change="setResultDensity($event.target.value)" class="s-select">
                <option value="tight">{{ t.settings.resultDensityTight }}</option>
                <option value="normal">{{ t.settings.resultDensityNormal }}</option>
                <option value="spacious">{{ t.settings.resultDensitySpacious }}</option>
              </select>
            </div>

            <div class="divider"></div>

            <button class="reset-btn" @click="resetAppearance">{{ t.settings.resetDefaults }}</button>
          </div>
        </section>

        <section v-else-if="activeTab === 'search'" class="panel">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>
              {{ t.settings.searchSettings }}
            </h2>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.aiMode }}</label>
              <p class="setting-desc">{{ t.settings.aiModeDesc }}</p>
              <label class="setting-label" style="margin-top: 4px; font-size: 0.78rem; color: var(--text-muted);">{{ t.settings.searchType }}</label>
              <select v-model="aiMode" @change="setAIMode($event.target.value)" class="s-select">
                <option value="auto">{{ t.settings.aiModeAuto }}</option>
                <option value="full_text">{{ t.settings.aiModeFullText }}</option>
                <option value="graphrag">{{ t.settings.aiModeGraphRAG }}</option>
                <option value="hybrid">{{ t.settings.aiModeHybrid }}</option>
              </select>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.safeSearch }}</label>
              <p class="setting-desc">{{ t.settings.safeSearchDesc }}</p>
              <select v-model="safeSearch" @change="setSafeSearch($event.target.value)" class="s-select">
                <option value="strict">{{ t.settings.strict }}</option>
                <option value="moderate">{{ t.settings.moderate }}</option>
                <option value="off">{{ t.settings.off }}</option>
              </select>
            </div>

            <div class="setting-row">
              <label class="setting-label">{{ t.settings.resultsPerPage }}</label>
              <select v-model="resultsPerPage" @change="setResultsPerPage($event.target.value)" class="s-select">
                <option value="10">10</option>
                <option value="15">15</option>
                <option value="25">25</option>
                <option value="50">50</option>
              </select>
            </div>

            <div class="divider"></div>

            <div class="toggle-row" @click="setOpenInNewTab(!openInNewTab)">
              <div class="toggle-info">
                <span class="toggle-title">{{ t.settings.openInNewTab }}</span>
                <span class="toggle-desc">{{ t.settings.openInNewTabDesc }}</span>
              </div>
              <div class="toggle-switch" :class="{ active: openInNewTab }"><div class="toggle-knob"></div></div>
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
                <input v-model="newBangShortcut" type="text" class="form-input" :placeholder="t.settings.bangShortcut" @keyup.enter="handleAddBang" />
                <input v-model="newBangUrl" type="text" class="form-input wide" placeholder="https://example.com/search?q={query}" @keyup.enter="handleAddBang" />
                <button class="add-btn" @click="handleAddBang">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/></svg>
                </button>
              </div>
              <p class="form-msg" v-if="bangError" :class="{ error: bangError }">{{ bangError }}</p>
            </div>
          </div>
        </section>

        <section v-else-if="activeTab === 'prisms'" class="panel">
          <div class="card">
            <h2 class="card-title">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polygon points="12 2 2 19 22 19"/>
                <line x1="12" y1="8" x2="12" y2="13"/>
                <line x1="12" y1="16" x2="12.01" y2="16"/>
              </svg>
              {{ t.settings.prisms }}
            </h2>
            <p class="card-desc">{{ t.settings.prismDesc }}</p>

            <div class="sub-section">
              <h3 class="sub-title">{{ t.settings.systemPrisms }}</h3>
              <div class="prism-grid">
                <div
                  v-for="prism in systemPrisms"
                  :key="prism.id"
                  class="prism-card prism-card-custom"
                  :class="{
                    active: activePrismId === prism.id,
                    'prism-card-hidden': hiddenSystemPrismIds.includes(prism.id)
                  }"
                  :style="activePrismId === prism.id ? { borderColor: prism.color, background: prism.color + '12' } : {}"
                >
                  <button class="prism-card-main" @click="handlePrismToggle(prism.id)" :disabled="hiddenSystemPrismIds.includes(prism.id)">
                    <span class="prism-name">{{ lang === 'tr' ? prism.nametr : prism.name }}</span>
                    <span v-if="activePrismId === prism.id" class="prism-active-badge" :style="{ color: prism.color, borderColor: prism.color + '40', background: prism.color + '15' }">{{ t.settings.activePrism }}</span>
                    <div class="prism-meta">
                      <span v-if="prism.tlds.length" class="prism-tag">{{ prism.tlds.slice(0,3).join(' ') }}</span>
                      <span v-if="prism.domains.length" class="prism-tag">{{ prism.domains.length }} sites</span>
                    </div>
                  </button>
                  <button
                    v-if="prism.id !== 'default' && !hiddenSystemPrismIds.includes(prism.id)"
                    class="prism-delete-btn"
                    @click="hideSystemPrism(prism.id)"
                    :title="t.settings.delete"
                  >
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
                  </button>
                </div>
              </div>
            </div>

            <div class="sub-section">
              <h3 class="sub-title">{{ t.settings.customPrisms }}</h3>
              <div class="prism-grid" v-if="customPrisms.length">
                <div
                  v-for="prism in customPrisms"
                  :key="prism.id"
                  class="prism-card prism-card-custom"
                  :class="{ active: activePrismId === prism.id }"
                  :style="activePrismId === prism.id ? { borderColor: prism.color, background: prism.color + '12' } : {}"
                >
                  <button class="prism-card-main" @click="handlePrismToggle(prism.id)">
                    <span class="prism-name">{{ prism.name }}</span>
                    <span v-if="activePrismId === prism.id" class="prism-active-badge" :style="{ color: prism.color, borderColor: prism.color + '40', background: prism.color + '15' }">{{ t.settings.activePrism }}</span>
                    <div class="prism-meta">
                      <span v-if="prism.tlds.length" class="prism-tag">{{ prism.tlds.slice(0,3).join(' ') }}</span>
                      <span v-if="prism.domains.length" class="prism-tag">{{ prism.domains.length }} sites</span>
                    </div>
                  </button>
                  <button class="prism-delete-btn" @click="handleRemovePrism(prism.id)" :title="t.settings.delete">
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
                  </button>
                </div>
              </div>
              <div class="empty-state" v-else>{{ t.settings.noPrisms }}</div>
            </div>

            <div class="add-form">
              <h3 class="sub-title">{{ t.settings.addPrism }}</h3>
              <div class="form-row">
                <input v-model="newPrismName" type="text" class="form-input wide" :placeholder="t.settings.prismName" @keyup.enter="handleAddPrism" />
                <div class="color-pick-row">
                  <button v-for="c in prismColors" :key="c" class="prism-color-dot" :class="{ active: newPrismColor === c }" :style="{ background: c }" @click="newPrismColor = c"></button>
                </div>
              </div>
              <div class="prism-filter-row">
                <div class="prism-filter-col">
                  <label class="setting-label">{{ t.settings.prismDomains }}</label>
                  <input v-model="newPrismDomains" type="text" class="form-input" :placeholder="t.settings.prismDomainsHint" />
                </div>
                <div class="prism-filter-col">
                  <label class="setting-label">{{ t.settings.prismTlds }}</label>
                  <input v-model="newPrismTlds" type="text" class="form-input" :placeholder="t.settings.prismTldsHint" />
                </div>
              </div>
              <p class="form-msg error" v-if="prismError">{{ prismError }}</p>
              <button class="create-prism-btn" @click="handleAddPrism" :style="{ background: newPrismColor }">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/></svg>
                {{ t.settings.addPrism }}
              </button>
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
                  <button v-if="site.type !== 'block'" class="icon-btn danger" @click="changeSiteType(idx, 'block')" title="Block">
                    <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"/></svg>
                  </button>
                  <button v-if="site.type !== 'restrict'" class="icon-btn warn" @click="changeSiteType(idx, 'restrict')" title="Restrict">
                    <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"/></svg>
                  </button>
                  <button class="icon-btn success" @click="promoteSite(site.url)" :title="t.settings.promoteSite">
                    <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="19" x2="12" y2="5"/><polyline points="5 12 12 5 19 12"/></svg>
                  </button>
                  <button class="icon-btn danger" @click="removeBlockedSite(site.url)" :title="t.settings.delete">
                    <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"/><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/></svg>
                  </button>
                </div>
              </div>
            </div>
            <div class="empty-state" v-else>{{ t.settings.noBlockedSites }}</div>

            <div class="add-form">
              <h3 class="sub-title">{{ t.settings.addSite }}</h3>
              <div class="form-row">
                <input v-model="newSiteUrl" type="text" class="form-input" :placeholder="t.settings.siteUrl" @keyup.enter="handleAddSite('block')" />
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
            <div class="lang-group">
              <button class="lang-opt" :class="{ active: lang === 'tr' }" @click="setLanguageSetting('tr')">
                <span class="lang-flag"></span> {{ t.settings.turkish }}
              </button>
              <button class="lang-opt" :class="{ active: lang === 'en' }" @click="setLanguageSetting('en')">
                <span class="lang-flag"></span> {{ t.settings.english }}
              </button>
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
import { usePrisms } from '../composables/usePrisms.js'
import { applyTheme, applyFontFamily, applyAccentColor, applyBgStyle, registerFontFace, loadCustomFonts } from '../composables/settings.js'

const { bangs, customBangs, deletedBangs, addBang, removeBang, restoreBang } = useBangs()
const { blockedSites, addBlockedSite, removeBlockedSite, promoteSite } = useBlockedSites()
const { allSystemPrisms, customPrisms, hiddenSystemPrismIds, activePrismId, setActivePrism, addPrism, removePrism, hideSystemPrism, PRISM_COLORS } = usePrisms()

const activeTab = ref('appearance')
const fontSize = ref('15px')
const theme = ref('dark')
const accentColor = ref('#3b82f6')
const fontFamily = ref('sans')
const borderRadius = ref('rounded')
const searchBarStyle = ref('default')
const homepageLayout = ref('centered')
const cardShadow = ref('medium')
const animationSpeed = ref('normal')
const backgroundStyle = ref('plain')
const density = ref('normal')
const logoSize = ref('medium')
const blurEnabled = ref(true)
const reducedMotion = ref(false)
const resultsPerPage = ref('15')
const safeSearch = ref('moderate')
const openInNewTab = ref(false)
const customFonts = ref([])
const newBangShortcut = ref('')
const newBangUrl = ref('')
const bangError = ref('')
const newSiteUrl = ref('')
const fontFileInput = ref(null)
const aiMode = ref('auto')
const highlightEnabled = ref(true)
const resultDensity = ref('normal')

const newPrismName = ref('')
const newPrismColor = ref(PRISM_COLORS[0])
const newPrismDomains = ref('')
const newPrismTlds = ref('')
const prismError = ref('')

const t = computed(() => lang.value === 'tr' ? tr : en)
const prismColors = PRISM_COLORS

const systemPrisms = computed(() => allSystemPrisms.value.filter(p => !hiddenSystemPrismIds.value.includes(p.id)))

function getContrastText(hex) {
  if (!hex || !hex.startsWith('#')) return '#ffffff'
  const c = hex.replace('#', '')
  const r = parseInt(c.substring(0, 2), 16) / 255
  const g = parseInt(c.substring(2, 4), 16) / 255
  const b = parseInt(c.substring(4, 6), 16) / 255
  return (0.299 * r + 0.587 * g + 0.114 * b) > 0.55 ? '#000000' : '#ffffff'
}

const accentPresets = [
  { value: '#3b82f6', name: 'Blue' },
  { value: '#6366f1', name: 'Indigo' },
  { value: '#8b5cf6', name: 'Violet' },
  { value: '#a855f7', name: 'Purple' },
  { value: '#d946ef', name: 'Fuchsia' },
  { value: '#ec4899', name: 'Pink' },
  { value: '#f43f5e', name: 'Rose' },
  { value: '#ef4444', name: 'Red' },
  { value: '#f97316', name: 'Orange' },
  { value: '#eab308', name: 'Yellow' },
  { value: '#22c55e', name: 'Green' },
  { value: '#14b8a6', name: 'Teal' },
  { value: '#06b6d4', name: 'Cyan' },
  { value: '#ffffff', name: 'White' },
]

const isCustomAccent = computed(() => !accentPresets.find(p => p.value === accentColor.value))

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

const activeDefaultBangs = computed(() => DEFAULT_BANGS_LIST.filter(b => !deletedBangs.value.includes(b.shortcut)))

const setFontSize = (v) => { fontSize.value = v; document.documentElement.style.fontSize = v; localStorage.setItem('synapic_font_size', v) }
const setTheme = (v) => { theme.value = v; localStorage.setItem('synapic_theme', v); applyTheme() }
const setBorderRadius = (v) => { borderRadius.value = v; localStorage.setItem('synapic_border_radius', v); document.documentElement.setAttribute('data-border-radius', v) }
const setSearchBarStyle = (v) => { searchBarStyle.value = v; localStorage.setItem('synapic_search_bar_style', v); document.documentElement.setAttribute('data-search-bar-style', v) }
const setHomepageLayout = (v) => { homepageLayout.value = v; localStorage.setItem('synapic_homepage_layout', v); document.documentElement.setAttribute('data-homepage-layout', v) }
const setCardShadow = (v) => { cardShadow.value = v; localStorage.setItem('synapic_card_shadow', v); document.documentElement.setAttribute('data-card-shadow', v) }
const setAnimationSpeed = (v) => { animationSpeed.value = v; localStorage.setItem('synapic_animation_speed', v); document.documentElement.setAttribute('data-animation-speed', v) }
const setLogoSize = (v) => { logoSize.value = v; localStorage.setItem('synapic_logo_size', v); document.documentElement.setAttribute('data-logo-size', v) }
const setDensity = (v) => { density.value = v; localStorage.setItem('synapic_density', v); document.documentElement.setAttribute('data-density', v) }
const setBlur = (v) => { blurEnabled.value = v; localStorage.setItem('synapic_blur', v ? '1' : '0'); document.documentElement.setAttribute('data-blur', v ? '1' : '0') }
const setReducedMotion = (v) => { reducedMotion.value = v; localStorage.setItem('synapic_reduced_motion', v ? '1' : '0'); document.documentElement.setAttribute('data-reduced-motion', v ? '1' : '0') }
const setResultsPerPage = (v) => { resultsPerPage.value = v; localStorage.setItem('synapic_results_per_page', v) }
const setSafeSearch = (v) => { safeSearch.value = v; localStorage.setItem('synapic_safe_search', v) }
const setOpenInNewTab = (v) => { openInNewTab.value = v; localStorage.setItem('synapic_open_new_tab', v ? '1' : '0') }
const setAIMode = (v) => { aiMode.value = v; localStorage.setItem('synapic_ai_mode', v) }
const setHighlight = (v) => { highlightEnabled.value = v; localStorage.setItem('synapic_highlight', v ? '1' : '0'); document.documentElement.setAttribute('data-highlight', v ? '1' : '0') }
const setResultDensity = (v) => { resultDensity.value = v; localStorage.setItem('synapic_result_density', v); document.documentElement.setAttribute('data-result-density', v) }

const setAccentColor = (color) => {
  accentColor.value = color
  localStorage.setItem('synapic_accent_color', color)
  applyAccentColor()
}

const setBackgroundStyle = (v) => {
  backgroundStyle.value = v
  localStorage.setItem('synapic_bg_style', v)
  applyBgStyle()
}

const setFontFamily = (family) => {
  fontFamily.value = family
  localStorage.setItem('synapic_font_family', family)
  applyFontFamily()
}

const handleFontUpload = (event) => {
  const file = event.target.files[0]
  if (!file) return
  const maxSize = 2 * 1024 * 1024
  if (file.size > maxSize) { alert(t.value.settings.fontUploadError); event.target.value = ''; return }
  const ext = file.name.split('.').pop().toLowerCase()
  if (!['ttf', 'otf', 'woff', 'woff2'].includes(ext)) { alert(t.value.settings.fontUploadError); event.target.value = ''; return }
  const fontName = file.name.replace(/\.[^/.]+$/, '')
  const fontValue = 'custom_' + fontName.toLowerCase().replace(/[^a-z0-9]/g, '_')
  if (customFonts.value.find(f => f.value === fontValue)) { alert(t.value.settings.fontNameExists); event.target.value = ''; return }
  const reader = new FileReader()
  reader.onload = (e) => {
    const base64 = e.target.result.split(',')[1]
    const fontData = { name: fontName, value: fontValue, format: ext, base64 }
    customFonts.value.push(fontData)
    try {
      localStorage.setItem('synapic_custom_fonts', JSON.stringify(customFonts.value))
    } catch (err) {
      customFonts.value.pop()
      alert(t.value.settings.fontUploadError + ' (LocalStorage limit exceeded)')
      return
    }
    registerFontFace(fontName, base64, ext)
    fontFamily.value = fontValue
    localStorage.setItem('synapic_font_family', fontValue)
    applyFontFamily()
  }
  reader.readAsDataURL(file)
  event.target.value = ''
}

const removeCustomFont = (value) => {
  customFonts.value = customFonts.value.filter(f => f.value !== value)
  try { localStorage.setItem('synapic_custom_fonts', JSON.stringify(customFonts.value)) } catch (e) { console.error(e) }
  if (fontFamily.value === value) setFontFamily('sans')
  const id = 'synapic-font-' + value.replace('custom_', '')
  const el = document.getElementById(id)
  if (el) el.remove()
}

const resetAppearance = () => {
  setTheme('dark'); setFontSize('15px'); setAccentColor('#3b82f6'); setFontFamily('sans')
  setBorderRadius('rounded'); setSearchBarStyle('default'); setHomepageLayout('centered')
  setCardShadow('medium'); setAnimationSpeed('normal'); setLogoSize('medium')
  setBackgroundStyle('plain'); setDensity('normal'); setBlur(true); setReducedMotion(false)
  setResultsPerPage('15'); setSafeSearch('moderate'); setOpenInNewTab(false); setAIMode('auto')
}

const handleAddBang = () => {
  bangError.value = ''
  const shortcut = newBangShortcut.value.trim()
  const url = newBangUrl.value.trim()
  if (!shortcut || !url) { bangError.value = lang.value === 'tr' ? 'Kısayol ve URL gerekli.' : 'Shortcut and URL are required.'; return }
  if (!url.includes('{query}')) { bangError.value = lang.value === 'tr' ? 'URL {query} içermeli.' : 'URL must contain {query}.'; return }
  if (bangs.value.find(b => b.shortcut === shortcut)) { bangError.value = lang.value === 'tr' ? 'Bu kısayol zaten mevcut.' : 'This shortcut already exists.'; return }
  try { addBang({ shortcut, url, site: new URL(url.replace('{query}', 'example')).hostname.replace('www.', '') }) }
  catch { addBang({ shortcut, url, site: url }) }
  newBangShortcut.value = ''; newBangUrl.value = ''
}

const handlePrismToggle = (id) => {
  if (activePrismId.value === id) {
    setActivePrism('default')
  } else {
    setActivePrism(id)
  }
}

const handleAddPrism = () => {
  prismError.value = ''
  const name = newPrismName.value.trim()
  if (!name) { prismError.value = t.value.settings.prismNameRequired; return }
  const rawDomains = newPrismDomains.value.trim()
  const rawTlds = newPrismTlds.value.trim()
  const domains = rawDomains ? rawDomains.split(/[\s,]+/).map(d => d.trim().replace(/^www\./, '').toLowerCase()).filter(Boolean) : []
  const tlds = rawTlds ? rawTlds.split(/[\s,]+/).map(t => t.trim().startsWith('.') ? t.trim() : '.' + t.trim()).filter(Boolean) : []
  if (domains.length === 0 && tlds.length === 0) { prismError.value = t.value.settings.prismNeedFilter; return }
  addPrism({ name, nametr: name, color: newPrismColor.value, domains, tlds })
  newPrismName.value = ''
  newPrismColor.value = PRISM_COLORS[0]
  newPrismDomains.value = ''
  newPrismTlds.value = ''
}

const handleRemovePrism = (id) => {
  removePrism(id)
}

const changeSiteType = (index, newType) => { const site = blockedSites.value[index]; if (site) addBlockedSite(site.url, newType) }
const handleAddSite = (type) => { const url = newSiteUrl.value.trim(); if (!url) return; addBlockedSite(url, type); newSiteUrl.value = '' }
const setLanguageSetting = (v) => setLanguage(v)

const AppearanceIcon = { render() { return h('svg', { width: 15, height: 15, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2, 'stroke-linecap': 'round', 'stroke-linejoin': 'round' }, [h('circle', { cx: 13.5, cy: 6.5, r: 0.5 }), h('circle', { cx: 17.5, cy: 10.5, r: 0.5 }), h('circle', { cx: 8.5, cy: 7.5, r: 0.5 }), h('circle', { cx: 6.5, cy: 12.5, r: 0.5 }), h('path', { d: 'M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10c.926 0 1.648-.746 1.648-1.688 0-.437-.18-.835-.437-1.125-.29-.289-.438-.652-.438-1.125a1.64 1.64 0 0 1 1.668-1.668h1.996c3.051 0 5.555-2.503 5.555-5.554C21.965 6.012 17.461 2 12 2z' })]) } }
const SearchIcon = { render() { return h('svg', { width: 15, height: 15, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2, 'stroke-linecap': 'round', 'stroke-linejoin': 'round' }, [h('circle', { cx: 11, cy: 11, r: 8 }), h('line', { x1: 21, y1: 21, x2: 16.65, y2: 16.65 })]) } }
const BangsIcon = { render() { return h('svg', { width: 15, height: 15, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2, 'stroke-linecap': 'round', 'stroke-linejoin': 'round' }, [h('polygon', { points: '13 2 3 14 12 14 11 22 21 10 12 10 13 2' })]) } }
const PrismsIcon = { render() { return h('svg', { width: 15, height: 15, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2, 'stroke-linecap': 'round', 'stroke-linejoin': 'round' }, [h('polygon', { points: '12 2 2 19 22 19' }), h('line', { x1: 12, y1: 8, x2: 12, y2: 13 }), h('line', { x1: 12, y1: 16, x2: 12.01, y2: 16 })]) } }
const BlockedIcon = { render() { return h('svg', { width: 15, height: 15, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2, 'stroke-linecap': 'round', 'stroke-linejoin': 'round' }, [h('circle', { cx: 12, cy: 12, r: 10 }), h('line', { x1: 4.93, y1: 4.93, x2: 19.07, y2: 19.07 })]) } }
const LanguageIcon = { render() { return h('svg', { width: 15, height: 15, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2, 'stroke-linecap': 'round', 'stroke-linejoin': 'round' }, [h('circle', { cx: 12, cy: 12, r: 10 }), h('line', { x1: 2, y1: 12, x2: 22, y2: 12 }), h('path', { d: 'M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z' })]) } }

const tabs = computed(() => [
  { id: 'appearance', label: t.value.settings.appearance, icon: AppearanceIcon },
  { id: 'search', label: t.value.settings.searchSettings, icon: SearchIcon },
  { id: 'bangs', label: t.value.settings.bangs, icon: BangsIcon },
  { id: 'prisms', label: t.value.settings.prisms, icon: PrismsIcon },
  { id: 'blocked', label: t.value.settings.blockedSites, icon: BlockedIcon },
  { id: 'language', label: t.value.settings.language, icon: LanguageIcon }
])

onMounted(async () => {
  await detectLanguage()
  const load = (key, fallback) => localStorage.getItem(key) || fallback

  fontSize.value = load('synapic_font_size', '15px')
  document.documentElement.style.fontSize = fontSize.value
  theme.value = load('synapic_theme', 'dark')
  accentColor.value = load('synapic_accent_color', '#3b82f6')
  fontFamily.value = load('synapic_font_family', 'sans')
  borderRadius.value = load('synapic_border_radius', 'rounded')
  searchBarStyle.value = load('synapic_search_bar_style', 'default')
  homepageLayout.value = load('synapic_homepage_layout', 'centered')
  cardShadow.value = load('synapic_card_shadow', 'medium')
  animationSpeed.value = load('synapic_animation_speed', 'normal')
  backgroundStyle.value = load('synapic_bg_style', 'plain')
  density.value = load('synapic_density', 'normal')
  logoSize.value = load('synapic_logo_size', 'medium')
  blurEnabled.value = load('synapic_blur', '1') === '1'
  reducedMotion.value = load('synapic_reduced_motion', '0') === '1'
  resultsPerPage.value = load('synapic_results_per_page', '15')
  safeSearch.value = load('synapic_safe_search', 'moderate')
  openInNewTab.value = load('synapic_open_new_tab', '0') === '1'
  aiMode.value = load('synapic_ai_mode', 'auto')
  highlightEnabled.value = load('synapic_highlight', '1') === '1'
  resultDensity.value = load('synapic_result_density', 'normal')

  if (typeof window !== 'undefined' && !window.__synapicFetchPatched) {
    window.__synapicFetchPatched = true
    const originalFetch = window.fetch
    window.fetch = async (input, init = {}) => {
      try {
        const url = typeof input === 'string' ? input : input.url
        if (url && url.includes('/search') && init.method !== 'GET') {
          if (init.body && typeof init.body === 'string') {
            const body = JSON.parse(init.body)
            body.mode = localStorage.getItem('synapic_ai_mode') || 'auto'
            init.body = JSON.stringify(body)
          }
        }
      } catch (e) { /* ignore parse errors */ }
      return originalFetch(input, init)
    }
  }

  try { customFonts.value = JSON.parse(localStorage.getItem('synapic_custom_fonts') || '[]') } catch (e) { console.error(e) }

  loadCustomFonts()
  applyTheme()
  applyAccentColor()
  applyFontFamily()
  applyBgStyle()

  document.documentElement.setAttribute('data-border-radius', borderRadius.value)
  document.documentElement.setAttribute('data-search-bar-style', searchBarStyle.value)
  document.documentElement.setAttribute('data-homepage-layout', homepageLayout.value)
  document.documentElement.setAttribute('data-card-shadow', cardShadow.value)
  document.documentElement.setAttribute('data-animation-speed', animationSpeed.value)
  document.documentElement.setAttribute('data-logo-size', logoSize.value)
  document.documentElement.setAttribute('data-density', density.value)
  document.documentElement.setAttribute('data-blur', blurEnabled.value ? '1' : '0')
  document.documentElement.setAttribute('data-reduced-motion', reducedMotion.value ? '1' : '0')
  document.documentElement.setAttribute('data-highlight', highlightEnabled.value ? '1' : '0')
  document.documentElement.setAttribute('data-result-density', resultDensity.value)
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=DM+Sans:opsz,wght@9..40,300;9..40,400;9..40,500;9..40,600&display=swap');

*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

.settings-page {
  min-height: 100vh;
  background: var(--bg-primary);
  font-family: var(--font-family, 'DM Sans', system-ui, -apple-system, sans-serif);
  color: var(--text-primary);
  -webkit-font-smoothing: antialiased;
  display: flex;
  flex-direction: column;
  transition: background-color 0.3s ease, color 0.3s ease;
}

.topbar { position: sticky; top: 0; z-index: 100; background: var(--bg-primary); border-bottom: 1px solid var(--border-color); padding: 12px 20px; transition: background-color 0.3s ease, border-color 0.3s ease; }
.topbar-inner { max-width: 800px; margin: 0 auto; display: flex; align-items: center; gap: 20px; }
.logo-link { display: flex; align-items: center; gap: 8px; text-decoration: none; flex-shrink: 0; color: var(--text-primary); }
.topbar-logo { width: 28px; height: 28px; object-fit: contain; border-radius: 6px; flex-shrink: 0; }
.logo-text { font-size: 1.1rem; font-weight: 600; color: var(--text-primary); letter-spacing: -0.01em; }
.page-title { font-size: 1.1rem; font-weight: 500; color: var(--text-secondary); }

.tab-nav { position: sticky; top: 49px; z-index: 99; background: var(--bg-primary); border-bottom: 1px solid var(--border-color); transition: background-color 0.3s ease, border-color 0.3s ease; }
.tab-nav-inner { max-width: 800px; margin: 0 auto; display: flex; align-items: center; padding: 0 20px; overflow-x: auto; scrollbar-width: none; }
.tab-nav-inner::-webkit-scrollbar { display: none; }
.tab-btn { display: flex; align-items: center; gap: 6px; padding: 12px 16px; background: none; border: none; border-bottom: 2px solid transparent; color: var(--text-muted); font-family: inherit; font-size: 0.85rem; font-weight: 500; cursor: pointer; transition: color 0.2s, border-color 0.2s; white-space: nowrap; flex-shrink: 0; }
.tab-btn:hover { color: var(--text-secondary); }
.tab-btn.active { color: var(--text-primary); border-bottom-color: var(--accent); }
.tab-btn svg { flex-shrink: 0; }

.content { flex: 1; max-width: 800px; width: 100%; margin: 0 auto; padding: 24px 28px 80px; }
.panel { animation: fadeUp 0.3s ease both; }
@keyframes fadeUp { from { opacity: 0; transform: translateY(8px); } to { opacity: 1; transform: translateY(0); } }

.card { background: var(--bg-card); border: 1px solid var(--border-color); border-radius: 12px; padding: 24px; transition: background-color 0.3s ease, border-color 0.3s ease; }
.card-title { font-size: 1.05rem; font-weight: 600; color: var(--text-primary); margin-bottom: 20px; display: flex; align-items: center; gap: 10px; }
.card-title svg { flex-shrink: 0; color: var(--text-muted); }
.card-desc { color: var(--text-muted); font-size: 0.87rem; line-height: 1.6; margin-bottom: 20px; }

.setting-row { margin-bottom: 20px; }
.setting-row:last-child { margin-bottom: 0; }
.setting-label { display: block; font-size: 0.87rem; font-weight: 500; color: var(--text-secondary); margin-bottom: 8px; }
.setting-desc { font-size: 0.78rem; color: var(--text-muted); margin-bottom: 8px; line-height: 1.5; }

.s-select {
  width: 100%; padding: 10px 36px 10px 14px; background: var(--bg-hover); border: 1px solid var(--border-color);
  border-radius: 8px; color: var(--text-primary); font-size: 0.85rem; font-family: inherit; cursor: pointer;
  transition: border-color 0.2s; outline: none; appearance: none;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='12' height='12' viewBox='0 0 24 24' fill='none' stroke='%23999' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpolyline points='6 9 12 15 18 9'%3E%3C/polyline%3E%3C/svg%3E");
  background-repeat: no-repeat; background-position: right 12px center;
}
.s-select:focus { border-color: var(--accent); }
.s-select optgroup { font-weight: 600; color: var(--text-muted); }
.s-select option { background: var(--bg-card); color: var(--text-primary); }

.color-swatches { display: flex; gap: 8px; flex-wrap: wrap; align-items: center; }
.color-swatch {
  width: 32px; height: 32px; border-radius: 50%; border: 2px solid transparent; cursor: pointer;
  display: flex; align-items: center; justify-content: center; transition: all 0.2s ease; flex-shrink: 0;
  outline: none; padding: 0;
}
.color-swatch:hover { transform: scale(1.15); }
.color-swatch.active { border-color: var(--text-primary); box-shadow: 0 0 0 2px var(--bg-primary), 0 0 0 4px var(--text-muted); }
.custom-swatch { background: conic-gradient(from 0deg, #ef4444, #f97316, #eab308, #22c55e, #06b6d4, #6366f1, #ec4899, #ef4444); position: relative; overflow: hidden; }
.custom-color-input { position: absolute; inset: 0; width: 100%; height: 100%; opacity: 0; cursor: pointer; border: none; padding: 0; }
.custom-swatch svg { color: #fff; filter: drop-shadow(0 1px 2px rgba(0,0,0,0.4)); position: relative; z-index: 1; }

.font-upload-area { display: flex; align-items: center; gap: 10px; margin-top: 8px; flex-wrap: wrap; }
.hidden-input { display: none; }
.upload-font-btn {
  display: flex; align-items: center; gap: 6px; padding: 8px 14px; background: var(--bg-hover);
  border: 1px solid var(--border-color); border-radius: 8px; color: var(--text-muted); font-size: 0.82rem;
  font-weight: 500; font-family: inherit; cursor: pointer; transition: all 0.2s ease; white-space: nowrap;
}
.upload-font-btn:hover { background: var(--border-color); border-color: var(--border-light); color: var(--text-primary); }
.upload-hint { font-size: 0.75rem; color: var(--text-muted); }
.custom-font-list { display: flex; flex-direction: column; gap: 4px; margin-top: 8px; }
.custom-font-item { display: flex; align-items: center; justify-content: space-between; padding: 8px 12px; background: var(--bg-hover); border-radius: 8px; }
.custom-font-name { font-size: 0.85rem; color: var(--text-secondary); }

.divider { height: 1px; background: var(--border-color); margin: 20px 0; }

.toggle-row { display: flex; align-items: center; justify-content: space-between; padding: 12px 0; cursor: pointer; gap: 16px; }
.toggle-row + .toggle-row { border-top: 1px solid var(--border-color); }
.toggle-info { display: flex; flex-direction: column; gap: 2px; min-width: 0; }
.toggle-title { font-size: 0.87rem; font-weight: 500; color: var(--text-secondary); }
.toggle-desc { font-size: 0.78rem; color: var(--text-muted); line-height: 1.4; }
.toggle-switch { width: 44px; height: 24px; background: var(--border-color); border-radius: 12px; padding: 2px; cursor: pointer; transition: background 0.2s ease; flex-shrink: 0; }
.toggle-switch.active { background: var(--accent); }
.toggle-knob { width: 20px; height: 20px; background: var(--accent-text, #fff); border-radius: 50%; transition: transform 0.2s ease; box-shadow: 0 1px 3px rgba(0,0,0,0.2); }
.toggle-switch.active .toggle-knob { transform: translateX(20px); }

.reset-btn {
  display: flex; align-items: center; justify-content: center; width: 100%; padding: 10px 16px;
  background: var(--bg-hover); border: 1px solid var(--border-color); border-radius: 8px; color: var(--text-secondary);
  font-size: 0.85rem; font-weight: 500; font-family: inherit; cursor: pointer; transition: all 0.2s ease;
}
.reset-btn:hover { background: rgba(239, 68, 68, 0.1); border-color: rgba(239, 68, 68, 0.25); color: var(--danger); }

.lang-group { display: flex; gap: 10px; }
.lang-opt {
  flex: 1; padding: 14px 18px; background: var(--bg-hover); border: 1px solid var(--border-color); border-radius: 10px;
  color: var(--text-muted); font-size: 0.9rem; font-weight: 500; font-family: inherit; cursor: pointer;
  transition: all 0.2s ease; display: flex; align-items: center; justify-content: center; gap: 8px;
}
.lang-opt:hover { border-color: var(--border-light); background: var(--border-color); }
.lang-opt.active { background: var(--accent); border-color: var(--accent); color: var(--accent-text, #fff); }
.lang-flag { font-size: 1.2rem; }

.sub-section { margin-bottom: 20px; padding-bottom: 20px; border-bottom: 1px solid var(--border-color); }
.sub-section:last-of-type { border-bottom: none; margin-bottom: 0; padding-bottom: 0; }
.sub-title { font-size: 0.9rem; font-weight: 600; color: var(--text-primary); margin-bottom: 12px; }
.sub-restore { color: var(--text-secondary); }

.item-list { display: flex; flex-direction: column; gap: 4px; max-height: 320px; overflow-y: auto; }
.item-list::-webkit-scrollbar { width: 4px; }
.item-list::-webkit-scrollbar-track { background: transparent; }
.item-list::-webkit-scrollbar-thumb { background: var(--scrollbar-thumb); border-radius: 10px; }

.item-row { display: flex; align-items: center; justify-content: space-between; padding: 10px 12px; background: var(--bg-hover); border-radius: 8px; border: 1px solid transparent; transition: all 0.15s ease; }
.item-row:hover { background: var(--border-color); border-color: var(--border-color); }
.item-dim { opacity: 0.5; }
.item-info { display: flex; align-items: center; gap: 10px; flex: 1; min-width: 0; }
.item-code { font-family: 'DM Mono', 'Courier New', monospace; font-size: 0.8rem; font-weight: 600; color: var(--text-primary); background: var(--bg-hover); padding: 3px 8px; border-radius: 4px; white-space: nowrap; flex-shrink: 0; }
.item-text { font-size: 0.85rem; color: var(--text-muted); overflow: hidden; text-overflow: ellipsis; white-space: nowrap; }

.site-url-text { font-size: 0.82rem; color: var(--text-muted); overflow: hidden; text-overflow: ellipsis; white-space: nowrap; font-family: 'DM Mono', 'Courier New', monospace; }
.type-badge { font-size: 0.68rem; font-weight: 600; text-transform: uppercase; padding: 2px 8px; border-radius: 20px; letter-spacing: 0.5px; flex-shrink: 0; }
.type-block { background: rgba(239, 68, 68, 0.12); color: var(--danger); border: 1px solid rgba(239, 68, 68, 0.25); }
.type-restrict { background: rgba(251, 191, 36, 0.12); color: var(--warning); border: 1px solid rgba(251, 191, 36, 0.25); }
.type-promote { background: rgba(34, 197, 94, 0.12); color: var(--success); border: 1px solid rgba(34, 197, 94, 0.25); }

.site-actions { display: flex; gap: 4px; flex-shrink: 0; margin-left: 8px; }
.icon-btn { width: 30px; height: 30px; display: flex; align-items: center; justify-content: center; background: var(--bg-hover); border: 1px solid var(--border-color); border-radius: 6px; color: var(--text-muted); cursor: pointer; transition: all 0.15s ease; flex-shrink: 0; }
.icon-btn:hover { background: var(--border-color); color: var(--text-primary); }
.icon-btn.danger:hover { background: rgba(239, 68, 68, 0.12); border-color: rgba(239, 68, 68, 0.25); color: var(--danger); }
.icon-btn.success:hover { background: rgba(34, 197, 94, 0.12); border-color: rgba(34, 197, 94, 0.25); color: var(--success); }
.icon-btn.warn:hover { background: rgba(251, 191, 36, 0.12); border-color: rgba(251, 191, 36, 0.25); color: var(--warning); }

.empty-state { text-align: center; color: var(--text-muted); font-size: 0.85rem; padding: 20px; }

.add-form { margin-top: 18px; padding-top: 18px; border-top: 1px solid var(--border-color); }
.form-row { display: flex; gap: 8px; flex-wrap: wrap; }
.form-input { flex: 1; min-width: 120px; padding: 9px 14px; background: var(--bg-hover); border: 1px solid var(--border-color); border-radius: 8px; color: var(--text-primary); font-size: 0.85rem; font-family: inherit; transition: all 0.2s ease; }
.form-input:focus { outline: none; border-color: var(--accent); background: var(--border-color); }
.form-input::placeholder { color: var(--text-muted); }
.form-input.wide { flex: 2; }
.add-btn { width: 38px; height: 38px; display: flex; align-items: center; justify-content: center; background: var(--bg-hover); border: 1px solid var(--border-color); border-radius: 8px; color: var(--text-primary); cursor: pointer; transition: all 0.2s ease; flex-shrink: 0; }
.add-btn:hover { background: var(--border-color); border-color: var(--border-light); }
.form-msg { font-size: 0.8rem; color: var(--text-muted); margin-top: 8px; }
.form-msg.error { color: var(--danger); }

.type-btn { padding: 9px 14px; background: var(--bg-hover); border: 1px solid var(--border-color); border-radius: 8px; color: var(--text-muted); font-size: 0.8rem; font-weight: 500; font-family: inherit; cursor: pointer; transition: all 0.2s ease; display: flex; align-items: center; gap: 5px; white-space: nowrap; flex-shrink: 0; }
.type-btn:hover { background: var(--border-color); border-color: var(--border-light); }
.block-type:hover { background: rgba(239, 68, 68, 0.12); border-color: rgba(239, 68, 68, 0.25); color: var(--danger); }
.restrict-type:hover { background: rgba(251, 191, 36, 0.12); border-color: rgba(251, 191, 36, 0.25); color: var(--warning); }
.promote-type:hover { background: rgba(34, 197, 94, 0.12); border-color: rgba(34, 197, 94, 0.25); color: var(--success); }

.site-row { flex-direction: column; align-items: flex-start; gap: 8px; }
.site-row .site-actions { margin-left: 0; width: 100%; justify-content: flex-end; }

.prism-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
  gap: 10px;
  margin-bottom: 4px;
}

.prism-card {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 6px;
  padding: 14px;
  background: var(--bg-hover);
  border: 1.5px solid var(--border-color);
  border-radius: 10px;
  cursor: pointer;
  transition: all 0.18s ease;
  text-align: left;
  font-family: inherit;
  color: var(--text-primary);
  width: 100%;
}

.prism-card:hover { border-color: var(--accent); background: var(--border-color); }

.prism-card-custom { padding: 0; overflow: hidden; flex-direction: row; align-items: stretch; }
.prism-card-main { flex: 1; display: flex; flex-direction: column; align-items: flex-start; gap: 6px; padding: 14px; background: none; border: none; cursor: pointer; color: var(--text-primary); font-family: inherit; text-align: left; }

.prism-name { font-size: 0.85rem; font-weight: 600; color: var(--text-primary); }
.prism-active-badge {
  font-size: 0.68rem; font-weight: 600; text-transform: uppercase; letter-spacing: 0.5px;
  padding: 2px 7px; border-radius: 20px;
}

.prism-meta { display: flex; flex-wrap: wrap; gap: 4px; }
.prism-tag {
  font-size: 0.68rem; color: var(--text-muted); background: var(--bg-primary);
  border: 1px solid var(--border-color); border-radius: 4px; padding: 1px 5px;
  font-family: 'DM Mono', 'Courier New', monospace;
}

.prism-delete-btn {
  display: flex; align-items: center; justify-content: center; width: 32px; flex-shrink: 0;
  background: none; border: none; border-left: 1px solid var(--border-color);
  color: var(--text-muted); cursor: pointer; transition: all 0.15s ease;
}
.prism-delete-btn:hover { background: rgba(239, 68, 68, 0.1); color: var(--danger); }

.prism-card-hidden { opacity: 0.45; }
.prism-card-hidden .prism-card-main { cursor: default; }

.prism-filter-row { display: flex; gap: 10px; margin-top: 10px; flex-wrap: wrap; }
.prism-filter-col { flex: 1; min-width: 180px; display: flex; flex-direction: column; gap: 6px; }

.color-pick-row { display: flex; gap: 6px; align-items: center; flex-shrink: 0; flex-wrap: wrap; }
.prism-color-dot {
  width: 22px;
  height: 22px;
  border-radius: 50%;
  border: 2px solid transparent;
  cursor: pointer;
  transition: all 0.15s ease;
  outline: none;
  padding: 0;
}
.prism-color-dot:hover { transform: scale(1.15); }
.prism-color-dot.active { border-color: var(--text-primary); box-shadow: 0 0 0 2px var(--bg-primary); }

.create-prism-btn {
  display: flex; align-items: center; justify-content: center; gap: 8px;
  width: 100%; margin-top: 12px; padding: 10px 16px;
  border: none; border-radius: 8px;
  color: #fff; font-size: 0.85rem; font-weight: 600;
  font-family: inherit; cursor: pointer; transition: opacity 0.2s ease;
}
.create-prism-btn:hover { opacity: 0.88; }

.night-hours-row { display: flex; gap: 10px; margin-top: 8px; }
.night-hour-col { flex: 1; min-width: 0; }

@media (max-width: 640px) {
  .tab-btn span { display: none; }
  .tab-btn { padding: 12px 14px; }
  .content { padding: 20px 16px 60px; }
  .card { padding: 18px; }
  .color-swatch { width: 28px; height: 28px; }
  .form-row { flex-direction: column; }
  .form-input, .form-input.wide { min-width: 100%; }
  .type-btn { width: 100%; justify-content: center; }
  .add-btn { width: 100%; height: auto; padding: 9px; }
  .font-upload-area { flex-direction: column; align-items: stretch; }
  .upload-font-btn { justify-content: center; }
  .lang-group { flex-direction: column; }
  .prism-grid { grid-template-columns: repeat(auto-fill, minmax(140px, 1fr)); }
  .prism-filter-row { flex-direction: column; }
  .color-pick-row { justify-content: flex-start; }
}

@media (max-width: 480px) {
  .page-title { font-size: 1rem; }
  .card { padding: 14px; }
  .color-swatch { width: 26px; height: 26px; }
  .color-swatches { gap: 6px; }
  .prism-grid { grid-template-columns: 1fr 1fr; }
}
</style>