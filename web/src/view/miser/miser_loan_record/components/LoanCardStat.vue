<script setup>
  import { ref, onMounted } from 'vue'
  import { getMiserLoanStatData } from '@/api/miser/miser_loan_record'
  import AmountCard from '@/components/AmountCard.vue'
  import { miserCfgMap } from '@/constants/miser'
  import { formatAmountCurrency } from '@/utils/format'

  defineOptions({ name: 'LoanCardStat' })
  const emits = defineEmits(['open'])

  // 响应式数据
  const cardRef = ref()
  const cardList = ref([
    {
      ...miserCfgMap.income,
      name: 'repaid_amount',
      label: '总归还'
    },
    {
      ...miserCfgMap.expense,
      name: 'lend_amount',
      label: '总借出'
    },
    {
      ...miserCfgMap.balance,
      name: 'repaying_amount',
      label: '待还款'
    }
  ])

  // 模板方法
  const handleClick = async (cardData) => {
    const { code, data } = await getMiserLoanStatData({ useGroup: true })
    if (code !== 0 || !Array.isArray(data)) {
      return
    }

    const { label, amount, color, colorTo, name: key } = cardData

    // 过滤零值项并降序排列
    const filterSorted = data
      .filter((item) => item[key] > 0)
      .sort((a, b) => b[key] - a[key])

    const chartData = {
      title: `『${label}』${formatAmountCurrency(amount)}`,
      xData: filterSorted.map(({ name }) => name),
      yData: filterSorted.map((item) => item[key]),
      itemColor: color,
      areaColor: colorTo
    }

    emits('open', chartData)
  }

  // 生命周期
  const fetchAndRender = async () => {
    const { code, data } = await getMiserLoanStatData()
    if (code !== 0 || !data) {
      return
    }

    cardList.value.forEach((item) => {
      item.amount = data[item.name] ?? 0
      item.current = 0
    })

    cardRef.value.doRender(cardList.value)
  }

  onMounted(fetchAndRender)

  defineExpose({ fetchAndRender })
</script>

<template>
  <amount-card ref="cardRef" :data-list="cardList" @click="handleClick" />
</template>
