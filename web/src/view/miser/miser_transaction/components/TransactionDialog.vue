<script setup>
  import { computed, onMounted, ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { createMiserTransaction } from '@/api/miser/miser_transaction'
  import { useMiserCategoryStore, useMiserTransactionTypeStore } from '@/pinia'
  import { miserCfgMap } from '@/constants/miser'

  defineOptions({ name: 'TransactionDialog' })
  const emits = defineEmits(['refresh'])

  // 状态管理
  const txnStore = useMiserTransactionTypeStore()
  const catStore = useMiserCategoryStore()

  // 响应式数据
  const createFormData = () => {
    return {
      date: (() => {
        const d = new Date()
        d.setMonth(d.getMonth() - 1) // 默认批量采集上一个月
        d.setDate(10)
        return d
      })(),
      categories: {}
    }
  }
  const formData = ref(createFormData())
  const income = computed(() => {
    const incomeCategories =
      catStore.txnCatsMap[miserCfgMap.income.transactionType] || []
    return incomeCategories.reduce((sum, c) => {
      const incomeAmount = Number(formData.value.categories[c.id]) || 0
      return sum + incomeAmount
    }, 0)
  })
  const expense = computed(() => {
    const expenseCategories =
      catStore.txnCatsMap[miserCfgMap.expense.transactionType] || []
    return expenseCategories.reduce((sum, c) => {
      const amount = Number(formData.value.categories[c.id]) || 0
      return sum + amount
    }, 0)
  })
  const balance = computed(() => income.value - expense.value)
  const disableSubmit = computed(() => {
    return income.value === 0 && expense.value === 0
  })

  // 对话框逻辑
  const showDialog = ref(false)
  const openDialog = () => {
    formData.value = createFormData()
    showDialog.value = true
  }
  const disabledDate = (time) => {
    return time.getTime() > Date.now()
  }
  const closeDialog = (needRefresh = false) => {
    if (typeof needRefresh === 'boolean' && needRefresh) {
      emits('refresh')
    }
    showDialog.value = false
  }
  const handleSubmit = async () => {
    const list = []
    const { date, categories } = formData.value
    Object.keys(categories).forEach((cid) => {
      list.push({
        categoryId: parseInt(cid),
        transactionType: catStore.catTxnMap[cid],
        date: date,
        amount: categories[cid]
      })
    })

    const { code } = await createMiserTransaction(list)
    if (code === 0) {
      ElMessage.success('批量新增交易流水成功')
      closeDialog(true)
    }
  }

  onMounted(async () => {
    await catStore.init()
    await txnStore.init()
  })

  defineExpose({ openDialog })
</script>

<template>
  <el-dialog
    v-model="showDialog"
    title="批量新增交易流水"
    width="800px"
    :close-on-click-modal="false"
  >
    <div class="date-picker-box">
      <el-date-picker
        v-model="formData.date"
        type="date"
        placeholder="请选择"
        :clearable="false"
        :disabled-date="disabledDate"
      />
    </div>

    <el-form :model="formData" inline label-width="70px">
      <!-- 按交易类型分别录入 -->
      <template
        v-for="(categories, transactionType) in catStore.txnCatsMap"
        :key="transactionType"
      >
        <el-divider>
          {{ txnStore.dataMap[transactionType] }}
        </el-divider>

        <el-form-item v-for="c in categories" :key="c.id" :label="c.name">
          <el-input-number
            v-model="formData.categories[c.id]"
            :min="0"
            :precision="2"
            placeholder="金额"
          />
        </el-form-item>
      </template>

      <!-- 金额合计 -->
      <el-divider>合计</el-divider>
      <el-form-item label="收入">
        <el-input-number :min="0" :precision="2" v-model="income" disabled />
      </el-form-item>
      <el-form-item label="支出">
        <el-input-number :min="0" :precision="2" v-model="expense" disabled />
      </el-form-item>
      <el-form-item label="结余">
        <el-input-number :precision="2" v-model="balance" disabled />
      </el-form-item>
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

<style scoped lang="scss">
  .date-picker-box {
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>
