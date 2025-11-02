<script setup>
  import { ref, onMounted, onBeforeUnmount } from 'vue'
  import { getRankingStat } from '@/api/miser/miser_stat'
  import { miserTxnCfgMap } from '@/constants/miser'
  import {
    useMiserStatStore,
    useMiserCategoryStore,
    useMiserTransactionTypeStore
  } from '@/pinia'
  import { formatDate, formatterAmount } from '@/utils/format'
  import { storeToRefs } from 'pinia'

  defineOptions({ name: 'TableRankingStat' })
  const emits = defineEmits(['open'])

  // 状态管理，用 storeToRefs 让解构后的值保持响应式
  const statStore = useMiserStatStore()
  const { fetchData, subscribe, unsubscribe } = statStore
  const { startMonth, endMonth } = storeToRefs(statStore)
  const catStore = useMiserCategoryStore()
  const txnStore = useMiserTransactionTypeStore()

  // 响应式数据
  const dataMap = ref({})

  // 模板方法
  const handleClick = async (rowData) => {
    const { categoryId } = rowData
    const { code, data } = await fetchData(
      getRankingStat.name,
      getRankingStat,
      { categoryId }
    )
    if (code !== 0 || !data) {
      return
    }

    const transactionType = catStore.catTxnMap[categoryId]
    const categoryName = catStore.dataMap[categoryId]
    const transactionTypeName = txnStore.dataMap[transactionType]
    const { color, colorTo } = miserTxnCfgMap[transactionType]

    const processed = data[transactionType].map(({ date, amount }) => ({
      date: formatDate(date, 'yy-MM-dd'),
      amount
    }))

    const chartData = {
      title: `${startMonth.value} 至 ${endMonth.value}『${categoryName}』${transactionTypeName}榜单趋势`,
      xData: processed.map(({ date }) => date),
      yData: processed.map(({ amount }) => amount),
      itemColor: color,
      areaColor: colorTo
    }

    emits('open', chartData)
  }

  // 生命周期
  const fetchAndRender = async () => {
    const { code, data } = await fetchData(getRankingStat.name, getRankingStat)
    dataMap.value = code === 0 && data ? data : {}
  }

  onMounted(async () => {
    await catStore.init()
    await txnStore.init()
    await fetchAndRender()
    subscribe(getRankingStat.name, fetchAndRender)
  })

  onBeforeUnmount(() => {
    unsubscribe(getRankingStat.name, fetchAndRender)
  })
</script>

<template>
  <div class="cards">
    <el-card
      class="table-card"
      v-for="txn in Object.keys(dataMap)"
      :key="txn"
      :header="`${txnStore.dataMap[txn]}榜单`"
    >
      <el-table
        :data="dataMap[txn]"
        border
        stripe
        max-height="250px"
        @row-click="handleClick"
      >
        <el-table-column type="index" label="#" align="center" width="55" />
        <el-table-column prop="name" label="名称" align="center" />
        <el-table-column
          prop="categoryId"
          label="分类"
          align="center"
          width="110"
          sortable
        >
          <template #default="{ row }">
            {{ catStore.dataMap[row.categoryId] }}
          </template>
        </el-table-column>
        <el-table-column
          prop="amount"
          label="金额"
          align="center"
          :formatter="formatterAmount"
          width="110"
          sortable
        />
        <el-table-column
          prop="date"
          label="日期"
          align="center"
          width="110"
          sortable
        >
          <template #default="{ row }">
            {{ formatDate(row.date, 'yyyy-MM-dd') }}
          </template>
        </el-table-column>
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
