// 自动生成模板MoneyCompoundInterestRecord
package money

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MoneyCompoundInterestRecord 结构体
type MoneyCompoundInterestRecord struct {
	global.GVA_MODEL_V2
	DepositTime       *time.Time `json:"depositTime" form:"depositTime" gorm:"column:deposit_time;"`                   // 存款时间
	ExpirationTime    *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;"`          // 到期时间
	Principal         *float64   `json:"principal" form:"principal" gorm:"column:principal;"`                          // 本金
	Interest          *float64   `json:"interest" form:"interest" gorm:"column:interest;"`                             // 利息
	PrincipalInterest *float64   `json:"principalInterest" form:"principalInterest" gorm:"column:principal_interest;"` // 本息和
}

// TableName MoneyCompoundInterestRecord 表名
func (MoneyCompoundInterestRecord) TableName() string {
	return "money_compound_interest_record"
}

type MoneyCompoundInterestRecordSummary struct {
	SumPrincipal         float64 `json:"sumPrincipal"`                  // 累计本金：当前用户存款时间最近十二个月的本金
	SumInterest          float64 `json:"sumInterest"`                   // 累计利息：当前用户所有记录产生的利息
	SumPrincipalInterest float64 `json:"sumPrincipalInterest" gorm:"-"` // 累计本息和：当前用户  累计本金+累计利息
}
