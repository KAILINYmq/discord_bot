package transpoint

type UpdateNickName struct {
	UUID     string `json:"uuid" binding:"required"`
	Contract string `json:"contract" binding:"required"`
	TokenId  int64  `json:"token_id" binding:"gte=1"`
	Nickname string `json:"nickname" binding:"required"`
}

type UpdateNickNameResp struct {
	Contract string `json:"contract"`
	TokenId  string `json:"token_id"`
	NickName string `json:"nickname"`
}

type NFTInfo struct {
	UUID             string `json:"uuid"`
	Contract         string `json:"contract"`
	TokenId          int64  `json:"token_id"`
	NftAddress       string `json:"nft_address"`
	Camp             string `json:"camp"`
	Energy           string `json:"energy"`
	ImagePath        string `json:"image_path"`
	Rarity           string `json:"rarity"`
	NickName         string `json:"nickname"`
	Favorites        int    `json:"favorites"`
	Owner            User   `json:"owner"`
	Price            string `json:"price"`
	USDPrice         string `json:"usd_price"`
	Listed           bool   `json:"listed"`
	InOwnerFavorites bool   `json:"in_owner_favorites"`
}

type User struct {
	OwnerAddress string `json:"owner_address"`
	Avatar       string `json:"avatar"`
	Username     string `json:"username"`
}

type Price struct {
	EthPrice  string `json:"eth_price"`
	UsdtPrice string `json:"usdt_price"`
}

type NFTRecommendResp struct {
	Recommend []MirrorBaseInfo `json:"recommend"`
}

type MirrorBaseInfo struct {
	Contract   string      `json:"contract"`
	TokenId    int64       `json:"token_id"`
	UUID       string      `json:"uuid"`
	NftAddress string      `json:"nft_address"`
	NftType    int         `json:"nft_type"`
	ImagePath  string      `json:"image_path"`
	Price      string      `json:"price"`
	Nickname   string      `json:"nickname"`   // 昵称
	Rarity     string      `json:"rarity"`     // 稀有度
	Favorites  int         `json:"favorites"`  // 收藏数
	Camp       string      `json:"camp"`       // 阵营
	GradeInfo  []GameGrade `json:"grade_info"` // 品级信息
}

type MirrorBaseInfo2 struct {
	MirrorBaseInfo
	InOwnerFavorite bool `json:"in_owner_favorite"`
}

type MirrorBaseInfoList []MirrorBaseInfo

func (mb MirrorBaseInfoList) MirrorOrEquipment() string {
	return "mirror"
}

type MirrorBaseInfoList2 []MirrorBaseInfo2

func (mb MirrorBaseInfoList2) MirrorOrEquipment() string {
	return "mirror"
}

type GameGrade struct {
	GameId string `json:"game_id"`
	Grade  int    `json:"grade"`
}

type RecommendSkill struct {
	SkillName string `json:"skill_name"`
	SkillType string `json:"skill_type"`
	ImagePath string `json:"image_path"`
	Locked    bool   `json:"locked"`
}

type MirrorGameProperty struct {
	Contract     string `json:"contract"`
	TokenId      int64  `json:"token_id"`
	CurrentLevel int64  `json:"current_level"`
	MaxLevel     int64  `json:"max_level"`
	GameId       string `json:"game_id"`
}

type SkillBaseInfo struct {
	SkillName   string `json:"skill_name"`
	SkillNameCN string `json:"skill_name_cn"`
	SkillType   string `json:"skill_type"`
	SkillDesc   string `json:"skill_desc"`
	SkillDescCN string `json:"skill_desc_cn"`
	SkillImage  string `json:"skill_image"`
}

type NftItemActivity struct {
	Event     string   `json:"event"`
	Price     string   `json:"price"`
	From      UserItem `json:"from"`
	To        UserItem `json:"to"`
	Date      string   `json:"date"`
	DateTag   string   `json:"date_tag"`
	Etherscan string   `json:"etherscan"`
}

type UserItem struct {
	Username     string `json:"username"`
	OwnerAddress string `json:"owner_address"`
}

type OwnerChanged struct {
	Contract     string `json:"contract" binding:"required"`
	TokenId      string `json:"token_id" binding:"required"`
	OwnerAddress string `json:"owner_address" binding:"required"`
}

type SetNftAddress struct {
	Contract   string `json:"contract" binding:"required"`
	TokenId    string `json:"token_id" binding:"required"`
	NftAddress string `json:"nft_address" binding:"required"`
}

type NftActivityResponse struct {
	Contract string            `json:"contract"`
	TokenId  string            `json:"token_id"`
	Activity []NftItemActivity `json:"activity"`
}

type NftUpgrade struct {
	Contract string `json:"contract" binding:"required"`
	TokenId  int64  `json:"token_id" binding:"gte=1"`
	GameId   string `json:"game_id" binding:"required"`
	Grade    int    `json:"grade" binding:"gte=1,lte=5"`
}

type NftUpgradeResp struct {
	Contract   string `json:"contract"`
	TokenId    int64  `json:"token_id"`
	GameId     string `json:"game_id"`
	NftAddress string `json:"nft_address"`
	Grade      int    `json:"grade"`
	GradeIcon  string `json:"grade_icon"`
	Level      int64  `json:"level"`
	MaxLevel   int64  `json:"max_level"`
}

// 技能搜索
type SearchSkillResp struct {
	Skills []SearchSkill `json:"skills"`
}

// 技能搜索信息
type SearchSkill struct {
	GameId    string `json:"game_id"`
	SkillId   int64  `json:"skill_id"`
	SkillName string `json:"skill_name"`
}

type GameFilter struct {
	GameId   string  `json:"game_id"`
	Grade    int     `json:"grade"`
	SkillIds []int64 `json:"skill_ids"`
}

type MirrorFilter struct {
	Sale      int          `json:"sale"`
	Rarity    string       `json:"rarity"`
	GameItems []GameFilter `json:"game_items" binding:"gte=0"`
}

// mirror page selector
type MarketMirrorSelector struct {
	Page     int64        `json:"page" binding:"gte=1"`
	PageSize int64        `json:"page_size" binding:"gte=1"`
	Filter   MirrorFilter `json:"filter"`
	Order    OrderOption  `json:"order"`
}

// equipment 过滤项
type MarketEquipmentFilter struct {
	Sale               int      `json:"sale"`
	Rarity             string   `json:"rarity"`
	Type               string   `json:"type"`
	StrengtheningLevel [2]int64 `json:"strengthening_level"`
}

// equipment page selector
type MarketEquipmentSelector struct {
	Page     int64                 `json:"page" binding:"gte=1"`
	PageSize int64                 `json:"page_size" binding:"gte=1"`
	Filter   MarketEquipmentFilter `json:"filter"`
	Order    OrderOption           `json:"order"`
}

type NFTMirror struct {
	Page      int64            `json:"page"`
	PageSize  int64            `json:"page_size"`
	TotalPage int64            `json:"total_page"`
	Quantity  int64            `json:"quantity"`
	Nfts      []MirrorBaseInfo `json:"nfts"`
}

type MirrorSearch struct {
	Quantity int              `json:"quantity"`
	Nfts     []MirrorBaseInfo `json:"nfts"`
}

// 技能搜索返回
type EquipmentSearch struct {
	Quantity int                 `json:"quantity"`
	Nfts     []EquipmentBaseInfo `json:"nfts"`
}

type EquipmentBaseInfo struct {
	UUID               string `json:"uuid"`
	Contract           string `json:"contract"`
	TokenId            int64  `json:"token_id"`
	NftType            int    `json:"nft_type"`
	Nickname           string `json:"nickname"`
	ImagePath          string `json:"image_path"`
	Price              string `json:"price"`
	Favorites          int    `json:"favorites"`
	StrengtheningLevel int64  `json:"strengthening_level"` // 	强化等级
	Rarity             string `json:"rarity"`
	GameId             string `json:"game_id"`
}

// 游戏属性信息
type GameProperty struct {
	TokenId       string     `json:"token_id"`
	NftRarity     string     `json:"nft_rarity"`
	GameId        string     `json:"game_id"`
	Grade         int        `json:"grade"`
	AttackKind    string     `json:"attack_kind"`
	GradeIcon     string     `json:"grade_icon"`
	NextGrade     int        `json:"next_grade"`
	NextGradeIcon string     `json:"next_grade_icon"`
	Level         int64      `json:"level"`
	MaxLevel      int64      `json:"max_level"`
	Skills        []Skill    `json:"skills"`
	Status        []Status   `json:"status"`
	Honor         *GameHonor `json:"honor"`
}

type Skill struct {
	SkillId   int64  `json:"skill_id"`
	SkillName string `json:"skill_name"`
	SkillType string `json:"skill_type"`
	ImagePath string `json:"image_path"`
	SkillDesc string `json:"skill_desc"`
	Locked    bool   `json:"locked"`
}

type Status struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
	Icon      string      `json:"icon"`
}

type GameHonor struct {
	Mainline  string      `json:"mainline"`
	Playtime  interface{} `json:"playtime"`
	Riftscore interface{} `json:"riftscore"`
	Rank      interface{} `json:"rank"`
}

type NftItem struct {
	Contract   string `json:"contract"`
	TokenId    string `json:"token_id"`
	NftAddress string `json:"nft_address"`
	Nickname   string `json:"nickname"`
	ImagePath  string `json:"image_path"`
}

type FilterOrder struct {
	Page     int64       `json:"page" binding:"gte=1"`
	PageSize int64       `json:"page_size" binding:"gte=1"`
	Filter   int         `json:"filter"`
	Order    OrderOption `json:"order"`
}

type OrderOption struct {
	OrderBy int  `json:"order_by"`
	Desc    bool `json:"desc"`
}

type FilterNfts struct {
	Page         int64                 `json:"page"`
	PageSize     int64                 `json:"page_size"`
	TotalPage    int64                 `json:"total_page"`
	Quantity     int64                 `json:"quantity"`
	OwnerAddress string                `json:"owner_address"`
	NFTS         MirrorOrEquipmentList `json:"nfts"`
}

type FilterNftsMirrorAndEquipment struct {
	Page         int64         `json:"page"`
	PageSize     int64         `json:"page_size"`
	TotalPage    int64         `json:"total_page"`
	Quantity     int64         `json:"quantity"`
	OwnerAddress string        `json:"owner_address"`
	NFTS         []interface{} `json:"nfts"`
}

type MirrorOrEquipmentList interface {
	MirrorOrEquipment() string
}

type HomePageResponse struct {
	FilterNftsMirrorAndEquipment
	UserInfo UserBaseInfo `json:"user_info"`
}

type UserBaseInfo struct {
	OwnerAddress string `json:"owner_address"`
	Username     string `json:"username"`
	Avatar       string `json:"avatar"`
}

type BuyNftResp struct {
	Contract     string `json:"contract"`
	TokenId      int64  `json:"token_id"`
	OwnerAddress string `json:"owner_address"`
}

type PullNftMetadataReq struct {
	TokenId []int64 `json:"token_id"`
}

type PullNftMetadataResp struct {
	Quantity    int               `json:"quantity"`
	NftMetadata []PullNftMetadata `json:"nft_metadata"`
}

type PullNftMetadata struct {
	TokenId      int64             `json:"token_id"`
	NickName     string            `json:"nick_name"`
	ImagePath    string            `json:"image_path"`
	Rarity       string            `json:"rarity"`
	Camp         string            `json:"camp"`
	GameProperty []PullNftProperty `json:"game_property"`
}

type PullNftProperty struct {
	GameId string `json:"game_id"`
	Grade  int    `json:"grade"`
	//AttackModel string   `json:"attack_model"`
	//Skills      []string `json:"skills"`
}

type OnSellResponse struct {
	OnSell              bool      `json:"on_sell"`
	WhitelistOnSell     bool      `json:"whitelist_on_sell"`
	OnSellTime          string    `json:"on_sell_time"`
	WhitelistOnSellTime string    `json:"whitelist_on_sell_time"`
	CountDown           CountDown `json:"count_down"`
}

type CountDown struct {
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}

type WhitelistResponse struct {
	InWhitelist bool `json:"in_whitelist"`
	Quantity    int  `json:"quantity"`
}

type SetSellTimeReq struct {
	Topic     string
	StartTime string
}

type UploadTwitterImage struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type Mirror2ExchangeMetadata struct {
	SrcTokenId int64 `json:"src_token_id"`
	DstTokenId int64 `json:"dst_token_id"`
}

type UploadS3 struct {
	Data         string `json:"data"`
	OwnerAddress string `json:"owner_address"`
	TokenId      int64  `json:"token_id"`
}
