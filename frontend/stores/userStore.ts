import { defineStore } from "pinia"

export const useUserStore = defineStore('user', {
    hydrate: async (storeState, initialState) => {
        if (localStorage.getItem("jwt-token")) {
            // if item exists check if token has not expired
            // if it expired try using refresh token
            // if it doesn't work delete jwt from localstorage and do nothing
            storeState.user = {
                id: "qweqwe",
                name: "Dawid"
            }
        }

    },
    state: () => {
        return {
            user: undefined as User | undefined
        }
    },
    getters: {
        getUser: (state) => state.user,
    },
    actions: {
        register() { },
        login() { },
        isLoggedIn() { },
    },

})

type User = {
    id: string
    name: string
}