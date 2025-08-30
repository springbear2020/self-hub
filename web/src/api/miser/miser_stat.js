import service from '@/utils/request'

export const getCardStat = (params) => {
  return service({
    url: 'miserStat/getCardStat',
    method: 'get',
    params
  })
}

export const getPieStat = (params) => {
  return service({
    url: 'miserStat/getPieStat',
    method: 'get',
    params
  })
}

export const getLineStat = (params) => {
  return service({
    url: 'miserStat/getLineStat',
    method: 'get',
    params
  })
}

export const getMonthStat = (params) => {
  return service({
    url: 'miserStat/getMonthStat',
    method: 'get',
    params
  })
}

export const getYearStat = (params) => {
  return service({
    url: 'miserStat/getYearStat',
    method: 'get',
    params
  })
}

export const getItemStat = (params) => {
  return service({
    url: 'miserStat/getItemStat',
    method: 'get',
    params
  })
}
