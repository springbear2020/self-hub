<script setup>
  import { computed, ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import {
    createMiserTransactionItem,
    listItemDistinctNames
  } from '@/api/miser/miser_transaction_item'

  const props = defineProps({
    // {categoryId: categoryName}
    categoryMap: {
      type: Object,
      required: true,
      default: () => {}
    }
  })

  const formData = ref({
    transactionId: null,
    categoryId: null,
    categoryAmount: null,
    items: []
  })
  const itemsAmountSum = computed(() => {
    return formData.value.items.reduce((sum, { amount }) => {
      return sum + Number(amount || 0)
    }, 0)
  })

  const showDialog = ref(false)
  const dialogTitle = ref('批量新增流水明细')
  const openDialog = async (rowData) => {
    const { id, categoryId, amount, date } = rowData

    // 表单数据赋值
    formData.value = {
      transactionId: id,
      categoryId: categoryId,
      categoryAmount: amount,
      items: [
        {
          name: null,
          amount: null
        }
      ]
    }

    // 设置对话框标题
    const dateStr = date.substring(0, 10)
    const categoryName = props.categoryMap[categoryId]
    dialogTitle.value = `批量新增『${dateStr}/${categoryName}/￥${amount}』流水明细`

    // 查询交易分类流水名称去重列表
    await fetchItemNameList(categoryId)

    showDialog.value = true
  }
  const closeDialog = () => {
    showDialog.value = false
  }

  const handleSubmit = async () => {
    const { transactionId, categoryId, categoryAmount } = formData.value

    // 各项之和校验
    const amountSum = Math.round(itemsAmountSum.value * 100) / 100
    if (amountSum !== categoryAmount) {
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
      closeDialog()
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
  const handleCopy = (index) => {
    formData.value.items.push({
      name: '',
      amount: formData.value.items[index].amount
    })
  }

  // 去重明细名称列表
  const distinctItemNames = ref([])
  const fetchItemNameList = async (categoryId) => {
    const { code, data } = await listItemDistinctNames({ categoryId })
    if (code === 0 && data && data.length > 0) {
      distinctItemNames.value = data
    } else {
      distinctItemNames.value = []
    }
  }

  defineExpose({ openDialog })
</script>

<template>
  <el-dialog
    v-model="showDialog"
    :title="dialogTitle"
    width="800px"
    :before-close="closeDialog"
    :close-on-click-modal="false"
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
              type="default"
              icon="copyDocument"
              @click="handleCopy($index)"
            />
            <el-button
              type="danger"
              icon="delete"
              @click="handleRemove($index)"
            />
          </template>
        </el-table-column>
      </el-table>

      <div class="mt-3 footer-operation-box">
        <el-button type="info" icon="circle-plus" @click="handleAdd" />
        <el-input-number
          v-model="itemsAmountSum"
          :min="0"
          :precision="2"
          disabled
        />
      </div>
    </el-form>

    <template #footer>
      <el-button @click="closeDialog">取消</el-button>
      <el-button type="primary" @click="handleSubmit">保存</el-button>
    </template>
  </el-dialog>
</template>

<style lang="scss" scoped>
  .footer-operation-box {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
</style>
