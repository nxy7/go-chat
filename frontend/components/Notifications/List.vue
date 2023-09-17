<template>
    <div class="absolute top-3 flex flex-col gap-3">
        <NotificationsRow v-for="n in notificationStore.notificationArr" v-bind="n"></NotificationsRow>
    </div>
</template>

<script setup lang="ts">
import { useNotificationStore } from '~/stores/notificationStore';

const notificationStore = useNotificationStore()
// if (process.client)
//     notificationStore.pushNotification({ Level: "INFO", Content: "start" }, 20)

</script>