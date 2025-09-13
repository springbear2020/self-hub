import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getDictFunc } from '@/utils/format'

export const useMiserTransactionTypeStore = defineStore(
  'miserTransactionTypeStore',
  () => {
    const dataList = ref([])
    const dataMap = ref({})

    const loading = ref(false)
    const initialized = ref(false)

    let fetchPromise = null

    const fetchData = async () => {
      if (fetchPromise) {
        return fetchPromise
      }

      loading.value = true
      fetchPromise = (async () => {
        try {
          const results = await getDictFunc('miser_transaction_type')

          const list = []
          const map = {}

          results.forEach(({ value, label }) => {
            const intVal = parseInt(value)
            list.push({ label, value: intVal })
            map[intVal] = label
          })

          dataList.value = list
          dataMap.value = map
          initialized.value = true
        } finally {
          loading.value = false
          fetchPromise = null
        }
      })()

      return fetchPromise
    }

    const init = async () => {
      if (!initialized.value) {
        await fetchData()
      }
    }

    const refresh = fetchData

    return {
      dataList,
      dataMap,
      loading,
      initialized,
      init,
      refresh
    }
  }
)
