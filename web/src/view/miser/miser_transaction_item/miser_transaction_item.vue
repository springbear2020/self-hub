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
        <el-form-item label="交易 ID" prop="transactionId">
          <el-input-number
            v-model="searchInfo.transactionId"
            placeholder="交易 ID"
          />
        </el-form-item>

        <el-form-item label="交易分类" prop="categoryId">
          <el-select v-model="searchInfo.categoryId" filterable>
            <el-option
              v-for="c in categoryList"
              :key="c.id"
              :label="c.name"
              :value="c.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="明细名称" prop="name">
          <el-select v-model="searchInfo.name" filterable>
            <el-option
              v-for="name in distinctItemNames"
              :key="name"
              :label="name"
              :value="name"
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
        <el-button type="primary" icon="plus" @click="openDialog()"
          >新增
        </el-button>
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
        <el-table-column align="center" type="selection" width="55" />

        <el-table-column align="center" label="ID" prop="id" width="55" />

        <el-table-column align="center" label="交易 ID" prop="transactionId" />

        <el-table-column align="center" label="交易分类" prop="categoryId">
          <template #default="{ row }">
            {{ categoryMap[row.categoryId] }}
          </template>
        </el-table-column>

        <el-table-column align="center" label="明细名称" prop="name" />

        <el-table-column
          align="center"
          label="交易金额"
          prop="amount"
          :formatter="formatterAmount"
        />

        <el-table-column align="center" label="交易日期" prop="date">
          <template #default="{ row }">
            {{ formatDate(row.date, 'yyyy-MM-dd') }}
          </template>
        </el-table-column>

        <el-table-column align="center" label="操作" fixed="right" width="210">
          <template #default="scope">
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
              @click="updateMiserTransactionItemFunc(scope.row)"
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
        <el-form-item label="交易 ID" prop="transactionId">
          <el-input-number
            v-model="formData.transactionId"
            :clearable="true"
            @blur="handleBlur"
            placeholder="交易 ID"
          />
        </el-form-item>
        <el-form-item label="交易分类" prop="categoryId">
          <el-select v-model="formData.categoryId" disabled>
            <el-option
              v-for="c in categoryList"
              :key="c.id"
              :label="c.name"
              :value="c.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="交易日期" prop="date">
          <el-input v-model="formData.date" placeholder="交易日期" disabled />
        </el-form-item>
        <el-form-item label="明细名称" prop="name">
          <el-select
            v-model="formData.name"
            allow-create
            filterable
            default-first-option
            placeholder="请输入明细名称"
          >
            <el-option
              v-for="name in distinctItemNames"
              :key="name"
              :label="name"
              :value="name"
            />
          </el-select>
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
        <el-descriptions-item label="用户 ID">
          {{ detailForm.userId }}
        </el-descriptions-item>
        <el-descriptions-item label="交易 ID">
          {{ detailForm.transactionId }}
        </el-descriptions-item>
        <el-descriptions-item label="交易分类">
          {{ categoryMap[detailForm.categoryId] }}
        </el-descriptions-item>
        <el-descriptions-item label="明细名称">
          {{ detailForm.name }}
        </el-descriptions-item>
        <el-descriptions-item label="交易金额">
          {{ formatAmount(detailForm.amount) }}
        </el-descriptions-item>
        <el-descriptions-item label="交易时间">
          {{ formatDate(detailForm.date, 'yyyy-MM-dd') }}
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
  import { formatAmount, formatDate, formatterAmount } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, computed, onMounted, watch } from 'vue'
  import { useAppStore } from '@/pinia'
  import { InfoFilled } from '@element-plus/icons-vue'
  import { findMiserTransaction } from '@/api/miser/miser_transaction'
  import { listMiserCategoryList } from '@/api/miser/miser_category'
  import {
    createMiserTransactionItem,
    deleteMiserTransactionItem,
    deleteMiserTransactionItemByIds,
    updateMiserTransactionItem,
    findMiserTransactionItem,
    getMiserTransactionItemList,
    listItemDistinctNames
  } from '@/api/miser/miser_transaction_item'

  defineOptions({
    name: 'MiserTransactionItem'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    transactionId: undefined,
    categoryId: undefined,
    name: '',
    amount: undefined,
    date: undefined
  })

  const handleBlur = async () => {
    const { transactionId } = formData.value
    if (transactionId) {
      const { categoryId, date } = await fetchTransactionInfo(transactionId)
      formData.value.categoryId = categoryId
      formData.value.date = date
    }
  }

  // 验证规则
  const rule = reactive({
    transactionId: [
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
  const searchInfo = ref({})
  const disabledDate = (time) => {
    return time.getTime() > Date.now()
  }
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
    const table = await getMiserTransactionItemList({
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
    const tip = `确定要删除『${row.name}/${formatDate(row.date, 'yyyy-MM-dd')}』吗？`
    ElMessageBox.confirm(tip, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteMiserTransactionItemFunc(row)
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
      const res = await deleteMiserTransactionItemByIds({ ids })
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
  const updateMiserTransactionItemFunc = async (row) => {
    const res = await findMiserTransactionItem({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteMiserTransactionItemFunc = async (row) => {
    const res = await deleteMiserTransactionItem({ id: row.id })
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
      transactionId: undefined,
      categoryId: undefined,
      name: '',
      amount: undefined,
      date: undefined
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
          res = await createMiserTransactionItem([formData.value])
          break
        case 'update':
          res = await updateMiserTransactionItem(formData.value)
          break
        default:
          res = await createMiserTransactionItem(formData.value)
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
  const fetchTransactionInfo = async (id) => {
    const { code, data } = await findMiserTransaction({ id: id })
    if (code === 0) {
      return data
    }
    return {}
  }
  const getDetails = async (row) => {
    // 打开弹窗
    const res = await findMiserTransactionItem({ id: row.id })
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

  // 分类列表
  const categoryList = ref([])
  const categoryMap = computed(() => {
    const resultMap = {}
    categoryList.value.forEach(({ id, name }) => {
      resultMap[id] = name
    })
    return resultMap
  })
  const fetchCategoryList = async () => {
    const { code, data } = await listMiserCategoryList()
    if (code === 0 && data && data.length > 0) {
      categoryList.value = data
    } else {
      categoryList.value = []
    }
  }

  // 明细名称列表
  const distinctItemNames = ref([])
  const fetchItemNameList = async () => {
    const { code, data } = await listItemDistinctNames()
    if (code === 0 && data && data.length > 0) {
      distinctItemNames.value = data
    } else {
      distinctItemNames.value = []
    }
  }

  watch(
    total,
    async () => {
      await fetchItemNameList()
    },
    { immediate: true }
  )

  onMounted(() => {
    fetchCategoryList()
  })
</script>
