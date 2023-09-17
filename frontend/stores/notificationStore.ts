import { defineStore } from "pinia";
import { ChatAppNotification } from "./notification";


export let useNotificationStore = defineStore('notificationStore', {
    state: () => {
        return {
            notificationArr: [] as ChatAppNotification[]
        }
    },
    actions: {
        pushNotification(n: ChatAppNotification, duration?: number) {
            this.notificationArr.push(n)
            let displayDuration = (duration ?? 15) * 1000

            setTimeout(() => {
                let ind = this.notificationArr.findIndex((v) => v.Content == n.Content)
                this.notificationArr.splice(ind, 1)
            }, displayDuration)
        },
        clearNotifications() {
            this.notificationArr = []
        }
    }
})

