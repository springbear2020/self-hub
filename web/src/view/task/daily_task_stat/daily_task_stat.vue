<script setup>
  import { ref } from 'vue'
  import CalendarHeatmap from './components/CalendarHeatmap.vue'
  import { getDailyTaskStatList } from '@/api/task/daily_task_completion'

  const list = ref([])
  const noMore = ref(false)
  const loading = ref(false)
  const minYear = 2025
  const currentYear = new Date().getFullYear()
  const formData = ref({
    page: 1,
    pageSize: 5,
    year: currentYear,
    startDate: `${currentYear}-01-01`,
    endDate: `${currentYear}-12-31`
  })

  const changeYear = (delta) => {
    const newYear = formData.value.year + delta
    formData.value.year = newYear
    formData.value.page = 1
    formData.value.startDate = `${newYear}-01-01`
    formData.value.endDate = `${newYear}-12-31`
    list.value = []
    noMore.value = false
    doLoad()
  }
  const doLoad = async () => {
    if (loading.value || noMore.value) {
      return
    }
    loading.value = true

    const { code, data } = await getDailyTaskStatList(formData.value)
    if (code === 0) {
      if (Array.isArray(data.list) && data.list.length > 0) {
        const mapped = data.list.map((item) => {
          const { task, completions } = item
          const heatmapData = completions.map(({ finishDate, countValue }) => {
            return [finishDate, countValue > 0 ? countValue : null]
          })
          return {
            ...task,
            heatmapData
          }
        })
        list.value.push(...mapped)
        formData.value.page += 1
      } else {
        noMore.value = true
      }
    } else {
      noMore.value = true
    }

    loading.value = false
  }
</script>

<template>
  <div class="year-adjuster">
    <el-button :disabled="formData.year <= minYear" @click="changeYear(-1)">
      &lt;
    </el-button>

    <span class="year-display">{{ formData.year }}</span>

    <el-button :disabled="formData.year >= currentYear" @click="changeYear(1)">
      &gt;
    </el-button>
  </div>

  <div
    class="infinite-list"
    v-infinite-scroll="doLoad"
    :infinite-scroll-disabled="loading || noMore"
    :infinite-scroll-distance="100"
  >
    <div v-for="item in list" :key="item.id" class="infinite-list-item">
      <calendar-heatmap
        :year="formData.year"
        :title="item.name"
        :data="item.heatmapData"
      />
    </div>

    <div v-if="loading" class="loading-text">加载中...</div>
    <div v-if="noMore" class="no-more-text">没有更多了</div>
  </div>
</template>

<style lang="scss" scoped>
  .year-adjuster {
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 7px 0;
    gap: 7px;

    .year-display {
      font-weight: 600;
      font-size: 18px;
      min-width: 60px;
      text-align: center;
    }
  }

  .infinite-list {
    height: calc(100vh - 100px);
    overflow: auto;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    box-sizing: border-box;

    .infinite-list-item {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 100%;
      margin-bottom: 7px;
    }

    .loading-text,
    .no-more-text {
      text-align: center;
      padding: 4px;
      color: var(--el-text-color-secondary);
    }
  }
</style>
