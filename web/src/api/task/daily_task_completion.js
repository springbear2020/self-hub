import service from '@/utils/request'

export const createDailyTaskCompletion = (data) => {
  return service({
    url: '/dailyTaskCompletion/createDailyTaskCompletion',
    method: 'post',
    data
  })
}

export const deleteDailyTaskCompletion = (params) => {
  return service({
    url: '/dailyTaskCompletion/deleteDailyTaskCompletion',
    method: 'delete',
    params
  })
}

export const deleteDailyTaskCompletionByIds = (params) => {
  return service({
    url: '/dailyTaskCompletion/deleteDailyTaskCompletionByIds',
    method: 'delete',
    params
  })
}

export const updateDailyTaskCompletion = (data) => {
  return service({
    url: '/dailyTaskCompletion/updateDailyTaskCompletion',
    method: 'put',
    data
  })
}

export const findDailyTaskCompletion = (params) => {
  return service({
    url: '/dailyTaskCompletion/findDailyTaskCompletion',
    method: 'get',
    params
  })
}

export const getDailyTaskCompletionList = (params) => {
  return service({
    url: '/dailyTaskCompletion/getDailyTaskCompletionList',
    method: 'get',
    params
  })
}

export const getDailyTaskStatList = (params) => {
  return service({
    url: '/dailyTaskCompletion/getDailyTaskStatList',
    method: 'get',
    params
  })
}
