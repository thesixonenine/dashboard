<script setup lang="ts">
defineOptions({name: 'Top'})
import {useRouter} from "vue-router";

let $router = useRouter();
const home = () => {
  $router.push({path: '/home'});
};
</script>

<template>
  <div class="top">
    <div class="container">
      <div class="left" @click="home">
        <img class="logo" src="@/assets/img/ts.jpg" alt="">
        <p>Dashboard</p>
      </div>
      <div class="right">
        <p class="github">GitHub</p>
        <p class="blog">Blog</p>
      </div>
    </div>

  </div>
</template>

<style scoped lang="scss">
.top {
  position: fixed;
  z-index: 999;
  width: 100%;
  height: 70px;
  background-color: white;
  display: flex;
  justify-content: center;

  .container {
    width: 1200px;
    height: 70px;
    background-color: white;
    display: flex;
    justify-content: space-between;

    .left {
      display: flex;
      justify-content: center;
      align-items: center;
      cursor: pointer;

      .logo {
        width: 50px;
        height: 50px;
        margin-right: 10px;
      }

      p {
        font-size: 20px;
        color: #55a6fe;
      }
    }

    .right {
      display: flex;
      justify-content: center;
      align-items: center;
      font-size: 14px;
      color: #bbb;

      .github {
        margin-right: 10px;
      }
    }
  }
}
</style>