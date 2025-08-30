<script setup>
  import { computed, ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { createMiserTransaction } from '@/api/miser/miser_transaction'

  const emits = defineEmits(['closed'])
  const props = defineProps({
    // {transactionType: transactionName}
    transactionTypeMap: {
      type: Object,
      required: true,
      default: () => {
      }
    },
    // {categoryId: transactionType}
    typedCategoryMap: {
      type: Object,
      required: true,
      default: () => {
      }
    },
    // {transactionType: categoryArray}
    groupedCategoryMap: {
      type: Object,
      required: true,
      default: () => {
      }
    }
  })

  const TRANSACTION_TYPE_INCOME = 1
  const TRANSACTION_TYPE_EXPENSE = 2

  const formData = ref({
    date: (() => {
      const d = new Date()
      d.setMonth(d.getMonth())
      d.setDate(10)
      return d
    })(),
    categories: {}
  })
  const income = computed(() => {
    const incomeCategories = props.groupedCategoryMap[TRANSACTION_TYPE_INCOME] || []
    return incomeCategories.reduce((sum, c) => {
      const incomeAmount = Number(formData.value.categories[c.id]) || 0
      return sum + incomeAmount
    }, 0)
  })
  const expense = computed(() => {
    const expenseCategories = props.groupedCategoryMap[TRANSACTION_TYPE_EXPENSE] || []
    return expenseCategories.reduce((sum, c) => {
      const amount = Number(formData.value.categories[c.id]) || 0
      return sum + amount
    }, 0)
  })
  const balance = computed(() => income.value - expense.value)

  const showDialog = ref(false)
  const openDialog = () => {
    formData.value = {
      date: (() => {
        const d = new Date()
        d.setMonth(d.getMonth())
        d.setDate(10)
        return d
      })(),
      categories: {}
    }

    showDialog.value = true
  }
  const closeDialog = (needUpdate = false) => {
    if (typeof needUpdate === 'boolean' && needUpdate) {
      emits('closed')
    }
    showDialog.value = false
  }
  const handleSubmit = async () => {
    const list = []
    const { date, categories } = formData.value
    Object.keys(categories).forEach((cid) => {
      list.push({
        categoryId: parseInt(cid),
        transactionType: props.typedCategoryMap[cid],
        date: date,
        amount: categories[cid]
      })
    })

    const { code } = await createMiserTransaction(list)
    if (code === 0) {
      ElMessage({
        type: 'success',
        message: '批量新增交易流水成功'
      })
      closeDialog(true)
    }
  }

  defineExpose({ openDialog })
</script>

<template>
  <el-dialog
    v-model="showDialog"
    title="批量新增交易流水"
    width="800px"
    :before-close="closeDialog"
  >
    <div class="date-picker-box">
      <el-date-picker
        v-model="formData.date"
        type="date"
        placeholder="选择日期"
        :clearable="false"
      />
    </div>

    <el-form :model="formData" inline label-width="70px">
      <!-- 按交易类型分别录入 -->
      <template v-for="(categories, transactionType) in groupedCategoryMap" :key="transactionType">
        <el-divider>
          {{ transactionTypeMap[transactionType] || '未知' }}
        </el-divider>

        <el-form-item
          v-for="c in categories"
          :key="c.id"
          :label="c.name"
        >
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
        <el-input-number :precision="2" v-model="income" disabled />
      </el-form-item>
      <el-form-item label="支出">
        <el-input-number :precision="2" v-model="expense" disabled />
      </el-form-item>
      <el-form-item label="结余">
        <el-input-number :precision="2" v-model="balance" disabled />
      </el-form-item>
    </el-form>

    <template #footer>
      <el-button @click="closeDialog">取消</el-button>
      <el-button type="primary" @click="handleSubmit" :disabled="income === 0 && expense === 0"
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