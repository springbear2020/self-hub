<script setup>
  import { ref, onMounted, onBeforeUnmount } from 'vue'
  import { getMonthStat } from '@/api/miser/miser_stat'
  import { miserTxnCfgMap } from '@/constants/miser'
  import { useMiserStatStore } from '@/pinia'
  import { formatAmountCurrency, formatterAmount } from '@/utils/format'

  defineOptions({ name: 'TableMonthStat' })
  const emits = defineEmits(['open'])

  // 状态管理
  const { fetchData, subscribe, unsubscribe } = useMiserStatStore()

  // 响应式数据
  const tableData = ref([])

  // 模板方法
  const handleClick = (row, headerId) => {
    let month = ''
    let amount = 0
    // 提取 + 过滤
    const filterSorted = Object.entries(row).reduce((acc, [label, value]) => {
      if (label === '月份') {
        month = value
      } else if (label === '金额') {
        amount = value
      } else if (value > 0) {
        acc.push({ label, value })
      }
      return acc
    }, [])
    // 排序
    filterSorted.sort((a, b) => b.value - a.value)

    const { label, color, colorTo } = miserTxnCfgMap[headerId]

    const chartData = {
      title: `${month}『${label}』${formatAmountCurrency(amount)}`,
      xData: filterSorted.map(({ label }) => label),
      yData: filterSorted.map(({ value }) => value),
      itemColor: color,
      areaColor: colorTo
    }

    emits('open', chartData)
  }

  // 生命周期
  const fetchAndRender = async () => {
    const { code, data } = await fetchData(getMonthStat.name, getMonthStat)
    if (code !== 0 || !Array.isArray(data)) {
      tableData.value = []
    } else {
      tableData.value = data
    }
  }

  onMounted(() => {
    fetchAndRender()
    subscribe(getMonthStat.name, fetchAndRender)
  })

  onBeforeUnmount(() => {
    unsubscribe(getMonthStat.name, fetchAndRender)
  })
</script>

<template>
  <div>
    <el-card
      v-for="{ headerId, header, rows, cols } in tableData"
      :header="header"
      :key="headerId"
      class="table-card"
    >
      <el-table
        :data="rows"
        border
        stripe
        show-summary
        max-height="250px"
        @row-click="(row) => handleClick(row, headerId)"
      >
        <el-table-column
          v-for="c in cols"
          :key="c.key"
          :prop="c.key"
          :label="c.label"
          :fixed="c.fixed"
          :sortable="c.sortable"
          :formatter="formatterAmount"
          align="center"
        />
      </el-table>
    </el-card>
  </div>
</template>

<style scoped>
  .table-card {
    margin-top: 8px;
    border-radius: 16px;
  }

  :deep(.el-table__body tr:hover > td) {
    cursor: pointer;
  }
</style>
