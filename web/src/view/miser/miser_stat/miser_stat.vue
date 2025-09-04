<script setup>
  import { computed, ref, defineAsyncComponent } from 'vue'
  import LazyWrapper from '@/components/LazyWrapper.vue'
  import { formatDate } from '@/utils/format'

  defineOptions({
    name: 'MiserStat'
  })

  // 月份选择框
  const minYear = 2023
  const minMonth = 7
  const currentYear = new Date().getFullYear()
  const monthRange = ref([new Date(minYear, minMonth), new Date()])
  const startMonth = computed(() => {
    return formatDate(monthRange.value[0], 'yyyy-MM')
  })
  const endMonth = computed(() => {
    return formatDate(monthRange.value[1], 'yyyy-MM')
  })
  const disabledDate = (time) => {
    return time.getTime() > Date.now()
  }
  const generateYearShortcuts = () => {
    const shortcuts = []
    shortcuts.push({
      text: '所有时间',
      value: () => {
        const now = new Date()
        return [new Date(minYear, minMonth), now]
      }
    })

    for (let year = currentYear; year >= minYear; year--) {
      shortcuts.push({
        text: `${year} 年`,
        value: () => {
          return [new Date(year, 0), new Date(year, 11, 31)]
        }
      })
    }

    return shortcuts
  }
  const shortcuts = generateYearShortcuts()

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

    <card-stat :start-month="startMonth" :end-month="endMonth" />

    <lazy-wrapper>
      <pie-stat :start-month="startMonth" :end-month="endMonth" />
    </lazy-wrapper>
    <lazy-wrapper>
      <line-stat :start-month="startMonth" :end-month="endMonth" />
    </lazy-wrapper>
    <lazy-wrapper>
      <table-month-stat :start-month="startMonth" :end-month="endMonth" />
    </lazy-wrapper>
    <lazy-wrapper>
      <table-summary-stat :start-month="startMonth" :end-month="endMonth" />
    </lazy-wrapper>
    <lazy-wrapper>
      <table-item-stat :start-month="startMonth" :end-month="endMonth" />
    </lazy-wrapper>

    <div class="no-more-text">没有更多了</div>
  </div>
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
