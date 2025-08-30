import service from '@/utils/request'

export const createMiserTransaction = (data) => {
  return service({
    url: '/miserTransaction/createMiserTransaction',
    method: 'post',
    data
  })
}

export const deleteMiserTransaction = (params) => {
  return service({
    url: '/miserTransaction/deleteMiserTransaction',
    method: 'delete',
    params
  })
}

export const deleteMiserTransactionByIds = (params) => {
  return service({
    url: '/miserTransaction/deleteMiserTransactionByIds',
    method: 'delete',
    params
  })
}

export const updateMiserTransaction = (data) => {
  return service({
    url: '/miserTransaction/updateMiserTransaction',
    method: 'put',
    data
  })
}

export const findMiserTransaction = (params) => {
  return service({
    url: '/miserTransaction/findMiserTransaction',
    method: 'get',
    params
  })
}

export const getMiserTransactionList = (params) => {
  return service({
    url: '/miserTransaction/getMiserTransactionList',
    method: 'get',
    params
  })
}
