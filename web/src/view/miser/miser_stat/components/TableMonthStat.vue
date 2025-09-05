<script setup>
  import { ref, watch, onMounted } from 'vue'
  import { getMonthStat } from '@/api/miser/miser_stat'
  import { formatterAmount } from '@/utils/format'

  const props = defineProps({
    startMonth: { type: String, required: true },
    endMonth: { type: String, required: true }
  })

  const dataList = ref([])
  const fetchData = async () => {
    const params = {
      startMonth: props.startMonth,
      endMonth: props.endMonth
    }
    const { code, data } = await getMonthStat(params)
    dataList.value = code === 0 ? data : []
  }

  onMounted(() => {
    fetchData()
  })

  watch(
    () => [props.startMonth, props.endMonth],
    () => fetchData()
  )
</script>

<template>
  <div>
    <el-card
      :header="item.header"
      v-for="item in dataList"
      :key="item.header"
      class="table-card"
    >
      <el-table :data="item.rows" border stripe show-summary max-height="250px">
        <el-table-column
          v-for="c in item.cols"
          :key="c.key"
          :prop="c.key"
          :label="c.label"
          :fixed="c.fixed"
          :sortable="c.sortable"
          :formatter="formatterAmount"
          align="center"
        />
      </el-table>
    </el-card>
  </div>
</template>

<style scoped>
  .table-card {
    margin-top: 8px;
    border-radius: 16px;
  }
</style>
