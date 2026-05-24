import service from '@/utils/request'

export const createMineBooks = (data) => {
  return service({
    url: '/mineBooks/createMineBooks',
    method: 'post',
    data
  })
}

export const deleteMineBooks = (params) => {
  return service({
    url: '/mineBooks/deleteMineBooks',
    method: 'delete',
    params
  })
}

export const deleteMineBooksByIds = (params) => {
  return service({
    url: '/mineBooks/deleteMineBooksByIds',
    method: 'delete',
    params
  })
}

export const updateMineBooks = (data) => {
  return service({
    url: '/mineBooks/updateMineBooks',
    method: 'put',
    data
  })
}

export const findMineBooks = (params) => {
  return service({
    url: '/mineBooks/findMineBooks',
    method: 'get',
    params
  })
}

export const getMineBooksList = (params) => {
  return service({
    url: '/mineBooks/getMineBooksList',
    method: 'get',
    params
  })
}
