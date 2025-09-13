import { defineStore } from 'pinia'
import { ref } from 'vue'
import { listItemDistinctNames } from '@/api/miser/miser_transaction_item'

export const useMiserTransactionItemStore = defineStore(
  'miserTransactionItemStore',
  () => {
    const dataList = ref([])

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
          const { code, data } = await listItemDistinctNames()
          if (code === 0 && Array.isArray(data) && data.length) {
            dataList.value = data
          } else {
            dataList.value = []
          }
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
      loading,
      initialized,
      init,
      refresh
    }
  }
)
