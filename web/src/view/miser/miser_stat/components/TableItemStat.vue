<script setup>
  import { ref, onMounted, onBeforeUnmount } from 'vue'
  import { getCategoryItemStat, getItemStat } from '@/api/miser/miser_stat'
  import { miserTxnCfgMap } from '@/constants/miser'
  import { useMiserCategoryStore, useMiserStatStore } from '@/pinia'
  import { formatAmountCurrency, formatterAmount } from '@/utils/format'
  import { storeToRefs } from 'pinia'

  defineOptions({ name: 'TableItemStat' })
  const emits = defineEmits(['open'])

  // 状态管理，用 storeToRefs 让解构后的值保持响应式
  const statStore = useMiserStatStore()
  const { fetchData, subscribe, unsubscribe } = statStore
  const { startMonth, endMonth } = storeToRefs(statStore)
  const catStore = useMiserCategoryStore()

  // 响应式数据
  const dataMap = ref({})

  // 模板方法
  const handleClick = async (rowData) => {
    const { categoryId, name, amount } = rowData
    const { code, data } = await fetchData(
      getCategoryItemStat.name,
      getCategoryItemStat,
      { name, categoryId }
    )
    if (code !== 0 || !Array.isArray(data)) {
      return
    }

    const transactionType = catStore.catTxnMap[categoryId]
    const { color, colorTo } = miserTxnCfgMap[transactionType]

    const chartData = {
      title: `${startMonth.value} 至 ${endMonth.value}『${name}』${formatAmountCurrency(amount)}`,
      xData: data.map(({ month }) => month),
      yData: data.map(({ amount }) => amount),
      itemColor: color,
      areaColor: colorTo
    }

    emits('open', chartData)
  }

  // 生命周期
  const fetchAndRender = async () => {
    const { code, data } = await fetchData(getItemStat.name, getItemStat)
    dataMap.value = code === 0 && data ? data : {}
  }

  onMounted(async () => {
    await catStore.init()
    await fetchAndRender()
    subscribe(getItemStat.name, fetchAndRender)
  })

  onBeforeUnmount(() => {
    unsubscribe(getItemStat.name, fetchAndRender)
  })
</script>

<template>
  <div class="cards">
    <el-card
      class="table-card"
      v-for="c in Object.keys(dataMap)"
      :key="c"
      :header="`明细汇总 - ${catStore.dataMap[c]}`"
    >
      <el-table
        :data="dataMap[c]"
        border
        stripe
        show-summary
        max-height="250px"
        @row-click="handleClick"
      >
        <el-table-column prop="name" label="名称" sortable align="center" />
        <el-table-column
          prop="amount"
          label="金额"
          sortable
          align="center"
          :formatter="formatterAmount"
        />
      </el-table>
    </el-card>
  </div>
</template>

<style scoped lang="scss">
  .cards {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 8px;
    margin-top: 8px;

    .table-card {
      border-radius: 16px;
    }
  }

  :deep(.el-table__body tr:hover > td) {
    cursor: pointer;
  }
</style>
