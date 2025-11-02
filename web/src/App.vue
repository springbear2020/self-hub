<template>
  <div id="app" class="bg-gray-50 text-slate-700 dark:text-slate-500 dark:bg-slate-800">
      <router-view />
      <Application />
  </div>
</template>

<script setup>
import Application from '@/components/application/index.vue'
import { useAppStore } from '@/pinia'

useAppStore()
defineOptions({
  name: 'App'
})
</script>
<style lang="scss">
// 引入初始化样式
#app {
  height: 100vh;
  overflow: hidden;
  font-weight: 400 !important;
}

.el-button {
  font-weight: 400 !important;
}

.gva-body-h {
  min-height: calc(100% - 3rem);
}

.gva-container {
  /* 100vh - headerHeight */
  height: calc(100vh - 56px);
}

.gva-container2 {
  /* 100vh - headerHeight - tabsHeight */
  height: calc(100vh - 56px - 41px);
}

.router-view-box {
  padding: 0 8px;
}
</style>
