import service from '@/utils/request'

export const createMiserCategory = (data) => {
  return service({
    url: '/miserCategory/createMiserCategory',
    method: 'post',
    data
  })
}

export const deleteMiserCategory = (params) => {
  return service({
    url: '/miserCategory/deleteMiserCategory',
    method: 'delete',
    params
  })
}

export const deleteMiserCategoryByIds = (params) => {
  return service({
    url: '/miserCategory/deleteMiserCategoryByIds',
    method: 'delete',
    params
  })
}

export const updateMiserCategory = (data) => {
  return service({
    url: '/miserCategory/updateMiserCategory',
    method: 'put',
    data
  })
}

export const findMiserCategory = (params) => {
  return service({
    url: '/miserCategory/findMiserCategory',
    method: 'get',
    params
  })
}

export const getMiserCategoryList = (params) => {
  return service({
    url: '/miserCategory/getMiserCategoryList',
    method: 'get',
    params
  })
}

export const listMiserCategoryList = () => {
  return service({
    url: '/miserCategory/listMiserCategoryList',
    method: 'get'
  })
}
