package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `json:"id" form:"id" gorm:"column:id;primarykey"`                                    // 主键ID
	CreatedAt time.Time      `json:"createdAt" form:"createdAt" gorm:"autoCreateTime:milli" swaggerignore:"true"` // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" form:"updatedAt" gorm:"autoUpdateTime:milli" swaggerignore:"true"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" form:"deletedAt" gorm:"index"`                                             // 删除时间
}

type GVA_MODEL_V2 struct {
	GVA_MODEL
	CreatedBy uint `json:"createdBy" form:"createdBy" swaggerignore:"true"`
	UpdatedBy uint `json:"updatedBy" form:"updatedBy" swaggerignore:"true"`
}
