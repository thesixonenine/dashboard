<script setup lang="ts">
// import HelloWorld from './components/HelloWorld.vue'
// import WishList from '@/components/WishList.vue'
import {onMounted} from "vue";

onMounted(() => {
  console.log('App.vue Mounted');
});
</script>

<template>
  <div class="container">
    <Top/>
    <div class="content">
      <router-view/>
    </div>
    <Bottom/>
  </div>
  <!-- <HelloWorld msg="Dashboard" /> -->
</template>

<style scoped lang="scss">
.container {
  display: flex;
  align-items: center;
  flex-direction: column;

  .content {
    margin-top: 70px;
    width: 1200px;
    min-height: 700px;
  }
}
</style>
