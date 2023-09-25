import {defineStore} from "pinia";
import {APIClient, APIResponse} from "../api/client";

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
        }
    }
})
