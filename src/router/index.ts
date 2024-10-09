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
            path: '/dashboard',
            component: () => import('@/pages/dashboard/index.vue'),
        },
        {
            path: '/auth',
            component: () => import('@/pages/auth/index.vue'),
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
