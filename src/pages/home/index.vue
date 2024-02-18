<script setup lang="ts">
defineOptions({name: 'Home'});
// 引入轮播图
import Carousel from './carousel/index.vue';
// 引入搜索
import Search from './search/index.vue';
// 引入等级
import Level from './level/index.vue';
// 引入地区
import Region from './region/index.vue';
// 引入角色卡片
import Card from './card/index.vue';
// 引入右侧静态
import Tip from './tip/index.vue';
// 引入组合式API函数
import {onMounted, ref} from 'vue';
import {cardListReq} from '@/api/home';
import type {Content, Hospital} from "@/api/home_type.ts";
import type {Response} from "@/api/common_type.ts";

let pageNo = ref<number>(1);
let pageSize = ref<number>(5);

let hasCardArr = ref<Content[]>([]);
let total = ref<number>(0);
// 子组件的等级数据
let hostype = ref<string>('');
let districtCode = ref<string>('');
onMounted(() => {
  getCardList()
});

const getCardList = async () => {
  let result: Response<Hospital> = await cardListReq(pageNo.value, pageSize.value, hostype.value, districtCode.value);
  console.log('getCardList:\n', result);
  if (result.code == 200) {
    hasCardArr.value = result.data.content;
    total.value = result.data.totalElements;
  }
};

// 分页
const currentChange = () => {
  console.log('页码发生变化');
  getCardList();
};
const sizeChange = () => {
  console.log('页大小发生变化');
  pageNo.value = 1;
  getCardList();
};
// 子组件自定义事件, 获取子组件给当前父组件传递的数据
const getLevel = (levelCode: string) => {
  console.log('已接收到子组件传递的等级数据', levelCode);
  hostype.value = levelCode;
  getCardList();
};
const getRegion = (region: string) => {
  console.log('已接收到子组件传递的地区数据', region);
  districtCode.value = region;
  getCardList();
};
</script>

<template>
  <div>
    <!-- 首页轮播图 -->
    <Carousel/>
    <!-- 首页搜索 -->
    <Search/>
    <!-- 内容展示 -->
    <el-row :gutter="20">
      <el-col :span="20">
        <!-- 星级, 需要绑定自定义事件, 以便接收该子组件上传的数据 -->
        <Level @getLevel="getLevel"/>
        <!-- 地区 -->
        <Region @getRegion="getRegion"/>
        <!-- 角色 -->
        <div class="cardList" v-if="hasCardArr.length>0">
          <Card class="card" v-for="(item,index) in hasCardArr" :key="index" :itemInfo="item"/>
        </div>
        <el-empty v-if="hasCardArr.length==0" description="这里啥也没"/>
        <!-- 分页器 -->
        <el-pagination
            v-model:current-page="pageNo"
            v-model:page-size="pageSize"
            :page-sizes="[5, 6, 7, 8]"
            :small="true"
            :background="true"
            layout="prev, pager, next, jumper, ->, total, sizes"
            :total="total"
            @current-change="currentChange"
            @size-change="sizeChange"
        />
      </el-col>
      <el-col :span="4">
        <Tip/>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped lang="scss">
.cardList {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;

  .card {
    width: 48%;
    margin: 10px 0;
  }
}
</style>