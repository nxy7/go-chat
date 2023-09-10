<!-- Displays channel and allows user to join it -->
<template>
    <div class="shadow-md relative flex flex-col justify-between items-center w-full">
        <div class="flex p-4 pb-1 justify-between items-center w-full">
            <div>Row <span class="font-semibold">{{ name }}</span> </div>
            <NuxtLink :to="name" v-if="usersConnected < capacity" class="p-2 shadow-sm rounded-md bg-green-400">Connect
            </NuxtLink>
            <NuxtLink v-else class="p-2 shadow-sm rounded-md bg-gray-300">Channel full</NuxtLink>
        </div>
        <div class="flex w-full justify-end p-2 pt-0">
            <span class="text-sm">
                {{ usersConnected ?? 0 }}/{{ capacity ?? 10 }}
            </span>
        </div>
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    name: string,
    capacity: number,
    usersConnected: number,
    channelDescription?: string,
}>();
</script>