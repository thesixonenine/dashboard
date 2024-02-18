<script setup lang="ts">
import {onMounted} from "vue";
import wishJson from '@/assets/data/star-rail-wish.json'

class WishItem {
  count: string
  gacha_id: string
  gacha_type: string
  id: string
  item_id: string
  item_type: string
  lang: string
  name: string
  rank_type: string
  time: string
  uid: string

  constructor(count: string, gacha_id: string, gacha_type: string, id: string, item_id: string, item_type: string, lang: string, name: string, rank_type: string, time: string, uid: string) {
    this.count = count;
    this.gacha_id = gacha_id;
    this.gacha_type = gacha_type;
    this.id = id;
    this.item_id = item_id;
    this.item_type = item_type;
    this.lang = lang;
    this.name = name;
    this.rank_type = rank_type;
    this.time = time;
    this.uid = uid;
  }
}

let wishList = new Array<WishItem>();
for (const idx in wishJson[1]) {
  const dx = <WishItem>(wishJson[1][idx]);
  wishList.push(dx);
}
for (const idx in wishJson[2]) {
  const dx = <WishItem>(wishJson[2][idx]);
  wishList.push(dx);
}
for (const idx in wishJson[11]) {
  const dx = <WishItem>(wishJson[11][idx]);
  wishList.push(dx);
}
for (const idx in wishJson[12]) {
  const dx = <WishItem>(wishJson[12][idx]);
  wishList.push(dx);
}
console.log("length = " + wishList.length);
onMounted(() => {
  console.log("WishList.vue is now mounted.")
})
</script>

<template>
  <div class="wishes">
    <table>
      <thead>
      <tr>
        <th>时间</th>
        <th>名称</th>
        <th>类型</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="wishItem in wishList">
        <td id="wishTime">{{ wishItem.time }}</td>
        <td v-if="wishItem.rank_type === '3'" id="wishName">{{ wishItem.name }}</td>
        <td v-if="wishItem.rank_type === '4'" style="color: blue" id="wishName">{{ wishItem.name }}</td>
        <td v-if="wishItem.rank_type === '5'" style="color: red" id="wishName">{{ wishItem.name }}</td>
        <td id="wishStar">{{ wishItem.item_type }}</td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.wishes {
  display: flex;
  flex-direction: column;
  background-color: white;
  padding: 10px;
  width: 100%;
}

table {
  background-color: wheat;
  border: solid 1px;
}

th {
  padding: 5px 0;
  border-bottom: solid 1px;
  margin-bottom: 10px;
}

td {
  text-align: center;
}
</style>