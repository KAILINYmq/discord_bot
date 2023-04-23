package model

import (
	"gorm.io/gorm"
)

type SubAccount struct {
	gorm.Model
	SubAccount        string `json:"sub_account" gorm:"uniqueIndex"`
	MainAccount       string `json:"main_account" gorm:"uniqueIndex:account_id"`
	SubAccountId      int    `json:"sub_account_id" gorm:"uniqueIndex:account_id"` // 子账户的 id，1 ～ 300
	AccountUserId     int64  `json:"account_user_id"`                              // 根据此 id 去 sso 查询子账户信息
	MainUserId        int64  `json:"main_user_id"`
	Nickname          string `json:"nickname"`
	Email             string `json:"email"`
	EnableConsumption []byte `json:"enable_consumption" gorm:"type:BLOB"`
}

func (sa *SubAccount) GetMaxAccountId(db *gorm.DB, mainAccount string) (int, error) {
	var result SubAccount
	err := db.Model(&SubAccount{}).Where("main_account = ?", mainAccount).Order("sub_account_id DESC").First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	return result.SubAccountId, nil
}

func (sa *SubAccount) GetSubAccountToMain(db *gorm.DB, ownerAddress, walletAddress string) (*SubAccount, error) {
	var result SubAccount
	err := db.Model(&SubAccount{}).Where("sub_account = ? and main_account = ?", walletAddress, ownerAddress).First(&result).Error
	return &result, err
}

func (sa *SubAccount) GetSubAccount(db *gorm.DB, walletAddress string) (*SubAccount, error) {
	var result SubAccount
	err := db.Model(&SubAccount{}).Where("sub_account = ?", walletAddress).First(&result).Error
	return &result, err
}

func (sa *SubAccount) SelectToSubAccountId(db *gorm.DB, ownerAddress string, id int) error {
	return db.Model(&SubAccount{}).Where("sub_account_id = ? and main_account = ?", id, ownerAddress).Find(&sa).Error
}

func (sa *SubAccount) SelectBatch(db *gorm.DB, address []string, ownerAddress string) ([]SubAccount, error) {
	var results []SubAccount
	err := db.Model(&SubAccount{}).Where("sub_account in (?) and main_account = ?", address, ownerAddress).Find(&results).Error
	return results, err
}

func (sa *SubAccount) SelectBatchToUserId(db *gorm.DB, userId []int64) ([]SubAccount, error) {
	var results []SubAccount
	err := db.Model(&SubAccount{}).Where("account_user_id in (?)", userId).Find(&results).Error
	return results, err
}

func (sa *SubAccount) GetSubAccountsToMainAddress(db *gorm.DB, mainAddress string, pageSize, limit int) ([]SubAccount, error) {
	var data []SubAccount
	err := db.Model(&SubAccount{}).Where("main_account = ?", mainAddress).Offset(pageSize).Limit(limit).Find(&data).Error
	return data, err
}

func (sa *SubAccount) GetAllSubAccounts(db *gorm.DB, mainAccount string) ([]SubAccount, error) {
	var result []SubAccount
	err := db.Model(&SubAccount{}).Where("main_account = ?", mainAccount).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
