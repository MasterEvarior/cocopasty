import { createApp } from 'vue'
import Toaster from '@meforma/vue-toaster';
import App from './App.vue'

let app = createApp(App)

app
    .use(Toaster)
    .mount('#app')
