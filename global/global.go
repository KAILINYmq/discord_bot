package global

import (
	"DiscordRolesBot/pkg/settings"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	Logger    *logrus.Entry
	Config    *settings.Config
	SkillInfo = map[string]*SkillBaseInfo{}

	TopicTime map[string]time.Time
)

type SkillBaseInfo struct {
	SkillName   string `json:"skill_name"`
	SkillNameCN string `json:"skill_name_cn"`
	SkillType   string `json:"skill_type"`
	SkillDesc   string `json:"skill_desc"`
	SkillDescCN string `json:"skill_desc_cn"`
	SkillImage  string `json:"skill_image"`
	SkillSchool string `json:"skill_school"`
}

type GameId string

func (gd GameId) String() string {
	return string(gd)
}

const (
	ListEvent     = "list"
	BuyEvent      = "buy"
	SaleEvent     = "sale"
	CancelEvent   = "cancel"
	TransferEvent = "transfer"
	WithdrawEvent = "Withdraw"
	DepositEvent  = "Deposit"
)

type UTSourceType string
type DifficultType string

const (
	NormalLevel    DifficultType = "normal"
	NightmareLevel DifficultType = "nightmare"
)

const (
	TaskClear UTSourceType = "task clear"

	Recharge UTSourceType = "recharge"
	Gifts    UTSourceType = "gifts"
	Claim    UTSourceType = "claim"
	Consume  UTSourceType = "consume"
	Bonus    UTSourceType = "bonus"
)

func (u UTSourceType) String() string {
	return string(u)
}

var (
	MRM GameId = "MRM"
	BOM GameId = "BOM"

	Genesis = "Genesis"

	Vida = "Vida"
	Nova = "Nova"
	Xeon = "Xeon"
)

const (
	DbActionQuery  = "query"
	DbActionUpdate = "update"
	DbActionDelete = "delete"
	DbActionCreate = "create"
	UserAddress    = "metamask_solana"
	UserEmail      = "email"
)

func LogDbMessage(table string, action string, message string) {
	msg := fmt.Sprintf("table: %v, action: %v, error: %v", table, action, message)
	Logger.Error(msg)
}
