export function applyTheme() {
  const theme = localStorage.getItem('synapic_theme') || 'dark'
  document.documentElement.setAttribute('data-theme', theme)
}

export function applyFontSize() {
  const size = localStorage.getItem('synapic_font_size')
  if (size) {
    document.documentElement.style.fontSize = size
  }
}

export function initSettings() {
  applyTheme()
  applyFontSize()
}
