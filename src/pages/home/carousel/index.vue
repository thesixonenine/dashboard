<script setup lang="ts">
defineOptions({name: 'Carousel'});
</script>

<template>
  <div class="carousel">
    <el-carousel height="400px">
      <el-carousel-item>
        <img src="@/assets/img/home/banner/banner1.png" alt="">
      </el-carousel-item>
      <el-carousel-item>
        <img src="@/assets/img/home/banner/banner2.png" alt="">
      </el-carousel-item>
      <el-carousel-item>
        <img src="@/assets/img/home/banner/banner3.png" alt="">
      </el-carousel-item>
      <el-carousel-item>
        <img src="@/assets/img/home/banner/banner4.png" alt="">
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<style scoped lang="scss">
.carousel {
  width: 70%;
  height: 400px;

  img {
    width: 100%;
    height: 100%;
  }
}

</style>