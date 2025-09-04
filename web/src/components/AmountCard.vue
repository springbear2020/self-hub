<script setup>
  import { onMounted, ref, nextTick } from 'vue'

  const props = defineProps({
    apiFunc: {
      type: Function,
      required: true
    },
    apiParams: {
      type: Object,
      default: () => ({})
    },
    cardConfig: {
      type: Array,
      required: true
    }
  })

  const dataList = ref([])

  const formatAmount = (val) =>
    new Intl.NumberFormat('zh-CN', {
      style: 'currency',
      currency: 'CNY'
    }).format(val ?? 0)

  const startCount = (item) => {
    const start = 0
    const end = item.amount
    const duration = 800
    const startTime = performance.now()

    const step = (now) => {
      const p = Math.min((now - startTime) / duration, 1)
      item.current = parseFloat((start + (end - start) * p).toFixed(2))
      if (p < 1) requestAnimationFrame(step)
    }
    requestAnimationFrame(step)
  }

  const doRender = () => {
    nextTick(() => {
      const cards = document.querySelectorAll('.stat-card')
      cards.forEach((card, index) => {
        const observer = new IntersectionObserver(
          ([entry]) => {
            if (entry.isIntersecting) {
              startCount(dataList.value[index])
              observer.disconnect()
            }
          },
          { threshold: 0.3, rootMargin: '80px' }
        )
        observer.observe(card)
      })
    })
  }

  const fetchAndRender = async () => {
    const { code, data } = await props.apiFunc(props.apiParams)
    if (code === 0 && data) {
      dataList.value.forEach((item) => {
        item.amount = data[item.title] ?? 0
        item.current = 0
      })
      doRender()
    }
  }

  onMounted(() => {
    dataList.value = props.cardConfig.map((it) => ({
      ...it,
      amount: 0,
      current: 0
    }))
    fetchAndRender()
  })

  defineExpose({ fetchAndRender })
</script>

<template>
  <div class="card-wrapper">
    <el-card
      v-for="item in dataList"
      :key="item.title"
      shadow="hover"
      class="stat-card"
      :style="{
        background: `linear-gradient(135deg, ${item.colorFrom}, ${item.colorTo})`
      }"
    >
      <div class="stat-content">
        <div class="icon-wrap">
          <el-icon :size="26" color="#fff">
            <component :is="item.icon" />
          </el-icon>
        </div>
        <div class="stat-info">
          <div class="stat-label">{{ item.label }}</div>
          <div class="stat-amount">{{ formatAmount(item.current) }}</div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<style scoped lang="scss">
  .card-wrapper {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
  }

  .stat-card {
    flex: 1;
    min-width: 220px;
    border: none;
    color: #fff;
    border-radius: 16px;
    transition:
      transform 0.2s ease,
      box-shadow 0.2s ease,
      opacity 0.25s ease;
    opacity: 0.98;

    &:hover {
      transform: translateY(-5px);
      box-shadow: 0 12px 28px rgba(0, 0, 0, 0.15);
      opacity: 1;
    }

    .stat-content {
      display: flex;
      align-items: center;
      gap: 12px;
      padding: 20px 10px;
    }

    .icon-wrap {
      background: rgba(255, 255, 255, 0.18);
      padding: 10px;
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      backdrop-filter: blur(2px);
    }

    .stat-info {
      display: flex;
      flex-direction: column;
    }

    .stat-label {
      font-size: 14px;
      opacity: 0.9;
      letter-spacing: 0.2px;
    }

    .stat-amount {
      font-size: 24px;
      font-weight: 800;
      margin-top: 4px;
      line-height: 1.2;
      text-shadow: 0 1px 0 rgba(0, 0, 0, 0.08);
    }
  }
</style>
