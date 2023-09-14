import request from '@/utils/request'

// 查询通知公告表列表
export function listNotice(query) {
  return request({
    url: '/sys/notice/page',
    method: 'get',
    params: query
  })
}

// 查询通知公告表详细
export function getNotice(id) {
  return request({
    url: '/sys/notice/get/' + id,
    method: 'get'
  })
}

// 新增通知公告表
export function addNotice(data) {
  return request({
    url: '/sys/notice/create',
    method: 'post',
    data: data
  })
}

// 修改通知公告表
export function updateNotice(data) {
  return request({
    url: '/sys/notice/update',
    method: 'put',
    data: data
  })
}

// 删除通知公告表
export function delNotice(id) {
  return request({
    url: '/sys/notice/delete/' + id,
    method: 'delete'
  })
}
