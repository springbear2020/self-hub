import service from '@/utils/request'

export const createMineSentences = (data) => {
  return service({
    url: '/mineSentences/createMineSentences',
    method: 'post',
    data
  })
}

export const deleteMineSentences = (params) => {
  return service({
    url: '/mineSentences/deleteMineSentences',
    method: 'delete',
    params
  })
}

export const deleteMineSentencesByIds = (params) => {
  return service({
    url: '/mineSentences/deleteMineSentencesByIds',
    method: 'delete',
    params
  })
}

export const updateMineSentences = (data) => {
  return service({
    url: '/mineSentences/updateMineSentences',
    method: 'put',
    data
  })
}

export const findMineSentences = (params) => {
  return service({
    url: '/mineSentences/findMineSentences',
    method: 'get',
    params
  })
}

export const getMineSentencesList = (params) => {
  return service({
    url: '/mineSentences/getMineSentencesList',
    method: 'get',
    params
  })
}

export const getMineSentencesStat = () => {
  return service({
    url: '/mineSentences/getMineSentencesStat',
    method: 'get'
  })
}
