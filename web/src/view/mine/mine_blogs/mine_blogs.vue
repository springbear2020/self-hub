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
        <el-form-item label="标题" prop="title">
          <el-input v-model="searchInfo.title" placeholder="标题" />
        </el-form-item>

        <el-form-item label="分类" prop="categoryId">
          <el-select v-model="searchInfo.categoryId">
            <el-option
              v-for="c in catList"
              :label="c.label"
              :value="c.value"
              :key="c.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="标签" prop="tagId">
          <el-select v-model="searchInfo.tagId">
            <el-option
              v-for="c in tagList"
              :label="c.label"
              :value="c.value"
              :key="c.id"
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

        <el-table-column align="center" label="封面" prop="cover">
          <template #default="{ row }">
            <custom-pic preview pic-type="file" :pic-src="row.cover" />
          </template>
        </el-table-column>

        <el-table-column align="center" label="标题" prop="title">
          <template #default="{ row }">
            <el-link :href="row.link" target="_blank">{{ row.title }}</el-link>
          </template>
        </el-table-column>

        <el-table-column align="center" label="分类" prop="categoryId">
          <template #default="{ row }">
            {{ catMap[row.categoryId] }}
          </template>
        </el-table-column>

        <el-table-column align="center" label="标签" prop="tagId">
          <template #default="{ row }">
            {{ tagMap[row.tagId] }}
          </template>
        </el-table-column>

        <el-table-column align="center" label="发布时间" prop="postTime">
          <template #default="scope"
            >{{ formatDate(scope.row.postTime) }}
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
              @click="updateMineBlogsFunc(row)"
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
        <el-form-item label="标题" prop="title">
          <el-input
            v-model="formData.title"
            :clearable="true"
            placeholder="请输入标题"
          />
        </el-form-item>
        <el-form-item label="链接" prop="link">
          <el-input
            v-model="formData.link"
            :clearable="true"
            placeholder="请输入链接"
          />
        </el-form-item>
        <el-form-item label="封面" prop="cover">
          <el-input
            v-model="formData.cover"
            :clearable="true"
            placeholder="请输入封面"
          />
        </el-form-item>
        <el-form-item label="发布时间" prop="postTime">
          <el-date-picker
            v-model="formData.postTime"
            type="datetime"
            placeholder="请选择"
            :clearable="false"
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
        <el-form-item label="分类" prop="categoryId">
          <el-select v-model="formData.categoryId">
            <el-option
              v-for="c in catList"
              :label="c.label"
              :value="parseInt(c.value)"
              :key="c.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="标签" prop="tagId">
          <el-select v-model="formData.tagId">
            <el-option
              v-for="c in tagList"
              :label="c.label"
              :value="parseInt(c.value)"
              :key="c.id"
            />
          </el-select>
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
        <el-descriptions-item label="标题">
          {{ detailForm.title }}
        </el-descriptions-item>
        <el-descriptions-item label="链接">
          {{ detailForm.link }}
        </el-descriptions-item>
        <el-descriptions-item label="封面">
          {{ detailForm.cover }}
        </el-descriptions-item>
        <el-descriptions-item label="分类">
          {{ catMap[detailForm.categoryId] }}
        </el-descriptions-item>
        <el-descriptions-item label="标签">
          {{ tagMap[detailForm.tagId] }}
        </el-descriptions-item>
        <el-descriptions-item label="排序值">
          {{ detailForm.sortValue }}
        </el-descriptions-item>
        <el-descriptions-item label="发布时间">
          {{ formatDate(detailForm.postTime) }}
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
    createMineBlogs,
    deleteMineBlogs,
    deleteMineBlogsByIds,
    updateMineBlogs,
    findMineBlogs,
    getMineBlogsList
  } from '@/api/mine/mine_blogs'
  import { formatDate, getDictFunc } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, onMounted } from 'vue'
  import { useAppStore } from '@/pinia'
  import CustomPic from '@/components/customPic/index.vue'

  defineOptions({
    name: 'MineBlogs'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    title: '',
    link: '',
    cover: '',
    postTime: new Date(),
    sortValue: undefined,
    categoryId: undefined,
    tagId: undefined
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
    link: [
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
    sortValue: [
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
    tagId: [
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
    title: undefined,
    categoryId: undefined,
    tagId: undefined
  })
  // 重置
  const onReset = () => {
    searchInfo.value = {
      title: undefined,
      categoryId: undefined,
      tagId: undefined
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
    const table = await getMineBlogsList({
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
      deleteMineBlogsFunc(row)
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
      const res = await deleteMineBlogsByIds({ ids })
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
  const updateMineBlogsFunc = async (row) => {
    const res = await findMineBlogs({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteMineBlogsFunc = async (row) => {
    const res = await deleteMineBlogs({ id: row.id })
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
      link: '',
      cover: '',
      postTime: new Date(),
      sortValue: undefined,
      categoryId: undefined,
      tagId: undefined
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
          res = await createMineBlogs(formData.value)
          break
        case 'update':
          res = await updateMineBlogs(formData.value)
          break
        default:
          res = await createMineBlogs(formData.value)
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
    const res = await findMineBlogs({ id: row.id })
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
  // 博客分类
  const catList = ref([])
  const catMap = ref({})
  const fetchCategoryList = async () => {
    const dataList = await getDictFunc('csdn_blog_category')

    const list = []
    const map = {}
    dataList.forEach(({ value, label }) => {
      list.push({ value, label })
      map[value] = label
    })

    catList.value = list
    catMap.value = map
  }

  // 博客标签
  const tagList = ref([])
  const tagMap = ref({})
  const fetchTagList = async () => {
    const dataList = await getDictFunc('csdn_blog_tag')

    const list = []
    const map = {}
    dataList.forEach(({ value, label }) => {
      list.push({ value, label })
      map[value] = label
    })

    tagList.value = list
    tagMap.value = map
  }

  onMounted(async () => {
    await fetchCategoryList()
    await fetchTagList()
  })
</script>
