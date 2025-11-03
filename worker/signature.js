import nacl from "https://esm.sh/tweetnacl@1.0.3";

/**
 * Ed25519 签名函数（Cloudflare Workers 兼容）
 * @param {string} seedStr - 种子字符串
 * @param {string} et - 消息前半部分
 * @param {string} pt - 消息后半部分
 * @returns {string} - 签名（十六进制字符串）
 */
export function generateSignature(seedStr, et, pt) {
  const SEED_SIZE = 32;
  let seed = seedStr;
  while (seed.length < SEED_SIZE) seed = seed.repeat(2);
  seed = seed.slice(0, SEED_SIZE);

  const encoder = new TextEncoder();
  const seedBytes = encoder.encode(seed);

  const keyPair = nacl.sign.keyPair.fromSeed(seedBytes);
  const msg = encoder.encode(et + pt);
  const sig = nacl.sign.detached(msg, keyPair.secretKey);

  // 转成十六进制字符串
  return Array.from(sig)
    .map((b) => b.toString(16).padStart(2, "0"))
    .join("");
}
