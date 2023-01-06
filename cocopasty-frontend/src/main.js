import { createApp } from 'vue'
import Toaster from '@meforma/vue-toaster';
import "@fontsource/open-sans/400.css"
import "@fontsource/open-sans/700.css"
import App from './App.vue'

let app = createApp(App)

app
    .use(Toaster)
    .mount('#app')
