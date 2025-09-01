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
        <el-form-item label="交易分类" prop="categoryId">
          <el-select v-model="searchInfo.categoryId">
            <el-option
              v-for="c in categoryList"
              :label="c.name"
              :value="c.id"
              :key="c.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="交易类型" prop="transactionType">
          <el-select v-model="searchInfo.transactionType">
            <el-option
              v-for="t in transactionTypeList"
              :label="t.label"
              :value="t.value"
              :key="t.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="交易日期" prop="date">
          <el-date-picker
            v-model="searchInfo.date"
            type="date"
            placeholder="选择日期"
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
        <el-button
          type="primary"
          icon="circle-plus"
          @click="handleTransactionsBatch"
          >流水
        </el-button>
        <el-button icon="plus" @click="openDialog()">新增</el-button>
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
          type="danger"
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
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="ID" prop="id" width="110" />

        <el-table-column align="left" label="交易分类" prop="categoryId">
          <template #default="{ row }">
            {{ categoryMap[row.categoryId] }}
          </template>
        </el-table-column>

        <el-table-column align="left" label="交易类型" prop="transactionType">
          <template #default="{ row }">
            {{ transactionTypeMap[row.transactionType] }}
          </template>
        </el-table-column>

        <el-table-column align="left" label="交易日期" prop="date">
          <template #default="{ row }"
            >{{ formatDate(row.date, 'yyyy-MM-dd') }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="交易金额" prop="amount" />

        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="250"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              class="table-button"
              @click="handleItemsBatch(scope.row)"
            >
              <el-icon style="margin-right: 5px">
                <CirclePlus />
              </el-icon>
              明细
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
              @click="updateMiserTransactionFunc(scope.row)"
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
        <el-form-item label="交易分类" prop="categoryId">
          <el-select
            v-model="formData.categoryId"
            @change="handleCategoryChange"
          >
            <el-option
              v-for="c in categoryList"
              :key="c.id"
              :label="c.name"
              :value="c.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="交易类型" prop="transactionType">
          <el-select v-model="formData.transactionType" disabled>
            <el-option
              v-for="t in transactionTypeList"
              :label="t.label"
              :value="t.value"
              :key="t.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="交易日期" prop="date">
          <el-date-picker
            v-model="formData.date"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            :clearable="false"
            :disabled-date="disabledDate"
          />
        </el-form-item>
        <el-form-item label="交易金额" prop="amount">
          <el-input-number
            v-model="formData.amount"
            style="width: 100%"
            :precision="2"
            :clearable="true"
            placeholder="请输入交易金额"
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
      title="查看"
    >
      <el-descriptions :column="1" border>
        <el-descriptions-item label="ID">
          {{ detailForm.id }}
        </el-descriptions-item>
        <el-descriptions-item label="UID">
          {{ detailForm.userId }}
        </el-descriptions-item>
        <el-descriptions-item label="交易分类">
          {{ categoryMap[detailForm.categoryId] }}
        </el-descriptions-item>
        <el-descriptions-item label="交易类型">
          {{ transactionTypeMap[detailForm.transactionType] }}
        </el-descriptions-item>
        <el-descriptions-item label="交易日期">
          {{ formatDate(detailForm.date, 'yyyy-MM-dd') }}
        </el-descriptions-item>
        <el-descriptions-item label="交易金额">
          {{ detailForm.amount }}
        </el-descriptions-item>
        <el-descriptions-item label="创建时间">
          {{ formatDate(detailForm.createdAt) }}
        </el-descriptions-item>
        <el-descriptions-item label="更新时间">
          {{ formatDate(detailForm.updatedAt) }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

    <transaction-dialog
      ref="transactionDialogRef"
      :typed-category-map="typedCategoryMap"
      :transaction-type-map="transactionTypeMap"
      :grouped-category-map="groupedCategoryMap"
      @closed="handleDialogClosed"
    />
    <transaction-items-dialog
      ref="transactionItemsDialogRef"
      :category-map="categoryMap"
    />
  </div>
</template>

<script setup>
  import { formatDate, getDictFunc } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, computed, onMounted } from 'vue'
  import { useAppStore } from '@/pinia'
  import { CirclePlus, InfoFilled } from '@element-plus/icons-vue'
  import { listMiserCategoryList } from '@/api/miser/miser_category'
  import {
    createMiserTransaction,
    deleteMiserTransaction,
    deleteMiserTransactionByIds,
    updateMiserTransaction,
    findMiserTransaction,
    getMiserTransactionList
  } from '@/api/miser/miser_transaction'
  import TransactionDialog from '@/view/miser/miser_transaction/components/TransactionDialog.vue'
  import TransactionItemsDialog from '@/view/miser/miser_transaction/components/TransactionItemsDialog.vue'

  defineOptions({
    name: 'MiserTransaction'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    categoryId: undefined,
    transactionType: undefined,
    date: new Date(),
    amount: undefined
  })

  // 验证规则
  const rule = reactive({
    categoryId: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    transactionType: [
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
    ],
    amount: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ]
  })

  const elFormRef = ref()
  const elSearchFormRef = ref()
  const disabledDate = (time) => {
    return time.getTime() > Date.now()
  }
  const handleCategoryChange = (value) => {
    const { transactionType } = categoryList.value.find(
      ({ id }) => id === value
    )
    formData.value.transactionType = transactionType
  }

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
    const table = await getMiserTransactionList({
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
    const tip = `确定要删除『${categoryMap.value[row.categoryId]}/${formatDate(row.date, 'yyyy-MM-dd')}』吗？`
    ElMessageBox.confirm(tip, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteMiserTransactionFunc(row)
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
      const res = await deleteMiserTransactionByIds({ ids })
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
  const updateMiserTransactionFunc = async (row) => {
    const res = await findMiserTransaction({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteMiserTransactionFunc = async (row) => {
    const res = await deleteMiserTransaction({ id: row.id })
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
      categoryId: undefined,
      transactionType: undefined,
      date: new Date(),
      amount: undefined
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
          res = await createMiserTransaction([formData.value])
          break
        case 'update':
          res = await updateMiserTransaction(formData.value)
          break
        default:
          res = await createMiserTransaction(formData.value)
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
    const res = await findMiserTransaction({ id: row.id })
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

  // 交易类型
  const transactionTypeList = ref([])
  const transactionTypeMap = computed(() => {
    const resultMap = {}
    transactionTypeList.value.forEach(({ value, label }) => {
      resultMap[value] = label
    })
    return resultMap
  })
  const fetchTransactionTypeList = async () => {
    const dataList = await getDictFunc('miser_transaction_type')
    transactionTypeList.value = dataList.map(({ value, label }) => {
      return {
        label: label,
        value: parseInt(value)
      }
    })
  }

  // 分类列表
  const categoryList = ref([])
  const { categoryMap, typedCategoryMap, groupedCategoryMap } = (() => {
    // {categoryId: categoryName}
    const categoryMap = computed(() =>
      categoryList.value.reduce((map, { id, name }) => {
        map[id] = name
        return map
      }, {})
    )
    // {categoryId: transactionType}
    const typedCategoryMap = computed(() =>
      categoryList.value.reduce((map, { id, transactionType }) => {
        map[id] = transactionType
        return map
      }, {})
    )
    // {transactionType: categoryArray}
    const groupedCategoryMap = computed(() =>
      categoryList.value.reduce((map, item) => {
        const { id, name, sort, transactionType } = item
        if (!map[transactionType]) {
          map[transactionType] = []
        }
        map[transactionType].push({ id, name, sort, transactionType })
        return map
      }, {})
    )
    return { categoryMap, typedCategoryMap, groupedCategoryMap }
  })()
  const fetchCategoryList = async () => {
    const { code, data } = await listMiserCategoryList()
    if (code === 0 && data && data.length > 0) {
      categoryList.value = data
    } else {
      categoryList.value = []
    }
  }

  // 批量新增交易流水
  const transactionDialogRef = ref()
  const handleTransactionsBatch = () => {
    transactionDialogRef.value.openDialog()
  }
  const handleDialogClosed = async () => {
    await getTableData()
  }
  // 批量新增流水明细
  const transactionItemsDialogRef = ref()
  const handleItemsBatch = (row) => {
    transactionItemsDialogRef.value.openDialog(row)
  }

  onMounted(() => {
    fetchTransactionTypeList()
    fetchCategoryList()
  })
</script>
