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
        <el-form-item label="书名" prop="title">
          <el-input v-model="searchInfo.title" placeholder="书名" />
        </el-form-item>

        <el-form-item label="作者" prop="author">
          <el-input v-model="searchInfo.author" placeholder="作者" />
        </el-form-item>

        <el-form-item label="名句" prop="iconicQuote">
          <el-input v-model="searchInfo.iconicQuote" placeholder="名句" />
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

        <el-table-column align="center" label="封面" prop="cover">
          <template #default="{ row }">
            <custom-pic preview pic-type="file" :pic-src="row.cover" />
          </template>
        </el-table-column>

        <el-table-column align="center" label="书名" prop="title">
          <template #default="{ row }">
            <el-link :href="row.link" target="_blank">{{ row.title }}</el-link>
          </template>
        </el-table-column>

        <el-table-column align="center" label="作者" prop="author" />

        <el-table-column align="center" label="排序值" prop="sortValue" />

        <el-table-column align="center" label="首次读完" prop="firstCompletion">
          <template #default="{ row }"
            >{{ formatDate(row.firstCompletion, 'yyyy-MM-dd') }}
          </template>
        </el-table-column>

        <el-table-column align="center" label="阅读时长" prop="readMinute">
          <template #default="{ row }">
            {{ formatReadMinute(row.readMinute) }}
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
              @click="updateMineBooksFunc(row)"
            >
              编辑
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
        <el-form-item label="书名" prop="title">
          <el-input
            v-model="formData.title"
            :clearable="true"
            placeholder="请输入书名"
          />
        </el-form-item>
        <el-form-item label="作者" prop="author">
          <el-input
            v-model="formData.author"
            :clearable="true"
            placeholder="请输入作者"
          />
        </el-form-item>
        <el-form-item label="封面" prop="cover">
          <el-input
            v-model="formData.cover"
            :clearable="true"
            placeholder="请输入封面"
          />
        </el-form-item>
        <el-form-item label="链接" prop="link">
          <el-input
            v-model="formData.link"
            :clearable="true"
            placeholder="请输入链接"
          />
        </el-form-item>
        <el-form-item label="首次读完" prop="firstCompletion">
          <el-date-picker
            v-model="formData.firstCompletion"
            type="date"
            placeholder="请选择"
            :clearable="false"
          />
        </el-form-item>
        <el-form-item label="阅读时长" prop="readMinute">
          <el-input-number
            v-model.number="formData.readMinute"
            :min="0"
            :precision="0"
            placeholder="单位(分钟)"
          />
        </el-form-item>
        <el-form-item label="排序值" prop="sortValue">
          <el-input-number
            v-model.number="formData.sortValue"
            :min="0"
            :precision="0"
            placeholder="排序值"
          />
        </el-form-item>
        <el-form-item label="名句" prop="iconicQuote">
          <el-input
            v-model="formData.iconicQuote"
            placeholder="请输入名句"
            type="textarea"
            maxlength="250"
            show-word-limit
            autosize
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
        <el-descriptions-item label="书名">
          {{ detailForm.title }}
        </el-descriptions-item>
        <el-descriptions-item label="作者">
          {{ detailForm.author }}
        </el-descriptions-item>
        <el-descriptions-item label="封面">
          {{ detailForm.cover }}
        </el-descriptions-item>
        <el-descriptions-item label="链接">
          {{ detailForm.link }}
        </el-descriptions-item>
        <el-descriptions-item label="名句">
          <multiline-text :text="detailForm.iconicQuote" />
        </el-descriptions-item>
        <el-descriptions-item label="排序值">
          {{ detailForm.sortValue }}
        </el-descriptions-item>
        <el-descriptions-item label="阅读时长">
          {{ formatReadMinute(detailForm.readMinute) }}
        </el-descriptions-item>
        <el-descriptions-item label="首次读完">
          {{ formatDate(detailForm.firstCompletion, 'yyyy-MM-dd') }}
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
    createMineBooks,
    deleteMineBooks,
    deleteMineBooksByIds,
    updateMineBooks,
    findMineBooks,
    getMineBooksList
  } from '@/api/mine/mine_books'
  import { formatDate } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive } from 'vue'
  import { useAppStore } from '@/pinia'
  import CustomPic from '@/components/customPic/index.vue'
  import MultilineText from '@/components/MultilineText.vue'

  defineOptions({
    name: 'MineBooks'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    title: '',
    author: '',
    cover: '',
    link: '',
    firstCompletion: new Date(),
    readMinute: undefined,
    iconicQuote: '',
    sortValue: undefined
  })

  // 验证规则
  const rule = reactive({
    title: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    author: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    cover: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    link: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    readMinute: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    iconicQuote: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      }
    ],
    sortValue: [
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
    author: undefined,
    title: undefined,
    iconicQuote: undefined
  })
  // 重置
  const onReset = () => {
    searchInfo.value = {
      author: undefined,
      title: undefined,
      iconicQuote: undefined
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
    const table = await getMineBooksList({
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
    const tip = `确定要删除『${row.title}』吗？`
    ElMessageBox.confirm(tip, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteMineBooksFunc(row)
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
      const res = await deleteMineBooksByIds({ ids })
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
  const updateMineBooksFunc = async (row) => {
    const res = await findMineBooks({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteMineBooksFunc = async (row) => {
    const res = await deleteMineBooks({ id: row.id })
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
      title: '',
      author: '',
      cover: '',
      link: '',
      firstCompletion: new Date(),
      readMinute: undefined,
      iconicQuote: '',
      sortValue: undefined
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
          res = await createMineBooks(formData.value)
          break
        case 'update':
          res = await updateMineBooks(formData.value)
          break
        default:
          res = await createMineBooks(formData.value)
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
    const res = await findMineBooks({ id: row.id })
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
  const formatReadMinute = (minute) => {
    const hours = Math.floor(minute / 60)
    if (hours < 1) return `${minute} 分钟`

    const minutes = minute % 60
    return minutes === 0 ? `${hours} 小时` : `${hours} 小时 ${minutes} 分钟`
  }
</script>
