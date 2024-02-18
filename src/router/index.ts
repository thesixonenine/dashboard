import {createRouter, createWebHistory} from 'vue-router'
// 创建路由器实例
export default createRouter({
    // 路由模式
    history: createWebHistory(),
    // 管理路由
    routes: [
        {
            path: '/home',
            component: () => import('@/pages/home/index.vue'),
        },
        {
            path: '/sr',
            component: () => import('@/pages/sr/index.vue'),
        },
        {
            path: '/',
            redirect: '/home',
        }
    ],
    // 控制滚动条的位置
    scrollBehavior() {
        return {
            left: 0,
            top: 0
        }
    }
})