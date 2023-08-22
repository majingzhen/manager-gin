import request from '@/utils/request'

// 查询岗位列表
export function listPost(query) {
  return request({
    url: '/sys/sysPost/page',
    method: 'get',
    params: query
  })
}

// 查询岗位详细
export function getPost(postId) {
  return request({
    url: '/sys/sysPost/get/' + postId,
    method: 'get'
  })
}

// 新增岗位
export function addPost(data) {
  return request({
    url: '/sys/sysPost/create',
    method: 'post',
    data: data
  })
}

// 修改岗位
export function updatePost(data) {
  return request({
    url: '/sys/sysPost/update',
    method: 'post',
    data: data
  })
}

// 删除岗位
export function delPost(postId) {
  return request({
    url: '/sys/sysPost/delete/' + postId,
    method: 'delete'
  })
}
