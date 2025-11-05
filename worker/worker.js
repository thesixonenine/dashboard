/**
 * Welcome to Cloudflare Workers! This is your first worker.
 *
 * - Run "npm run dev" in your terminal to start a development server
 * - Open a browser tab at http://localhost:8787/ to see your worker in action
 * - Run "npm run deploy" to publish your worker
 *
 * Learn more at https://developers.cloudflare.com/workers/
 */

import { generateSignature } from "./signature.js";
import { resend } from "./email.js";

export default {
  async fetch(request, env, ctx) {
    return await handleRequest(request, env, ctx);
  },
};

async function handleRequest(request, env, ctx) {
//   if (request.method === 'OPTIONS') {
//     return handleOptions(request);
//   }
  // 邮件处理
  if (request.url.includes('/email/notify')){
    // 验证签名
    if (request.method === 'POST') {
      return await resend(env, `<h1>Hello!</h1><p>This email was sent via Cloudflare Worker and Resend API.</p>`);
    }
  }
  // 公众号相关的业务
  else if (request.url.includes('/qbot/notify')){
    // 验证签名
    if (request.method === 'GET') {
      return handleMPSign(request);
    } 
    // QBot消息
    else if (request.method === 'POST') {
      return await handleQBotSign(request, env);
    }
  }
  // google drive
//   else if (request.method === 'POST' ) {
//     return await exchangeToken(request, env);
//   }
  return env.ASSETS.fetch(request);
}

async function handleQBotSign(request, env) {
  const { d, op } = await request.json();
  console.log(d);
  console.log(op);
  const pt = d.plain_token;
  const et = d.event_ts;
  const seed = env.QBOT_SEED;

  const signature = generateSignature(seed, et, pt);
  return new Response('{"plain_token": "'+pt+'", "signature": "'+signature+'"}', {status: 200});
}



function handleMPSign(request) {
  const url = new URL(request.url);
  const signature = url.searchParams.get('signature');
  const timestamp = url.searchParams.get('timestamp');
  const nonce = url.searchParams.get('nonce');
  const echostr = url.searchParams.get('echostr');
  return new Response(echostr, {status: 200});
}

// 处理Authorization Code交换为Access Token的请求
async function exchangeToken(request, env) {
  const { code, redirectUri } = await request.json();
  const clientId = env.GOOGLE_CLIENT_ID;
  const clientSecret = env.GOOGLE_CLIENT_SECRET;

  const tokenUrl = 'https://oauth2.googleapis.com/token';
  const tokenParams = new URLSearchParams();
  tokenParams.append('code', code);
  tokenParams.append('client_id', clientId);
  tokenParams.append('client_secret', clientSecret);
  tokenParams.append('redirect_uri', redirectUri);
  tokenParams.append('grant_type', 'authorization_code');

  try {
    const tokenResponse = await fetch(tokenUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      body: tokenParams
    });

    if (!tokenResponse.ok) {
      return new Response('{"msg": "Failed to exchange token"}', { headers:{...corsHeaders()},status: 400 });
    }

    const tokenData = await tokenResponse.json();

    // 返回Access Token给前端
    return new Response(JSON.stringify(tokenData), {
      headers: { 'Content-Type': 'application/json',...corsHeaders() },
      status: 200
    });

  } catch (error) {
    return new Response('{"msg": "Internal Server Error"}', { headers:{...corsHeaders()},status: 500 });
  }
}

// 处理OPTIONS请求
function handleOptions(request) {
  return new Response(null, {
    headers: corsHeaders()
  });
}

// 添加CORS头的辅助函数
function corsHeaders() {
  return {
    'Access-Control-Allow-Origin': '*',
    'Access-Control-Allow-Methods': '*',
    'Access-Control-Allow-Headers': '*',
    'Access-Control-Max-Age': '86400'
  };
}

