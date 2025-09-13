<script setup>
  import { computed, onMounted, ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import {
    createMiserTransactionItem,
    listItemDistinctNames
  } from '@/api/miser/miser_transaction_item'
  import { useMiserCategoryStore } from '@/pinia'
  import { formatAmountCurrency } from '@/utils/format'

  defineOptions({ name: 'TransactionItemsDialog' })

  // 状态管理
  const catStore = useMiserCategoryStore()

  // 响应式数据
  const formData = ref({})
  const itemsAmountSum = computed(() => {
    return formData.value.items.reduce((sum, { amount }) => {
      return sum + Number(amount || 0)
    }, 0)
  })
  const disableSubmit = computed(() => {
    return itemsAmountSum.value <= 0
  })

  // 对话框逻辑
  const showDialog = ref(false)
  const dialogTitle = ref('批量新增流水明细')
  const openDialog = async (rowData) => {
    const { id, categoryId, amount, date } = rowData

    // 表单数据赋值
    formData.value = {
      transactionId: id,
      categoryId: categoryId,
      categoryAmount: amount,
      categoryDate: date,
      items: [
        {
          name: null,
          amount: null
        }
      ]
    }

    // 设置对话框标题
    const dateStr = date.substring(0, 10)
    const categoryName = catStore.dataMap[categoryId]
    dialogTitle.value = `批量新增『${dateStr}/${categoryName}/${formatAmountCurrency(amount)}』流水明细`

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
      ElMessage.error(
        `明细金额之和 ${amountSum} 与流水金额 ${categoryAmount} 不等`
      )
      return
    }

    const list = formData.value.items.map(({ name, amount }) => {
      return {
        transactionId,
        categoryId,
        name,
        amount,
        date: formData.value.categoryDate
      }
    })

    const { code } = await createMiserTransactionItem(list)
    if (code === 0) {
      ElMessage.success('批量新增流水明细成功')
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
    if (code === 0 && Array.isArray(data)) {
      distinctItemNames.value = data
    } else {
      distinctItemNames.value = []
    }
  }

  onMounted(async () => {
    await catStore.init()
  })

  defineExpose({ openDialog })
</script>

<template>
  <el-dialog
    v-model="showDialog"
    :title="dialogTitle"
    width="800px"
    :close-on-click-modal="false"
  >
    <el-form :model="formData" label-width="70px">
      <el-table :data="formData.items" border stripe>
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
      <el-button @click="closeDialog" icon="CloseBold">取消</el-button>
      <el-button
        type="primary"
        icon="Select"
        @click="handleSubmit"
        :disabled="disableSubmit"
        >保存
      </el-button>
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
