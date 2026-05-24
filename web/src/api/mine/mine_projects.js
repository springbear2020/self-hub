import service from '@/utils/request'

export const createMineProjects = (data) => {
  return service({
    url: '/mineProjects/createMineProjects',
    method: 'post',
    data
  })
}

export const deleteMineProjects = (params) => {
  return service({
    url: '/mineProjects/deleteMineProjects',
    method: 'delete',
    params
  })
}

export const deleteMineProjectsByIds = (params) => {
  return service({
    url: '/mineProjects/deleteMineProjectsByIds',
    method: 'delete',
    params
  })
}

export const updateMineProjects = (data) => {
  return service({
    url: '/mineProjects/updateMineProjects',
    method: 'put',
    data
  })
}

export const findMineProjects = (params) => {
  return service({
    url: '/mineProjects/findMineProjects',
    method: 'get',
    params
  })
}

export const getMineProjectsList = (params) => {
  return service({
    url: '/mineProjects/getMineProjectsList',
    method: 'get',
    params
  })
}
