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
        <el-form-item label="交易分类" prop="name">
          <el-input v-model="searchInfo.name" placeholder="交易分类" />
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
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="ID" prop="id" width="110" />

        <el-table-column align="left" label="交易类型" prop="transactionType">
          <template #default="{ row }">
            {{ transactionTypeMap[row.transactionType] }}
          </template>
        </el-table-column>

        <el-table-column align="left" label="交易分类" prop="name" />

        <el-table-column align="left" label="分类描述" prop="description" />

        <el-table-column align="left" label="排序值" prop="sort" />

        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="200"
        >
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
              @click="updateMiserCategoryFunc(scope.row)"
            >
              编辑
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

    <!-- 新增/编辑 -->
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
        <el-form-item label="交易类型" prop="transactionType">
          <el-select v-model="formData.transactionType">
            <el-option
              v-for="t in transactionTypeList"
              :label="t.label"
              :value="t.value"
              :key="t.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="交易分类" prop="name">
          <el-input
            v-model="formData.name"
            :clearable="true"
            placeholder="请输入交易分类"
            @blur="handleBlur"
          />
        </el-form-item>
        <el-form-item label="分类描述" prop="description">
          <el-input
            v-model="formData.description"
            :clearable="true"
            placeholder="请输入分类描述"
          />
        </el-form-item>
        <el-form-item label="排序值" prop="sort">
          <el-input-number v-model="formData.sort" placeholder="请输入排序值" />
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
        <el-descriptions-item label="交易类型">
          {{ transactionTypeMap[detailForm.transactionType] }}
        </el-descriptions-item>
        <el-descriptions-item label="交易分类">
          {{ detailForm.name }}
        </el-descriptions-item>
        <el-descriptions-item label="分类描述">
          {{ detailForm.description }}
        </el-descriptions-item>
        <el-descriptions-item label="排序值">
          {{ detailForm.sort }}
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
    createMiserCategory,
    deleteMiserCategory,
    deleteMiserCategoryByIds,
    findMiserCategory,
    getMiserCategoryList,
    updateMiserCategory
  } from '@/api/miser/miser_category'

  // 全量引入格式化工具 请按需保留
  import { formatDate, getDictFunc } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { computed, onMounted, reactive, ref } from 'vue'
  import { useAppStore } from '@/pinia'
  import { InfoFilled } from '@element-plus/icons-vue'

  defineOptions({
    name: 'MiserCategory'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    transactionType: undefined,
    name: '',
    description: '',
    sort: 0
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
    sort: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ]
  })

  const handleBlur = () => {
    if (type.value === 'create') {
      formData.value.description = formData.value.name
    }
  }

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
    const table = await getMiserCategoryList({
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
      deleteMiserCategoryFunc(row)
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
      const res = await deleteMiserCategoryByIds({ ids })
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
  const updateMiserCategoryFunc = async (row) => {
    const res = await findMiserCategory({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteMiserCategoryFunc = async (row) => {
    const res = await deleteMiserCategory({ id: row.id })
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
    formData.value.sort = total.value + 1
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
      transactionType: undefined,
      name: '',
      description: '',
      sort: total.value + 1
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
          res = await createMiserCategory(formData.value)
          break
        case 'update':
          res = await updateMiserCategory(formData.value)
          break
        default:
          res = await createMiserCategory(formData.value)
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
    const res = await findMiserCategory({ id: row.id })
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

  onMounted(() => {
    fetchTransactionTypeList()
  })
</script>
