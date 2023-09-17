<template>
    <div class="flex flex-col w-fit">
        <div class="flex">
            <div class="grow text-center italic text-sm max-w-[80%] m-auto">{{ channelDetails.description }}</div>
        </div>
        <div class="flex w-full">
            <div class="flex flex-col justify-between h-96 min-w-[30rem] grow">
                <div class="grow overflow-scroll px-4 flex flex-col-reverse">
                    <div class="grow px-4 flex flex-col gap-1 mb-2">
                        <ChatRoomMessage v-for="m, ind in messages" :key="m.CreatedAt" :authorName="m.AuthorName"
                            :content="m.Content"
                            :previousMessageAuthor="ind >= 1 ? messages[ind - 1].AuthorName : undefined">
                        </ChatRoomMessage>
                    </div>

                </div>
                <div v-if="user.jwt" class="flex items-center justify-center shadow-md w-full">
                    <input class="p-2 grow" type=" text" v-model="messageContent" placeholder="Type your message here.."
                        @keydown.enter="sendMessage" autofocus>
                    <button class="bg-slate-100 hover:bg-slate-200 duration-200 p-2 px-5" @click="sendMessage">send</button>
                </div>
                <div v-else class="w-fill text-center p-3">
                    You have to <NuxtLink to="/login" class="text-primary">log in</NuxtLink> or <NuxtLink to="/register"
                        class="text-primary">
                        register</NuxtLink>
                    before writing messages
                </div>

            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useNotificationStore } from '~/stores/notificationStore';
import { useUserStore } from '~/stores/userStore';

const router = useRouter()
const user = useUserStore()
const props = defineProps<{
    channelId: string
}>();

let channelDetailsRequest = await fetch("/api/channel/details/" + props.channelId)
let channelDetails = await channelDetailsRequest.json()
console.log(channelDetails)

type Message = {
    AuthorName: string,
    Content: string,
    CreatedAt: number
}
let messages = ref([] as Message[])
let messageContent = ref("")


const socket = new WebSocket(WebsocketRelativeAddr('/api/channel/eventStream/' + props.channelId));
socket.onerror = (e) => {
    console.log(e)

    useNotificationStore().pushNotification({ Level: "WARNING", Content: "Selected channel is full" })
    router.push("/")
}
socket.addEventListener("message", (event) => {
    let parsed = JSON.parse(event.data)
    messages.value.push(parsed)
    console.log(parsed)
    setTimeout(() => {
        let ind = messages.value.findIndex((v) => v.CreatedAt == parsed.CreatedAt)
        messages.value.splice(ind, 1)
    }, channelDetails.MessageDuration / 1000000);
})
router.beforeEach(() => {
    socket.close()
})

async function sendMessage() {
    if (!user.jwt) {
        console.log("not logged in, cannot send message")
        return
    }
    socket.send(JSON.stringify({
        "AuthorName": user.jwt.user,
        "Content": messageContent.value,
        "CreatedAt": Date.now(),
        "Hidden": false,
    }))
    messageContent.value = ""
}
</script>