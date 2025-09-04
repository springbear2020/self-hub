<script setup>
  import { markRaw, ref, watch } from 'vue'
  import { Money, TrendCharts, Wallet } from '@element-plus/icons-vue'
  import { getCardStat } from '@/api/miser/miser_stat'
  import config from '@/core/config'
  import AmountCard from '@/components/AmountCard.vue'

  const props = defineProps({
    startMonth: { type: String, required: true },
    endMonth: { type: String, required: true }
  })

  const amountCardRef = ref()
  const cardConfig = [
    {
      title: 'income',
      label: '总收入',
      icon: markRaw(Money),
      colorFrom: config.miser.income.color,
      colorTo: config.miser.income.colorTo
    },
    {
      title: 'expense',
      label: '总支出',
      icon: markRaw(TrendCharts),
      colorFrom: config.miser.expense.color,
      colorTo: config.miser.expense.colorTo
    },
    {
      title: 'balance',
      label: '总结余',
      icon: markRaw(Wallet),
      colorFrom: config.miser.balance.color,
      colorTo: config.miser.balance.colorTo
    }
  ]
  const fetchAndRender = () => {
    amountCardRef.value.fetchAndRender()
  }

  watch(() => [props.startMonth, props.endMonth], fetchAndRender)
</script>

<template>
  <amount-card
    ref="amountCardRef"
    :api-func="getCardStat"
    :api-params="{ startMonth, endMonth }"
    :card-config="cardConfig"
  />
</template>
