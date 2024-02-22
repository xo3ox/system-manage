package money

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/money"
	moneyReq "github.com/flipped-aurora/gin-vue-admin/server/model/money/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/function"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MoneyCompoundInterestRecordApi struct {
}

var moneyCompoundInterestRecordService = service.ServiceGroupApp.MoneyServiceGroup.MoneyCompoundInterestRecordService

// CreateMoneyCompoundInterestRecord 创建
// @Tags MoneyCompoundInterestRecord
// @Summary 创建
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - body money.MoneyCompoundInterestRecord true "传递object"
// @Success 200 {object} response.Response "返回结果"
// @Router /moneyCompoundInterestRecord/createMoneyCompoundInterestRecord [post]
func (moneyCompoundInterestRecordApi *MoneyCompoundInterestRecordApi) CreateMoneyCompoundInterestRecord(c *gin.Context) {
	var moneyCompoundInterestRecord money.MoneyCompoundInterestRecord
	err := c.ShouldBindJSON(&moneyCompoundInterestRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Principal": {utils.NotEmpty()},
	}
	if err := utils.Verify(moneyCompoundInterestRecord, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	moneyCompoundInterestRecord.CreatedBy = utils.GetUserID(c)
	if err := moneyCompoundInterestRecordService.CreateMoneyCompoundInterestRecord(&moneyCompoundInterestRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMoneyCompoundInterestRecord 删除
// @Tags MoneyCompoundInterestRecord
// @Summary 删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query int true "传递id"
// @Success 200 {object} response.Response "返回结果"
// @Router /moneyCompoundInterestRecord/deleteMoneyCompoundInterestRecord [delete]
func (moneyCompoundInterestRecordApi *MoneyCompoundInterestRecordApi) DeleteMoneyCompoundInterestRecord(c *gin.Context) {
	var moneyCompoundInterestRecord money.MoneyCompoundInterestRecord
	err := c.ShouldBindJSON(&moneyCompoundInterestRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyCompoundInterestRecordService.DeleteMoneyCompoundInterestRecord(moneyCompoundInterestRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMoneyCompoundInterestRecordByIds 批量删除
// @Tags MoneyCompoundInterestRecord
// @Summary 批量删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - body request.IdsReq true "传递object"
// @Success 200 {object} response.Response "返回结果"
// @Router /moneyCompoundInterestRecord/deleteMoneyCompoundInterestRecordByIds [delete]
func (moneyCompoundInterestRecordApi *MoneyCompoundInterestRecordApi) DeleteMoneyCompoundInterestRecordByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyCompoundInterestRecordService.DeleteMoneyCompoundInterestRecordByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMoneyCompoundInterestRecord 更新
// @Tags MoneyCompoundInterestRecord
// @Summary 更新
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - body money.MoneyCompoundInterestRecord true "传递object"
// @Success 200 {object} response.Response "返回结果"
// @Router /moneyCompoundInterestRecord/updateMoneyCompoundInterestRecord [put]
func (moneyCompoundInterestRecordApi *MoneyCompoundInterestRecordApi) UpdateMoneyCompoundInterestRecord(c *gin.Context) {
	var moneyCompoundInterestRecord money.MoneyCompoundInterestRecord
	err := c.ShouldBindJSON(&moneyCompoundInterestRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Principal": {utils.NotEmpty()},
	}
	if err := utils.Verify(moneyCompoundInterestRecord, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if moneyCompoundInterestRecord.PrincipalInterest != nil {
		temp := function.Decimal(*moneyCompoundInterestRecord.PrincipalInterest - *moneyCompoundInterestRecord.Principal)
		moneyCompoundInterestRecord.Interest = &temp
	}
	moneyCompoundInterestRecord.UpdatedBy = utils.GetUserID(c)
	if err := moneyCompoundInterestRecordService.UpdateMoneyCompoundInterestRecord(moneyCompoundInterestRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMoneyCompoundInterestRecord 用ID查询
// @Tags MoneyCompoundInterestRecord
// @Summary 用id查询
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query int true "id"
// @Success 200 {object} response.Response{data=money.MoneyCompoundInterestRecord} "返回结果"
// @Router /moneyCompoundInterestRecord/findMoneyCompoundInterestRecord [get]
func (moneyCompoundInterestRecordApi *MoneyCompoundInterestRecordApi) FindMoneyCompoundInterestRecord(c *gin.Context) {
	var moneyCompoundInterestRecord money.MoneyCompoundInterestRecord
	err := c.ShouldBindQuery(&moneyCompoundInterestRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remoneyCompoundInterestRecord, err := moneyCompoundInterestRecordService.GetMoneyCompoundInterestRecord(moneyCompoundInterestRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "查询失败", c)
	} else {
		response.OkWithData(gin.H{"remoneyCompoundInterestRecord": remoneyCompoundInterestRecord}, c)
	}
}

// GetMoneyCompoundInterestRecordList 分页获取列表
// @Tags MoneyCompoundInterestRecord
// @Summary 分页获取列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - query moneyReq.MoneyCompoundInterestRecordSearch true "传递object"
// @Success 200 {object} response.Response{data=[]money.MoneyCompoundInterestRecord} "返回结果"
// @Router /moneyCompoundInterestRecord/getMoneyCompoundInterestRecordList [get]
func (moneyCompoundInterestRecordApi *MoneyCompoundInterestRecordApi) GetMoneyCompoundInterestRecordList(c *gin.Context) {
	var info moneyReq.MoneyCompoundInterestRecordSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info.CreatedBy = utils.GetUserID(c)
	if list, total, err := moneyCompoundInterestRecordService.GetMoneyCompoundInterestRecordInfoList(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, "获取成功", c)
	}
}

// GetMoneyCompoundInterestRecordSummary 获取复利存款记录累计概况
// @Tags MoneyCompoundInterestRecord
// @Summary 累计概况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=money.MoneyCompoundInterestRecordSummary} "返回结果"
// @Router /moneyCompoundInterestRecord/getMoneyCompoundInterestRecordSummary [get]
func (moneyCompoundInterestRecordApi *MoneyCompoundInterestRecordApi) GetMoneyCompoundInterestRecordSummary(c *gin.Context) {
	userId := utils.GetUserID(c)
	if result, err := moneyCompoundInterestRecordService.GetMoneyCompoundInterestRecordSummary(userId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "查询失败", c)
	} else {
		response.OkWithData(gin.H{"result": result}, c)
	}
}
