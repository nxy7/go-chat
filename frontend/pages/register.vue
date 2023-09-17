<template>
    <div class="flex flex-col grow p-3 w-full">
        <LoginForm :OnSubmit="u.register" ButtonText="Register"></LoginForm>
        <div>Already have account? <NuxtLink to="/login" class="text-primary">Log in!</NuxtLink>
        </div>
    </div>
</template>

<script setup lang="ts">
import LoginForm from '~/components/LoginForm.vue';
import { useUserStore } from '~/stores/userStore';
const u = useUserStore()
</script>