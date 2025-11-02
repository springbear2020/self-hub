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
        <el-form-item label="项目名称" prop="name">
          <el-input v-model="searchInfo.name" placeholder="项目名称" />
        </el-form-item>

        <el-form-item label="交易日期" prop="date">
          <el-date-picker
            v-model="searchInfo.date"
            type="date"
            placeholder="请选择"
            :disabled-date="disabledDate"
          />
        </el-form-item>

        <el-form-item label="交易分类" prop="categoryId">
          <el-select v-model="searchInfo.categoryId">
            <el-option
              v-for="c in catStore.dataList"
              :label="c.name"
              :value="c.id"
              :key="c.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="交易类型" prop="transactionType">
          <el-select v-model="searchInfo.transactionType">
            <el-option
              v-for="t in txnStore.dataList"
              :label="t.label"
              :value="t.value"
              :key="t.value"
            />
          </el-select>
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
        <el-button type="primary" icon="plus" @click="openDialog()"
          >新增
        </el-button>
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

        <el-table-column align="center" label="项目名称" prop="name" />

        <el-table-column align="center" label="交易日期" prop="date">
          <template #default="{ row }"
            >{{ formatDate(row.date, 'yyyy-MM-dd') }}
          </template>
        </el-table-column>

        <el-table-column
          align="center"
          label="交易金额"
          prop="amount"
          :formatter="formatterAmount"
        />

        <el-table-column align="center" label="交易分类" prop="categoryId">
          <template #default="{ row }">
            {{ catStore.dataMap[row.categoryId] }}
          </template>
        </el-table-column>

        <el-table-column align="center" label="交易类型" prop="transactionType">
          <template #default="{ row }">
            <el-tag :type="miserTxnCfgMap[row.transactionType]?.tagType"
              >{{ txnStore.dataMap[row.transactionType] }}
            </el-tag>
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
              @click="updateMiserRankingRecordFunc(row)"
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

      <el-alert type="info" show-icon :closable="false" class="alert-tip"
        >重点记录单笔金额超过 500
        元的非刚性收支，并参与榜单排名，助您对大额收支洞若观火
      </el-alert>

      <el-form
        :model="formData"
        label-position="top"
        ref="elFormRef"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="交易分类" prop="categoryId">
          <el-select
            v-model="formData.categoryId"
            @change="handleCategoryChange"
          >
            <el-option
              v-for="c in catStore.dataList"
              :key="c.id"
              :label="c.name"
              :value="c.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="交易类型" prop="transactionType">
          <el-select v-model="formData.transactionType" disabled>
            <el-option
              v-for="t in txnStore.dataList"
              :label="t.label"
              :value="t.value"
              :key="t.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="项目名称" prop="name">
          <el-input
            v-model="formData.name"
            :clearable="true"
            placeholder="请输入项目名称"
          />
        </el-form-item>
        <el-form-item label="交易日期" prop="date">
          <el-date-picker
            v-model="formData.date"
            type="date"
            placeholder="请选择"
            :clearable="false"
            :disabled-date="disabledDate"
          />
        </el-form-item>
        <el-form-item label="交易金额" prop="amount">
          <el-input-number
            v-model="formData.amount"
            :min="0"
            :precision="2"
            placeholder="交易金额"
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
        <el-descriptions-item label="交易分类">
          {{ catStore.dataMap[detailForm.categoryId] }}
        </el-descriptions-item>
        <el-descriptions-item label="交易类型">
          <el-tag :type="miserTxnCfgMap[detailForm.transactionType]?.tagType"
            >{{ txnStore.dataMap[detailForm.transactionType] }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="项目名称">
          {{ detailForm.name }}
        </el-descriptions-item>
        <el-descriptions-item label="交易日期">
          {{ formatDate(detailForm.date, 'yyyy-MM-dd') }}
        </el-descriptions-item>
        <el-descriptions-item label="交易金额">
          {{ formatAmount(detailForm.amount) }}
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
    createMiserRankingRecord,
    deleteMiserRankingRecord,
    deleteMiserRankingRecordByIds,
    updateMiserRankingRecord,
    findMiserRankingRecord,
    getMiserRankingRecordList
  } from '@/api/miser/miser_ranking_record'
  import { formatDate, formatterAmount, formatAmount } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, onMounted } from 'vue'
  import {
    useAppStore,
    useMiserCategoryStore,
    useMiserTransactionTypeStore
  } from '@/pinia'
  import { miserTxnCfgMap } from '@/constants/miser'

  defineOptions({
    name: 'MiserRankingRecord'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    transactionType: undefined,
    categoryId: undefined,
    name: '',
    amount: 0,
    date: new Date()
  })

  // 验证规则
  const rule = reactive({
    transactionType: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    categoryId: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
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
    amount: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    date: [
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
    date: undefined,
    categoryId: undefined,
    transactionType: undefined
  })
  // 重置
  const onReset = () => {
    searchInfo.value = {
      name: undefined,
      date: undefined,
      categoryId: undefined,
      transactionType: undefined
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
    const table = await getMiserRankingRecordList({
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
    const tradeDate = formatDate(row.date, 'yyyy-MM-dd')
    const tip = `确定要删除『${tradeDate}/${row.name}』吗？`
    ElMessageBox.confirm(tip, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteMiserRankingRecordFunc(row)
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
      const res = await deleteMiserRankingRecordByIds({ ids })
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
  const updateMiserRankingRecordFunc = async (row) => {
    const res = await findMiserRankingRecord({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteMiserRankingRecordFunc = async (row) => {
    const res = await deleteMiserRankingRecord({ id: row.id })
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
      transactionType: undefined,
      categoryId: undefined,
      name: '',
      amount: 0,
      date: new Date()
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
          res = await createMiserRankingRecord(formData.value)
          break
        case 'update':
          res = await updateMiserRankingRecord(formData.value)
          break
        default:
          res = await createMiserRankingRecord(formData.value)
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
    const res = await findMiserRankingRecord({ id: row.id })
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
  // 禁用日期
  const disabledDate = (time) => {
    return time.getTime() > Date.now()
  }

  // 交易分类 -> 交易类型
  const handleCategoryChange = (categoryId) => {
    formData.value.transactionType = catStore.catTxnMap[categoryId]
  }

  // 交易分类
  const catStore = useMiserCategoryStore()

  // 交易类型
  const txnStore = useMiserTransactionTypeStore()

  onMounted(async () => {
    await catStore.init()
    await txnStore.init()
  })
</script>

<style scoped lang="scss">
  .alert-tip {
    margin-bottom: 20px;
  }
</style>
