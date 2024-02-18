// 封装 axios
// 使用请求/响应拦截器

import axios from 'axios';
import {ElMessage} from "element-plus";

const request = axios.create({
    baseURL: '/api',
    timeout: 5000,
});
// 请求拦截器
request.interceptors.request.use((config) => {
    return config;
});

// 响应拦截器
request.interceptors.response.use(
    (resp) => {
        // 成功的回调
        return resp.data;
    }, (error) => {
        // 失败的回调
        let status = error.response.status;
        switch (status) {
            case 400: {
                ElMessage({
                    type: 'error',
                    message: '400-' + error.message,
                });
                break;
            }
            case 404: {
                ElMessage({
                    type: 'error',
                    message: '404-' + error.message,
                });
                break;
            }
            case 500:
            case 501:
            case 502:
            case 503:
            case 504:
            case 505: {
                ElMessage({
                    type: 'error',
                    message: '50x-' + error.message,
                });
                break;
            }
        }
        return Promise.reject(new Error(error.message));
    });

export default request;
