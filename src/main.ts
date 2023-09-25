import { createApp } from 'vue'
// @ts-ignore
import App from './App.vue'
import "@/assets/scss/main.scss";
import VueCookies from 'vue-cookies'
import {createPinia} from "pinia";

const pinia = createPinia()
createApp(App)
    .use(pinia)
    .use(VueCookies)
    .mount('#app-container')
