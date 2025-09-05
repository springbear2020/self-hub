<script setup>
  import { markRaw, nextTick, ref } from 'vue'
  import { Money, TrendCharts, Wallet } from '@element-plus/icons-vue'
  import config from '@/core/config'
  import { getMiserLoanStatData } from '@/api/miser/miser_loan_record'
  import AmountCard from '@/components/AmountCard.vue'

  const amountCardRef = ref()
  const cardConfig = [
    {
      title: 'lend_amount',
      label: '总借出',
      icon: markRaw(Money),
      colorFrom: config.miser.expense.color,
      colorTo: config.miser.expense.colorTo
    },
    {
      title: 'repaid_amount',
      label: '总归还',
      icon: markRaw(TrendCharts),
      colorFrom: config.miser.balance.color,
      colorTo: config.miser.balance.colorTo
    },
    {
      title: 'repaying_amount',
      label: '待还款',
      icon: markRaw(Wallet),
      colorFrom: config.miser.income.color,
      colorTo: config.miser.income.colorTo
    }
  ]
  const fetchAndRender = () => {
    nextTick(() => {
      amountCardRef.value.fetchAndRender()
    })
  }

  defineExpose({ fetchAndRender })
</script>

<template>
  <amount-card
    ref="amountCardRef"
    :api-func="getMiserLoanStatData"
    :card-config="cardConfig"
  />
</template>
