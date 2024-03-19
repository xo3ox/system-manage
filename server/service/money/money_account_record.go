package money

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/money"
	moneyReq "github.com/flipped-aurora/gin-vue-admin/server/model/money/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/function"
	"gorm.io/gorm"
)

type MoneyAccountRecordService struct {
}

// CreateMoneyAccountRecord 创建MoneyAccountRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyAccountRecordService *MoneyAccountRecordService) CreateMoneyAccountRecord(moneyAccountRecord *money.MoneyAccountRecord) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(moneyAccountRecord).Error; err != nil {
			return err
		}
		return nil
	})
	return
}

// DeleteMoneyAccountRecord 删除MoneyAccountRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyAccountRecordService *MoneyAccountRecordService) DeleteMoneyAccountRecord(moneyAccountRecord money.MoneyAccountRecord) (err error) {
	if err = global.GVA_DB.Delete(&moneyAccountRecord).Error; err != nil {
		return
	}
	return
}

// DeleteMoneyAccountRecordByIds 批量删除MoneyAccountRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyAccountRecordService *MoneyAccountRecordService) DeleteMoneyAccountRecordByIds(ids request.IdsReq) (err error) {
	if err = global.GVA_DB.Where("id IN ?", ids.Ids).Delete(&money.MoneyAccountRecord{}).Error; err != nil {
		return
	}
	return
}

// UpdateMoneyAccountRecord 更新MoneyAccountRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyAccountRecordService *MoneyAccountRecordService) UpdateMoneyAccountRecord(moneyAccountRecord money.MoneyAccountRecord) (err error) {
	err = global.GVA_DB.Updates(&moneyAccountRecord).Error
	return
}

// GetMoneyAccountRecord 根据id获取MoneyAccountRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyAccountRecordService *MoneyAccountRecordService) GetMoneyAccountRecord(id uint) (moneyAccountRecord money.MoneyAccountRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).Find(&moneyAccountRecord).Error
	return
}

// GetMoneyAccountRecordInfoList 分页获取MoneyAccountRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyAccountRecordService *MoneyAccountRecordService) GetMoneyAccountRecordInfoList(info moneyReq.MoneyAccountRecordSearch) (list []*money.MoneyAccountRecord, total int64, err error) {

	// 创建db
	db := global.GVA_DB.Model(&money.MoneyAccountRecord{}).Where("created_by = ?", info.CreatedBy)
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.InOut != nil {
		db = db.Where("in_out = ?", info.InOut)
	}
	if info.CategoryId != nil {
		db = db.Where("category_id = ?", info.CategoryId)
	}
	if info.StartOccurTime != nil && info.EndOccurTime != nil {
		db = db.Where("occur_time BETWEEN ? AND ? ", info.StartOccurTime, info.EndOccurTime)
	}
	if info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+info.Remark+"%")
	}

	if info.Page > 0 && info.PageSize > 0 {
		db = db.Limit(info.PageSize).Offset((info.Page - 1) * info.PageSize)
	}

	if err = db.Order("id DESC").Find(&list).Limit(-1).Count(&total).Offset(-1).Error; err != nil {
		return
	}
	return
}

// GetMoneyAccountRecordSummary 根据id获取概况
func (moneyAccountRecordService *MoneyAccountRecordService) GetMoneyAccountRecordSummary(userId uint) (result money.MoneyAccountRecordSummary, err error) {
	sql := `SELECT
	ROUND(SUM(IF(in_out = 1, money, 0)), 2) AS sum_in,
	ROUND(SUM(IF(in_out = 2, money, 0)), 2) AS sum_out
	FROM money_account_record
	WHERE deleted_at IS NULL AND created_by = ?`
	if err = global.GVA_DB.Raw(sql, userId).Scan(&result).Error; err != nil {
		return
	}
	result.SumBalance = function.Decimal(result.SumIn - result.SumOut)
	return
}
