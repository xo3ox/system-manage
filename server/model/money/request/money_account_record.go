package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/money"
)

type MoneyAccountRecordSearch struct {
	money.MoneyAccountRecord
	StartOccurTime *time.Time `json:"startOccurTime" form:"startOccurTime"`
	EndOccurTime   *time.Time `json:"endOccurTime" form:"endOccurTime"`
	request.PageInfo
}
