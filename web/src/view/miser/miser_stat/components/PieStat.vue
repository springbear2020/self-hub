<script setup>
  import { computed, onBeforeUnmount, onMounted, ref, watch } from 'vue'
  import * as echarts from 'echarts'
  import { listMiserCategoryList } from '@/api/miser/miser_category'
  import { getCardStat, getPieStat } from '@/api/miser/miser_stat'
  import config from '@/core/config'

  const props = defineProps({
    startMonth: { type: String, required: true },
    endMonth: { type: String, required: true }
  })

  let pieIncomeInstance, pieExpenseInstance, pieBalanceInstance
  const pieIncomeRef = ref()
  const pieExpenseRef = ref()
  const pieBalanceRef = ref()

  // 分类列表
  const categoryList = ref([])
  const categoryMap = computed(() => {
    const resultMap = {}
    categoryList.value.forEach(({ id, name }) => {
      resultMap[id] = name
    })
    return resultMap
  })
  const fetchCategories = async () => {
    const { code, data } = await listMiserCategoryList()
    categoryList.value = code === 0 && Array.isArray(data) ? data : []
  }

  // 饼图数据
  const pieData = ref({})
  const fetchAndRender = async () => {
    const params = {
      startMonth: props.startMonth,
      endMonth: props.endMonth
    }
    const [pieResp, cardResp] = await Promise.all([
      getPieStat(params),
      getCardStat(params)
    ])

    pieData.value = pieResp.code === 0 ? pieResp.data : {}
    if (cardResp.code === 0) {
      pieData.value.income = cardResp.data?.income ?? 0
      pieData.value.expense = cardResp.data?.expense ?? 0
      pieData.value.balance = cardResp.data?.balance ?? 0
    }

    doRender()
  }

  // 图表渲染
  const doRender = () => {
    pieIncomeInstance.clear()
    pieExpenseInstance.clear()
    pieBalanceInstance.clear()

    const incomeData = mapCategoryData(pieData.value.incomes)
    const expenseData = mapCategoryData(pieData.value.expenses)
    const balanceData = [
      { name: '支出', value: pieData.value.expense ?? 0 },
      { name: '结余', value: pieData.value.balance ?? 0 }
    ]

    pieIncomeInstance.setOption(createPieOption('收入占比', incomeData))
    pieExpenseInstance.setOption(createPieOption('支出占比', expenseData))
    pieBalanceInstance.setOption(
      createPieOption('支出/结余', balanceData, [
        config.miser.expense.color,
        config.miser.balance.color
      ])
    )
  }
  const mapCategoryData = (list) =>
    (list || []).map((i) => ({
      name: categoryMap.value[i.category] || i.category,
      value: i.amount
    }))
  const createPieOption = (title, data, color = []) => ({
    title: {
      text: title,
      left: 'center',
      top: 10,
      textStyle: { fontSize: 16, fontWeight: 600 }
    },
    tooltip: {
      trigger: 'item',
      formatter: ({ name, value, percent }) =>
        `${name}: ${Number(value).toFixed(2)} (${percent}%)`
    },
    legend: { bottom: 0, type: 'scroll' },
    series: [
      {
        type: 'pie',
        radius: ['40%', '70%'],
        center: ['50%', '50%'],
        avoidLabelOverlap: false,
        itemStyle: { borderRadius: 6, borderColor: '#fff', borderWidth: 2 },
        label: { show: false, position: 'center' },
        emphasis: { label: { show: true, fontSize: 16, fontWeight: 'bold' } },
        labelLine: { show: false },
        color,
        data
      }
    ]
  })

  onMounted(async () => {
    pieIncomeInstance = echarts.init(pieIncomeRef.value)
    pieExpenseInstance = echarts.init(pieExpenseRef.value)
    pieBalanceInstance = echarts.init(pieBalanceRef.value)

    await fetchCategories()
    await fetchAndRender()
  })

  onBeforeUnmount(() => {
    pieIncomeInstance?.dispose()
    pieExpenseInstance?.dispose()
    pieBalanceInstance?.dispose()
  })

  watch(
    () => [props.startMonth, props.endMonth],
    () => {
      fetchAndRender()
    }
  )
</script>

<template>
  <div class="pie-wrapper">
    <div ref="pieIncomeRef" class="pie-chart" />
    <div ref="pieExpenseRef" class="pie-chart" />
    <div ref="pieBalanceRef" class="pie-chart" />
  </div>
</template>

<style scoped lang="scss">
  .pie-wrapper {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    margin-top: 8px;
  }

  .pie-chart {
    flex: 1;
    min-width: 300px;
    height: 320px;
    background: #fff;
    border-radius: 16px;
    box-shadow: 0 4px 14px rgba(0, 0, 0, 0.05);
    padding: 10px;
  }
</style>
