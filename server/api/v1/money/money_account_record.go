package money

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/money"
	moneyReq "github.com/flipped-aurora/gin-vue-admin/server/model/money/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MoneyAccountRecordApi struct {
}

var moneyAccountRecordService = service.ServiceGroupApp.MoneyServiceGroup.MoneyAccountRecordService

// CreateMoneyAccountRecord 创建
// @Tags MoneyAccountRecord
// @Summary 创建
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - body money.MoneyAccountRecord true "传递object"
// @Success 200 {object} response.Response "返回结果"
// @Router /moneyAccountRecord/createMoneyAccountRecord [post]
func (moneyAccountRecordApi *MoneyAccountRecordApi) CreateMoneyAccountRecord(c *gin.Context) {
	var moneyAccountRecord money.MoneyAccountRecord
	err := c.ShouldBindJSON(&moneyAccountRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Type":      {utils.NotEmpty()},
		"Category":  {utils.NotEmpty()},
		"OccurTime": {utils.NotEmpty()},
		"Money":     {utils.NotEmpty()},
	}
	if err := utils.Verify(moneyAccountRecord, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	moneyAccountRecord.CreatedBy = utils.GetUserID(c)
	if err := moneyAccountRecordService.CreateMoneyAccountRecord(&moneyAccountRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMoneyAccountRecord 删除
// @Tags MoneyAccountRecord
// @Summary 删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query int true "传递id"
// @Success 200 {object} response.Response "返回结果"
// @Router /moneyAccountRecord/deleteMoneyAccountRecord [delete]
func (moneyAccountRecordApi *MoneyAccountRecordApi) DeleteMoneyAccountRecord(c *gin.Context) {
	var moneyAccountRecord money.MoneyAccountRecord
	err := c.ShouldBindJSON(&moneyAccountRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyAccountRecordService.DeleteMoneyAccountRecord(moneyAccountRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMoneyAccountRecordByIds 批量删除
// @Tags MoneyAccountRecord
// @Summary 批量删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - body request.IdsReq true "传递object"
// @Success 200 {object} response.Response "返回结果"
// @Router /moneyAccountRecord/deleteMoneyAccountRecordByIds [delete]
func (moneyAccountRecordApi *MoneyAccountRecordApi) DeleteMoneyAccountRecordByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyAccountRecordService.DeleteMoneyAccountRecordByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMoneyAccountRecord 更新
// @Tags MoneyAccountRecord
// @Summary 更新
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - body money.MoneyAccountRecord true "传递object"
// @Success 200 {object} response.Response "返回结果"
// @Router /moneyAccountRecord/updateMoneyAccountRecord [put]
func (moneyAccountRecordApi *MoneyAccountRecordApi) UpdateMoneyAccountRecord(c *gin.Context) {
	var moneyAccountRecord money.MoneyAccountRecord
	err := c.ShouldBindJSON(&moneyAccountRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Type":      {utils.NotEmpty()},
		"Category":  {utils.NotEmpty()},
		"OccurTime": {utils.NotEmpty()},
		"Money":     {utils.NotEmpty()},
	}
	if err := utils.Verify(moneyAccountRecord, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	moneyAccountRecord.UpdatedBy = utils.GetUserID(c)
	if err := moneyAccountRecordService.UpdateMoneyAccountRecord(moneyAccountRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMoneyAccountRecord 用ID查询
// @Tags MoneyAccountRecord
// @Summary 用id查询
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query int true "id"
// @Success 200 {object} response.Response{data=money.MoneyAccountRecord} "返回结果"
// @Router /moneyAccountRecord/findMoneyAccountRecord [get]
func (moneyAccountRecordApi *MoneyAccountRecordApi) FindMoneyAccountRecord(c *gin.Context) {
	var moneyAccountRecord money.MoneyAccountRecord
	err := c.ShouldBindQuery(&moneyAccountRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remoneyAccountRecord, err := moneyAccountRecordService.GetMoneyAccountRecord(moneyAccountRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "查询失败", c)
	} else {
		response.OkWithData(gin.H{"remoneyAccountRecord": remoneyAccountRecord}, c)
	}
}

// GetMoneyAccountRecordList 分页获取列表
// @Tags MoneyAccountRecord
// @Summary 分页获取列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - query moneyReq.MoneyAccountRecordSearch true "传递object"
// @Success 200 {object} response.Response{data=[]money.MoneyAccountRecord} "返回结果"
// @Router /moneyAccountRecord/getMoneyAccountRecordList [get]
func (moneyAccountRecordApi *MoneyAccountRecordApi) GetMoneyAccountRecordList(c *gin.Context) {
	var info moneyReq.MoneyAccountRecordSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := moneyAccountRecordService.GetMoneyAccountRecordInfoList(info); err != nil {
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
