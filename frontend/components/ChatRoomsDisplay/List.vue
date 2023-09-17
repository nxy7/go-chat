<template>
  <div class="w-full h-96 overflow-y-scroll">
    <div v-if="pending" class="animate-pulse w-full h-full flex items-center justify-center bg-gray-100">
      <img
        src="https://raw.githubusercontent.com/n3r4zzurr0/svg-spinners/abfa05c49acf005b8b1e0ef8eb25a67a7057eb20/svg-css/ring-resize.svg"
        class="h-12 w-12 text-blue-700" alt="">
    </div>
    <ChatRoomsDisplayRow v-else v-for="i, ind in data" :isOdd="ind % 2 == 0" :id="i.id" :name="i.name"
      :capacity=i.capacity :listenerCount=i.listenerCount>
    </ChatRoomsDisplayRow>
  </div>
</template>

<script setup lang="ts">
type Channel = {
  id: string,
  name: string,
  capacity: number,
  listenerCount: number
}

const { pending, data } = useFetch("/api/channel", {
  server: false,
  transform: (d: string) => {
    let parsed: any[] = JSON.parse(d)
    console.log(parsed)
    return parsed
  },
})

watch(data, (d) => {
  console.log(d)
})

</script>