<script setup lang="ts">
defineOptions({name: 'Card'});
import {useRouter} from "vue-router";
// 接受传入的props itemInfo
let props = defineProps(['itemInfo']);
// 路由对象
let $router = useRouter();
// 点击卡片
const selected = (t: any) => {
  console.log(t.value);
  console.log('props:', props);
  console.log('itemInfo:', props.itemInfo);
  $router.push({path: '/hospital/reg', query: {hoscode: props.itemInfo.hoscode}});
};
</script>

<template>
  <el-card class="box-card" shadow="hover" @click="selected">
    <div class="content">
      <div class="left">
        <div class="name">{{ itemInfo.hosname }}</div>
        <div class="info">
          <div class="star">❤❤❤❤❤</div>
          <div class="type">{{ itemInfo.param?.hostypeString }}</div>
        </div>
      </div>
      <div class="right">
        <!-- <img src="@/assets/img/home/card/kafka.png" alt=""/> -->
        <img :src="`data:image/jpeg;base64,${itemInfo.logoData}`" alt=""/>
      </div>
    </div>
  </el-card>
</template>

<style scoped lang="scss">
.content {
  display: flex;
  justify-content: space-between;

  .left {
    width: 60%;

    .info {
      color: #7f7f7f;
      margin-top: 20px;
      display: flex;
      justify-content: space-between;
    }
  }

  .right {
    img {
      width: 50px;
      height: 50px;
    }
  }
}
</style>
