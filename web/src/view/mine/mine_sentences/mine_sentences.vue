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

        <el-form-item label="标题" prop="title">
          <el-input v-model="searchInfo.title" placeholder="标题" />
        </el-form-item>

        <el-form-item label="内容" prop="content">
          <el-input v-model="searchInfo.content" placeholder="内容" />
        </el-form-item>

        <template v-if="showAllQuery">
          <el-form-item label="作者" prop="author">
            <el-input v-model="searchInfo.author" placeholder="作者" />
          </el-form-item>

          <el-form-item label="朝代" prop="dynasty">
            <el-input v-model="searchInfo.dynasty" placeholder="朝代" />
          </el-form-item>
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询
          </el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button
            link
            type="primary"
            icon="arrow-down"
            @click="showAllQuery = true"
            v-if="!showAllQuery"
            >展开
          </el-button>
          <el-button
            link
            type="primary"
            icon="arrow-up"
            @click="showAllQuery = false"
            v-else
            >收起
          </el-button>
        </el-form-item>
      </el-form>

      <!-- 统计区域 -->
      <div class="flex items-center gap-3 mb-5">
        <span class="font-medium text-gray-600 shrink-0">统计</span>
        <div class="flex flex-wrap gap-2">
          <el-tag type="info" v-for="(item, index) in catStat" :key="index">
            {{ catMap[String(item.label)] }}（{{ item.total }}）
          </el-tag>
        </div>
      </div>
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

        <el-table-column align="center" label="标题" prop="title">
          <template #default="{ row }">
            <el-link :href="row.link" target="_blank">{{ row.title }}</el-link>
          </template>
        </el-table-column>

        <el-table-column align="center" label="作者" prop="author">
          <template #default="{ row }">
            <span>{{ row.dynasty }} · {{ row.author }}</span>
          </template>
        </el-table-column>

        <el-table-column align="center" label="内容" prop="content">
          <template #default="{ row }">
            <el-text truncated>{{ row.content }}</el-text>
          </template>
        </el-table-column>

        <el-table-column align="center" label="分类" prop="categoryId">
          <template #default="{ row }">
            {{ catMap[row.categoryId] }}
          </template>
        </el-table-column>

        <el-table-column align="center" label="排序值" prop="sortValue" />

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
              @click="updateMineSentencesFunc(row)"
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
        <el-form-item label="作者" prop="author">
          <el-input
            v-model="formData.author"
            :clearable="true"
            placeholder="请输入作者"
          />
        </el-form-item>
        <el-form-item label="朝代" prop="dynasty">
          <el-input
            v-model="formData.dynasty"
            :clearable="true"
            placeholder="请输入朝代"
          />
        </el-form-item>
        <el-form-item label="链接" prop="link">
          <el-input
            v-model="formData.link"
            :clearable="true"
            placeholder="请输入链接"
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
        <el-form-item label="排序值" prop="sortValue">
          <el-input-number
            v-model.number="formData.sortValue"
            :min="0"
            :precision="0"
            placeholder="排序值"
          />
        </el-form-item>

        <el-form-item label="内容" prop="content">
          <el-input
            v-model="formData.content"
            placeholder="请输入内容"
            type="textarea"
            maxlength="1000"
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
        <el-descriptions-item label="标题">
          {{ detailForm.title }}
        </el-descriptions-item>
        <el-descriptions-item label="作者">
          {{ detailForm.author }}
        </el-descriptions-item>
        <el-descriptions-item label="朝代">
          {{ detailForm.dynasty }}
        </el-descriptions-item>
        <el-descriptions-item label="内容">
          <multiline-text :text="detailForm.content" />
        </el-descriptions-item>
        <el-descriptions-item label="链接">
          {{ detailForm.link }}
        </el-descriptions-item>
        <el-descriptions-item label="分类">
          {{ catMap[detailForm.categoryId] }}
        </el-descriptions-item>
        <el-descriptions-item label="排序值">
          {{ detailForm.sortValue }}
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
    createMineSentences,
    deleteMineSentences,
    deleteMineSentencesByIds,
    updateMineSentences,
    findMineSentences,
    getMineSentencesList,
    getMineSentencesStat
  } from '@/api/mine/mine_sentences'
  import { formatDate, getDictFunc } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, onMounted } from 'vue'
  import { useAppStore } from '@/pinia'
  import MultilineText from '@/components/MultilineText.vue'

  defineOptions({
    name: 'MineSentences'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    title: '',
    author: '',
    dynasty: '',
    content: '',
    link: '',
    categoryId: undefined,
    sortValue: undefined
  })

  // 验证规则
  const rule = reactive({
    title: [
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
    author: [
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
    dynasty: [
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
    content: [
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
    link: [
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
    categoryId: [
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
    title: undefined,
    author: undefined,
    dynasty: undefined,
    content: undefined,
    categoryId: undefined
  })
  // 重置
  const onReset = () => {
    searchInfo.value = {
      title: undefined,
      author: undefined,
      dynasty: undefined,
      content: undefined,
      categoryId: undefined
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
    const table = await getMineSentencesList({
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
      deleteMineSentencesFunc(row)
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
      const res = await deleteMineSentencesByIds({ ids })
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
  const updateMineSentencesFunc = async (row) => {
    const res = await findMineSentences({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteMineSentencesFunc = async (row) => {
    const res = await deleteMineSentences({ id: row.id })
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
      dynasty: '',
      content: '',
      link: '',
      categoryId: undefined,
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
          res = await createMineSentences(formData.value)
          break
        case 'update':
          res = await updateMineSentences(formData.value)
          break
        default:
          res = await createMineSentences(formData.value)
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
    const res = await findMineSentences({ id: row.id })
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
  // 句子分类
  const catList = ref([])
  const catMap = ref({})
  const fetchCategoryList = async () => {
    const dataList = await getDictFunc('sentence_category')

    const list = []
    const map = {}
    dataList.forEach(({ value, label }) => {
      list.push({ value, label })
      map[value] = label
    })

    catList.value = list
    catMap.value = map
  }

  // 分类统计
  const catStat = ref([])
  const fetchStat = async () => {
    const { code, data } = await getMineSentencesStat()
    if (code === 0) {
      catStat.value = data
    }
  }

  onMounted(async () => {
    await fetchCategoryList()
    await fetchStat()
  })
</script>
