import service from '@/utils/request'

export const createMineWebsites = (data) => {
  return service({
    url: '/mineWebsites/createMineWebsites',
    method: 'post',
    data
  })
}

export const deleteMineWebsites = (params) => {
  return service({
    url: '/mineWebsites/deleteMineWebsites',
    method: 'delete',
    params
  })
}

export const deleteMineWebsitesByIds = (params) => {
  return service({
    url: '/mineWebsites/deleteMineWebsitesByIds',
    method: 'delete',
    params
  })
}

export const updateMineWebsites = (data) => {
  return service({
    url: '/mineWebsites/updateMineWebsites',
    method: 'put',
    data
  })
}

export const findMineWebsites = (params) => {
  return service({
    url: '/mineWebsites/findMineWebsites',
    method: 'get',
    params
  })
}

export const getMineWebsitesList = (params) => {
  return service({
    url: '/mineWebsites/getMineWebsitesList',
    method: 'get',
    params
  })
}
