<script setup>
  import { onMounted, onBeforeUnmount, ref, nextTick } from 'vue'
  import * as echarts from 'echarts'
  import { formatAmountCurrency } from '@/utils/format'

  defineOptions({ name: 'LineBarDialog' })

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

    const { xData, yData, title, itemColor, areaColor } = chartData

    chartInstance.clear()
    chartInstance.setOption({
      tooltip: {
        trigger: 'axis',
        backgroundColor: '#fff',
        borderColor: '#eee',
        borderWidth: 1,
        textStyle: { color: '#333' },
        axisPointer: { type: 'shadow' },
        valueFormatter: (value) => formatAmountCurrency(value)
      },
      toolbox: {
        feature: { magicType: { show: true, type: ['line', 'bar'] } },
        top: 10,
        right: 10,
        emphasis: {
          iconStyle: {
            color: itemColor
          }
        }
      },
      title: {
        text: title,
        left: 'center',
        top: 20,
        textStyle: { fontSize: 16, fontWeight: 600 }
      },
      grid: { left: '2%', right: '2%', bottom: '4%', containLabel: true },
      xAxis: {
        type: 'category',
        data: xData,
        axisLabel: { rotate: 30 },
        boundaryGap: false
      },
      yAxis: {
        type: 'value',
        splitLine: { lineStyle: { type: 'dashed', color: '#ddd' } }
      },
      series: [
        {
          type: 'line',
          smooth: true,
          data: yData,
          symbol: 'circle',
          symbolSize: 4,
          itemStyle: { color: itemColor },
          lineStyle: { width: 2 },
          areaStyle: { color: areaColor }
        }
      ]
    })
  }

  // ==== dialog 逻辑 ====
  const showDialog = ref(false)
  const openDialog = async (data) => {
    showDialog.value = true
    await nextTick() // 确保 DOM 已存在
    initChart()
    updateChart(data)
  }

  // ==== 生命周期 ====
  const handleResize = () => {
    chartInstance?.resize()
  }

  onMounted(() => {
    window.addEventListener('resize', handleResize)
  })

  onBeforeUnmount(() => {
    window.removeEventListener('resize', handleResize)
    disposeChart()
  })

  defineExpose({ openDialog })
</script>

<template>
  <el-dialog width="80%" v-model="showDialog" align-center draggable>
    <div class="chart-wrapper">
      <div ref="chartRef" class="chart" />
    </div>
  </el-dialog>
</template>

<style scoped lang="scss">
  .chart-wrapper {
    .chart {
      height: 340px;
      border-radius: 8px;
      box-shadow: 0 6px 18px rgba(0, 0, 0, 0.06);
    }
  }
</style>
