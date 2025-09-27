<script setup>
  import { onMounted, onBeforeUnmount, ref, watch } from 'vue'
  import * as echarts from 'echarts'

  defineOptions({ name: 'CalendarHeatmap' })

  // === 响应式数据 ===
  const props = defineProps({
    year: { type: Number, default: new Date().getFullYear() },
    task: { type: Object, required: true, default: () => ({}) },
    completions: { type: Array, default: () => [] }
  })

  // ==== 图表逻辑 ====
  const chartRef = ref()
  let chartInstance = null

  const initChart = () => {
    if (!chartRef.value) {
      return
    }

    if (chartInstance) {
      disposeChart()
    }

    chartInstance = echarts.init(chartRef.value)
  }

  const disposeChart = () => {
    if (!chartInstance) {
      return
    }

    chartInstance.dispose()
    chartInstance = null
  }

  const updateChart = (chartData = {}) => {
    if (!chartInstance) {
      return
    }

    const { title, min, max, data } = chartData

    chartInstance.clear()
    chartInstance.setOption({
      title: { top: 20, left: 'center', text: title },
      tooltip: {
        formatter: ({ value, data }) => {
          const [date = '', count = 0] = value
          const dateStr = date ? date.slice(5) : ''
          const remark = data.remark?.replace(/\n/g, '<br/>') ?? ''
          return `${dateStr}: ${count}<br/>${remark}`
        }
      },
      visualMap: {
        min: min,
        max: max,
        type: 'piecewise',
        orient: 'horizontal',
        left: 'center',
        top: 60,
        inRange: {
          color: ['#e0f9f8', '#a6eeeb', '#00ccc6', '#009e9a', '#006865']
        }
      },
      calendar: {
        top: 120,
        left: 30,
        right: 30,
        cellSize: ['auto', 15],
        range: props.year,
        itemStyle: { borderWidth: 0.5 },
        dayLabel: { firstDay: 1 }
      },
      series: {
        type: 'heatmap',
        coordinateSystem: 'calendar',
        data: data
      }
    })
  }

  const renderChart = () => {
    const { name, minValue, maxValue } = props.task
    const data = getFullYearData(props.year, props.completions)
    updateChart({
      title: name,
      min: minValue,
      max: maxValue,
      data: data
    })
  }

  // ==== 工具函数 ====
  function getFullYearData(year, rawData) {
    const start = new Date(year, 0, 1)
    const end = new Date(year, 11, 31)
    const mapData = new Map(
      rawData.map((item) => [item[0], { count: item[1], remark: item[2] }])
    )

    const results = []
    let cur = new Date(start)
    while (cur <= end) {
      const dateStr = cur.toISOString().slice(0, 10)
      const { count, remark } = mapData.get(dateStr) || { count: null, remark: null }
      results.push({ value: [dateStr, count], remark })
      cur.setDate(cur.getDate() + 1)
    }

    return results
  }

  // ==== 生命周期 ====
  const handleResize = () => {
    chartInstance?.resize()
  }

  watch(() => props.year, renderChart)

  onMounted(() => {
    initChart()
    renderChart()
    window.addEventListener('resize', handleResize)
  })

  onBeforeUnmount(() => {
    window.removeEventListener('resize', handleResize)
    disposeChart()
  })
</script>

<template>
  <div class="task-card">
    <div ref="chartRef" class="chart" />
  </div>
</template>

<style scoped lang="scss">
  .task-card {
    border: 1px solid var(--el-border-color-dark);
    border-radius: 8px;
    background: #fff;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
    padding: 8px;

    .chart {
      height: 250px;
      width: 100%;
    }
  }
</style>
