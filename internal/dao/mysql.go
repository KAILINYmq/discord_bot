package dao

import (
	"DiscordRolesBot/internal/model"
	"DiscordRolesBot/pkg/db_engine/sql"
	"errors"
)

func GetAllAccount(account string) (mainAccount string, data []string, err error) {
	var user model.User
	var subAccount model.SubAccount
	user.Address = account
	userData, err := user.FindUserByAddress(sql.GetMarketDbEngine())
	if err != nil {
		return mainAccount, data, err
	}
	if userData == nil {
		subAccountData, err := subAccount.GetSubAccount(sql.GetMarketDbEngine(), account)
		if err != nil {
			return mainAccount, data, err
		}
		if subAccountData == nil {
			return mainAccount, data, errors.New("no such account")
		}
		user.Address = subAccountData.MainAccount
	}
	userData, err = user.FindUserByAddress(sql.GetMarketDbEngine())
	if err != nil {
		return mainAccount, data, err
	}
	subAccounts, err := subAccount.GetAllSubAccounts(sql.GetMarketDbEngine(), userData.Address)
	if err != nil {
		return mainAccount, data, err
	}
	mainAccount = user.Address
	data = append(data, userData.Address)
	for _, v := range subAccounts {
		data = append(data, v.SubAccount)
	}
	return mainAccount, data, err
}
