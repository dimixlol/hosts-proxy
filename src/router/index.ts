import {createRouter, createWebHistory} from "vue-router";
// @ts-ignore
import Home from "@/views/Home.vue";
// @ts-ignore
import Site from "@/views/Site.vue";

export const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: "/", name: "homeView", component: Home},
        {path: "/site", name: "siteView", component: Site},
    ]
})