import service from '@/utils/request'

// @Tags MoneyCategory
// @Summary 创建MoneyCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoneyCategory true "创建MoneyCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyCategory/createMoneyCategory [post]
export const createMoneyCategory = (data) => {
  return service({
    url: '/moneyCategory/createMoneyCategory',
    method: 'post',
    data
  })
}

// @Tags MoneyCategory
// @Summary 删除MoneyCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoneyCategory true "删除MoneyCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moneyCategory/deleteMoneyCategory [delete]
export const deleteMoneyCategory = (data) => {
  return service({
    url: '/moneyCategory/deleteMoneyCategory',
    method: 'delete',
    data
  })
}

// @Tags MoneyCategory
// @Summary 删除MoneyCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MoneyCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moneyCategory/deleteMoneyCategory [delete]
export const deleteMoneyCategoryByIds = (data) => {
  return service({
    url: '/moneyCategory/deleteMoneyCategoryByIds',
    method: 'delete',
    data
  })
}

// @Tags MoneyCategory
// @Summary 更新MoneyCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MoneyCategory true "更新MoneyCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /moneyCategory/updateMoneyCategory [put]
export const updateMoneyCategory = (data) => {
  return service({
    url: '/moneyCategory/updateMoneyCategory',
    method: 'put',
    data
  })
}

// @Tags MoneyCategory
// @Summary 用id查询MoneyCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MoneyCategory true "用id查询MoneyCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /moneyCategory/findMoneyCategory [get]
export const findMoneyCategory = (params) => {
  return service({
    url: '/moneyCategory/findMoneyCategory',
    method: 'get',
    params
  })
}

// @Tags MoneyCategory
// @Summary 分页获取MoneyCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MoneyCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyCategory/getMoneyCategoryList [get]
export const getMoneyCategoryList = (params) => {
  return service({
    url: '/moneyCategory/getMoneyCategoryList',
    method: 'get',
    params
  })
}
