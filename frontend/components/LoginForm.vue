<template>
    <div v-if="!pending">
        <input type="text" id="username" placeholder="Put in your username" :value="username">
        <input type="password" id="password" placeholder="Put in your password" :value="password">
        <button :onclick="RunOnSubmit()">Register</button>
    </div>
    <div v-else>
        Pending
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    OnSubmit: (username: string, password: string) => void
}>();

async function RunOnSubmit() {
    pending = true
    await props.OnSubmit(username, password)
    pending = false
}

let username = ""
let password = ""
let pending = false
</script>