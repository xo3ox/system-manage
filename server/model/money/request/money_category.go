package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/money"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MoneyCategorySearch struct{
    money.MoneyCategory
    request.PageInfo
}
