import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../pages/HomeView.vue'
import ResultsView from '@/pages/ResultsView.vue'
import SettingsView from '@/pages/SettingsView.vue'
import ApiView from '@/pages/ApiView.vue'
import TermsView from '@/pages/TermsView.vue'

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
    path: '/terms',
    name: 'terms',
    component: TermsView
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router