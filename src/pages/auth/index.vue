<template>
    <div class="flex justify-center gap-20 my-20">
        <button @click="index">回主页</button>
        <button @click="authorize">{{ accessToken ? '已' : '未' }}登录Google</button>
        <input type="file" @change="handleFileChange"/>
        <button @click="uploadFile">上传文件到Google Drive</button>
    </div>
</template>
<script setup lang="ts">
defineOptions({name: 'Auth'});
import {onMounted, ref} from "vue";
let file = ref();
let accessToken = ref<string>('');
let item = localStorage.getItem("accessToken");
accessToken.value = item ? item : '';

const redirectUri = window.location.origin + window.location.pathname.replace(/\/+$/, '');

const index = () => {
    window.location.href = window.location.origin;
};
// 引导用户进行Google OAuth 2.0授权
const authorize = () => {
    if (accessToken.value) {
        console.log("authed")
        return
    }
    const clientId = '410659159953-8laduca307mq64f9u8pn6ebfgfvsl9ii.apps.googleusercontent.com';
    const scope = 'https://www.googleapis.com/auth/drive.file';
    window.location.href = `https://accounts.google.com/o/oauth2/v2/auth?response_type=code&client_id=${clientId}&redirect_uri=${redirectUri}&scope=${scope}&access_type=offline`;
};
// 获取URL中的Authorization Code
const getAuthorizationCode = () => {
    const params = new URLSearchParams(window.location.search);
    return params.get('code');
};
// 调用后端（Cloudflare Workers）以交换Authorization Code为Access Token
const exchangeAuthorizationCode = async () => {
    const authorizationCode = getAuthorizationCode();
    if (accessToken.value) {
        console.log("authed")
        if (authorizationCode) {
            console.log("has authCode")
            index();
        }
        return
    }

    if (authorizationCode) {
        try {
            const response = await fetch('https://dashboard.thesixonenine.workers.dev/', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({code: authorizationCode, redirectUri: redirectUri}),
            });
            const data = await response.json();
            if (data.msg) {
                console.warn('auth failed: ', data.msg);
                return
            }
            accessToken.value = data.access_token;
            localStorage.setItem("accessToken", accessToken.value);
            console.log('Access Token:', accessToken.value);
        } catch (error) {
            console.error('Failed to exchange token:', error);
        }
    }
};
const handleFileChange = (event: any) => {
    file.value = event.target.files[0];
}

// 上传文件到Google Drive
const uploadFile = async () => {
    if (!accessToken.value) {
        alert('请先登录Google并获取访问令牌');
        return;
    }

    const metadata = {
        name: file.value.name,
        mimeType: file.value.type
    };

    const form = new FormData();
    form.append('metadata', new Blob([JSON.stringify(metadata)], {type: 'application/json'}));
    form.append('file', file.value);

    try {
        const response = await fetch('https://www.googleapis.com/upload/drive/v3/files?uploadType=multipart', {
            method: 'POST',
            headers: {
                Authorization: `Bearer ${accessToken.value}`
            },
            body: form
        });
        const result = await response.json();
        console.log('文件上传成功，文件ID:', result.id);
    } catch (error) {
        console.error('文件上传失败:', error);
    }
}
onMounted(() => {
    // 在页面加载时检查是否有Authorization Code并交换为Access Token
    exchangeAuthorizationCode();
});
</script>
