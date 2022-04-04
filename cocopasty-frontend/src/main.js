import { createApp } from 'vue'
import Toaster from '@meforma/vue-toaster';
import App from './App.vue'

let app = createApp(App)

app.config.globalProperties.backendHost = 'localhost'
app.config.globalProperties.backendPort = '8081'

app
    .use(Toaster)
    .mount('#app')
