import request from '@/utils/request'

const prefix = '/v1/admin'

const api = {
  profile: `${prefix}/profile`,
  env: `${prefix}/env`,
  res: `${prefix}/res`,
  vms: `${prefix}/vms`,
  vmTempls: `${prefix}/vmTempls`,
  containers: `${prefix}/containers`,

  plans: `${prefix}/plans`,

  user: `${prefix}/user`,
  role: `${prefix}/role`,
  service: `${prefix}/service`,
  permission: `${prefix}/permission`,
  permissionNoPager: `${prefix}/permission/no-pager`,
  orgTree: `${prefix}/org/tree`
}

export const WsApi = 'ws://127.0.0.1:8085/api/v1/ws'

export function getProfile (parameter) {
  return request({
    url: api.profile,
    method: 'get',
    data: parameter
  })
}

export function listPlan () {
  return request({
    url: api.plans,
    method: 'get',
    params: {}
  })
}

export function listEnv () {
  return request({
    url: api.env,
    method: 'get',
    params: {}
  })
}

export function listVm () {
  return request({
    url: api.res + '/listVm',
    method: 'get',
    params: {}
  })
}
export function loadVmTempl (data) {
  return request({
    url: api.vmTempls,
    method: 'post',
    data: data
  })
}
export function saveVmTempl (model) {
  return request({
    url: api.vmTempls,
    method: 'put',
    data: model
  })
}

export function listContainer (parameter) {
  return request({
    url: api.res + '/listContainer',
    method: 'get',
    params: parameter
  })
}

export function getUserList (parameter) {
  return request({
    url: api.user,
    method: 'get',
    params: parameter
  })
}

export function getRoleList (parameter) {
  return request({
    url: api.role,
    method: 'get',
    params: parameter
  })
}

export function getServiceList (parameter) {
  return request({
    url: api.service,
    method: 'get',
    params: parameter
  })
}

export function getPermissions (parameter) {
  return request({
    url: api.permissionNoPager,
    method: 'get',
    params: parameter
  })
}

export function getOrgTree (parameter) {
  return request({
    url: api.orgTree,
    method: 'get',
    params: parameter
  })
}

// id == 0 add     post
// id != 0 update  put
export function saveService (parameter) {
  return request({
    url: api.service,
    method: parameter.id === 0 ? 'post' : 'put',
    data: parameter
  })
}

export function saveSub (sub) {
  return request({
    url: '/sub',
    method: sub.id === 0 ? 'post' : 'put',
    data: sub
  })
}
