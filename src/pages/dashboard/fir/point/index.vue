<script setup lang="ts">

</script>

<template>
  <div class="point">
    <div class="one mp"></div>
    <div class="two mp"></div>
    <div class="three mp"></div>
    <div class="four mp"></div>
  </div>
</template>

<style scoped lang="scss">
.point {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  height: 10px;
  width: 10px;
  .mp {
    width: 5px;
    height: 5px;
    border: black solid 1px;
  }
  //.one {
  //  border-right-style: solid;
  //  border-right-width: 1px;
  //  border-bottom-style: solid;
  //  border-bottom-width: 1px;
  //}
  //.two {
  //  border-left-style: solid;
  //  border-left-width: 1px;
  //  border-bottom-style: solid;
  //  border-bottom-width: 1px;
  //}
  //.three {
  //  border-top-style: solid;
  //  border-top-width: 1px;
  //  border-right-style: solid;
  //  border-right-width: 1px;
  //}
  //.four {
  //  border-top-style: solid;
  //  border-top-width: 1px;
  //  border-left-style: solid;
  //  border-left-width: 1px;
  //}
}
</style>