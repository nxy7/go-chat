import { StoreDefinition, defineStore } from "pinia"
import { Buffer } from "buffer"
import { useNotificationStore } from "./notificationStore"

export const useUserStore = defineStore('user', {
    hydrate: async (storeState, initialState) => {
        let token = localStorage.getItem("jwt")
        if (!token) {
            console.log("token not found")
            return
        }
        storeState.jwt = DecodeJwt(token)
        storeState.rawToken = token
        KeepTokenFresh()
        watch(storeState, () => {
            console.log("Updating localstoreage token")
            localStorage.setItem("jwt", storeState.rawToken ?? "")
        })

    },
    state: () => {
        return {
            jwt: null as JwtData | null,
            rawToken: null as string | null
        }
    },
    // getters: {
    //     getUser: (state) => state.user,
    // },
    actions: {
        async register(username: string, password: string) {
            let response = await fetch("/api/register", {
                method: "POST",
                body: JSON.stringify({ "username": username, "password": password }),
            })
            if (response.status != 200) {
                let notifications = useNotificationStore()
                notifications.pushNotification({ Level: "WARNING", Content: await response.text() }, 7)
                return
            }
            let token = await response.text()
            let u = DecodeJwt(token)
            this.rawToken = token
            this.jwt = u
            localStorage.setItem("jwt", token)
            RedirectToUrl("/")
        },
        async login(username: string, password: string) {

            let response = await fetch("/api/login", {
                method: "POST",
                body: JSON.stringify({ "username": username, "password": password }),
            })
            if (response.status != 200) {
                let notifications = useNotificationStore()
                notifications.pushNotification({ Level: "WARNING", Content: await response.text() }, 7)
                return
            }
            let token = await response.text()
            let u = DecodeJwt(token)
            this.rawToken = token
            this.jwt = u
            localStorage.setItem("jwt", token)
            RedirectToUrl("/")

        },
        isLoggedIn() { },
        async logout() {
            this.rawToken = null
            this.jwt = null
            localStorage.removeItem("jwt")
            await fetch(`/api/logout`)
            RedirectToUrl("/")
        }
    },
})


type JwtData = {
    user: string
    exp: number
}

export function RedirectToUrl(url: string) {
    let r = useRouter()
    r.push(url)
}

export function DecodeJwt(token: string): JwtData | null {
    let claimsPart = token.split(".").at(1)
    if (!claimsPart) {
        console.log("No claims part in ", token)
        return null
    }
    let buff = Buffer.from(claimsPart, 'base64')
    let asString = buff.toString()
    console.log(asString)
    return JSON.parse(asString)
}

async function KeepTokenFresh(
    // storeState: ReturnType<typeof useUserStore>["$state"]
) {
    let storeState = useUserStore()

    if (!storeState.rawToken || !storeState.jwt)
        return

    while (true) {
        let exp = storeState.jwt.exp
        let waitTime = exp * 1000 - Date.now()
        console.log("wait for ", waitTime, "exp: ", exp)

        console.log("Waiting for token to expire for ", waitTime, " ms")
        await new Promise((r) => setTimeout(r, waitTime))

        console.log("Refreshing token")
        let newToken = await RefreshToken()
        let newUserObj = DecodeJwt(newToken)
        if (!newUserObj) {
            console.log("Couldn't decode token")
            return
        }
        console.log("New token set")
        console.log("Old: ", storeState)
        storeState.rawToken = newToken
        storeState.jwt = newUserObj
        console.log("New: ", storeState)
    }


}

// Get new token from backend, refesh token should be in cookie so just make request i guess
async function RefreshToken() {
    let t = await (await fetch(`/api/token`)).text()

    return t
}