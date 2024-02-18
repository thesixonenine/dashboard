<script setup lang="ts">
defineOptions({name: 'Level'});
import {onMounted, ref} from "vue";
import {cardLevelReq} from "@/api/home.ts";
import {HospitalDict} from "@/api/home_type.ts";
import type {Response} from "@/api/common_type.ts";
// 等级数据
let levelList = ref<HospitalDict[]>([]);
// 高亮控制
let activated = ref<string>('');
onMounted(() => {
  levelData();
});
// 获取医院的等级数据
const levelData = async () => {
  let resp: Response<HospitalDict[]> = await cardLevelReq('Hostype');
  console.log("cardLevel:\n", resp);
  if (resp.code == 200) {
    levelList.value = resp.data;
  }
};
// 点击事件
const changeLevel = (levelCode: string) => {
  // 改变选中状态
  activated.value = levelCode;
  // 触发自定义事件, 将level的参数传递给父组件
  $emit('getLevel', levelCode);
};
// 绑定父组件中的自定义事件
let $emit = defineEmits(['getLevel']);
</script>

<template>
  <div class="level">
    <h1>角色筛选</h1>
    <div class="content">
      <div class="left">星级:</div>
      <ul class="star">
        <li :class="{active:activated==''}" @click="changeLevel('')">全部</li>
        <li v-for="level in levelList"
            :key="level.value"
            @click="changeLevel(level.value)"
            :class="{active:activated==level.value}">{{ level.name }}
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped lang="scss">
.level {
  color: #7f7f7f;

  h1 {
    font-weight: 900;
    margin: 10px 0;
  }

  .content {
    display: flex;

    .left {
      margin-right: 10px;
    }

    .star {
      display: flex;

      li {
        margin-right: 10px;

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