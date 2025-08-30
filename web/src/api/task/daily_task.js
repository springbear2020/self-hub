import service from '@/utils/request'

export const createDailyTask = (data) => {
  return service({
    url: '/dailyTask/createDailyTask',
    method: 'post',
    data
  })
}

export const deleteDailyTask = (params) => {
  return service({
    url: '/dailyTask/deleteDailyTask',
    method: 'delete',
    params
  })
}

export const deleteDailyTaskByIds = (params) => {
  return service({
    url: '/dailyTask/deleteDailyTaskByIds',
    method: 'delete',
    params
  })
}

export const updateDailyTask = (data) => {
  return service({
    url: '/dailyTask/updateDailyTask',
    method: 'put',
    data
  })
}

export const findDailyTask = (params) => {
  return service({
    url: '/dailyTask/findDailyTask',
    method: 'get',
    params
  })
}

export const getDailyTaskList = (params) => {
  return service({
    url: '/dailyTask/getDailyTaskList',
    method: 'get',
    params
  })
}

export const listActiveTaskList = () => {
  return service({
    url: '/dailyTask/listActiveTaskList',
    method: 'get'
  })
}