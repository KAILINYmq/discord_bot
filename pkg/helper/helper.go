package helper

import (
	"DiscordRolesBot/global"
	"DiscordRolesBot/internal/model"
	"DiscordRolesBot/pkg/transpoint"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type EthPriceResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  EthMsg `json:"result"`
}

type EthMsg struct {
	Ethbtc          string `json:"ethbtc"`
	EthbtcTimestamp string `json:"ethbtc_timestamp"`
	Ethusd          string `json:"ethusd"`
	EthusdTimestamp string `json:"ethusd_timestamp"`
}

type Recommend struct {
	Code    int     `json:"code"`
	Data    []int64 `json:"data"`
	Message string  `json:"message"`
}

func IsValidUsername(username string) bool {
	// [] 表示接受的字符 * 表示长度无限制 或者使用 {m, n} -> 表示长度 [m, n]
	ok, _ := regexp.Match("^[a-zA-Z0-9_-]*$", []byte(username))
	return ok
}

func RandomUsername() string {
	rand.Seed(time.Now().UnixNano())
	return "Mirror Collector #" + strconv.Itoa(rand.Intn(99999)+1)
}

// page size 参数是否合理
func ValidPage(totalCount, page, pageSize int64) (bool, int64) {
	if pageSize <= 0 || page <= 0 {
		return false, 0
	}
	pageNum := totalCount / pageSize
	if totalCount%pageSize != 0 {
		pageNum += 1
	}
	if pageNum < page && pageNum > 0 {
		return false, 0
	}
	return true, pageNum

}

func LowerEthAddress(address string) string {
	if strings.HasPrefix(address, "0x") {
		return strings.ToLower(address)
	}
	return address
}

func GetCurrentEthPrice() (bool, string) {
	url := "http://api.etherscan.io/api?module=stats&action=ethprice"

	reqOption := grequests.RequestOptions{
		DialTimeout:         10 * time.Second,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	response, err := grequests.DoRegularRequest(http.MethodGet, url, &reqOption)
	if err != nil {
		fmt.Println("do request error:", err.Error())
		return false, ""
	}

	body, err := ioutil.ReadAll(response.RawResponse.Body)
	defer response.RawResponse.Body.Close()

	var result EthPriceResp

	json.Unmarshal(body, &result)

	return true, result.Result.Ethusd
}

func GetUsdtPrice(price string) string {
	ok, ethusdt := GetCurrentEthPrice()
	if !ok {
		return ""
	}

	priceDeci, _ := decimal.NewFromString(price)
	ethusdtDeci, _ := decimal.NewFromString(ethusdt)

	return priceDeci.Mul(ethusdtDeci).String()
}

func MirrorRecommendNft(gameId, tokenId string) (bool, []int64) {
	// TODO 推荐系统那边写死了是 mce， 需要做下处理
	if strings.ToLower(gameId) == "mrm" {
		gameId = "MCE"
	}
	url := global.Config.Base.Recommend

	reqOption := grequests.RequestOptions{
		DialTimeout:         5 * time.Second,
		TLSHandshakeTimeout: 2 * time.Second,
		Headers: map[string]string{
			"Sso-Token": "kHlCbGpZKkJXZEmmy9Og3zEiKRyeEZDh",
		},
		Params: map[string]string{
			"game_id":  gameId,
			"token_id": tokenId,
		},
	}

	response, err := grequests.DoRegularRequest(http.MethodGet, url, &reqOption)
	if err != nil || response.StatusCode != http.StatusOK {
		return false, nil
	}

	body, err := ioutil.ReadAll(response.RawResponse.Body)
	if err != nil {
		return false, nil
	}
	defer response.RawResponse.Body.Close()
	var result Recommend
	json.Unmarshal(body, &result)

	return true, result.Data
}

func ParseSkill(skillList []model.SkillDesc, gameId string, grade model.GRADE_TYPE) []transpoint.Skill {
	var skills []transpoint.Skill
	var locked bool
	// 不同品级解锁不同的被动技能数
	gradeInt := grade.Int()
	if gradeInt > 5 || gradeInt <= 0 {
		return nil
	}

	for i, s := range skillList {
		info := global.SkillInfo[fmt.Sprintf("%v_%v", gameId, s.SkillId)]
		if gameId == global.MRM.String() {
			locked = false
		} else {
			if gradeInt-1 >= i {
				locked = false
			} else {
				locked = true
			}
		}
		var image, desc, skillName string
		if info != nil {
			image = info.SkillImage
			desc = info.SkillDesc
			skillName = info.SkillName
		}
		skills = append(skills, transpoint.Skill{
			SkillId:   s.SkillId,
			SkillName: skillName,
			SkillType: s.SkillType.String(),
			ImagePath: image,
			SkillDesc: desc,
			Locked:    locked,
		})
	}
	return skills
}

func GetNftSkill(gameInfo []model.GameProperty, gameId string) []transpoint.RecommendSkill {
	var game *model.GameProperty

	for _, v := range gameInfo {
		if v.GameID == gameId {
			game = &v
			break
		}
	}

	if game == nil {
		return nil
	}

	skills := ParseSkill(game.Skills, gameId, game.Grade)

	var result []transpoint.RecommendSkill

	for _, skill := range skills {
		result = append(result, transpoint.RecommendSkill{
			SkillName: skill.SkillName,
			SkillType: skill.SkillType,
			ImagePath: skill.ImagePath,
			Locked:    skill.Locked,
		})
	}

	return result
}

func GetGameGradeInfo(gameInfo []model.GameProperty) []transpoint.GameGrade {
	result := make([]transpoint.GameGrade, 0)
	for _, g := range gameInfo {
		result = append(result, transpoint.GameGrade{
			GameId: g.GameID,
			Grade:  g.Grade.Int(),
		})
	}
	return result
}

func copyWords(number float64, words string) string {
	if number == 1 {
		return "a " + words + " ago"
	}
	return fmt.Sprintf("%v %ss ago", number, words)
}

func GetGradeIcon(gameId string, rarity model.RARITY_TYPE, grade model.GRADE_TYPE) string {
	return fmt.Sprintf("https://xxxxxx/%v/grade/%v_%v.svg", gameId, rarity.String(), grade.Int())
}

func GetNextGradeInfo(gameId string, rarity model.RARITY_TYPE, grade model.GRADE_TYPE) (int, string) {
	nextGrade := GetNextGrade(grade)
	if nextGrade != nil {
		nextGradeStr := nextGrade.Int()
		nextGradeIcon := GetGradeIcon(gameId, rarity, *nextGrade)
		return nextGradeStr, nextGradeIcon
	}
	return 0, ""
}

func GetMRMStatusIcon(name string) string {
	iconPath := "https://xxxxxxxx.fun/MRM/market_icon"
	icons := map[string]string{
		"attack":   iconPath + "/attack.svg",
		"defences": iconPath + "/defence.svg",
		"blood":    iconPath + "/blood.svg",
	}
	return icons[name]
}

func ValidGrade(grade int) *model.GRADE_TYPE {
	if grade > 5 || grade < 1 {
		return nil
	}
	gt := model.GRADE_TYPE(grade)
	return &gt
}

func GradeMaxLevel(grade model.GRADE_TYPE) int64 {
	gradeLevel := []int64{80, 100, 150, 180, 200}

	if grade.Int() > len(gradeLevel) {
		return 0
	}
	return gradeLevel[grade-1]
}

func GetNextGrade(currentGrade model.GRADE_TYPE) *model.GRADE_TYPE {
	nextGrade := currentGrade.Int() + 1
	if nextGrade > 0 && nextGrade <= 5 {
		ng := model.GRADE_TYPE(nextGrade)
		return &ng
	}
	return nil
}

func ConvertMongoDecimal(value primitive.Decimal128) string {
	if value.IsZero() {
		return "0"
	}
	return value.String()
}

func Min(v1 int64, v2 int64) int64 {
	if v1 <= v2 {
		return v1
	}
	return v2
}

func getSkillName(gameId string, skills []model.SkillDesc) []string {
	skillNames := make([]string, 0)
	for _, v := range skills {
		info := global.SkillInfo[fmt.Sprintf("%v_%v", gameId, v.SkillId)]
		skillNames = append(skillNames, info.SkillName)
	}
	return skillNames
}

func PullNftProperty(gameProperty []model.GameProperty) []transpoint.PullNftProperty {
	result := make([]transpoint.PullNftProperty, 0)
	for _, v := range gameProperty {
		result = append(result, transpoint.PullNftProperty{
			GameId: v.GameID,
			Grade:  v.Grade.Int(),
			//AttackModel: v.AttackModel,
			//Skills:      getSkillName(v.GameID, v.Skills),
		})
	}
	return result
}

func GetConsumption(data []byte) bool {
	if len(data) > 0 && data[0] == 1 {
		return true
	} else {
		return false
	}
}
