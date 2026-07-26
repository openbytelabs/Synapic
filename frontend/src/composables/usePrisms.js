import { ref, computed } from 'vue'

const STORAGE_KEY = 'synapic_prisms'

const PRISM_COLORS = [
  '#3b82f6', '#6366f1', '#8b5cf6', '#a855f7', '#d946ef',
  '#ec4899', '#f43f5e', '#ef4444', '#f97316', '#eab308',
  '#22c55e', '#14b8a6', '#06b6d4', '#0ea5e9', '#64748b'
]

const DEFAULT_PRISMS = [
  {
    id: 'default',
    name: 'All Web',
    nametr: 'Tüm Web',
    label: 'AW',
    color: '#64748b',
    domains: [],
    tlds: [],
    isDefault: true,
    isSystem: true
  },
  {
    id: 'academic',
    name: 'Academic',
    nametr: 'Akademik',
    label: 'AC',
    color: '#8b5cf6',
    domains: [],
    tlds: ['.edu', '.ac.uk', '.ac.jp', '.ac.au', '.ac.nz'],
    isDefault: false,
    isSystem: true
  },
  {
    id: 'government',
    name: 'Government',
    nametr: 'Devlet',
    label: 'GV',
    color: '#f59e0b',
    domains: [],
    tlds: ['.gov', '.gov.tr', '.gov.uk', '.gov.au', '.mil'],
    isDefault: false,
    isSystem: true
  },
  {
    id: 'nonprofit',
    name: 'Non-Profit & Open',
    nametr: 'Sivil & Açık',
    label: 'NP',
    color: '#22c55e',
    domains: [],
    tlds: ['.org', '.ngo', '.ong'],
    isDefault: false,
    isSystem: true
  },
  {
    id: 'tech',
    name: 'Tech & Dev',
    nametr: 'Teknoloji',
    label: 'TD',
    color: '#06b6d4',
    domains: ['github.com', 'stackoverflow.com', 'developer.mozilla.org', 'docs.microsoft.com', 'devdocs.io', 'npmjs.com', 'pypi.org', 'hackernews.com', 'dev.to', 'medium.com', 'codeproject.com'],
    tlds: ['.dev', '.io'],
    isDefault: false,
    isSystem: true
  },
  {
    id: 'news',
    name: 'News',
    nametr: 'Haberler',
    label: 'NW',
    color: '#ef4444',
    domains: ['bbc.com', 'reuters.com', 'apnews.com', 'theguardian.com', 'nytimes.com', 'cnn.com', 'aljazeera.com', 'dw.com', 'euronews.com', 'hurriyet.com.tr', 'cumhuriyet.com.tr', 'bianet.org'],
    tlds: [],
    isDefault: false,
    isSystem: true
  },
  {
    id: 'science',
    name: 'Science',
    nametr: 'Bilim',
    label: 'SC',
    color: '#6366f1',
    domains: ['nature.com', 'sciencedirect.com', 'pubmed.ncbi.nlm.nih.gov', 'scholar.google.com', 'arxiv.org', 'researchgate.net', 'semanticscholar.org', 'jstor.org', 'plos.org', 'springer.com'],
    tlds: [],
    isDefault: false,
    isSystem: true
  }
]

function loadPrisms() {
  try {
    const raw = localStorage.getItem(STORAGE_KEY)
    return raw ? JSON.parse(raw) : []
  } catch {
    return []
  }
}

function savePrisms(prisms) {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(prisms))
}

function normalizeUrl(raw) {
  if (!raw || typeof raw !== 'string') return ''
  return raw.match(/^https?:\/\//) ? raw : 'https://' + raw
}

const HIDDEN_KEY = 'synapic_hidden_prisms'

function loadHiddenPrisms() {
  try {
    const raw = localStorage.getItem(HIDDEN_KEY)
    return raw ? JSON.parse(raw) : []
  } catch { return [] }
}

function saveHiddenPrisms(ids) {
  localStorage.setItem(HIDDEN_KEY, JSON.stringify(ids))
}

const customPrisms = ref(loadPrisms())
const hiddenSystemPrismIds = ref(loadHiddenPrisms())
const activePrismId = ref(localStorage.getItem('synapic_active_prism') || 'default')

export function usePrisms() {
  const allPrisms = computed(() => [
    ...DEFAULT_PRISMS.filter(p => !hiddenSystemPrismIds.value.includes(p.id)),
    ...customPrisms.value
  ])

  const allSystemPrisms = computed(() => DEFAULT_PRISMS)

  const activePrism = computed(() =>
    allPrisms.value.find(p => p.id === activePrismId.value) || DEFAULT_PRISMS[0]
  )

  function setActivePrism(id) {
    activePrismId.value = id
    localStorage.setItem('synapic_active_prism', id)
  }

  function addPrism(prism) {
    const id = 'custom_' + Date.now()
    const label = (prism.name || '').slice(0, 2).toUpperCase()
    const color = prism.color || PRISM_COLORS[customPrisms.value.length % PRISM_COLORS.length]
    customPrisms.value.push({ ...prism, id, label, color, isSystem: false })
    savePrisms(customPrisms.value)
    return id
  }

  function removePrism(id) {
    customPrisms.value = customPrisms.value.filter(p => p.id !== id)
    savePrisms(customPrisms.value)
    if (activePrismId.value === id) setActivePrism('default')
  }

  function hideSystemPrism(id) {
    if (id === 'default') return
    if (!hiddenSystemPrismIds.value.includes(id)) {
      hiddenSystemPrismIds.value.push(id)
      saveHiddenPrisms(hiddenSystemPrismIds.value)
      if (activePrismId.value === id) setActivePrism('default')
    }
  }

  function updatePrism(id, updates) {
    const idx = customPrisms.value.findIndex(p => p.id === id)
    if (idx !== -1) {
      customPrisms.value[idx] = { ...customPrisms.value[idx], ...updates }
      savePrisms(customPrisms.value)
    }
  }

  function filterByPrism(results) {
    const prism = activePrism.value
    if (!prism || prism.id === 'default') return results
    if (prism.domains.length === 0 && prism.tlds.length === 0) return results

    return results.filter(result => {
      const rawUrl = result.url || ''
      if (!rawUrl) return false
      try {
        const urlStr = normalizeUrl(rawUrl)
        const parsed = new URL(urlStr)
        const hostname = parsed.hostname.replace(/^www\./, '').toLowerCase()

        const domainMatch = prism.domains.length > 0 && prism.domains.some(d => {
          const clean = d.replace(/^www\./, '').toLowerCase()
          return hostname === clean || hostname.endsWith('.' + clean)
        })

        const tldMatch = prism.tlds.length > 0 && prism.tlds.some(tld => {
          const cleanTld = tld.startsWith('.') ? tld.toLowerCase() : '.' + tld.toLowerCase()
          return hostname.endsWith(cleanTld)
        })

        return domainMatch || tldMatch
      } catch {
        return false
      }
    })
  }

  return {
    allPrisms,
    allSystemPrisms,
    customPrisms,
    hiddenSystemPrismIds,
    activePrism,
    activePrismId,
    setActivePrism,
    addPrism,
    removePrism,
    updatePrism,
    hideSystemPrism,
    filterByPrism,
    PRISM_COLORS
  }
}