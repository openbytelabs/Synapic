import { ref, computed } from 'vue'

const DEFAULT_BANGS = [
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

function loadFromStorage(key, fallback) {
  try {
    const raw = localStorage.getItem(key)
    return raw ? JSON.parse(raw) : fallback
  } catch {
    return fallback
  }
}

function saveToStorage(key, value) {
  localStorage.setItem(key, JSON.stringify(value))
}

const customBangs = ref(loadFromStorage('synapic_custom_bangs', []))
const deletedBangs = ref(loadFromStorage('synapic_deleted_bangs', []))

export function useBangs() {
  const bangs = computed(() => {
    const activeDefaults = DEFAULT_BANGS.filter(
      (b) => !deletedBangs.value.includes(b.shortcut)
    )
    return [...activeDefaults, ...customBangs.value]
  })

  function addBang(bang) {
    const exists = bangs.value.some((b) => b.shortcut === bang.shortcut)
    if (exists) return
    customBangs.value.push(bang)
    saveCustomBangs()
  }

  function removeBang(shortcut) {
    if (isDefault(shortcut)) {
      if (!deletedBangs.value.includes(shortcut)) {
        deletedBangs.value.push(shortcut)
        saveDeletedBangs()
      }
    } else {
      customBangs.value = customBangs.value.filter((b) => b.shortcut !== shortcut)
      saveCustomBangs()
    }
  }

  function restoreBang(shortcut) {
    deletedBangs.value = deletedBangs.value.filter((s) => s !== shortcut)
    saveDeletedBangs()
  }

  function isDefault(shortcut) {
    return DEFAULT_BANGS.some((b) => b.shortcut === shortcut)
  }

  function checkBang(query) {
    const match = query.match(/^!(\S+)\s+(.*)$/)
    if (!match) return null
    const shortcut = match[1]
    const remaining = match[2]
    const bang = bangs.value.find((b) => b.shortcut === shortcut)
    if (!bang) return null
    return { bang, query: remaining }
  }

  function saveCustomBangs() {
    saveToStorage('synapic_custom_bangs', customBangs.value)
  }

  function saveDeletedBangs() {
    saveToStorage('synapic_deleted_bangs', deletedBangs.value)
  }

  return {
    bangs,
    customBangs,
    deletedBangs,
    addBang,
    removeBang,
    restoreBang,
    isDefault,
    checkBang,
    saveCustomBangs,
    saveDeletedBangs
  }
}
