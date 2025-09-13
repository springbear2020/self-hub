<script setup>
  import { computed, ref } from 'vue'
  import CalendarHeatmap from './components/CalendarHeatmap.vue'
  import { getDailyTaskStatList } from '@/api/task/daily_task_completion'
  import { TASK_MIN_YEAR } from '@/constants/miser'

  defineOptions({ name: 'DailyTaskStat' })

  // === 常量 ===
  const MIN_YEAR = TASK_MIN_YEAR
  const CURRENT_YEAR = new Date().getFullYear()
  const PAGE_SIZE = 5

  // === 响应式数据 ===
  const list = ref([])
  const noMore = ref(false)
  const loading = ref(false)
  const formData = ref(createFormData(CURRENT_YEAR))
  const disablePrevYear = computed(() => {
    return formData.value.year <= MIN_YEAR
  })
  const disableNextYear = computed(() => {
    return formData.value.year >= CURRENT_YEAR
  })

  // === 模板方法 ===
  const doLoad = async () => {
    if (loading.value || noMore.value) {
      return
    }

    loading.value = true
    try {
      const { code, data } = await getDailyTaskStatList(formData.value)
      if (code === 0 && data && Array.isArray(data.list)) {
        list.value.push(...mapTaskData(data.list))
        if (data.list.length >= formData.value.pageSize) {
          formData.value.page += 1
        } else {
          noMore.value = true
        }
      } else {
        noMore.value = true
      }
    } finally {
      loading.value = false
    }
  }
  const changeYear = (delta = 0) => {
    const newYear = formData.value.year + delta
    formData.value = createFormData(newYear) // 重置分页和时间范围
    list.value = []
    noMore.value = false
    doLoad()
  }

  // === 工具函数 ===
  function createFormData(year, page = 1, pageSize = PAGE_SIZE) {
    return {
      page,
      pageSize,
      year,
      startDate: `${year}-01-01`,
      endDate: `${year}-12-31`
    }
  }

  function mapTaskData(data) {
    return data.map(({ task, completions }) => ({
      ...task,
      completions: completions.map(({ finishDate, countValue }) => [
        finishDate,
        countValue > 0 ? countValue : null
      ])
    }))
  }
</script>

<template>
  <div class="year-adjuster">
    <el-button
      circle
      :disabled="disablePrevYear"
      @click="changeYear(-1)"
      icon="ArrowLeftBold"
    >
    </el-button>

    <span class="year-display">{{ formData.year }}</span>

    <el-button
      circle
      :disabled="disableNextYear"
      @click="changeYear(1)"
      icon="ArrowRightBold"
    >
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
        :task="item"
        :completions="item.completions"
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
    padding: 16px 0;
    gap: 8px;

    .year-display {
      font-weight: bold;
      font-size: 20px;
      color: #333;
      min-width: 80px;
      text-align: center;
    }
  }

  .infinite-list {
    display: grid;
    gap: 8px;

    .infinite-list-item {
      background: #fff;
      border-radius: 16px;
      box-shadow: 0 2px 6px rgba(0, 0, 0, 0.06);
      padding: 16px;
      transition:
        transform 0.2s ease,
        box-shadow 0.2s ease;
    }

    .infinite-list-item:hover {
      transform: translateY(-2px);
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
    }
  }

  .loading-text,
  .no-more-text {
    text-align: center;
    padding: 8px 0 16px;
    color: #999;
    font-size: 14px;
  }
</style>
