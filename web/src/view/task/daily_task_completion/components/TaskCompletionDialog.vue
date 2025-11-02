<script setup>
  import { ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { createDailyTaskCompletion } from '@/api/task/daily_task_completion'

  defineOptions({ name: 'TaskCompletionDialog' })

  const props = defineProps({
    dailyTasks: { type: Array, required: true }
  })
  const emits = defineEmits(['refresh'])

  const formData = ref([])
  const showDialog = ref(false)

  const openDialog = async () => {
    const today = new Date()
    formData.value = props.dailyTasks.map(
      ({ id, name, minValue, maxValue }) => {
        return {
          taskId: id,
          finishDate: today,
          countValue: minValue,
          remark: '',

          name,
          minValue,
          maxValue
        }
      }
    )

    showDialog.value = true
  }
  const closeDialog = () => {
    showDialog.value = false
  }

  const handleSubmit = async () => {
    const { code } = await createDailyTaskCompletion(formData.value)
    if (code === 0) {
      ElMessage.success('批量新增今日任务完成情况成功')
      emits('refresh')
      closeDialog()
    }
  }
  const handleRemove = (index) => {
    formData.value.splice(index, 1)
  }

  defineExpose({ openDialog })
</script>

<template>
  <el-dialog
    v-model="showDialog"
    title="批量新增今日任务完成情况"
    :close-on-click-modal="false"
  >
    <el-form :model="formData">
      <el-table :data="formData" border>
        <el-table-column label="任务名称" align="center" prop="name" />

        <el-table-column label="完成说明" align="center" prop="remark">
          <template #default="{ row }">
            <el-input
              v-model="row.remark"
              placeholder="请输入任务完成说明"
              type="textarea"
              maxlength="250"
              show-word-limit
              autosize
            />
          </template>
        </el-table-column>

        <el-table-column
          label="计数值"
          align="center"
          prop="countValue"
          width="180"
        >
          <template #default="{ row }">
            <el-input-number
              v-model="row.countValue"
              :min="row.minValue"
              :max="row.maxValue"
              :precision="0"
              placeholder="计数值"
            />
          </template>
        </el-table-column>

        <el-table-column label="操作" align="center" width="80">
          <template #default="{ $index }">
            <el-button
              type="danger"
              icon="delete"
              @click="handleRemove($index)"
            />
          </template>
        </el-table-column>
      </el-table>
    </el-form>

    <template #footer>
      <el-button @click="closeDialog">取消</el-button>
      <el-button type="primary" @click="handleSubmit">保存</el-button>
    </template>
  </el-dialog>
</template>
