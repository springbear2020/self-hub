<script setup>
  import { computed, onMounted, ref, watch } from 'vue'
  import { listMiserCategoryList } from '@/api/miser/miser_category'
  import { getItemStat } from '@/api/miser/miser_stat'

  const props = defineProps({
    startMonth: { type: String, required: true },
    endMonth: { type: String, required: true }
  })

  // 分类列表
  const categoryList = ref([])
  const categoryMap = computed(() => {
    return Object.fromEntries(
      categoryList.value.map(({ id, name }) => [String(id), name])
    )
  })
  const fetchCategories = async () => {
    const { code, data } = await listMiserCategoryList()
    categoryList.value = code === 0 && Array.isArray(data) ? data : []
  }

  // 表格数据
  const dataMap = ref({})
  const fetchData = async () => {
    const params = { startMonth: props.startMonth, endMonth: props.endMonth }
    const { code, data } = await getItemStat(params)
    dataMap.value = code === 0 && data ? data : {}
  }

  onMounted(async () => {
    await fetchCategories()
    await fetchData()
  })

  watch(() => [props.startMonth, props.endMonth], fetchData)
</script>

<template>
  <div class="cards">
    <el-card
      class="table-card"
      v-for="c in Object.keys(dataMap)"
      :key="c"
      :header="`明细汇总 - ${categoryMap[c]}`"
    >
      <el-table
        :data="dataMap[c]"
        border
        stripe
        show-summary
        max-height="250px"
      >
        <el-table-column prop="name" label="名称" sortable align="center" />
        <el-table-column prop="amount" label="金额" sortable align="center" />
      </el-table>
    </el-card>
  </div>
</template>

<style scoped lang="scss">
  .cards {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 8px;
    margin-top: 8px;

    .table-card {
      border-radius: 16px;
    }
  }
</style>
