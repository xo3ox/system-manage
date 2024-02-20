package money

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MoneyCompoundInterestRecordRouter struct {
}

// InitMoneyCompoundInterestRecordRouter 初始化 MoneyCompoundInterestRecord 路由信息
func (s *MoneyCompoundInterestRecordRouter) InitMoneyCompoundInterestRecordRouter(Router *gin.RouterGroup) {
	moneyCompoundInterestRecordRouter := Router.Group("moneyCompoundInterestRecord").Use(middleware.OperationRecord())
	moneyCompoundInterestRecordRouterWithoutRecord := Router.Group("moneyCompoundInterestRecord")
	var moneyCompoundInterestRecordApi = v1.ApiGroupApp.MoneyApiGroup.MoneyCompoundInterestRecordApi
	{
		moneyCompoundInterestRecordRouter.POST("createMoneyCompoundInterestRecord", moneyCompoundInterestRecordApi.CreateMoneyCompoundInterestRecord)   // 新建MoneyCompoundInterestRecord
		moneyCompoundInterestRecordRouter.DELETE("deleteMoneyCompoundInterestRecord", moneyCompoundInterestRecordApi.DeleteMoneyCompoundInterestRecord) // 删除MoneyCompoundInterestRecord
		moneyCompoundInterestRecordRouter.DELETE("deleteMoneyCompoundInterestRecordByIds", moneyCompoundInterestRecordApi.DeleteMoneyCompoundInterestRecordByIds) // 批量删除MoneyCompoundInterestRecord
		moneyCompoundInterestRecordRouter.PUT("updateMoneyCompoundInterestRecord", moneyCompoundInterestRecordApi.UpdateMoneyCompoundInterestRecord)    // 更新MoneyCompoundInterestRecord
	}
	{
		moneyCompoundInterestRecordRouterWithoutRecord.GET("findMoneyCompoundInterestRecord", moneyCompoundInterestRecordApi.FindMoneyCompoundInterestRecord)        // 根据ID获取MoneyCompoundInterestRecord
		moneyCompoundInterestRecordRouterWithoutRecord.GET("getMoneyCompoundInterestRecordList", moneyCompoundInterestRecordApi.GetMoneyCompoundInterestRecordList)  // 获取MoneyCompoundInterestRecord列表
	}
}
