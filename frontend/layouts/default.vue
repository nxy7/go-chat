<template>
    <div class="w-screen h-screen flex justify-center items-center relative">
        <NotificationsList></NotificationsList>
        <AboutOverlay v-if="showAboutOverlay" @close-overlay="showAboutOverlay = false"></AboutOverlay>
        <div
            class="duration-300 min-w-[20rem] min-h-[20rem] flex flex-col items-center justify-center shadow-lg border-t-4 border-primary">
            <div class="flex justify-center relative p-3 w-full">
                <div v-if="showGoBack" class="absolute left-3">
                    <NuxtLink to="/">Go back</NuxtLink>
                </div>
                <div class=" text-primary font-bold">Chat App</div>
                <div class="absolute right-3">
                    <button class="w-6 bg-primary rounded-full flex items-center justify-center font-bold text-white "
                        @click="showAboutOverlay = true">?</button>
                </div>

            </div>
            <div class="flex grow w-full h-full">
                <slot />
            </div>

        </div>

    </div>
</template>

<script setup lang="ts">
import { useNotificationStore } from '~/stores/notificationStore';
import { useUserStore } from '~/stores/userStore';

const userStore = useUserStore()
const notificationStore = useNotificationStore()
const router = useRouter()
let showGoBack = ref(false)
let showAboutOverlay = ref(false)

function ShowGoBack() {
    if (router.currentRoute.value.path != "/")
        showGoBack.value = true
    else
        showGoBack.value = false
}
ShowGoBack()
router.afterEach(() => {
    ShowGoBack()
})
console.log(userStore.jwt)
</script>