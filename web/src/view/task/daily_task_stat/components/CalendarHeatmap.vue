<script setup>
  import * as echarts from 'echarts'
  import { onMounted, onBeforeUnmount, ref } from 'vue'

  const props = defineProps({
    title: {
      type: String,
      required: true,
      default: ''
    },
    year: {
      type: Number,
      required: true,
      default: 0
    },
    data: {
      type: Array,
      required: true,
      default: () => []
    }
  })

  const chartRef = ref(null)
  let chartInstance = null

  const renderChart = () => {
    if (!chartInstance) {
      chartInstance = echarts.init(chartRef.value)
    }

    chartInstance.setOption({
      title: {
        top: 20,
        left: 'center',
        text: props.title
      },
      tooltip: {
        formatter: function(params) {
          const dateStr = params.value[0].substring(5)
          return `${dateStr}：${params.value[1]}`
        }
      },
      visualMap: {
        min: 0,
        max: 100,
        type: 'piecewise',
        orient: 'horizontal',
        left: 'center',
        top: 60
      },
      calendar: {
        top: 120,
        left: 30,
        right: 30,
        cellSize: ['auto', 15],
        range: props.year,
        itemStyle: {
          borderWidth: 0.5
        },
        dayLabel: {
          firstDay: 1
        }
      },
      series: {
        type: 'heatmap',
        coordinateSystem: 'calendar',
        data: props.data
      }
    })
  }

  onMounted(renderChart)

  onBeforeUnmount(() => {
    if (chartInstance) {
      chartInstance.dispose()
      chartInstance = null
    }
  })
</script>

<template>
  <div ref="chartRef" class="chart" />
</template>

<style scoped lang="scss">
  .chart {
    height: 250px;
    width: 100%;
    border: 1px solid var(--el-border-color-dark);
    background: #fff;
  }
</style>
