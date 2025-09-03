<script setup>
  import * as echarts from 'echarts'
  import { onMounted, onBeforeUnmount, ref } from 'vue'
  import { formatDate } from '@/utils/format'
  import { ElMessage } from 'element-plus'
  import { createDailyTaskCompletion } from '@/api/task/daily_task_completion'

  const emits = defineEmits(['refresh'])
  const props = defineProps({
    year: { type: Number, default: new Date().getFullYear() },
    task: { type: Object, required: true, default: () => ({}) },
    completions: { type: Array, default: () => [] }
  })

  const chartRef = ref()
  let chartInstance = null
  const formData = ref({})
  const showForm = ref(false)
  const today = new Date()
  const todayStr = formatDate(today, 'yyyy-MM-dd')

  const handleSubmit = async () => {
    const { code } = await createDailyTaskCompletion(formData.value)
    if (code === 0) {
      ElMessage.success('任务完成情况已保存')
      showForm.value = false
      emits('refresh')
    }
  }
  const getFullYearData = (year, rawData) => {
    const start = new Date(year, 0, 1)
    const end = new Date(year, 11, 31)
    const mapData = new Map(rawData.map((item) => [item[0], item[1]]))

    const results = []
    let cur = new Date(start)
    while (cur <= end) {
      // 数据库存在则使用数据库数据
      // 数据库不存在时，若日期是今天则赋值为 0，否则为 null
      const dateStr = cur.toISOString().slice(0, 10)
      const value = mapData.get(dateStr) ?? (dateStr === todayStr ? 0 : null)
      results.push([dateStr, value])
      cur.setDate(cur.getDate() + 1)
    }

    return results
  }
  const renderChart = () => {
    // 初始化图表并绑定事件
    if (!chartInstance) {
      chartInstance = echarts.init(chartRef.value)
      chartInstance.on('click', ({ seriesType, value }) => {
        if (seriesType === 'heatmap') {
          const [date, count] = value
          if (date === todayStr && count === 0) {
            // 仅允许添加今日尚未添加的数据
            showForm.value = true
          }
        }
      })
    }

    // 设置 options 配置数据
    const fullData = getFullYearData(props.year, props.completions)
    const { id, name, minValue, maxValue } = props.task
    chartInstance.setOption({
      title: { top: 20, left: 'center', text: name },
      tooltip: {
        formatter: ({ value }) => `${value[0]}：${value[1]}`
      },
      visualMap: {
        min: minValue,
        max: maxValue,
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
        data: fullData
      }
    })

    // 表单数据赋值
    formData.value = {
      taskId: id,
      finishDate: today,
      countValue: 1,
      remark: ''
    }
  }

  const handleResize = () => {
    chartInstance?.resize()
  }

  onMounted(() => {
    renderChart()

    window.addEventListener('resize', handleResize)
  })
  onBeforeUnmount(() => {
    window.removeEventListener('resize', handleResize)

    chartInstance?.dispose()
    chartInstance = null
  })
</script>

<template>
  <div class="task-card">
    <div ref="chartRef" class="chart" />

    <!-- 内联输入表单 -->
    <transition name="fade">
      <div v-if="showForm" class="inline-form">
        <el-form :model="formData" label-width="80px">
          <el-form-item label="完成日期">
            <span>{{ formatDate(formData.finishDate, 'yyyy-MM-dd') }}</span>
          </el-form-item>

          <el-form-item label="计数值" required>
            <el-input-number
              v-model="formData.countValue"
              :min="props.task.minValue"
              :max="props.task.maxValue"
              placeholder="请输入任务完成计数值"
            />
          </el-form-item>

          <el-form-item label="完成说明" required>
            <el-input
              v-model="formData.remark"
              placeholder="请填写任务完成说明"
            />
          </el-form-item>

          <div class="form-actions">
            <el-button @click="showForm = false">取消</el-button>
            <el-button type="primary" @click="handleSubmit"> 保存</el-button>
          </div>
        </el-form>
      </div>
    </transition>
  </div>
</template>

<style scoped lang="scss">
  .task-card {
    border: 1px solid var(--el-border-color-dark);
    border-radius: 8px;
    background: #fff;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
    padding: 8px;
  }

  .chart {
    height: 250px;
    width: 100%;
  }

  .inline-form {
    padding: 16px;
    border-top: 1px solid #eee;
    background: #fafafa;
    border-radius: 8px;
  }

  .form-actions {
    text-align: right;
  }

  .fade-enter-active,
  .fade-leave-active {
    transition: all 0.3s ease;
  }

  .fade-enter-from,
  .fade-leave-to {
    opacity: 0;
    transform: translateY(-6px);
  }
</style>
