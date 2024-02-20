package money

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MoneyCategoryRouter struct {
}

// InitMoneyCategoryRouter 初始化 MoneyCategory 路由信息
func (s *MoneyCategoryRouter) InitMoneyCategoryRouter(Router *gin.RouterGroup) {
	moneyCategoryRouter := Router.Group("moneyCategory").Use(middleware.OperationRecord())
	moneyCategoryRouterWithoutRecord := Router.Group("moneyCategory")
	var moneyCategoryApi = v1.ApiGroupApp.MoneyApiGroup.MoneyCategoryApi
	{
		moneyCategoryRouter.POST("createMoneyCategory", moneyCategoryApi.CreateMoneyCategory)   // 新建MoneyCategory
		moneyCategoryRouter.DELETE("deleteMoneyCategory", moneyCategoryApi.DeleteMoneyCategory) // 删除MoneyCategory
		moneyCategoryRouter.DELETE("deleteMoneyCategoryByIds", moneyCategoryApi.DeleteMoneyCategoryByIds) // 批量删除MoneyCategory
		moneyCategoryRouter.PUT("updateMoneyCategory", moneyCategoryApi.UpdateMoneyCategory)    // 更新MoneyCategory
	}
	{
		moneyCategoryRouterWithoutRecord.GET("findMoneyCategory", moneyCategoryApi.FindMoneyCategory)        // 根据ID获取MoneyCategory
		moneyCategoryRouterWithoutRecord.GET("getMoneyCategoryList", moneyCategoryApi.GetMoneyCategoryList)  // 获取MoneyCategory列表
	}
}
