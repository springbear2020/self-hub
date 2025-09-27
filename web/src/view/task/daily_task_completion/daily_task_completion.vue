<template>
  <div>
    <!-- 搜索框 -->
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="任务名称" prop="taskId">
          <el-select v-model="searchInfo.taskId">
            <el-option
              v-for="task in dailyTaskList"
              :key="task.id"
              :label="task.name"
              :value="task.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="完成日期" prop="finishDate">
          <el-date-picker
            v-model="searchInfo.finishDate"
            type="date"
            placeholder="请选择"
            :disabled-date="disabledDate"
          />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询
          </el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 表格 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="CirclePlus" @click="handleToday"
          >今日
        </el-button>
        <el-button icon="plus" @click="openDialog()">新增</el-button>
        <el-button
          icon="delete"
          :disabled="!multipleSelection.length"
          @click="onDelete"
          type="danger"
        >
          删除
        </el-button>
      </div>

      <el-table
        ref="multipleTable"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
      >
        <el-table-column align="center" type="selection" width="55" />
        <el-table-column align="center" label="ID" prop="id" width="55" />
        <el-table-column align="center" label="任务名称" prop="taskId">
          <template #default="{ row }"
            >{{ dailyTaskMap[row.taskId] }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="完成日期" prop="finishDate">
          <template #default="{ row }"
            >{{ formatDate(row.finishDate, 'yyyy-MM-dd') }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="计数值" prop="countValue" />
        <el-table-column align="center" label="完成说明" prop="remark">
          <template #default="{ row }">
            <multiline-text :text="row.remark" />
          </template>
        </el-table-column>

        <el-table-column align="center" label="操作" fixed="right" width="210">
          <template #default="{ row }">
            <el-button
              type="info"
              icon="InfoFilled"
              link
              @click="getDetails(row)"
            >
              详情
            </el-button>
            <el-button
              type="warning"
              link
              icon="edit"
              @click="updateDailyTaskCompletionFunc(row)"
              >编辑
            </el-button>
            <el-button type="danger" link icon="delete" @click="deleteRow(row)"
              >删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- 新增 / 编辑 -->
    <el-drawer
      destroy-on-close
      :size="appStore.drawerSize"
      v-model="dialogFormVisible"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog"
              >确 定
            </el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form
        :model="formData"
        label-position="top"
        ref="elFormRef"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="任务名称" prop="taskId">
          <el-select v-model="formData.taskId">
            <el-option
              v-for="task in dailyTaskList"
              :key="task.id"
              :label="task.name"
              :value="task.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="完成说明" prop="remark">
          <el-input
            v-model="formData.remark"
            placeholder="请输入完成说明"
            type="textarea"
            maxlength="250"
            show-word-limit
            autosize
          />
        </el-form-item>
        <el-form-item label="完成日期" prop="finishDate">
          <el-date-picker
            v-model="formData.finishDate"
            type="date"
            placeholder="请选择"
            :disabled-date="disabledDate"
            :clearable="false"
          />
        </el-form-item>
        <el-form-item label="计数值" prop="countValue">
          <el-input-number
            v-model="formData.countValue"
            :min="countValueRange[0]"
            :max="countValueRange[1]"
            :precision="0"
            placeholder="计数值"
          />
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 详情 -->
    <el-drawer
      destroy-on-close
      :size="appStore.drawerSize"
      v-model="detailShow"
      :show-close="true"
      :before-close="closeDetailShow"
      title="详情"
    >
      <el-descriptions :column="1" border>
        <el-descriptions-item label="ID">
          {{ detailForm.id }}
        </el-descriptions-item>
        <el-descriptions-item label="任务名称">
          {{ dailyTaskMap[detailForm.taskId] }}
        </el-descriptions-item>
        <el-descriptions-item label="完成日期">
          {{ formatDate(detailForm.finishDate, 'yyyy-MM-dd') }}
        </el-descriptions-item>
        <el-descriptions-item label="计数值">
          {{ detailForm.countValue }}
        </el-descriptions-item>
        <el-descriptions-item label="完成说明">
          <multiline-text :text="detailForm.remark" />
        </el-descriptions-item>
        <el-descriptions-item label="创建时间">
          {{ formatDate(detailForm.createdAt) }}
        </el-descriptions-item>
        <el-descriptions-item label="更新时间">
          {{ formatDate(detailForm.updatedAt) }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

    <!-- 批量新增 -->
    <task-completion-dialog
      ref="completionRef"
      :daily-tasks="dailyTaskList"
      @refresh="getTableData"
    />
  </div>
</template>

<script setup>
  import {
    createDailyTaskCompletion,
    deleteDailyTaskCompletion,
    deleteDailyTaskCompletionByIds,
    updateDailyTaskCompletion,
    findDailyTaskCompletion,
    getDailyTaskCompletionList
  } from '@/api/task/daily_task_completion'
  import { listActiveTaskList } from '@/api/task/daily_task'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, onMounted, computed } from 'vue'
  import { useAppStore } from '@/pinia'
  import { formatDate } from '@/utils/format'
  import TaskCompletionDialog from './components/TaskCompletionDialog.vue'
  import MultilineText from '@/components/MultilineText.vue'

  defineOptions({
    name: 'DailyTaskCompletion'
  })

  // 提交按钮
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    taskId: undefined,
    finishDate: new Date(),
    countValue: 1,
    remark: ''
  })

  // 验证规则
  const rule = reactive({
    taskId: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    finishDate: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    countValue: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    remark: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      },
      {
        whitespace: true,
        message: '不能只输入空格',
        trigger: ['input', 'blur']
      }
    ]
  })

  const elFormRef = ref()
  const elSearchFormRef = ref()

  // =========== 表格控制部分 ===========
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({
    taskId: undefined,
    finishDate: undefined
  })
  // 重置
  const onReset = () => {
    searchInfo.value = {
      taskId: undefined,
      finishDate: undefined
    }
    getTableData()
  }

  // 搜索
  const onSubmit = () => {
    elSearchFormRef.value?.validate(async (valid) => {
      if (!valid) return
      page.value = 1
      await getTableData()
    })
  }

  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  // 修改页面容量
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  // 查询
  const getTableData = async () => {
    const table = await getDailyTaskCompletionList({
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  getTableData()

  // ============== 表格控制部分结束 ===============

  // 获取需要的字典 可能为空 按需保留
  const setOptions = async () => {}

  // 获取需要的字典 可能为空 按需保留
  setOptions()

  // 多选数据
  const multipleSelection = ref([])
  // 多选
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }

  // 删除行
  const deleteRow = (row) => {
    const tip = `确定要删除『${dailyTaskMap.value[row.taskId]}/${formatDate(row.finishDate, 'yyyy-MM-dd')}』吗？`
    ElMessageBox.confirm(tip, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteDailyTaskCompletionFunc(row)
    })
  }

  // 多选删除
  const onDelete = async () => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map((item) => {
          ids.push(item.id)
        })
      const res = await deleteDailyTaskCompletionByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        await getTableData()
      }
    })
  }

  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')

  // 更新行
  const updateDailyTaskCompletionFunc = async (row) => {
    const res = await findDailyTaskCompletion({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteDailyTaskCompletionFunc = async (row) => {
    const res = await deleteDailyTaskCompletion({ id: row.id })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      await getTableData()
    }
  }

  // 弹窗控制标记
  const dialogFormVisible = ref(false)

  // 打开弹窗
  const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
      taskId: undefined,
      finishDate: new Date(),
      countValue: 1,
      remark: ''
    }
  }
  // 弹窗确定
  const enterDialog = async () => {
    btnLoading.value = true
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return (btnLoading.value = false)
      let res
      switch (type.value) {
        case 'create':
          res = await createDailyTaskCompletion([formData.value])
          break
        case 'update':
          res = await updateDailyTaskCompletion(formData.value)
          break
        default:
          res = await createDailyTaskCompletion(formData.value)
          break
      }
      btnLoading.value = false
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        closeDialog()
        await getTableData()
      }
    })
  }

  const detailForm = ref({})

  // 查看详情控制标记
  const detailShow = ref(false)

  // 打开详情弹窗
  const openDetailShow = () => {
    detailShow.value = true
  }

  // 打开详情
  const getDetails = async (row) => {
    // 打开弹窗
    const res = await findDailyTaskCompletion({ id: row.id })
    if (res.code === 0) {
      detailForm.value = res.data
      openDetailShow()
    }
  }

  // 关闭详情弹窗
  const closeDetailShow = () => {
    detailShow.value = false
    detailForm.value = {}
  }

  /********************************************************************************************************************/

  // 批量新增今日任务完成情况
  const completionRef = ref()
  const handleToday = () => {
    completionRef.value.openDialog()
  }

  // 计数值范围
  const countValueRange = computed(() => {
    const target = dailyTaskList.value.find(
      ({ id }) => id === formData.value.taskId
    )
    const { minValue = 1, maxValue = 100 } = target || {}
    return [minValue, maxValue]
  })

  // 日期可选范围
  const disabledDate = (time) => {
    return time.getTime() > Date.now()
  }

  // 任务列表
  const dailyTaskList = ref([])
  const dailyTaskMap = ref({})
  const fetchTaskList = async () => {
    const { code, data } = await listActiveTaskList()
    if (code === 0) {
      dailyTaskList.value = data
      dailyTaskMap.value = data.reduce((map, { id, name }) => {
        map[id] = name
        return map
      }, {})
    } else {
      dailyTaskList.value = []
      dailyTaskMap.value = {}
    }
  }

  onMounted(async () => {
    await fetchTaskList()
  })
</script>
