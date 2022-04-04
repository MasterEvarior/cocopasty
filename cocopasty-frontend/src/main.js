import { createApp } from 'vue'
import Toaster from '@meforma/vue-toaster';
import App from './App.vue'

createApp(App)
    .use(Toaster)
    .mount('#app')
