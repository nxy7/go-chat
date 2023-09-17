<template>
    <div class="flex justify-center items-center rounded-md p-3 w-60 min-h-12 shadow-md bg-slate-50 border-b-4 duration-200 z-50"
        :class="[{ 'scale-0': hidden }, { 'border-yellow-400': Level == 'INFO' }, { 'border-orange-400': Level == 'WARNING' }, { 'border-red-400': Level == 'ERROR' }]">
        {{
            props.Content }}</div>
</template>

<script setup lang="ts">


const props = defineProps<{
    Level: "INFO" | "WARNING" | "ERROR"
    Content: string
}>()
let hidden = ref(true)
setTimeout(() => {
    hidden.value = false
}, 5)
</script>