// @ts-ignore
import {createStore} from "vuex";
import {APIClient} from "../api/client";

export const store = createStore({
    state: {
        appName: "knowYourWebsite",
        client: new APIClient({
            API_URL: import.meta.env.VITE_API_URL,
            CLIENT_TIMEOUT: import.meta.env.VITE_CLIENT_TIMEOUT,
            HEADERS: import.meta.env.VITE_HEADERS,
        }),
        testCookie: import.meta.env.VITE_COOKIE_FOR_TEST,
        notificationVisible: false,
        notificationMessage: "Smth went wrong",
        slugData: {},
    },
    getters: {
        appName: (state: any) => state.appName,
        client: (state: any) => state.client,
        copyRightString: (state:any) => {
            const startYear = 2023;
            let yearSlug = String(startYear);
            const current = new Date().getFullYear();
            if (startYear !== current) {
                yearSlug += "-"+String(current);
            }
            return `Copyright © ${yearSlug} - ${state.appName} | All Rights Reserved `;
        },
        testCookie: (state: any) => state.testCookie,
        notificationVisible: (state: any) => state.notificationVisible,
        notificationMessage: (state: any) => state.notificationMessage,
        slugData: (state: any) => state.slugData,
    },
    mutations: {
        setNotificationMessage(state: any, message: string) {
            state.notificationMessage = message
        },
        toggleNotification(state: any) {
            state.notificationVisible = !state.notificationVisible
        },
        setSlugData(state: any, slug: any) { state.slugData = slug }
    },
    actions: {
        toggleNotification({commit}: any, message: string) {
            commit("setNotificationMessage", message)
            commit("toggleNotification")
            setTimeout(() => commit("toggleNotification"), 2000)
        }
    }
})


