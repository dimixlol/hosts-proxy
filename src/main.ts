import { createApp } from 'vue'
// @ts-ignore
import App from './App.vue'
import {router} from "./router";
import "@/assets/scss/main.scss";
import VueCookies from 'vue-cookies'
import {createPinia} from "pinia";

const pinia = createPinia()
createApp(App)
    .use(pinia)
    .use(router)
    .use(VueCookies)
    .mount('#app-container')
