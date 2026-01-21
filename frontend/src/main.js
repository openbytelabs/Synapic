import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index'
import { inject } from '@vercel/analytics';

inject();

const app = createApp(App) 

app.use(router)

app.mount('#app')
