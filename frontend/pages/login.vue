<template>
    <div class="flex flex-col grow p-3 w-full">
        <LoginForm :OnSubmit="user.login" ButtonText="Login"></LoginForm>
        <div>Don't have account yet? <NuxtLink to="/register" class="text-primary">Register!</NuxtLink>
        </div>
    </div>
</template>

<script setup lang="ts">
import LoginForm from '~/components/LoginForm.vue';
import { useUserStore } from '~/stores/userStore';

const user = useUserStore()
</script>