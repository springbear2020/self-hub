<script setup>
  import { computed, ref, defineAsyncComponent, watch, onBeforeUnmount } from 'vue'
  import LazyWrapper from '@/components/LazyWrapper.vue'
  import LineBarDialog from '@/components/LineBarDialog.vue'
  import { formatDate } from '@/utils/format'
  import { useMiserStatStore } from '@/pinia'
  import { MISER_MIN_MONTH, MISER_MIN_YEAR } from '@/constants/miser'

  defineOptions({ name: 'MiserStat' })

  // 常量
  const MIN_YEAR = MISER_MIN_YEAR
  const MIN_MONTH = MISER_MIN_MONTH - 1

  // 组件懒加载
  const CardStat = defineAsyncComponent(
    () => import('./components/CardStat.vue')
  )
  const PieStat = defineAsyncComponent(() => import('./components/PieStat.vue'))
  const LineStat = defineAsyncComponent(
    () => import('./components/LineStat.vue')
  )
  const TableMonthStat = defineAsyncComponent(
    () => import('./components/TableMonthStat.vue')
  )
  const TableSummaryStat = defineAsyncComponent(
    () => import('./components/TableSummaryStat.vue')
  )
  const TableItemStat = defineAsyncComponent(
    () => import('./components/TableItemStat.vue')
  )

  // 状态管理
  const { setRange, clearCache } = useMiserStatStore()

  // 响应式数据
  const monthRange = ref([new Date(MIN_YEAR, MIN_MONTH), new Date()])
  const startMonth = computed(() => formatDate(monthRange.value[0], 'yyyy-MM'))
  const endMonth = computed(() => formatDate(monthRange.value[1], 'yyyy-MM'))

  watch(
    [startMonth, endMonth],
    ([s, e]) => {
      if (!s || !e) {
        return
      }
      setRange(s, e)
    },
    { immediate: true }
  )

  // 模板方法
  const disabledDate = (time) => {
    return time.getTime() > Date.now()
  }
  const shortcuts = ((minYear, minMonth) => {
    const list = []

    list.push({
      text: '所有时间',
      value: () => {
        const now = new Date()
        return [new Date(minYear, minMonth), now]
      }
    })

    for (let cur = new Date().getFullYear(); cur >= minYear; cur--) {
      list.push({
        text: `${cur} 年`,
        value: () => [new Date(cur, 0), new Date(cur, 11, 31)]
      })
    }

    return list
  })(MIN_YEAR, MIN_MONTH)

  // 对话框逻辑
  const dialogRef = ref()
  const handleOpen = (chartData) => {
    dialogRef.value.openDialog(chartData)
  }

  // 生命周期
  onBeforeUnmount(clearCache)
</script>

<template>
  <div>
    <div class="month-picker">
      <el-form>
        <el-form-item>
          <el-date-picker
            v-model="monthRange"
            :clearable="false"
            :disabled-date="disabledDate"
            :shortcuts="shortcuts"
            type="monthrange"
            range-separator="至"
            start-placeholder="开始月份"
            end-placeholder="结束月份"
          />
        </el-form-item>
      </el-form>
    </div>

    <card-stat @open="handleOpen" />

    <lazy-wrapper>
      <pie-stat @open="handleOpen" />
    </lazy-wrapper>

    <lazy-wrapper>
      <line-stat @open="handleOpen" />
    </lazy-wrapper>

    <lazy-wrapper>
      <table-month-stat @open="handleOpen" />
    </lazy-wrapper>

    <lazy-wrapper>
      <table-summary-stat @open="handleOpen" />
    </lazy-wrapper>

    <lazy-wrapper>
      <table-item-stat @open="handleOpen" />
    </lazy-wrapper>

    <div class="no-more-text">没有更多了</div>
  </div>

  <line-bar-dialog ref="dialogRef" />
</template>

<style scoped lang="scss">
  .month-picker {
    display: flex;
    align-content: center;
    justify-content: center;
    padding: 16px 0;

    .el-form-item {
      margin-bottom: 0;
    }
  }

  .no-more-text {
    text-align: center;
    padding: 16px 0;
    color: #999;
    font-size: 14px;
  }
</style>
