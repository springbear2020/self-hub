import { defineStore } from 'pinia'
import { ref } from 'vue'
import { listMiserCategoryList } from '@/api/miser/miser_category'

export const useMiserCategoryStore = defineStore('miserCategoryStore', () => {
  const dataList = ref([])
  const dataMap = ref({}) // categoryId: categoryName
  const catTxnMap = ref({}) // categoryId: transactionType
  const txnCatsMap = ref({}) // transactionType: []categoryId

  const loading = ref(false)
  const initialized = ref(false)

  let fetchPromise = null

  const resetData = () => {
    dataList.value = []
    dataMap.value = {}
    catTxnMap.value = {}
    txnCatsMap.value = {}
  }

  const buildMaps = (data) => {
    const dMap = {}
    const cMap = {}
    const tMap = {}

    data.forEach(({ id, name, sort, transactionType }) => {
      dMap[id] = name
      cMap[id] = transactionType
      if (!tMap[transactionType]) {
        tMap[transactionType] = []
      }
      tMap[transactionType].push({ id, name, sort, transactionType })
    })

    dataMap.value = dMap
    catTxnMap.value = cMap
    txnCatsMap.value = tMap
  }

  const fetchData = async () => {
    if (fetchPromise) {
      return fetchPromise
    }

    loading.value = true
    fetchPromise = (async () => {
      try {
        const { code, data } = await listMiserCategoryList()
        if (code === 0 && Array.isArray(data) && data.length) {
          dataList.value = data
          buildMaps(data)
        } else {
          resetData()
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
    dataMap,
    catTxnMap,
    txnCatsMap,
    loading,
    initialized,
    init,
    refresh
  }
})
