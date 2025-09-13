import { createPinia } from 'pinia'
import { useAppStore } from '@/pinia/modules/app'
import { useUserStore } from '@/pinia/modules/user'
import { useDictionaryStore } from '@/pinia/modules/dictionary'
import { useMiserCategoryStore } from '@/pinia/modules/miser_category'
import { useMiserTransactionTypeStore } from '@/pinia/modules/miser_transaction_type'
import { useMiserTransactionItemStore } from '@/pinia/modules/miser_transaction_item'
import { useMiserStatStore } from '@/pinia/modules/miser_stat'

const store = createPinia()

export {
  store,
  useAppStore,
  useUserStore,
  useDictionaryStore,
  useMiserCategoryStore,
  useMiserTransactionTypeStore,
  useMiserTransactionItemStore,
  useMiserStatStore
}
