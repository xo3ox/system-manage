import service from '@/utils/request'

// @Tags MoneyCompoundInterestRecord
// @Summary 创建MoneyCompoundInterestRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoneyCompoundInterestRecord true "创建MoneyCompoundInterestRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyCompoundInterestRecord/createMoneyCompoundInterestRecord [post]
export const createMoneyCompoundInterestRecord = (data) => {
  return service({
    url: '/moneyCompoundInterestRecord/createMoneyCompoundInterestRecord',
    method: 'post',
    data
  })
}

// @Tags MoneyCompoundInterestRecord
// @Summary 删除MoneyCompoundInterestRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoneyCompoundInterestRecord true "删除MoneyCompoundInterestRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moneyCompoundInterestRecord/deleteMoneyCompoundInterestRecord [delete]
export const deleteMoneyCompoundInterestRecord = (data) => {
  return service({
    url: '/moneyCompoundInterestRecord/deleteMoneyCompoundInterestRecord',
    method: 'delete',
    data
  })
}

// @Tags MoneyCompoundInterestRecord
// @Summary 删除MoneyCompoundInterestRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MoneyCompoundInterestRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moneyCompoundInterestRecord/deleteMoneyCompoundInterestRecord [delete]
export const deleteMoneyCompoundInterestRecordByIds = (data) => {
  return service({
    url: '/moneyCompoundInterestRecord/deleteMoneyCompoundInterestRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags MoneyCompoundInterestRecord
// @Summary 更新MoneyCompoundInterestRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoneyCompoundInterestRecord true "更新MoneyCompoundInterestRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /moneyCompoundInterestRecord/updateMoneyCompoundInterestRecord [put]
export const updateMoneyCompoundInterestRecord = (data) => {
  return service({
    url: '/moneyCompoundInterestRecord/updateMoneyCompoundInterestRecord',
    method: 'put',
    data
  })
}

// @Tags MoneyCompoundInterestRecord
// @Summary 用id查询MoneyCompoundInterestRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MoneyCompoundInterestRecord true "用id查询MoneyCompoundInterestRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /moneyCompoundInterestRecord/findMoneyCompoundInterestRecord [get]
export const findMoneyCompoundInterestRecord = (params) => {
  return service({
    url: '/moneyCompoundInterestRecord/findMoneyCompoundInterestRecord',
    method: 'get',
    params
  })
}

// @Tags MoneyCompoundInterestRecord
// @Summary 分页获取MoneyCompoundInterestRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MoneyCompoundInterestRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyCompoundInterestRecord/getMoneyCompoundInterestRecordList [get]
export const getMoneyCompoundInterestRecordList = (params) => {
  return service({
    url: '/moneyCompoundInterestRecord/getMoneyCompoundInterestRecordList',
    method: 'get',
    params
  })
}

// @Tags getMoneyCompoundInterestRecordSummary
// @Summary 复利存款记录概括
export const getMoneyCompoundInterestRecordSummary = (params) => {
  return service({
    url: '/moneyCompoundInterestRecord/getMoneyCompoundInterestRecordSummary',
    method: 'get',
    params
  })
}