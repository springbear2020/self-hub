import service from '@/utils/request'

export const createMineArchives = (data) => {
  return service({
    url: '/mineArchives/createMineArchives',
    method: 'post',
    data
  })
}

export const deleteMineArchives = (params) => {
  return service({
    url: '/mineArchives/deleteMineArchives',
    method: 'delete',
    params
  })
}

export const deleteMineArchivesByIds = (params) => {
  return service({
    url: '/mineArchives/deleteMineArchivesByIds',
    method: 'delete',
    params
  })
}

export const updateMineArchives = (data) => {
  return service({
    url: '/mineArchives/updateMineArchives',
    method: 'put',
    data
  })
}

export const findMineArchives = (params) => {
  return service({
    url: '/mineArchives/findMineArchives',
    method: 'get',
    params
  })
}

export const getMineArchivesList = (params) => {
  return service({
    url: '/mineArchives/getMineArchivesList',
    method: 'get',
    params
  })
}
