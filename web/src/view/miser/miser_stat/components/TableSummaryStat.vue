<script setup>
  import { ref, onMounted, onBeforeUnmount } from 'vue'
  import { getLineStat, getYearStat } from '@/api/miser/miser_stat'
  import { miserCfgMap, miserTxnCfgMap } from '@/constants/miser'
  import { useMiserStatStore } from '@/pinia'
  import { formatterAmount } from '@/utils/format'

  defineOptions({ name: 'TableSummaryStat' })
  const emits = defineEmits(['open'])

  // 状态管理
  const { fetchData, subscribe, unsubscribe } = useMiserStatStore()

  // 响应式数据
  const tableData = ref([
    {
      header: '结余汇总 - 按月统计',
      headerId: miserCfgMap.balance.transactionType,
      cols: [
        { name: 'month', label: '月份' },
        { ...miserCfgMap.income },
        { ...miserCfgMap.expense },
        { ...miserCfgMap.balance }
      ],
      rows: [],
      fetcher: getLineStat
    },
    {
      header: '结余汇总 - 按年统计',
      headerId: miserCfgMap.balance.transactionType,
      cols: [
        { name: 'year', label: '年份' },
        { ...miserCfgMap.income },
        { ...miserCfgMap.expense },
        { ...miserCfgMap.balance }
      ],
      rows: [],
      fetcher: getYearStat
    }
  ])

  // 模板方法
  const handleClick = (row, headerId) => {
    let subTitle = ''
    // 提取 + 过滤
    const filtered = Object.entries(row).reduce((acc, [label, value]) => {
      if (label === 'month' || label === 'year') {
        subTitle = value
      } else if (value > 0) {
        acc.push({ label: miserCfgMap[label].label, value })
      }
      return acc
    }, [])

    const { color, colorTo } = miserTxnCfgMap[headerId]

    const chartData = {
      title: subTitle,
      xData: filtered.map(({ label }) => label),
      yData: filtered.map(({ value }) => value),
      itemColor: color,
      areaColor: colorTo
    }

    emits('open', chartData)
  }

  // 生命周期
  const fetchAndRender = async () => {
    for (const tableItem of tableData.value) {
      const { code, data } = await fetchData(
        tableItem.fetcher.name,
        tableItem.fetcher
      )
      // data 先解构再赋值，避免 reverse 影响 store 及折线图数据
      tableItem.rows = code === 0 && Array.isArray(data) ? [...data] : []
      tableItem.rows.reverse()
    }
  }

  onMounted(() => {
    fetchAndRender()
    subscribe(getLineStat.name, fetchAndRender)
    subscribe(getYearStat.name, fetchAndRender)
  })

  onBeforeUnmount(() => {
    unsubscribe(getLineStat.name, fetchAndRender)
    unsubscribe(getYearStat.name, fetchAndRender)
  })
</script>

<template>
  <div class="table-wrapper">
    <el-card
      v-for="{ headerId, header, rows, cols } in tableData"
      :key="headerId"
      :header="header"
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
          :key="c.name"
          :prop="c.name"
          :label="c.label"
          :formatter="formatterAmount"
          sortable
          align="center"
        />
      </el-table>
    </el-card>
  </div>
</template>

<style scoped lang="scss">
  .table-wrapper {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    margin-top: 8px;

    .table-card {
      flex: 1;
      border-radius: 16px;
    }
  }

  :deep(.el-table__body tr:hover > td) {
    cursor: pointer;
  }
</style>
