// 自动生成模板MoneyAccountRecord
package money

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MoneyAccountRecord 结构体
type MoneyAccountRecord struct {
	global.GVA_MODEL_V2
	InOut      *int       `json:"inOut" form:"inOut" gorm:"column:in_out;"`                // 类型;1：收入（钱进） 2：支出（钱出）
	CategoryId *int       `json:"categoryId" form:"categoryId" gorm:"column:category_id;"` // 分类;1:存取款
	OccurTime  *time.Time `json:"occurTime" form:"occurTime" gorm:"column:occur_time;"`    // 发生时间
	Money      *float64   `json:"money" form:"money" gorm:"column:money;"`                 // 金额
	Remark     string     `json:"remark" form:"remark" gorm:"column:remark;"`              // 备注
}

// TableName MoneyAccountRecord 表名
func (MoneyAccountRecord) TableName() string {
	return "money_account_record"
}
