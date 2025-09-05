<script setup>
  import { ref, watch } from 'vue'
  import { getLineStat, getYearStat } from '@/api/miser/miser_stat'
  import { formatterAmount } from '@/utils/format'

  const props = defineProps({
    startMonth: { type: String, required: true },
    endMonth: { type: String, required: true }
  })

  const tableConfigs = [
    {
      header: '结余明细 - 按月统计',
      cols: [
        { key: 'month', label: '月份' },
        { key: 'income', label: '收入' },
        { key: 'expense', label: '支出' },
        { key: 'balance', label: '结余' }
      ],
      fetcher: getLineStat
    },
    {
      header: '结余明细 - 按年统计',
      cols: [
        { key: 'year', label: '年份' },
        { key: 'income', label: '收入' },
        { key: 'expense', label: '支出' },
        { key: 'balance', label: '结余' }
      ],
      fetcher: getYearStat
    }
  ]

  const dataList = ref(tableConfigs.map((cfg) => ({ ...cfg, rows: [] })))
  const fetchData = async () => {
    const params = {
      startMonth: props.startMonth,
      endMonth: props.endMonth
    }

    try {
      const results = await Promise.all(
        dataList.value.map((cfg) => cfg.fetcher(params))
      )
      results.forEach((res, idx) => {
        dataList.value[idx].rows = res.code === 0 ? res.data : []
        dataList.value[idx].rows.reverse()
      })
    } catch (e) {
      dataList.value.forEach((item) => (item.rows = []))
    }
  }

  watch(() => [props.startMonth, props.endMonth], fetchData, {
    immediate: true
  })
</script>

<template>
  <div class="table-wrapper">
    <el-card
      v-for="item in dataList"
      :key="item.header"
      :header="item.header"
      class="table-card"
    >
      <el-table :data="item.rows" border stripe show-summary max-height="250px">
        <el-table-column
          v-for="c in item.cols"
          :key="c.key"
          :prop="c.key"
          :label="c.label"
          :formatter="formatterAmount"
          sortable
          align="center"
        />
      </el-table>
    </el-card>
  </div>
</template>

<style scoped>
  .table-wrapper {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    margin-top: 8px;
  }

  .table-card {
    flex: 1;
    border-radius: 16px;
  }
</style>
