<script setup lang="ts">
defineOptions({name: 'Search'});
import {cardByNameReq} from "@/api/home.ts";
// 引入图标
import {Search} from "@element-plus/icons-vue";
import {ref} from "vue";
import {Response} from "@/api/common_type.ts";
import {Content} from "@/api/home_type.ts";
import {useRouter} from "vue-router";
// 路由对象
let $router = useRouter();
// 定义搜索字段
let inputName = ref<string>('');
// 查询搜索建议
const fetchSuggest = async (keyword: string, cb: any) => {
  console.log("keyword:", keyword);
  const resp: Response<Content[]> = await cardByNameReq(keyword);
  if (resp.code == 200) {
    console.log('suggestions:', resp.data.map(t => t.hosname).join('\n'));
    cb(resp.data.map(t => {
      return {value: t.hosname, hoscode: t.hoscode}
    }));
  }
};
// 点击推荐项
const selected = (t: any) => {
  console.log(t.value);
  $router.push({path: '/hospital/reg', query: {hoscode: t.hoscode}});
};
</script>

<template>
  <div class="search">
    <el-autocomplete
        clearable
        class="inline-input w-50"
        placeholder="请输入角色名称"
        v-model="inputName"
        :fetch-suggestions="fetchSuggest"
        :trigger-on-focus="false"
        @select="selected"
    />
    <el-button type="primary" size="default" @click="" :icon="Search">搜索</el-button>
  </div>
</template>

<style scoped lang="scss">
.search {
  width: 100%;
  height: 50px;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 10px 0;
  // 深度选择器(父组件改变子组件的样式)
  :deep(.el-input__wrapper) {
    width: 300px;
    margin-right: 10px;
  }
}
</style>
