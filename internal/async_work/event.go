package async_work

import (
	"DiscordRolesBot/global"
	"DiscordRolesBot/internal/dao"
	"DiscordRolesBot/internal/model"
	"DiscordRolesBot/pkg/db_engine/sql"
	discord "DiscordRolesBot/pkg/discord"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var UserTableLength int64 = 0

func eventTopic(message *kafka.Message) {
	fmt.Println("event nft:", string(message.Value))
	var msg map[string]interface{}

	err := json.Unmarshal(message.Value, &msg)
	if err != nil {
		errMsg := fmt.Sprintf("kafka message unmarshal error, topic: %v, error: %v", global.Config.Kafka.EventTopic, err.Error())
		global.Logger.Error(errMsg)
		return
	}
	if msg["nft"] == nil {
		fmt.Println("not event")
		return
	}
	receiptType := msg["receipt_type"]

	switch receiptType {
	case "listing_receipt":
		listNft(msg)
	case "purchase_receipt":
		purchaseNft(msg)
	default:
		fmt.Println("unknown message: ", msg)
	}
}

func listNft(message map[string]interface{}) {
	defer func() {
		err := recover()
		if err != nil {
			global.Logger.Error(fmt.Sprintf("list nft panic: %v", err))
		}
	}()
	fmt.Println("list nft:", message)
	err := AddRoles(message["seller"].(string))
	fmt.Println(message["seller"].(string))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func purchaseNft(message map[string]interface{}) {
	defer func() {
		err := recover()
		if err != nil {
			global.Logger.Error(fmt.Sprintf("purchase panic: %v", err))
		}
	}()
	fmt.Println("cancel list nft:", message)
	err := AddRoles(message["buyer"].(string))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func AddRoles(address string) error {
	mainAccount, accounts, err := dao.GetAllAccount(address)
	if err != nil {
		fmt.Println("err ", err.Error())
		return err
	}

	u := model.SsoUser{}
	u.SolAddress = mainAccount
	user, err := u.FindUserByAddress(sql.GetSsoDbEngine())
	if err != nil {
		fmt.Println("err ", err.Error())
		return err
	}

	NftNums, err := dao.GetAllNFTNum(accounts)
	if err != nil {
		fmt.Println("err ", err.Error())
		fmt.Println(err.Error())
	}

	err = discord.RoleMatch(NftNums, user.DiscordId)
	if err != nil {
		fmt.Println("err ", err.Error())
	}
	return err
}

func ScheduledJob() {
	var user model.SsoUser
	count, err := user.GetCount(sql.GetSsoDbEngine())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if UserTableLength > count {
		UserTableLength = 1
	}

	users, err := user.FindUserSlice(sql.GetSsoDbEngine(), UserTableLength, 100)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range users {
		if v != nil && len(v.SolAddress) > 0 && len(v.DiscordId) > 0 {
			if err = AddRoles(v.SolAddress); err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	UserTableLength += int64(len(users))
	fmt.Println("ScheduledJob: user id ", UserTableLength)
}
