<template>
    <div v-if="!pending" class="flex flex-col grow justify-center items-center gap-3">
        <input class="p-2 rounded-lg bg-slate-100" type="text" id="username" placeholder="Put in your username"
            v-model="username" @keydown.enter="RunOnSubmit">
        <input class="p-2 rounded-lg bg-slate-100" type="password" id="password" placeholder="Put in your password"
            v-model="password" @keydown.enter="RunOnSubmit">
        <button @click="RunOnSubmit" class="rounded-sm bg-primary text-white font-bold w-fit px-3 py-1">{{ ButtonText
        }}</button>
    </div>
    <div v-else>
        Pending
    </div>
</template>

<script setup lang="ts">
import { useUserStore } from '~/stores/userStore';

const userStore = useUserStore()
const props = defineProps<{
    OnSubmit: (username: string, password: string) => void,
    ButtonText: string
}>();

async function RunOnSubmit() {
    console.log(username.value, password.value)
    pending = true
    await props.OnSubmit(username.value, password.value)
    pending = false
}

let username = ref("")
let password = ref("")
let pending = false
</script>