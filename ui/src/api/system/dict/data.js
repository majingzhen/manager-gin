import request from '@/utils/request'

// 查询字典数据列表
export function listData(query) {
  return request({
    url: 'sys/sysDictData/page',
    method: 'get',
    params: query
  })
}

// 查询字典数据详细
export function getData(id) {
  return request({
    url: 'sys/sysDictData/get/' + id,
    method: 'get'
  })
}

// 根据字典类型查询字典数据信息
export function getDicts(dictType) {
  return request({
    url: 'sys/sysDictData/type/' + dictType,
    method: 'get'
  })
}

// 新增字典数据
export function addData(data) {
  return request({
    url: 'sys/sysDictData/create',
    method: 'post',
    data: data
  })
}

// 修改字典数据
export function updateData(data) {
  return request({
    url: 'sys/sysDictData/update',
    method: 'post',
    data: data
  })
}

// 删除字典数据
export function delData(id) {
  return request({
    url: 'sys/sysDictData/delete/' + id,
    method: 'delete'
  })
}
