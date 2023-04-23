package model

import (
	"DiscordRolesBot/global"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        `json:"-"`
	Username          string `json:"username" gorm:"unique, not null"`
	LowerCaseUsername string `json:"lower_case_username" gorm:"unique, not null"` // 不区分大小写的用户名，用来查重
	Address           string `json:"address" gorm:"unique, not null"`
	Avatar            string `json:"avatar"`
}

type UserInCache struct {
	Username string `json:"username"`
	Address  string `json:"address"`
	Avatar   string `json:"avatar"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) FindUserByAddress(db *gorm.DB) (*User, error) {
	var user User
	err := db.Model(&User{}).Where("address = ?", u.Address).First(&user).Error
	if err != nil {
		global.LogDbMessage(u.TableName(), global.DbActionQuery, err.Error())
		return nil, err
	}
	return &user, nil
}

func (u *User) IsDuplicateUsername(db *gorm.DB) bool {
	user := &User{}
	err := db.Model(u).Where("lower_case_username = ?", u.LowerCaseUsername).First(user).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (u *User) GetUserCount(db *gorm.DB) int64 {
	var count int64

	err := db.Model(&User{}).Count(&count).Error
	if err != nil {
		return count
	}
	return count
}

func (u *User) QueryUserNameBatch(db *gorm.DB, page, batch int) ([]User, error) {
	result := make([]User, 0)

	err := db.Model(&User{}).Order("id").Limit(batch).Offset((page - 1) * batch).Find(&result).Error
	return result, err
}
