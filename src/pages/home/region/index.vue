<script setup lang="ts">
defineOptions({name: 'Region'});
import {onMounted, ref} from "vue";
import {cardLevelReq} from "@/api/home.ts";
import {HospitalDict} from "@/api/home_type.ts";
import type {Response} from "@/api/common_type.ts";
// 区域数据
let regionList = ref<HospitalDict[]>([]);
// 高亮控制
let activated = ref<string>('')
onMounted(() => {
  regionData();
});
// 获取医院的区域数据
const regionData = async () => {
  let resp: Response<HospitalDict[]> = await cardLevelReq('beijin');
  console.log("cardRegion:\n", resp);
  if (resp.code == 200) {
    regionList.value = resp.data;
  }
};
// 点击事件
const changeRegion = (region: string) => {
  // 改变选中状态
  activated.value = region;
  $emit('getRegion', region);
};
let $emit = defineEmits(['getRegion']);
</script>

<template>
  <div class="region">
    <div class="content">
      <div class="left">地区:</div>
      <ul class="star">
        <li :class="{active:activated==''}" @click="changeRegion('')">全部</li>
        <li v-for="region in regionList"
            :key="region.value"
            @click="changeRegion(region.value)"
            :class="{active:activated==region.value}">{{ region.name }}
        </li>
        <!--<li>银河</li>-->
        <!--<li>仙舟「罗浮」</li>-->
        <!--<li>贝洛伯格</li>-->
        <!--<li>空间站「黑塔」</li>-->
        <!--<li>星穹列车</li>-->
        <!--<li>星核猎手</li>-->
      </ul>
    </div>
  </div>
</template>

<style scoped lang="scss">
.region {
  color: #7f7f7f;
  margin-top: 10px;

  h1 {
    font-weight: 900;
    margin: 10px 0;
  }

  .content {
    display: flex;

    .left {
      margin-right: 10px;
      width: 36px;
    }

    .star {
      display: flex;
      flex-wrap: wrap;

      li {
        margin-right: 10px;
        margin-bottom: 10px;

        &.active {
          color: #55a6fe;
        }
      }

      li:hover {
        color: #55a6fe;
        cursor: pointer;
      }
    }
  }
}
</style>