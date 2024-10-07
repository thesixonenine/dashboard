<script setup lang="ts">
import {onMounted} from "vue";

onMounted(() => {
    console.log('App.vue Mounted');
});
</script>

<template>
    <main class="relative overflow-hidden">
<!--        <Top/>-->
<!--        <div class="content">-->
<!--            <router-view/>-->
<!--        </div>-->
        <footer class="padding-x padding-t bg-gray-300 pb-8">
            <Bottom/>
        </footer>
    </main>
</template>
