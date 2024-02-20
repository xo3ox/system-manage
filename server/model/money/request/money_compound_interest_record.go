package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/money"
)

type MoneyCompoundInterestRecordSearch struct {
	money.MoneyCompoundInterestRecord
	StartDepositTime *time.Time `json:"startDepositTime" form:"startDepositTime"`
	EndDepositTime   *time.Time `json:"endDepositTime" form:"endDepositTime"`
	request.PageInfo
}
