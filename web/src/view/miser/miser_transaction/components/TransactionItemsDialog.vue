<script setup>
  import { onMounted, ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { createMiserTransactionItem, listItemDistinctNames } from '@/api/miser/miser_transaction_item'

  const props = defineProps({
    // {categoryId: categoryName}
    categoryMap: {
      type: Object,
      required: true,
      default: () => {
      }
    }
  })

  const formData = ref({
    transactionId: null,
    categoryId: null,
    categoryAmount: null,
    items: []
  })

  const showDialog = ref(false)
  const dialogTitle = ref('批量新增流水明细')
  const openDialog = (rowData) => {
    const { id, categoryId, amount, date } = rowData

    // 表单数据赋值
    formData.value = {
      transactionId: id,
      categoryId: categoryId,
      categoryAmount: amount,
      items: [{
        name: null,
        amount: null
      }]
    }

    // 设置对话框标题
    const dateStr = date.substring(0, 10)
    const categoryName = props.categoryMap[categoryId]
    dialogTitle.value = `批量新增『${dateStr}/${categoryName}/￥${amount}』流水明细`

    showDialog.value = true
  }
  const closeDialog = (needUpdate = false) => {
    if (typeof needUpdate === 'boolean' && needUpdate === true) {
      fetchItemNameList()
    }
    showDialog.value = false
  }

  const handleSubmit = async () => {
    const { transactionId, categoryId, categoryAmount } = formData.value

    // 各项之和校验
    const amountSum = formData.value.items.reduce((sum, { amount }) => {
      return sum + Number(amount || 0)
    }, 0)
    if (String(amountSum) !== String(categoryAmount)) {
      ElMessage({
        type: 'error',
        message: `明细金额之和 ${amountSum} 与流水金额 ${categoryAmount} 不等`
      })
      return
    }

    const list = []
    formData.value.items.forEach(({ name, amount }) => {
      list.push({
        transactionId,
        categoryId,
        name,
        amount
      })
    })

    const { code } = await createMiserTransactionItem(list)
    if (code === 0) {
      ElMessage({
        type: 'success',
        message: '批量新增流水明细成功'
      })
      closeDialog(true)
    }
  }
  const handleAdd = () => {
    formData.value.items.push({
      name: null,
      amount: null
    })
  }
  const handleRemove = (index) => {
    formData.value.items.splice(index, 1)
  }

  // 去重明细名称列表
  const distinctItemNames = ref([])
  const fetchItemNameList = async () => {
    const { code, data } = await listItemDistinctNames()
    if (code === 0 && data && data.length > 0) {
      distinctItemNames.value = data
    } else {
      distinctItemNames.value = []
    }
  }

  onMounted(() => {
    fetchItemNameList()
  })

  defineExpose({ openDialog })
</script>

<template>
  <el-dialog
    v-model="showDialog"
    :title="dialogTitle"
    width="800px"
    :before-close="closeDialog"
  >
    <el-form :model="formData" label-width="70px">
      <el-table :data="formData.items" border>
        <el-table-column label="明细名称" align="center">
          <template #default="{ row }">
            <el-select
              v-model="row.name"
              allow-create
              filterable
              default-first-option
            >
              <el-option
                v-for="name in distinctItemNames"
                :key="name"
                :label="name"
                :value="name"
              />
            </el-select>
          </template>
        </el-table-column>

        <el-table-column label="交易金额" align="center">
          <template #default="{ row }">
            <el-input-number
              v-model="row.amount"
              :min="0"
              :precision="2"
              placeholder="金额"
            />
          </template>
        </el-table-column>

        <el-table-column label="操作" align="center">
          <template #default="{ $index }">
            <el-button
              type="danger"
              icon="delete"
              @click="handleRemove($index)"
            />
          </template>
        </el-table-column>
      </el-table>

      <div class="mt-3">
        <el-button type="primary" icon="circle-plus" @click="handleAdd" />
      </div>
    </el-form>

    <template #footer>
      <el-button @click="closeDialog">取消</el-button>
      <el-button type="primary" @click="handleSubmit">保存</el-button>
    </template>
  </el-dialog>
</template>
