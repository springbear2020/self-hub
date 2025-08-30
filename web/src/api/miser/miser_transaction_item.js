import service from '@/utils/request'

export const createMiserTransactionItem = (data) => {
  return service({
    url: '/miserTransactionItem/createMiserTransactionItem',
    method: 'post',
    data
  })
}

export const deleteMiserTransactionItem = (params) => {
  return service({
    url: '/miserTransactionItem/deleteMiserTransactionItem',
    method: 'delete',
    params
  })
}

export const deleteMiserTransactionItemByIds = (params) => {
  return service({
    url: '/miserTransactionItem/deleteMiserTransactionItemByIds',
    method: 'delete',
    params
  })
}

export const updateMiserTransactionItem = (data) => {
  return service({
    url: '/miserTransactionItem/updateMiserTransactionItem',
    method: 'put',
    data
  })
}

export const findMiserTransactionItem = (params) => {
  return service({
    url: '/miserTransactionItem/findMiserTransactionItem',
    method: 'get',
    params
  })
}

export const getMiserTransactionItemList = (params) => {
  return service({
    url: '/miserTransactionItem/getMiserTransactionItemList',
    method: 'get',
    params
  })
}

export const listItemDistinctNames = () => {
  return service({
    url: '/miserTransactionItem/listItemDistinctNames',
    method: 'get'
  })
}