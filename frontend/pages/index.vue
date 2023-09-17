<template>
  <div class="flex flex-col w-full">
    <ChatRoomsDisplayList></ChatRoomsDisplayList>
    <div class="p-2 w-full flex justify-between">
      <div>
        <span v-if="!userStore.jwt">
          <NuxtLink to="/login" class="text-primary">Log in</NuxtLink> or <NuxtLink to="/register" class="text-primary">
            Register</NuxtLink>
        </span>
        <div v-else>Hello <NuxtLink :to="'/profile/' + userStore.jwt.user">{{ userStore.jwt.user }}</NuxtLink>! <button
            class="text-primary" @click="userStore.logout">Logout</button>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '~/stores/userStore';

const userStore = useUserStore()
</script>