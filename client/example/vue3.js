import { createApp } from 'vue'
import App from './index.vue'
import LuckyWheel from '../src/index.js'

// import axios from 'axios'
// import VueAxios from 'vue-axios'

console.log('vue3')

const app = createApp(App)

// axios
// app.prototype.$axios = axios

app.use(LuckyWheel).mount('#app')



