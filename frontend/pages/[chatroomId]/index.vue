<template>
    <div>
        <ClientOnly>
            <ChatRoomContainer :channelId="chatroomId"></ChatRoomContainer>
        </ClientOnly>
    </div>
</template>

<script setup lang="ts">
const route = useRoute()
let chatroomId = route.params.chatroomId as string
</script>