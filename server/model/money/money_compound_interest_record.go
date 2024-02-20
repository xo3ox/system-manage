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
	PrincipalInterest *float64   `json:"principalInterest" form:"principalInterest" gorm:"column:principal_interest;"` // 本利和
}

// TableName MoneyCompoundInterestRecord 表名
func (MoneyCompoundInterestRecord) TableName() string {
	return "money_compound_interest_record"
}
