package model

import (
	"DiscordRolesBot/global"
	"gorm.io/gorm"
)

type SsoUser struct {
	ID            uint   `gorm:"primarykey"`
	Nonce         string `json:"nonce"`
	EthAddress    string `json:"eth_address"`
	SolAddress    string `json:"sol_address"`
	Email         string `json:"email"`
	EmailVerified byte   `json:"email_verified"`
	UserName      string `json:"user_name"`
	PassWord      string `json:"pass_word"`
	GoogleId      string `json:"google_id"`
	AppleId       string `json:"apple_id"`
	FaceBookId    string `json:"facebook_id"`
	MainUserId    int    `json:"main_user_id"`
	AllowSpend    byte   `json:"allow_spend"`
	DiscordId     string `json:"discord_id"`
}

func (u *SsoUser) TableName() string {
	return "users"
}

func (u *SsoUser) FindUserByAddress(db *gorm.DB) (*SsoUser, error) {
	var user SsoUser
	err := db.Model(&SsoUser{}).Where("sol_address = ?", u.SolAddress).First(&user).Error
	if err != nil {
		global.LogDbMessage(u.TableName(), global.DbActionQuery, err.Error())
		return nil, err
	}
	return &user, nil
}

func (u *SsoUser) GetCount(db *gorm.DB) (int64, error) {
	var count int64 = 0
	err := db.Model(&SsoUser{}).Count(&count).Error
	if err != nil {
		global.LogDbMessage(u.TableName(), global.DbActionQuery, err.Error())
		return count, err
	}
	return count, nil
}

func (u *SsoUser) FindUserSlice(db *gorm.DB, min int64, max int) ([]*SsoUser, error) {
	var users []*SsoUser
	err := db.Model(&SsoUser{}).Where("id > ?", min).Order("id").Limit(max).Find(&users).Error
	if err != nil {
		global.LogDbMessage(u.TableName(), global.DbActionQuery, err.Error())
		return nil, err
	}
	return users, nil
}
