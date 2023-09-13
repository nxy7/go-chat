<template>
    <LoginForm :OnSubmit="Register"></LoginForm>
</template>

<script setup lang="ts">
import LoginForm from '~/components/LoginForm.vue';

async function Register(username: string, password: string) {
    let response = await fetch("/api/register", {
        method: "POST",
        body: JSON.stringify({ "username": username, "password": password })
    })
    let json = await response.json()
    console.log(json)
}

</script>