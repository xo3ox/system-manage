// 自动生成模板MoneyCategory
package money

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// MoneyCategory 结构体
type MoneyCategory struct {
      global.GVA_MODEL_V2
      Tag  string `json:"tag" form:"tag" gorm:"column:tag;"` // 分类名称
}


// TableName MoneyCategory 表名
func (MoneyCategory) TableName() string {
  return "money_category"
}

