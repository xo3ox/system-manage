package money

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/money"
	moneyReq "github.com/flipped-aurora/gin-vue-admin/server/model/money/request"
	"gorm.io/gorm"
)

type MoneyCategoryService struct {
}

// CreateMoneyCategory 创建MoneyCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyCategoryService *MoneyCategoryService) CreateMoneyCategory(moneyCategory *money.MoneyCategory) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 不存在则创建
		if err = tx.Create(moneyCategory).Error; err != nil {
			return err
		}
		return nil
	})
	return
}

// DeleteMoneyCategory 删除MoneyCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyCategoryService *MoneyCategoryService) DeleteMoneyCategory(moneyCategory money.MoneyCategory) (err error) {
	if err = global.GVA_DB.Delete(&moneyCategory).Error; err != nil {
		return
	}
	return
}

// DeleteMoneyCategoryByIds 批量删除MoneyCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyCategoryService *MoneyCategoryService) DeleteMoneyCategoryByIds(ids request.IdsReq) (err error) {
	if err = global.GVA_DB.Where("id IN ?", ids.Ids).Delete(&money.MoneyCategory{}).Error; err != nil {
		return
	}
	return
}

// UpdateMoneyCategory 更新MoneyCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyCategoryService *MoneyCategoryService) UpdateMoneyCategory(moneyCategory money.MoneyCategory) (err error) {
	err = global.GVA_DB.Updates(&moneyCategory).Error
	return
}

// GetMoneyCategory 根据id获取MoneyCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyCategoryService *MoneyCategoryService) GetMoneyCategory(id uint) (moneyCategory money.MoneyCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).Find(&moneyCategory).Error
	return
}

// GetMoneyCategoryInfoList 分页获取MoneyCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyCategoryService *MoneyCategoryService) GetMoneyCategoryInfoList(info moneyReq.MoneyCategorySearch) (list []*money.MoneyCategory, total int64, err error) {

	// 创建db
	db := global.GVA_DB.Model(&money.MoneyCategory{}).Where("created_by = ?", info.CreatedBy)
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.Page > 0 && info.PageSize > 0 {
		db = db.Limit(info.PageSize).Offset((info.Page - 1) * info.PageSize)
	}

	if err = db.Order("id DESC").Find(&list).Limit(-1).Count(&total).Offset(-1).Error; err != nil {
		return
	}
	return
}
