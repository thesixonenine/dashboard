import {createApp} from 'vue';
// 引入清除默认样式
import '@/assets/style/reset.scss';
// 引入根组件App
import App from '@/App.vue';
// 引入顶部和底部的全局组件
import Top from '@/components/top/index.vue';
import Bottom from '@/components/bottom/index.vue';
// 引入 vue-router
import router from '@/router/index.ts';
// 引入 element-plus
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
// 国际化组件
//@ts-ignore
import zhCn from 'element-plus/dist/locale/zh-cn.mjs';
// 引入 pinia
import pinia from '@/store';
import '@/tailwind.css'

const app = createApp(App);
app.component('Top', Top);
app.component('Bottom', Bottom);
// 安装 vue-router
app.use(router);
// 安装 element-plus
app.use(ElementPlus, {
    locale: zhCn,
})
// 安装 pinia
app.use(pinia);
// 挂载
app.mount('#app');
