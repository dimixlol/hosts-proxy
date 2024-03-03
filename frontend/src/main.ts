import {createApp} from 'vue'
import VueCookies from 'vue-cookies'
import {createPinia} from "pinia";
import App from './App.vue'
import "@/assets/scss/main.scss";

const pinia = createPinia()
createApp(App)
    .use(pinia)
    .use(VueCookies)
    .mount('#app-container')
