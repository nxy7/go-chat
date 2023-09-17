<template>
    <div
        class="absolute flex flex-col justify-center items-center h-screen w-screen bg-white z-30 bg-opacity-20 backdrop-blur-sm">
        <div class="bg-white flex flex-col justify-center items-center gap-3 w-[40rem] py-7 shadow-md relative p-4">
            <button
                class="absolute top-2 right-2 flex items-center justify-center h-6 w-6 bg-primary text-white rounded-full font-extrabold"
                @click="$emit('close-overlay')">X</button>

            <div class="flex flex-col gap-4">
                <div>
                    <span class="font-semibold">Features:</span>
                    <ul class="list-disc ml-5">
                        <li>Send messages to channels (like in messenger group chat)</li>
                        <li>Messages dissapear after some time (display duration depends on channel that the emssage was
                            sent
                            to)</li>
                        <li>Show user message count (click on any nickname in conversation)</li>
                        <li>Channels have limited number of connected users</li>
                        <li>Custom authentications using JWTs</li>
                        <li>Backend and frontend can be scaled independently (backend is stateless)</li>
                    </ul>
                </div>
                <div>
                    <span class="mt-6 font-semibold">
                        Tech used:
                    </span>
                    <ul class="list-disc ml-5">
                        <li>
                            Nuxt (Vue3) - Frontend is using Nuxt SSG to generate html and take some pressure of the server.
                        </li>
                        <li>
                            Golang - Backend is built entirely in golang. Backend is responsible for authentication,
                            authorization and serving data to frontend. I've designed backend to be totally stateless, so
                            many replicas can be spun up to serve the traffic.
                        </li>
                        <li>
                            Redis - database associated with most chat functionality. New chat messages are posted through
                            pubsub, so other users can see them and they're also saved separately so they're persisted for
                            desired amount of time (duration depends on channel).
                        </li>
                        <li>
                            MongoDB - Stores user profile, including username, password hash and count of messages sent by
                            user.
                        </li>
                        <li>
                            Kubernetes - the app operates inside kubernetes cluster for easy deployment process.
                        </li>
                    </ul>


                </div>
            </div>
        </div>

    </div>
</template>

<script setup lang="ts">
</script>