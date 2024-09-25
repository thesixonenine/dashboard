<script setup lang="ts">
defineOptions({name: 'Board'});
import point from '../point/index.vue';
</script>

<template>
<div class="board">
  <div v-for="n in 25">
    {{n}}
    <point/>
  </div>
</div>
</template>

<style scoped lang="scss">
.board{
  //border: #55a6fe solid 1px;
  height: 100px;
  width: 100px;
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}
</style>