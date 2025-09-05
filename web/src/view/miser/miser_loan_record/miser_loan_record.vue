<template>
  <div>
    <!-- 统计卡片 -->
    <div class="card-stat-box">
      <loan-card-stat ref="cardStatRef" />
    </div>

    <!-- 搜索框 -->
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="借还对象" prop="name">
          <el-select
            v-model="searchInfo.name"
            allow-create
            filterable
            default-first-option
            placeholder="借还对象"
          >
            <el-option
              v-for="name in distinctNames"
              :key="name"
              :label="name"
              :value="name"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="资金状态" prop="fundStatus">
          <el-select v-model="searchInfo.fundStatus">
            <el-option
              v-for="{ value, label } in loanStatusList"
              :key="value"
              :label="label"
              :value="value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="借出日期" prop="lendDate">
          <el-date-picker
            v-model="searchInfo.lendDate"
            type="date"
            placeholder="请选择"
            :disabled-date="disabledDate"
          />
        </el-form-item>

        <el-form-item label="归还日期" prop="repayDate">
          <el-date-picker
            v-model="searchInfo.repayDate"
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
        <el-button type="primary" icon="CirclePlus" @click="handleLoan()"
          >借出
        </el-button>
        <el-button icon="Plus" @click="openDialog()">新增</el-button>
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
        >
          删除
        </el-button>
      </div>

      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
      >
        <el-table-column align="center" type="selection" width="55" />

        <el-table-column align="center" label="ID" prop="id" width="55" />

        <el-table-column align="center" label="借还对象" prop="name" />

        <el-table-column align="center" label="借出日期" prop="lendDate">
          <template #default="scope"
            >{{ formatDate(scope.row.lendDate, 'yyyy-MM-dd') }}
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="借出金额"
          prop="lendAmount"
          :formatter="formatterAmount"
        />

        <el-table-column align="center" label="归还日期" prop="repayDate">
          <template #default="scope"
            >{{ formatDate(scope.row.repayDate, 'yyyy-MM-dd') }}
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="归还金额"
          prop="repayAmount"
          :formatter="formatterAmount"
        />

        <el-table-column align="center" label="资金状态" prop="fundStatus">
          <template #default="{ row }">
            <el-tag :type="loanStatusMap[row.fundStatus]?.color"
              >{{ loanStatusMap[row.fundStatus]?.label }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="结清耗时"
          prop="settlementDuration"
        >
          <template #default="{ row }">
            {{ formatDuration(row) }}
          </template>
        </el-table-column>

        <el-table-column align="center" label="操作" fixed="right" width="280">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="CirclePlus"
              class="table-button"
              @click="handleRepay(scope.row)"
              :disabled="scope.row.fundStatus === LOAN_FUND_STATUS_REPAID"
              >归还
            </el-button>
            <el-button
              type="info"
              link
              class="table-button"
              @click="getDetails(scope.row)"
            >
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              详情
            </el-button>
            <el-button
              type="warning"
              link
              icon="edit"
              class="table-button"
              @click="updateMiserLoanRecordFunc(scope.row)"
              >编辑
            </el-button>
            <el-button
              type="danger"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
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
          <span class="text-lg">{{ drawerTitle }}</span>
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
        <el-form-item label="借还对象" prop="name">
          <el-select
            v-model="formData.name"
            allow-create
            filterable
            default-first-option
            placeholder="请输入借还对象"
            :disabled="!renderLoanItem"
          >
            <el-option
              v-for="name in distinctNames"
              :key="name"
              :label="name"
              :value="name"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="借出日期" prop="lendDate">
          <el-date-picker
            v-model="formData.lendDate"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            :clearable="false"
            :disabled-date="disabledDate"
            :disabled="!renderLoanItem"
          />
        </el-form-item>
        <el-form-item label="借出金额" prop="lendAmount">
          <el-input-number
            v-model="formData.lendAmount"
            style="width: 100%"
            :min="0"
            :precision="2"
            placeholder="借出金额"
            :disabled="!renderLoanItem"
          />
        </el-form-item>
        <el-form-item label="借还说明" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            placeholder="请输入借还说明"
            maxlength="500"
            show-word-limit
            autosize
          />
        </el-form-item>
        <el-form-item label="归还日期" prop="repayDate">
          <el-date-picker
            v-model="formData.repayDate"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            :clearable="false"
            :disabled-date="disabledDate"
            :disabled="!renderRepayItem"
          />
        </el-form-item>
        <el-form-item label="归还金额" prop="repayAmount">
          <el-input-number
            v-model="formData.repayAmount"
            style="width: 100%"
            :min="0"
            :max="formData.lendAmount"
            :precision="2"
            :disabled="!renderRepayItem"
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
        <el-descriptions-item label="UID">
          {{ detailForm.userId }}
        </el-descriptions-item>
        <el-descriptions-item label="借还对象">
          {{ detailForm.name }}
        </el-descriptions-item>
        <el-descriptions-item label="借出日期">
          {{ formatDate(detailForm.lendDate, 'yyyy-MM-dd') }}
        </el-descriptions-item>
        <el-descriptions-item label="借出金额">
          {{ formatAmount(detailForm.lendAmount) }}
        </el-descriptions-item>
        <el-descriptions-item label="归还日期">
          {{ formatDate(detailForm.repayDate, 'yyyy-MM-dd') }}
        </el-descriptions-item>
        <el-descriptions-item label="归还金额">
          {{ formatAmount(detailForm.repayAmount) }}
        </el-descriptions-item>
        <el-descriptions-item label="借还说明">
          <multiline-text :text="detailForm.description" />
        </el-descriptions-item>
        <el-descriptions-item label="资金状态">
          <el-tag :type="loanStatusMap[detailForm.fundStatus]?.color"
            >{{ loanStatusMap[detailForm.fundStatus]?.label }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="结清耗时">
          {{ formatDuration(detailForm) }}
        </el-descriptions-item>
        <el-descriptions-item label="创建时间">
          {{ formatDate(detailForm.createdAt) }}
        </el-descriptions-item>
        <el-descriptions-item label="更新时间">
          {{ formatDate(detailForm.updatedAt) }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    createMiserLoanRecord,
    deleteMiserLoanRecord,
    deleteMiserLoanRecordByIds,
    updateMiserLoanRecord,
    findMiserLoanRecord,
    getMiserLoanRecordList,
    listMiserLoanNameList
  } from '@/api/miser/miser_loan_record'
  import {
    formatDate,
    getDictFunc,
    formatAmount,
    formatterAmount
  } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, computed, onMounted, watch } from 'vue'
  import { useAppStore } from '@/pinia'
  import { InfoFilled } from '@element-plus/icons-vue'
  import MultilineText from '@/components/MultilineText.vue'
  import LoanCardStat from '@/view/miser/miser_loan_record/components/LoanCardStat.vue'

  defineOptions({
    name: 'MiserLoanRecord'
  })

  const cardStatRef = ref()

  const LOAN_FUND_STATUS_LOANING = 1 // 待还款
  const LOAN_FUND_STATUS_REPAYING = 2 // 部分还
  const LOAN_FUND_STATUS_REPAID = 3 // 已结清

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    name: '',
    lendDate: new Date(),
    lendAmount: null,
    repayDate: new Date(),
    repayAmount: 0,
    description: ''
  })

  const disabledDate = (time) => {
    return time.getTime() > Date.now()
  }
  const formatDuration = ({ fundStatus, settlementDuration }) => {
    const totalDays = settlementDuration
    if (fundStatus !== LOAN_FUND_STATUS_REPAID || totalDays < 0) return '-'
    if (totalDays === 0) return '0 天'

    const years = Math.floor(totalDays / 360)
    const months = Math.floor((totalDays % 360) / 30)
    const days = totalDays % 30

    const parts = []
    if (years) parts.push(`${years} 年`)
    if (months) parts.push(` ${months} 个月`)
    if (days) {
      if (years || months) parts.push('零')
      parts.push(` ${days} 天`)
    }

    return parts.join('').trim()
  }

  // 验证规则
  const rule = reactive({
    name: [
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
    ],
    lendDate: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    lendAmount: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    description: [
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
    ],
    repayDate: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    repayAmount: [
      {
        required: true,
        message: '',
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
  const searchInfo = ref({})
  // 重置
  const onReset = () => {
    searchInfo.value = {}
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
    const table = await getMiserLoanRecordList({
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
    const tip = `确定要删除『${row.name}』吗？`
    ElMessageBox.confirm(tip, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteMiserLoanRecordFunc(row)
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
      const res = await deleteMiserLoanRecordByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        await getTableData()
        await cardStatRef.value.fetchAndRender()
      }
    })
  }

  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')
  const renderLoanItem = computed(() => {
    return type.value !== 'repay'
  })
  const renderRepayItem = computed(() => {
    return type.value !== 'loan'
  })
  const drawerTitle = computed(() => {
    switch (type.value) {
      case 'create':
        return '新增'
      case 'update':
        return '编辑'
      case 'loan':
        return '借出'
      case 'repay':
        return '归还'
      default:
        return '未知'
    }
  })

  // 更新行
  const updateMiserLoanRecordFunc = async (row) => {
    const res = await findMiserLoanRecord({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 资金归还
  const handleRepay = async (row) => {
    type.value = 'repay'
    const res = await findMiserLoanRecord({ id: row.id })
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteMiserLoanRecordFunc = async (row) => {
    const res = await deleteMiserLoanRecord({ id: row.id })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      await getTableData()
      await cardStatRef.value.fetchAndRender()
    }
  }

  // 弹窗控制标记
  const dialogFormVisible = ref(false)

  // 打开弹窗
  const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
  }

  const handleLoan = () => {
    type.value = 'loan'
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
      name: '',
      lendDate: new Date(),
      lendAmount: null,
      repayDate: new Date(),
      repayAmount: 0,
      description: ''
    }
  }
  // 弹窗确定
  const enterDialog = async () => {
    btnLoading.value = true
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return (btnLoading.value = false)
      let res
      switch (type.value) {
        case 'loan':
          res = await createMiserLoanRecord(formData.value)
          break
        case 'create':
          res = await createMiserLoanRecord(formData.value)
          break
        case 'repay':
          res = await updateMiserLoanRecord(formData.value)
          break
        case 'update':
          res = await updateMiserLoanRecord(formData.value)
          break
        default:
          res = await createMiserLoanRecord(formData.value)
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
        await cardStatRef.value.fetchAndRender()
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
    const res = await findMiserLoanRecord({ id: row.id })
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

  // 借还资金状态
  const loanStatusList = ref([])
  const loanStatusMap = computed(() => {
    const resultMap = {}
    loanStatusList.value.forEach(({ value, label, extend }) => {
      resultMap[value] = {
        label: label,
        color: extend
      }
    })
    return resultMap
  })
  const fetchLoadStatusList = async () => {
    const dataList = await getDictFunc('miser_loan_fund_status')
    loanStatusList.value = dataList.map(({ value, label, extend }) => {
      return {
        label: label,
        value: parseInt(value),
        extend: extend
      }
    })
  }

  // 借还对象列表
  const distinctNames = ref([])
  const fetchDistinctNameList = async () => {
    const { code, data } = await listMiserLoanNameList()
    if (code === 0 && data && data.length > 0) {
      distinctNames.value = data
    } else {
      distinctNames.value = []
    }
  }

  watch(
    total,
    async () => {
      await fetchDistinctNameList()
    },
    { immediate: true }
  )

  onMounted(() => {
    fetchLoadStatusList()
  })
</script>

<style lang="scss" scoped>
  .card-stat-box {
    margin-top: 8px;
  }
</style>
