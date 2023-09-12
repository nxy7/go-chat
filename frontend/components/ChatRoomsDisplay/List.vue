<template>
  <div class="duration-300 w-80 flex flex-col items-center justify-center shadow-lg">
    <div class="p-3 font-semibold">Simple Chat App</div>
    <div class="w-full h-96 overflow-y-scroll">
      <div v-if="pending" class="animate-pulse w-full h-full flex items-center justify-center bg-gray-100">
        <img
          src="https://raw.githubusercontent.com/n3r4zzurr0/svg-spinners/abfa05c49acf005b8b1e0ef8eb25a67a7057eb20/svg-css/ring-resize.svg"
          class="h-12 w-12 text-blue-700" alt="">
      </div>
      <ChatRoomsDisplayRow v-else v-for="i, ind in data" :isOdd="ind % 2 == 0" :id="i.id" :name="i.name"
        :capacity=i.capacity :usersConnected=i.usersConnected>
      </ChatRoomsDisplayRow>
    </div>
    <div class="p-2 w-full flex justify-between">
      <div>
        <a v-if="!userStore.user" href="/login">Click here to log in..</a>
        <div v-else>Elo {{ userStore.user?.name }}</div>
      </div>
      <div class="self-end justify-self-end place-self-end">
        <button class="w-6 bg-gray-300 rounded-full flex items-center justify-center font-bold text-white ">?</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '~/stores/userStore';

const userStore = useUserStore()
console.log(userStore.user)

type Channel = {
  id: string,
  name: string,
  capacity: number,
  usersConnected: number
}
const { pending, data } = useFetch("/api/channel", {
  server: false,
  transform: (d: string) => {
    let parsed: any[] = JSON.parse(d)


    return parsed.map<Channel>((o) => {
      o.usersConnected = o.activeUsers.length
      return o
    })
  },
})
watch(data, (d) => {
  console.log(d)
})

</script>