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

type MoneyCategoryApi struct {
}

var moneyCategoryService = service.ServiceGroupApp.MoneyServiceGroup.MoneyCategoryService

// CreateMoneyCategory 创建
// @Tags MoneyCategory
// @Summary 创建
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - body money.MoneyCategory true "传递object"
// @Success 200 {object} response.Response "返回结果"
// @Router /moneyCategory/createMoneyCategory [post]
func (moneyCategoryApi *MoneyCategoryApi) CreateMoneyCategory(c *gin.Context) {
	var moneyCategory money.MoneyCategory
	err := c.ShouldBindJSON(&moneyCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	moneyCategory.CreatedBy = utils.GetUserID(c)
	if err := moneyCategoryService.CreateMoneyCategory(&moneyCategory); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMoneyCategory 删除
// @Tags MoneyCategory
// @Summary 删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query int true "传递id"
// @Success 200 {object} response.Response "返回结果"
// @Router /moneyCategory/deleteMoneyCategory [delete]
func (moneyCategoryApi *MoneyCategoryApi) DeleteMoneyCategory(c *gin.Context) {
	var moneyCategory money.MoneyCategory
	err := c.ShouldBindJSON(&moneyCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyCategoryService.DeleteMoneyCategory(moneyCategory); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMoneyCategoryByIds 批量删除
// @Tags MoneyCategory
// @Summary 批量删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - body request.IdsReq true "传递object"
// @Success 200 {object} response.Response "返回结果"
// @Router /moneyCategory/deleteMoneyCategoryByIds [delete]
func (moneyCategoryApi *MoneyCategoryApi) DeleteMoneyCategoryByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyCategoryService.DeleteMoneyCategoryByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMoneyCategory 更新
// @Tags MoneyCategory
// @Summary 更新
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - body money.MoneyCategory true "传递object"
// @Success 200 {object} response.Response "返回结果"
// @Router /moneyCategory/updateMoneyCategory [put]
func (moneyCategoryApi *MoneyCategoryApi) UpdateMoneyCategory(c *gin.Context) {
	var moneyCategory money.MoneyCategory
	err := c.ShouldBindJSON(&moneyCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	moneyCategory.UpdatedBy = utils.GetUserID(c)
	if err := moneyCategoryService.UpdateMoneyCategory(moneyCategory); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMoneyCategory 用ID查询
// @Tags MoneyCategory
// @Summary 用id查询
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query int true "id"
// @Success 200 {object} response.Response{data=money.MoneyCategory} "返回结果"
// @Router /moneyCategory/findMoneyCategory [get]
func (moneyCategoryApi *MoneyCategoryApi) FindMoneyCategory(c *gin.Context) {
	var moneyCategory money.MoneyCategory
	err := c.ShouldBindQuery(&moneyCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remoneyCategory, err := moneyCategoryService.GetMoneyCategory(moneyCategory.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(), "查询失败", c)
	} else {
		response.OkWithData(gin.H{"remoneyCategory": remoneyCategory}, c)
	}
}

// GetMoneyCategoryList 分页获取列表
// @Tags MoneyCategory
// @Summary 分页获取列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - query moneyReq.MoneyCategorySearch true "传递object"
// @Success 200 {object} response.Response{data=[]money.MoneyCategory} "返回结果"
// @Router /moneyCategory/getMoneyCategoryList [get]
func (moneyCategoryApi *MoneyCategoryApi) GetMoneyCategoryList(c *gin.Context) {
	var info moneyReq.MoneyCategorySearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info.CreatedBy = utils.GetUserID(c)
	if list, total, err := moneyCategoryService.GetMoneyCategoryInfoList(info); err != nil {
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
