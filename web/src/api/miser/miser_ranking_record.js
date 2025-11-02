import service from '@/utils/request'

export const createMiserRankingRecord = (data) => {
  return service({
    url: '/miserRankingRecord/createMiserRankingRecord',
    method: 'post',
    data
  })
}

export const deleteMiserRankingRecord = (params) => {
  return service({
    url: '/miserRankingRecord/deleteMiserRankingRecord',
    method: 'delete',
    params
  })
}

export const deleteMiserRankingRecordByIds = (params) => {
  return service({
    url: '/miserRankingRecord/deleteMiserRankingRecordByIds',
    method: 'delete',
    params
  })
}

export const updateMiserRankingRecord = (data) => {
  return service({
    url: '/miserRankingRecord/updateMiserRankingRecord',
    method: 'put',
    data
  })
}

export const findMiserRankingRecord = (params) => {
  return service({
    url: '/miserRankingRecord/findMiserRankingRecord',
    method: 'get',
    params
  })
}

export const getMiserRankingRecordList = (params) => {
  return service({
    url: '/miserRankingRecord/getMiserRankingRecordList',
    method: 'get',
    params
  })
}
