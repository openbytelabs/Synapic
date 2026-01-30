import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index'
import { inject } from '@vercel/analytics';

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

// Icon setleri
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'

// Hepsini ekle
library.add(fas, far, fab)

inject();

const app = createApp(App) 

app.component('font-awesome-icon', FontAwesomeIcon)

app.use(router)

app.mount('#app')
