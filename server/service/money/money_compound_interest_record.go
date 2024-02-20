package money

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/money"
	moneyReq "github.com/flipped-aurora/gin-vue-admin/server/model/money/request"
	"gorm.io/gorm"
)

type MoneyCompoundInterestRecordService struct {
}

// CreateMoneyCompoundInterestRecord 创建MoneyCompoundInterestRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyCompoundInterestRecordService *MoneyCompoundInterestRecordService) CreateMoneyCompoundInterestRecord(moneyCompoundInterestRecord *money.MoneyCompoundInterestRecord) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(moneyCompoundInterestRecord).Error; err != nil {
			return err
		}
		return nil
	})
	return
}

// DeleteMoneyCompoundInterestRecord 删除MoneyCompoundInterestRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyCompoundInterestRecordService *MoneyCompoundInterestRecordService) DeleteMoneyCompoundInterestRecord(moneyCompoundInterestRecord money.MoneyCompoundInterestRecord) (err error) {
	if err = global.GVA_DB.Delete(&moneyCompoundInterestRecord).Error; err != nil {
		return
	}
	return
}

// DeleteMoneyCompoundInterestRecordByIds 批量删除MoneyCompoundInterestRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyCompoundInterestRecordService *MoneyCompoundInterestRecordService) DeleteMoneyCompoundInterestRecordByIds(ids request.IdsReq) (err error) {
	if err = global.GVA_DB.Where("id IN ?", ids.Ids).Delete(&money.MoneyCompoundInterestRecord{}).Error; err != nil {
		return
	}
	return
}

// UpdateMoneyCompoundInterestRecord 更新MoneyCompoundInterestRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyCompoundInterestRecordService *MoneyCompoundInterestRecordService) UpdateMoneyCompoundInterestRecord(moneyCompoundInterestRecord money.MoneyCompoundInterestRecord) (err error) {
	err = global.GVA_DB.Updates(&moneyCompoundInterestRecord).Error
	return
}

// GetMoneyCompoundInterestRecord 根据id获取MoneyCompoundInterestRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyCompoundInterestRecordService *MoneyCompoundInterestRecordService) GetMoneyCompoundInterestRecord(id uint) (moneyCompoundInterestRecord money.MoneyCompoundInterestRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).Find(&moneyCompoundInterestRecord).Error
	return
}

// GetMoneyCompoundInterestRecordInfoList 分页获取MoneyCompoundInterestRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyCompoundInterestRecordService *MoneyCompoundInterestRecordService) GetMoneyCompoundInterestRecordInfoList(info moneyReq.MoneyCompoundInterestRecordSearch) (list []*money.MoneyCompoundInterestRecord, total int64, err error) {

	// 创建db
	db := global.GVA_DB.Model(&money.MoneyCompoundInterestRecord{}).Where("created_by = ?", info.CreatedBy)
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartDepositTime != nil && info.EndDepositTime != nil {
		db = db.Where("deposit_time BETWEEN ? AND ? ", info.StartDepositTime, info.EndDepositTime)
	}

	if info.Page > 0 && info.PageSize > 0 {
		db = db.Limit(info.PageSize).Offset((info.Page - 1) * info.PageSize)
	}

	if err = db.Order("id DESC").Find(&list).Limit(-1).Count(&total).Offset(-1).Error; err != nil {
		return
	}
	return
}
