<script setup lang="ts">

</script>

<template>
  <div class="bottom">
    <div class="content">
      <div class="left">
        <a href="https://www.bilibili.com/video/BV1CN41167AC?p=21" target="_blank">参考视频</a>
      </div>
      <div class="right">
        <span>我的Blog</span>
        <span>我的GitHub</span>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.bottom {
  margin-top: 10px;
  width: 100%;
  height: 50px;
  background-color: #f0f2f5;
  display: flex;
  justify-content: center;

  .content {
    width: 1200px;
    height: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 14px;
    color: black;

    .right {
      span {
        margin: 0 5px;
      }
    }
  }
}
</style>
