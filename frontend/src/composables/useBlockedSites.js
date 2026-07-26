import { ref } from 'vue'

const STORAGE_KEY = 'synapic_blocked_sites'

const blockedSites = ref([])

function loadBlockedSites() {
  try {
    const raw = localStorage.getItem(STORAGE_KEY)
    if (raw) {
      blockedSites.value = JSON.parse(raw)
    }
  } catch {
    blockedSites.value = []
  }
}

function saveBlockedSites() {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(blockedSites.value))
}

function extractDomain(url) {
  try {
    const u = new URL(url)
    return u.hostname.replace(/^www\./, '')
  } catch {
    const cleaned = url.trim().replace(/^https?:\/\//, '').replace(/^www\./, '').split('/')[0]
    return cleaned
  }
}

function domainMatches(blockedDomain, resultDomain) {
  const bd = blockedDomain.toLowerCase().replace(/^www\./, '').replace(/^\./, '')
  const rd = resultDomain.toLowerCase().replace(/^www\./, '').replace(/^\./, '')
  if (bd === rd) return true
  if (rd.endsWith('.' + bd)) return true
  return false
}

function addBlockedSite(url, type) {
  const domain = extractDomain(url)
  if (!domain) return

  const existing = blockedSites.value.findIndex(
    (s) => domainMatches(extractDomain(s.url), domain)
  )
  if (existing !== -1) {
    blockedSites.value[existing] = { url: domain, type, addedAt: Date.now() }
  } else {
    blockedSites.value.push({ url: domain, type, addedAt: Date.now() })
  }
  saveBlockedSites()
}

function removeBlockedSite(url) {
  const domain = extractDomain(url)
  blockedSites.value = blockedSites.value.filter(
    (s) => !domainMatches(extractDomain(s.url), domain)
  )
  saveBlockedSites()
}

function isBlocked(url) {
  const domain = extractDomain(url)
  return blockedSites.value.some(
    (s) => domainMatches(extractDomain(s.url), domain) && s.type === 'block'
  )
}

function isRestricted(url) {
  const domain = extractDomain(url)
  return blockedSites.value.some(
    (s) => domainMatches(extractDomain(s.url), domain) && s.type === 'restrict'
  )
}

function isPromoted(url) {
  const domain = extractDomain(url)
  return blockedSites.value.some(
    (s) => domainMatches(extractDomain(s.url), domain) && s.type === 'promote'
  )
}

function promoteSite(url) {
  const domain = extractDomain(url)
  const existing = blockedSites.value.findIndex(
    (s) => domainMatches(extractDomain(s.url), domain)
  )
  if (existing !== -1) {
    blockedSites.value.splice(existing, 1)
  }
  blockedSites.value.unshift({ url: domain, type: 'promote', addedAt: Date.now() })
  saveBlockedSites()
}

function getFilteredResults(results) {
  const filtered = results.filter((r) => {
    const domain = extractDomain(r.url)
    return !blockedSites.value.some(
      (s) => domainMatches(extractDomain(s.url), domain) && s.type === 'block'
    )
  })

  const promoted = []
  const normal = []

  for (const item of filtered) {
    const domain = extractDomain(item.url)
    const site = blockedSites.value.find(
      (s) => domainMatches(extractDomain(s.url), domain) && s.type === 'promote'
    )
    if (site) {
      promoted.push(item)
    } else {
      normal.push(item)
    }
  }

  return [...promoted, ...normal]
}

loadBlockedSites()

export function useBlockedSites() {
  return {
    blockedSites,
    addBlockedSite,
    removeBlockedSite,
    isBlocked,
    isRestricted,
    isPromoted,
    promoteSite,
    getFilteredResults,
    saveBlockedSites
  }
}
