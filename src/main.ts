import { createApp } from 'vue'
// @ts-ignore
import App from './App.vue'
import {store} from "./store";
import {router} from "./router";
import "@/assets/scss/main.scss";
import VueCookies from 'vue-cookies'

createApp(App)
    .use(store)
    .use(router)
    .use(VueCookies)
    .mount('#app-container')
