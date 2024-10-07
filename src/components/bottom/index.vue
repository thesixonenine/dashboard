<script setup lang="ts">

</script>

<template>
    <div class="main-container">
        <div class="flex flex-wrap justify-start gap-20 lg:gap-10 border border-red-500">
            <div class="flex flex-col text-2xl items-center  border border-red-500">
                <ul>
                    <li>
                        <div>
                            <a href="https://www.bilibili.com/video/BV1CN41167AC?p=21" target="_blank">参考视频</a>
                        </div>
                    </li>
                    <li>
                        <div>
                            <a href="https://www.bilibili.com/video/BV1CN41167AC?p=21" target="_blank">参考音频</a>
                        </div>
                    </li>
                </ul>
            </div>
            <div v-for="i in 5" class="flex-1 bg-green-100 text-2xl font-medium text-nowrap pl-12  border border-red-500">我的Blog{{ i }}</div>
        </div>
    </div>
</template>
