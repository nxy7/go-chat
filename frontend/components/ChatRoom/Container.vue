<template>
    <div class="flex flex-col">
        <div class="flex">
            <NuxtLink to="/">Go Back</NuxtLink>
            <div>top bar</div>
        </div>
        <div class="flex">
            <div>sidebar</div>
            <div class="flex flex-col justify-between max-h-80">
                <div class="h-full overflow-scroll min-h-[10rem]">messages</div>
                <div>new message box</div>
            </div>
        </div>
        <div>
            footer
        </div>
    </div>
</template>