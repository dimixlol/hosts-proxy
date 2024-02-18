import {defineStore} from "pinia";
import {APIClient, APISlugDataResponse} from "../api/client";
import {inject} from "vue";
import VueCookies from "vue-cookies";

export const useStore = defineStore( {
    id: "main",
    state: () => ({
        _appName: import.meta.env.VITE_APP_NAME,
        _client: new APIClient({
            API_URL: import.meta.env.VITE_API_URL,
            CLIENT_TIMEOUT: import.meta.env.VITE_CLIENT_TIMEOUT,
            HEADERS: import.meta.env.VITE_HEADERS,
            API_VERSION: import.meta.env.VITE_API_VERSION
        }),
        _testCookie: import.meta.env.VITE_COOKIE_FOR_TEST,
        _notificationMessage: "Something went wrong :(",
        _slugData: {} as APISlugDataResponse,
        _notificationVisible: false,
        _showEgg: false,
        _siteShown: false,
        _$cookies: inject<VueCookies>('$cookies')
    }),
    getters: {
        appName: (state: any) => state._appName,
        client: (state: any) => state._client,
        copyRightString: (state:any) => {
            const startYear = 2023;
            const current = new Date().getFullYear();
            const yearSlug = startYear === current ? String(startYear) : `${startYear}-${current}`;
            return `Copyright © ${yearSlug} - ${state._appName} | All Rights Reserved`;
        },
        testCookie: (state: any) => state._testCookie,
        notificationVisible: (state: any) => state._notificationVisible,
        notificationMessage: (state: any) => state._notificationMessage,
        slugData: (state: any) => state._slugData,
        showEgg: (state: any) => state._showEgg,
        siteShown: (state: any) => state._siteShown,
        csrfToken: (state: any) => state._$cookies.get(import.meta.env.VITE_CSRF_COOKIE_NAME)
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
            if (Math.floor(Math.random() * 100)%7 == 0 || this._$cookies.get("testCookie") === this.testCookie) {
                this._showEgg = !this._showEgg;
                setTimeout(() => this._showEgg = !this._showEgg, 2000);
            }
        },
        toggleSiteView() {
            this._siteShown = !this._siteShown;
        },
    }
})
