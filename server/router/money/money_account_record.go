package money

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MoneyAccountRecordRouter struct {
}

// InitMoneyAccountRecordRouter 初始化 MoneyAccountRecord 路由信息
func (s *MoneyAccountRecordRouter) InitMoneyAccountRecordRouter(Router *gin.RouterGroup) {
	moneyAccountRecordRouter := Router.Group("moneyAccountRecord").Use(middleware.OperationRecord())
	moneyAccountRecordRouterWithoutRecord := Router.Group("moneyAccountRecord")
	var moneyAccountRecordApi = v1.ApiGroupApp.MoneyApiGroup.MoneyAccountRecordApi
	{
		moneyAccountRecordRouter.POST("createMoneyAccountRecord", moneyAccountRecordApi.CreateMoneyAccountRecord)   // 新建MoneyAccountRecord
		moneyAccountRecordRouter.DELETE("deleteMoneyAccountRecord", moneyAccountRecordApi.DeleteMoneyAccountRecord) // 删除MoneyAccountRecord
		moneyAccountRecordRouter.DELETE("deleteMoneyAccountRecordByIds", moneyAccountRecordApi.DeleteMoneyAccountRecordByIds) // 批量删除MoneyAccountRecord
		moneyAccountRecordRouter.PUT("updateMoneyAccountRecord", moneyAccountRecordApi.UpdateMoneyAccountRecord)    // 更新MoneyAccountRecord
	}
	{
		moneyAccountRecordRouterWithoutRecord.GET("findMoneyAccountRecord", moneyAccountRecordApi.FindMoneyAccountRecord)        // 根据ID获取MoneyAccountRecord
		moneyAccountRecordRouterWithoutRecord.GET("getMoneyAccountRecordList", moneyAccountRecordApi.GetMoneyAccountRecordList)  // 获取MoneyAccountRecord列表
	}
}
