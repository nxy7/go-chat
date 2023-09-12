<!-- Displays channel and allows user to join it -->
<template>
    <div class="relative flex flex-col justify-between items-center w-full " :class="{ 'bg-gray-100': isOdd }">
        <div class="flex px-4 py-3 justify-between items-center w-full">
            <div class="flex flex-col items-center w-20 relative top-[5px]">
                <div><span class="font-semibold">{{ name }}</span> </div>
                <div>
                    <span class="text-xs relative bottom-3">
                        {{ usersConnected ?? 0 }}/{{ capacity ?? 10 }}
                    </span>
                </div>
            </div>
            <NuxtLink :to="id" v-if="usersConnected < capacity"
                class="p-2 shadow-sm rounded-md bg-emerald-700 font-semibold text-white">Connect
            </NuxtLink>
            <NuxtLink v-else class="p-2 shadow-sm rounded-md bg-gray-300">Channel full</NuxtLink>
        </div>
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    id: string,
    name: string,
    isOdd?: boolean,
    capacity: number,
    usersConnected: number,
    channelDescription?: string,
}>();
</script>