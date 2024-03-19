import service from '@/utils/request'

// @Tags MoneyAccountRecord
// @Summary 创建MoneyAccountRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoneyAccountRecord true "创建MoneyAccountRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyAccountRecord/createMoneyAccountRecord [post]
export const createMoneyAccountRecord = (data) => {
  return service({
    url: '/moneyAccountRecord/createMoneyAccountRecord',
    method: 'post',
    data
  })
}

// @Tags MoneyAccountRecord
// @Summary 删除MoneyAccountRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoneyAccountRecord true "删除MoneyAccountRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moneyAccountRecord/deleteMoneyAccountRecord [delete]
export const deleteMoneyAccountRecord = (data) => {
  return service({
    url: '/moneyAccountRecord/deleteMoneyAccountRecord',
    method: 'delete',
    data
  })
}

// @Tags MoneyAccountRecord
// @Summary 删除MoneyAccountRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MoneyAccountRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moneyAccountRecord/deleteMoneyAccountRecord [delete]
export const deleteMoneyAccountRecordByIds = (data) => {
  return service({
    url: '/moneyAccountRecord/deleteMoneyAccountRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags MoneyAccountRecord
// @Summary 更新MoneyAccountRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoneyAccountRecord true "更新MoneyAccountRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /moneyAccountRecord/updateMoneyAccountRecord [put]
export const updateMoneyAccountRecord = (data) => {
  return service({
    url: '/moneyAccountRecord/updateMoneyAccountRecord',
    method: 'put',
    data
  })
}

// @Tags MoneyAccountRecord
// @Summary 用id查询MoneyAccountRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MoneyAccountRecord true "用id查询MoneyAccountRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /moneyAccountRecord/findMoneyAccountRecord [get]
export const findMoneyAccountRecord = (params) => {
  return service({
    url: '/moneyAccountRecord/findMoneyAccountRecord',
    method: 'get',
    params
  })
}

// @Tags MoneyAccountRecord
// @Summary 分页获取MoneyAccountRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MoneyAccountRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyAccountRecord/getMoneyAccountRecordList [get]
export const getMoneyAccountRecordList = (params) => {
  return service({
    url: '/moneyAccountRecord/getMoneyAccountRecordList',
    method: 'get',
    params
  })
}
// @Tags getMoneyAccountRecordSummary
// @Summary 账目管理概括
export const getMoneyAccountRecordSummary = (params) => {
  return service({
    url: '/moneyAccountRecord/getMoneyAccountRecordSummary',
    method: 'get',
    params
  })
}