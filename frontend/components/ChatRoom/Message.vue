<template>
    <div class="flex flex-col">
        <div v-if="previousMessageAuthor != authorName" class="text-sm mx-3"
            :class="[authorMessage ? 'items-end place-items-end self-end' : 'items-start']">
            <NuxtLink :to="'/profile/' + authorName">
                {{ authorName }}
            </NuxtLink>
        </div>
        <div class="flex items-center gap-3" :class="[authorMessage ? 'justify-end' : 'justify-start']">
            <!-- <div class="h-9 w-9 rounded-full" :class="[authorMessage ? 'bg-accent' : 'bg-slate-200']"></div> -->
            <div class="flex px-3 py-1 rounded-full" :class="[authorMessage ? 'bg-accent' : 'bg-slate-200']">{{
                props.content }}</div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useUserStore } from '~/stores/userStore';

const props = defineProps<{
    authorName: string
    previousMessageAuthor?: string
    content: string
}>()

const user = useUserStore()
let authorMessage = user.jwt?.user == props.authorName ? true : false
</script>