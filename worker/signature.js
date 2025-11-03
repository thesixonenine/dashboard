import nacl from "tweetnacl";

/**
 * 基于字符串 seed（截断/重复至32字节）生成 ed25519 签名（十六进制）
 * @param {string} seedStr
 * @param {string} et
 * @param {string} pt
 * @returns {string}
 */
export function generateSignature(seedStr, et, pt) {
  const SEED_SIZE = 32;
  let seed = seedStr;
  while (seed.length < SEED_SIZE) seed = seed.repeat(2);
  seed = seed.slice(0, SEED_SIZE);

  const encoder = new TextEncoder();
  const seedBytes = encoder.encode(seed);

  // nacl 要求 seed 必须是 32 byte 的 Uint8Array
  // 如果 seedBytes 不是正好32（例如有多字节 Unicode），
  // 我们应该把 seed 转为 ASCII bytes — 但你的 seed 看起来是 ASCII，所以 encoder.encode 合适。
  // 为了保险起见，确保 seedBytes 长度为 32：
  let seed32;
  if (seedBytes.length === 32) {
    seed32 = seedBytes;
  } else if (seedBytes.length > 32) {
    seed32 = seedBytes.slice(0, 32);
  } else {
    // pad with zeros (一般不应该发生因为上面已按字符长度调整)
    seed32 = new Uint8Array(32);
    seed32.set(seedBytes);
  }

  const keyPair = nacl.sign.keyPair.fromSeed(seed32);
  const msg = encoder.encode(et + pt);
  const sig = nacl.sign.detached(msg, keyPair.secretKey);

  // 转 hex
  return Array.from(sig).map((b) => b.toString(16).padStart(2, "0")).join("");
}
