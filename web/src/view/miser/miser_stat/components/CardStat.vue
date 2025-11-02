<script setup>
  import { ref, onMounted, onBeforeUnmount } from 'vue'
  import { getCardStat, getLineStat } from '@/api/miser/miser_stat'
  import AmountCard from '@/components/AmountCard.vue'
  import { miserCfgMap, miserTxnCfgMap } from '@/constants/miser'
  import { useMiserStatStore } from '@/pinia'
  import { formatAmountCurrency } from '@/utils/format'
  import { storeToRefs } from 'pinia'

  defineOptions({ name: 'CardStat' })
  const emits = defineEmits(['open'])

  // 状态管理，用 storeToRefs 让解构后的值保持响应式
  const statStore = useMiserStatStore()
  const { fetchData, subscribe, unsubscribe } = statStore
  const { startMonth, endMonth } = storeToRefs(statStore)

  // 响应式数据
  const cardRef = ref()
  const cardList = ref([
    {
      ...miserCfgMap.income,
      label: '总收入'
    },
    {
      ...miserCfgMap.expense,
      label: '总支出'
    },
    {
      ...miserCfgMap.balance,
      label: '总结余'
    }
  ])

  // 模板方法
  const handleClick = async (cardData) => {
    const { code, data } = await fetchData(getLineStat.name, getLineStat)
    if (code !== 0 || !Array.isArray(data)) {
      return
    }

    const { label, amount, transactionType } = cardData
    const { color, colorTo, name: key } = miserTxnCfgMap[transactionType]

    const chartData = {
      title: `${startMonth.value} 至 ${endMonth.value}『${label}』${formatAmountCurrency(amount)}`,
      xData: data.map(({ month }) => month),
      yData: data.map((item) => item[key]),
      itemColor: color,
      areaColor: colorTo
    }

    emits('open', chartData)
  }

  // 生命周期
  const fetchAndRender = async () => {
    const { code, data } = await fetchData(getCardStat.name, getCardStat)
    if (code !== 0 || !data) {
      return
    }

    cardList.value.forEach((item) => {
      item.amount = data[item.name] ?? 0
      item.current = 0
    })

    cardRef.value.doRender(cardList.value)
  }

  onMounted(() => {
    fetchAndRender()
    subscribe(getCardStat.name, fetchAndRender)
  })

  onBeforeUnmount(() => {
    unsubscribe(getCardStat.name, fetchAndRender)
  })
</script>

<template>
  <amount-card ref="cardRef" :data-list="cardList" @click="handleClick" />
</template>
