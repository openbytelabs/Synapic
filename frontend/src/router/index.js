import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ResultsView from '../views/ResultsView.vue'
import ApiView from '../views/ApiView.vue'
import SettingsView from '../views/SettingsView.vue'
import TermsView from '../views/TermsView.vue'
import DownloadView from '../views/DownloadView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView 
  },
  {
    path: '/search',
    name: 'search',
    component: ResultsView
  },
  {
    path: '/terms',
    name: 'terms',
    component: TermsView
  },
  {
    path: '/settings',
    name: 'settings',
    component: SettingsView
  },
  {
    path: '/api',
    name: 'api',
    component: ApiView
  },
  {
    path: '/download',
    name: 'download',
    component: DownloadView 
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router