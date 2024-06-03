<script setup lang="ts">

</script>

<template>
  <a href="https://mp.weixin.qq.com/s/itAiycBBVPyudyU2ZJ2B_A" target="_blank">Flex布局</a>
  <br>
  <a href="https://github.com/dashboard" target="_blank">参考GitHub的Dashboard</a>
  <div class="container">
    <div class="l">Left</div>
    <div class="r">Right</div>
  </div>
</template>

<style scoped lang="scss">
.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  //align-items: center;
  .l {
    height: 100px;
    width: 100px;
  }
  .r {
    height: 100px;
    width: 100px;
    //line-height: 100px;
    //flex: 1;
  }
}
</style>