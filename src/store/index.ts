import {defineStore} from "pinia";
import {APIClient, APIResponse} from "../api/client";
import {inject} from "vue";
import VueCookies from "vue-cookies";

export const useMainStore = defineStore( {
    id: "main",

    state: () => ({
        _appName: "knowYourWebsite",
        _client: new APIClient({
            API_URL: import.meta.env.VITE_API_URL,
            CLIENT_TIMEOUT: import.meta.env.VITE_CLIENT_TIMEOUT,
            HEADERS: import.meta.env.VITE_HEADERS,
        }),
        _testCookie: import.meta.env.VITE_COOKIE_FOR_TEST,
        _notificationVisible: false,
        _notificationMessage: "Smth went wrong",
        _slugData: {} as APIResponse,
        _showEgg: false,
        _siteShown: false,
    }),

    getters: {
        appName: (state: any) => state._appName,
        client: (state: any) => state._client,
        copyRightString: (state:any) => {
            const startYear = 2023;
            let yearSlug = String(startYear);
            const current = new Date().getFullYear();
            if (startYear !== current) {
                yearSlug += "-"+String(current);
            }
            return `Copyright © ${yearSlug} - ${state._appName} | All Rights Reserved `;
        },
        testCookie: (state: any) => state._testCookie,
        notificationVisible: (state: any) => state._notificationVisible,
        notificationMessage: (state: any) => state._notificationMessage,
        slugData: (state: any) => state._slugData,
        showEgg: (state: any) => state._showEgg,
        siteShown: (state: any) => state._siteShown,
    },

    actions: {
        setNotificationMessage(message: string) {
            this._notificationMessage = message
        },
        _toggleNotification() {
            this._notificationVisible = !this._notificationVisible
        },
        setSlugData(slug: any) { this._slugData = slug },
        toggleNotification(message: string) {
            this.setNotificationMessage(message)
            this._toggleNotification()
            setTimeout(() => this._toggleNotification(), 2000)
        },
         toggleEgg() {
            // @ts-ignore
             const $cookies = inject<VueCookies>('$cookies')
            if (Math.floor(Math.random() * 100) == 50 || $cookies.get("testCookie") === this.testCookie) {
                this._showEgg = !this._showEgg;
                setTimeout(() => this._showEgg = !this._showEgg, 2000);
            }
        },
        toggleSiteView() {
            this._siteShown = !this._siteShown;
        },
    }
})
