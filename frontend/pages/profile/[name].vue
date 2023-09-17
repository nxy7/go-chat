<template>
    <div class="!h-full grow w-full m-auto">
        <div v-if="!pending" class="flex flex-col items-center h-full justify-around">
            <div class="text-4xl">
                <span class="font-semibold ">{{ data.Name }}</span>
            </div>
            <div>
                sent {{ data.MessageCount }} messages so far :3
            </div>
        </div>
        <div v-else class="animate-pulse w-full h-full flex items-center justify-center">
            <img src="https://raw.githubusercontent.com/n3r4zzurr0/svg-spinners/abfa05c49acf005b8b1e0ef8eb25a67a7057eb20/svg-css/ring-resize.svg"
                class="h-12 w-12 text-blue-700" alt="">
        </div>
    </div>
</template>

<script setup lang="ts">
import { useUserStore } from '~/stores/userStore';
const user = useUserStore()
const route = useRoute()
console.log(route.params.name)

const { pending, data } = useFetch(`/api/user/details/${route.params.name}`, {
    server: false,
    transform: (d: string) => {
        let parsed: any = JSON.parse(d)
        console.log(parsed)
        return parsed
    },
})
watch(data, () => {
    console.log(data)
})
</script>