package model

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type NFTType int
type SkillType string
type GRADE_TYPE int
type RARITY_TYPE string
type CAMP_TYPE string

const (
	MIRROR_NFT    NFTType = 1
	EQUIPMENT_NFT NFTType = 3
	PULL_NFT      NFTType = 2

	ACTIVE_SKILL     SkillType = "active skill"
	SUPPORTING_SKILL SkillType = "supporting skill"

	GRADE_1 GRADE_TYPE = 1
	GRADE_2 GRADE_TYPE = 2
	GRADE_3 GRADE_TYPE = 3
	GRADE_4 GRADE_TYPE = 4
	GRADE_5 GRADE_TYPE = 5

	RARITY_COMMOMN   RARITY_TYPE = "common"
	RARITY_RARE      RARITY_TYPE = "rare"
	RARIRT_ELITE     RARITY_TYPE = "elite"
	RARITY_LEGENDARY RARITY_TYPE = "legendary"
	RARITY_MYTHICAL  RARITY_TYPE = "mythical"

	Genesis CAMP_TYPE = "Genesis"
)

func (gt GRADE_TYPE) Int() int {
	return int(gt)
}

func (gt GRADE_TYPE) String() string {
	return strconv.Itoa(int(gt))
}

func (st SkillType) String() string {
	return string(st)
}

func (rt RARITY_TYPE) String() string {
	return string(rt)
}

func (ct CAMP_TYPE) String() string {
	return string(ct)
}

func (nt NFTType) Int() int {
	return int(nt)
}

type NFT struct {
	mgm.DefaultModel `bson:",inline"`
	NftAddress       string               `json:"nft_address" bson:"nft_address"`
	Contract         string               `json:"contract" bson:"contract"`
	TokenId          int64                `json:"token_id" bson:"token_id"`
	Camp             CAMP_TYPE            `json:"camp" bson:"camp"`
	Rarity           RARITY_TYPE          `json:"rarity" bson:"rarity"`
	OwnerAddress     string               `json:"owner_address" bson:"owner_address"`
	AILevel          int64                `json:"ai_level" bson:"ai_level"`
	Price            primitive.Decimal128 `json:"price" bson:"price"`
	ListTime         primitive.DateTime   `json:"list_time" bson:"list_time"`
	Listed           bool                 `json:"listed" bson:"listed"`
	NickName         string               `json:"nickname" bson:"nickname"`
	NftType          NFTType              `json:"nft_type" bson:"nft_type"`
	ImagePath        string               `json:"image_path" bson:"image_path"`
	Favorites        int                  `json:"favorites" bson:"favorites"`
	GameProperty     []GameProperty       `json:"game_property" bson:"game_property"`
	EquipmentInfo    *EquipmentDetail     `json:"equipment_info" bson:"equipment_info"`
}

type NftInCache struct {
	Contract      string
	TokenId       int64
	UUID          string
	NftAddress    string
	NftType       NFTType
	NickName      string
	ImagePath     string
	Camp          string
	Favorites     int
	Price         primitive.Decimal128
	Rarity        RARITY_TYPE
	GameProperty  []GameProperty
	EquipmentInfo *EquipmentDetail
	Expire        int64
}

func (n *NFT) CollectionName() string {
	return "nfts"
}

type Mirrors struct {
	mgm.DefaultModel `bson:",inline"`
	Generation       int            `json:"generation" bson:"generation"`
	TokenId          int64          `json:"token_id" bson:"token_id"`
	Camp             CAMP_TYPE      `json:"camp" bson:"camp"`
	Rarity           RARITY_TYPE    `json:"rarity" bson:"rarity"`
	IsSold           bool           `json:"is_sold" bson:"is_sold"`
	NickName         string         `json:"nickname" bson:"nickname"`
	ImagePath        string         `json:"image_path" bson:"image_path"`
	GameProperty     []GameProperty `json:"game_property" bson:"game_property"`
	Metadata         []Properties   `json:"metadata"`
}

func (m *Mirrors) CollectionName() string {
	return "yarns"
}

type GameProperty struct {
	GameID      string       `json:"game_id" bson:"game_id"`
	GameUUID    string       `json:"game_uuid" bson:"game_uuid"`
	Level       int64        `json:"level" bson:"level"`
	MaxLevel    int64        `json:"max_level" bson:"max_level"`
	AttackModel string       `json:"attack_model" bson:"attack_model"`
	Grade       GRADE_TYPE   `json:"grade" bson:"grade"`
	Property    []Properties `json:"property" bson:"property"`
	Skills      []SkillDesc  `json:"skills" bson:"skills"`
}

type EquipmentDetail struct {
	GameID             string `json:"game_id" bson:"game_id"`
	GameUUID           string `json:"game_uuid" bson:"game_uuid"`
	StrengtheningLevel int64  `json:"strengthening_level" bson:"strengthening_level"`
	EquipmentType      string `json:"equipment_type" bson:"equipment_type"`

	MainProperty []EquipmentProperty `json:"main_property" bson:"main_property"`

	EntryProperty []EquipmentProperty `json:"entry_property" bson:"entry_property"`
	WearLevel     int64               `json:"wear_level" bson:"wear_level"` // 佩戴等级
}

type EquipmentProperty struct {
	TraitType string `json:"trait_type" bson:"trait_type"`
	Value     int64  `json:"value" bson:"value"`
}

type Properties struct {
	TraitType string      `json:"trait_type" bson:"trait_type"`
	Value     interface{} `json:"value" bson:"value"`
}

type SkillDetail struct {
	ActiveSkill     []SkillDesc `json:"active_skill"`
	SupportingSkill []SkillDesc `json:"supporting_skill"`
}

type SkillDesc struct {
	SkillId int64 `json:"skill_id"`
	//SkillName string    `json:"skill_name" bson:"skill_name"`
	SkillType SkillType `json:"skill_type" bson:"skill_type"`
}

type SkillInfo struct {
	gorm.Model
	GameId      string `json:"game_id" gorm:"index:game_skill"`
	SkillId     int64  `json:"skill_id" gorm:"index:game_skill"`
	SkillName   string `json:"skill_name"`
	SkillNameCN string `json:"skill_name_cn"`

	SkillType   string `json:"skill_type"`
	CollingTime string `json:"colling_time"`

	SkillDesc   string `json:"skill_desc" gorm:"type:longtext"`
	SkillDescCN string `json:"skill_desc_cn" gorm:"type:longtext"`
	SkillSchool string `json:"skill_school"`
	SkillImage  string `json:"skill_image"`
	Rarity      string `json:"rarity"`
}

func (s *SkillInfo) GetAllSkill(db *gorm.DB) ([]SkillInfo, error) {
	var result []SkillInfo

	err := db.Model(&SkillInfo{}).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *SkillInfo) SearchSkill(db *gorm.DB, skillName string) ([]SkillInfo, error) {
	var result []SkillInfo

	sql := fmt.Sprintf("select * from skill_infos where LOWER(skill_name) like \"%%%v%%\"", strings.ToLower(skillName))
	err := db.Raw(sql).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
