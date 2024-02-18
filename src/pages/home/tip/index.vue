<script setup lang="ts">

</script>

<template>
  <div class="tip">
    <!-- 常见科室 -->
    <div class="department">
      <div class="header">
        <div class="left">
          <span>常见科室</span>
        </div>
        <div class="right">
          <span>全部</span>
        </div>
      </div>
      <div class="content">
        <ul>
          <li>内科</li>
          <li>外科</li>
          <li>小儿科</li>
          <li>妇科</li>
        </ul>
      </div>
    </div>
    <!-- 平台公告 -->
    <div class="notice">
      <div class="header">
        <div class="left">
          <span>平台公告</span>
        </div>
        <div class="right">
          <span>全部</span>
        </div>
      </div>
      <div class="content">
        <ul>
          <li>通知一</li>
          <li>通知二</li>
          <li>通知三</li>
        </ul>
      </div>
    </div>
    <!-- 停诊公告 -->
    <div></div>

  </div>
</template>

<style scoped lang="scss">
.tip {
  color: #7f7f7f;

  .department {
    .header {
      display: flex;
      justify-content: space-between;
    }

    .content {
      ul {
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;

        li {
          width: 40%;
          margin-top: 20px;
        }
      }
    }
  }

  .notice {
    margin-top: 20px;

    .header {
      display: flex;
      justify-content: space-between;
    }

    .content {
      ul {
        li {
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
          margin-top: 20px;
        }
      }
    }
  }
}
</style>
