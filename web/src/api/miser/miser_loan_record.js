import service from '@/utils/request'

export const createMiserLoanRecord = (data) => {
  return service({
    url: '/miserLoanRecord/createMiserLoanRecord',
    method: 'post',
    data
  })
}

export const deleteMiserLoanRecord = (params) => {
  return service({
    url: '/miserLoanRecord/deleteMiserLoanRecord',
    method: 'delete',
    params
  })
}

export const deleteMiserLoanRecordByIds = (params) => {
  return service({
    url: '/miserLoanRecord/deleteMiserLoanRecordByIds',
    method: 'delete',
    params
  })
}

export const updateMiserLoanRecord = (data) => {
  return service({
    url: '/miserLoanRecord/updateMiserLoanRecord',
    method: 'put',
    data
  })
}

export const findMiserLoanRecord = (params) => {
  return service({
    url: '/miserLoanRecord/findMiserLoanRecord',
    method: 'get',
    params
  })
}

export const getMiserLoanRecordList = (params) => {
  return service({
    url: '/miserLoanRecord/getMiserLoanRecordList',
    method: 'get',
    params
  })
}

export const listMiserLoanNameList = () => {
  return service({
    url: '/miserLoanRecord/listMiserLoanNameList',
    method: 'get'
  })
}

export const getMiserLoanStatData = () => {
  return service({
    url: '/miserLoanRecord/getMiserLoanStatData',
    method: 'get'
  })
}
