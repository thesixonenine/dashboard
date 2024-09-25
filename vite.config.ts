import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
// 引入node提供的path,获取绝对路径
import path from 'path';
import tailwindcss from 'tailwindcss'
import autoprefixer from 'autoprefixer'

// https://vitejs.dev/config/
export default defineConfig({
    base: '/',
    plugins: [vue()],
    resolve: {
        alias: {
            "@": path.resolve(__dirname, 'src')
        }
    },
    css: {
        postcss: {
            plugins: [tailwindcss, autoprefixer]
        }
    },
    // 配置代理跨域
    server: {
        proxy: {
            '/api': {
                target: 'http://syt.atguigu.cn',
                changeOrigin: true,
                // rewrite: (path) => path.replace(/^\/api/, ''),
            }
        }
    }
})
