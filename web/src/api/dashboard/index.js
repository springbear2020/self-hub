import service from '@/utils/request'

export const getDashboardStats = () => {
  return service({
    url: '/dashboard/stat',
    method: 'get'
  })
}
