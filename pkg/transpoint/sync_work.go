package transpoint

import (
	"DiscordRolesBot/global"
	"github.com/shopspring/decimal"
	"time"
)

type TransLogStruct struct {
	OwnerAddress string
	TradeId      string
	GameId       string
	TradeType    global.UTSourceType
	TradeItem    string
	Quantity     decimal.Decimal
	Current      decimal.Decimal
	TxTime       time.Time
	TxHash       string
	Desc         string
}

type StageClearStruct struct {
	OwnerAddress string
	Contract     string
	TokenId      int64
	GameId       string
	Difficulty   global.DifficultType
	Chapter      int64
	Stage        int64
	LogTime      time.Time
	//Tokens       string
	//CurrentToken string
}

type NftEventInfo struct {
	NftAddress  string
	Contract    string
	TokenId     int64
	Event       string
	Price       string
	FromAddress string
	ToAddress   string
	Date        time.Time
	Tx          string
}

type GameTokensEvent struct {
	TransId       string
	WalletAddress string
	Amount        decimal.Decimal
	Signature     string
	SubmitTime    int
}

type ReportTypeEnum int

var (
	Set     ReportTypeEnum = 0
	SetOnce ReportTypeEnum = 1
	UserSet ReportTypeEnum = 2
)

type MetricReportInfo struct {
	Topic        string                 `json:"topic"`
	ReportType   ReportTypeEnum         `json:"report_type"`
	ReportTime   string                 `json:"report_time"`
	TokenId      int64                  `json:"token_id"`
	OwnerAddress string                 `json:"owner_address"`
	Properties   map[string]interface{} `json:"properties"`
}
