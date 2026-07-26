function getLuminance(hex) {
  if (!hex || !hex.startsWith('#')) return 0.5
  const c = hex.replace('#', '')
  const r = parseInt(c.substring(0, 2), 16) / 255
  const g = parseInt(c.substring(2, 4), 16) / 255
  const b = parseInt(c.substring(4, 6), 16) / 255
  return 0.299 * r + 0.587 * g + 0.114 * b
}

function getContrastText(hex) {
  return getLuminance(hex) > 0.55 ? '#000000' : '#ffffff'
}

function adjustBrightness(hex, percent) {
  const num = parseInt(hex.replace('#', ''), 16)
  const amt = Math.round(2.55 * percent)
  const R = (num >> 16) + amt
  const G = (num >> 8 & 0x00FF) + amt
  const B = (num & 0x000FF) + amt
  const n = (0x1000000 + (R < 255 ? R < 1 ? 0 : R : 255) * 0x10000 + (G < 255 ? G < 1 ? 0 : G : 255) * 0x100 + (B < 255 ? B < 1 ? 0 : B : 255))
  return '#' + n.toString(16).slice(1)
}

export function applyTheme() {
  const theme = localStorage.getItem('synapic_theme') || 'dark'
  if (theme === 'auto') {
    const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
    document.documentElement.setAttribute('data-theme', prefersDark ? 'dark' : 'light')
  } else {
    document.documentElement.setAttribute('data-theme', theme)
  }
}

export function applyFontSize() {
  const size = localStorage.getItem('synapic_font_size')
  if (size) document.documentElement.style.fontSize = size
}

export function applyAccentColor() {
  const color = localStorage.getItem('synapic_accent_color')
  if (!color) return
  document.documentElement.style.setProperty('--accent', color)
  document.documentElement.style.setProperty('--accent-hover', adjustBrightness(color, 20))
  document.documentElement.style.setProperty('--accent-text', getContrastText(color))
  document.documentElement.style.setProperty('--dot-color', color)
  applyBgStyle()
}

const fontFamilyMap = {
  sans: 'system-ui, -apple-system, sans-serif',
  serif: 'Georgia, "Times New Roman", serif',
  mono: '"Courier New", monospace',
  inter: "'Inter', system-ui, sans-serif",
  geist: "'Geist', system-ui, sans-serif",
  dm: "'DM Sans', system-ui, sans-serif",
  fira: "'Fira Code', monospace",
  playfair: "'Playfair Display', serif",
}

export function applyFontFamily() {
  const family = localStorage.getItem('synapic_font_family') || 'sans'
  document.documentElement.setAttribute('data-font-family', family)

  let cssFont = fontFamilyMap[family]

  if (!cssFont) {
    try {
      const customFonts = JSON.parse(localStorage.getItem('synapic_custom_fonts') || '[]')
      const found = customFonts.find(f => f.value === family)
      if (found) cssFont = `'${found.name}', sans-serif`
    } catch (e) { console.error(e) }
  }

  if (!cssFont) cssFont = fontFamilyMap.sans
  document.documentElement.style.setProperty('--font-family', cssFont)

  let globalFontStyle = document.getElementById('synapic-global-font')
  if (!globalFontStyle) {
    globalFontStyle = document.createElement('style')
    globalFontStyle.id = 'synapic-global-font'
    document.head.appendChild(globalFontStyle)
  }
  globalFontStyle.textContent = `body, input, select, button, textarea, h1, h2, h3, h4, h5, h6, p, span, a, div, li, td, th, label, code { font-family: ${cssFont} !important; }`
}

export function registerFontFace(name, base64, format) {
  const formatMap = { ttf: 'truetype', otf: 'opentype', woff: 'woff', woff2: 'woff2' }
  const fmt = formatMap[format] || format
  const id = 'synapic-font-' + name.replace(/[^a-zA-Z0-9]/g, '-').toLowerCase()
  let styleEl = document.getElementById(id)
  if (styleEl) styleEl.remove()
  styleEl = document.createElement('style')
  styleEl.id = id
  styleEl.textContent = `@font-face { font-family: '${name}'; src: url(data:font/${format};base64,${base64}) format('${fmt}'); font-weight: normal; font-style: normal; font-display: swap; }`
  document.head.appendChild(styleEl)
}

export function loadCustomFonts() {
  try {
    const fonts = JSON.parse(localStorage.getItem('synapic_custom_fonts') || '[]')
    fonts.forEach(f => {
      if (f.base64 && f.format) {
        registerFontFace(f.name, f.base64, f.format)
      }
    })
  } catch (e) { console.error(e) }
}

export function applyBgStyle() {
  const style = localStorage.getItem('synapic_bg_style') || 'plain'
  document.documentElement.setAttribute('data-bg-style', style)

  let bgStyleEl = document.getElementById('synapic-bg-style')
  if (!bgStyleEl) {
    bgStyleEl = document.createElement('style')
    bgStyleEl.id = 'synapic-bg-style'
    document.head.appendChild(bgStyleEl)
  }

  if (style === 'dots') {
    bgStyleEl.textContent = `
      html[data-bg-style="dots"] body::after {
        content: '';
        position: fixed;
        inset: 0;
        background-image: radial-gradient(circle, var(--dot-color, #3b82f6) 0.8px, transparent 0.8px);
        background-size: 22px 22px;
        opacity: 0.09;
        pointer-events: none;
        z-index: 9999;
      }
    `
  } else {
    bgStyleEl.textContent = ''
  }
}

export function applyAllSettings() {
  loadCustomFonts()
  applyTheme()
  applyFontSize()
  applyAccentColor()
  applyFontFamily()
  applyBgStyle()

  const attrs = [
    ['synapic_border_radius', 'data-border-radius'],
    ['synapic_search_bar_style', 'data-search-bar-style'],
    ['synapic_homepage_layout', 'data-homepage-layout'],
    ['synapic_card_shadow', 'data-card-shadow'],
    ['synapic_animation_speed', 'data-animation-speed'],
    ['synapic_logo_size', 'data-logo-size'],
    ['synapic_density', 'data-density'],
    ['synapic_blur', 'data-blur'],
    ['synapic_reduced_motion', 'data-reduced-motion'],
    ['synapic_result_density', 'data-result-density']
  ]

  attrs.forEach(([key, attr]) => {
    const val = localStorage.getItem(key)
    if (val) document.documentElement.setAttribute(attr, val)
  })
}

export function initSettings() {
  applyAllSettings()
}