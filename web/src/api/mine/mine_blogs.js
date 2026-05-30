import service from '@/utils/request'

export const createMineBlogs = (data) => {
  return service({
    url: '/mineBlogs/createMineBlogs',
    method: 'post',
    data
  })
}

export const deleteMineBlogs = (params) => {
  return service({
    url: '/mineBlogs/deleteMineBlogs',
    method: 'delete',
    params
  })
}

export const deleteMineBlogsByIds = (params) => {
  return service({
    url: '/mineBlogs/deleteMineBlogsByIds',
    method: 'delete',
    params
  })
}

export const updateMineBlogs = (data) => {
  return service({
    url: '/mineBlogs/updateMineBlogs',
    method: 'put',
    data
  })
}

export const findMineBlogs = (params) => {
  return service({
    url: '/mineBlogs/findMineBlogs',
    method: 'get',
    params
  })
}

export const getMineBlogsList = (params) => {
  return service({
    url: '/mineBlogs/getMineBlogsList',
    method: 'get',
    params
  })
}

export const getMineBlogsStat = () => {
  return service({
    url: '/mineBlogs/getMineBlogsStat',
    method: 'get'
  })
}
