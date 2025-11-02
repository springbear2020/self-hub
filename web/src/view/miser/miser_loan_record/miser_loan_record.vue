<template>
  <div>
    <!-- 统计卡片 -->
    <div class="card-stat-box">
      <loan-card-stat ref="cardRef" @open="handleOpen" />
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
            filterable
            default-first-option
            placeholder="请选择"
          >
            <el-option
              v-for="name in distinctNameList"
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
        <el-button type="primary" icon="CirclePlus" @click="handleLoan"
          >借出
        </el-button>
        <el-button icon="Plus" @click="openDialog()">新增</el-button>
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
        <el-table-column align="center" label="借还对象" prop="name" />
        <el-table-column align="center" label="借出日期" prop="lendDate">
          <template #default="{ row }"
            >{{ formatDate(row.lendDate, 'yyyy-MM-dd') }}
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="借出金额"
          prop="lendAmount"
          :formatter="formatterAmount"
        />
        <el-table-column align="center" label="归还日期" prop="repayDate">
          <template #default="{ row }"
            >{{ formatDate(row.repayDate, 'yyyy-MM-dd') }}
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
          <template #default="{ row }">
            <el-button
              type="primary"
              icon="CirclePlus"
              link
              @click="handleRepay(row)"
              :disabled="disableRepay(row)"
              >归还
            </el-button>
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
              @click="updateMiserLoanRecordFunc(row)"
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
              v-for="name in distinctNameList"
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
            placeholder="请选择"
            :clearable="false"
            :disabled-date="disabledDate"
            :disabled="!renderLoanItem"
          />
        </el-form-item>
        <el-form-item label="借出金额" prop="lendAmount">
          <el-input-number
            v-model="formData.lendAmount"
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
            placeholder="请选择"
            :clearable="false"
            :disabled-date="disabledDate"
            :disabled="!renderRepayItem"
          />
        </el-form-item>
        <el-form-item label="归还金额" prop="repayAmount">
          <el-input-number
            v-model="formData.repayAmount"
            :min="0"
            :max="formData.lendAmount"
            :precision="2"
            :disabled="!renderRepayItem"
            placeholder="归还金额"
          />
          <el-button
            link
            type="primary"
            icon="DocumentCopy"
            title="全额归还"
            style="margin-left: 8px"
            @click="handleFullRepay"
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

    <line-bar-dialog ref="dialogRef" />
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
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, computed, onMounted } from 'vue'
  import { useAppStore } from '@/pinia'
  import {
    formatDate,
    getDictFunc,
    formatAmount,
    formatterAmount
  } from '@/utils/format'
  import MultilineText from '@/components/MultilineText.vue'
  import LoanCardStat from '@/view/miser/miser_loan_record/components/LoanCardStat.vue'
  import { miserLoanCfgMap } from '@/constants/miser'
  import LineBarDialog from '@/components/LineBarDialog.vue'

  defineOptions({
    name: 'MiserLoanRecord'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    name: '',
    lendDate: new Date(),
    lendAmount: undefined,
    repayDate: new Date(),
    repayAmount: 0,
    description: ''
  })

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
  const searchInfo = ref({
    name: undefined,
    fundStatus: undefined,
    lendDate: undefined,
    repayDate: undefined
  })
  // 重置
  const onReset = () => {
    searchInfo.value = {
      name: undefined,
      fundStatus: undefined,
      lendDate: undefined,
      repayDate: undefined
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
    const tip = `确定要删除『${row.name}/${formatDate(row.lendDate, 'yyyy-MM-dd')}』吗？`
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
        await fetchDistinctNameList()
        await cardRef.value.fetchAndRender()
      }
    })
  }

  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')

  // 更新行
  const updateMiserLoanRecordFunc = async (row) => {
    const res = await findMiserLoanRecord({ id: row.id })
    type.value = 'update'
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
      await fetchDistinctNameList()
      await cardRef.value.fetchAndRender()
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
      name: '',
      lendDate: new Date(),
      lendAmount: undefined,
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
        await fetchDistinctNameList()
        await cardRef.value.fetchAndRender()
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

  /********************************************************************************************************************/

  // 统计卡片
  const cardRef = ref()

  // 对话框
  const dialogRef = ref()
  const handleOpen = (data) => {
    dialogRef.value.openDialog(data)
  }

  // 禁用日期
  const disabledDate = (time) => {
    return time.getTime() > Date.now()
  }

  // 抽屉标题
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

  // 借出
  const handleLoan = () => {
    type.value = 'loan'
    dialogFormVisible.value = true
  }
  const renderLoanItem = computed(() => {
    return type.value !== 'repay'
  })

  // 归还
  const handleRepay = async (row) => {
    type.value = 'repay'
    const { code, data } = await findMiserLoanRecord({ id: row.id })
    if (code === 0) {
      formData.value = data
      formData.value.repayDate = new Date()
      dialogFormVisible.value = true
    }
  }
  const renderRepayItem = computed(() => {
    return type.value !== 'loan'
  })
  const disableRepay = ({ fundStatus }) => {
    return fundStatus === miserLoanCfgMap.repaid
  }
  const handleFullRepay = () => {
    formData.value.repayAmount = formData.value.lendAmount
  }

  // 结清耗时
  const formatDuration = ({ fundStatus, settlementDuration }) => {
    const totalDays = settlementDuration
    if (fundStatus !== miserLoanCfgMap.repaid || totalDays < 0) {
      return '-'
    }
    if (totalDays === 0) {
      return '0 天'
    }

    const years = Math.floor(totalDays / 360)
    const months = Math.floor((totalDays % 360) / 30)
    const days = totalDays % 30

    const parts = []
    if (years) {
      parts.push(`${years} 年`)
    }
    if (months) {
      parts.push(` ${months} 个月`)
    }
    if (days) {
      if (years || months) parts.push('零')
      parts.push(` ${days} 天`)
    }

    return parts.join('').trim()
  }

  // 资金状态
  const loanStatusList = ref([])
  const loanStatusMap = ref({})
  const fetchLoanStatusList = async () => {
    const dataList = await getDictFunc('miser_loan_fund_status')

    const list = []
    const map = {}
    dataList.forEach(({ value, label, extend }) => {
      const intValue = parseInt(value)
      list.push({ value: intValue, label, extend })
      map[intValue] = { label, color: extend }
    })

    loanStatusList.value = list
    loanStatusMap.value = map
  }

  // 借还对象
  const distinctNameList = ref([])
  const fetchDistinctNameList = async () => {
    const { code, data } = await listMiserLoanNameList()
    if (code === 0 && data && data.length > 0) {
      distinctNameList.value = data
    } else {
      distinctNameList.value = []
    }
  }

  onMounted(async () => {
    await fetchLoanStatusList()
    await fetchDistinctNameList()
  })
</script>

<style lang="scss" scoped>
  .card-stat-box {
    margin-top: 8px;
  }
</style>
