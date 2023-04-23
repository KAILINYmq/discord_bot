package discord

import (
	"DiscordRolesBot/global"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type search [][]int

var (
	Discord *discordgo.Session
	Roles   map[string]*discordgo.Role
	Match   map[string]search //  todo match 的增删改查
)

// 匹配规则： 二维数组里的每个元素代表一个规则
// 例：{2, 1}需要1个二代, {1, 3}需要3个一代
var (
	GenTwoMirrorCollector search = [][]int{{2, 1}}
	GenOneMirrorCollector search = [][]int{{1, 1}}
	MirrorKnight          search = [][]int{{1, 1}, {2, 1}}
	MirrorPaladin         search = [][]int{{1, 3}, {2, 5}}
	MirrorWhale           search = [][]int{{1, 10}, {2, 15}}
)

func init() {
	Roles = make(map[string]*discordgo.Role, 10)
	Match = make(map[string]search, 10)
	Match["Gen 2 Collector"] = GenTwoMirrorCollector
	Match["Gen 1 Collector"] = GenOneMirrorCollector
	Match["Knight"] = MirrorKnight
	Match["Paladin"] = MirrorPaladin
	Match["Whale"] = MirrorWhale
}

func NewDiscord() error {
	var err error
	Discord, err = discordgo.New("Bot " + global.Config.Discord.BotToken)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	Discord.AddHandler(ready)
	Discord.AddHandler(messageCreate)
	err = Discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return err
	}
	return err
}

func GetAllRoles() error {
	roles, err := Discord.GuildRoles(global.Config.Discord.GuildID)
	if err != nil {
		return err
	}
	for _, v := range roles {
		Roles[v.Name] = v
	}

	return err
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func ready(s *discordgo.Session, m *discordgo.Ready) {
	// 如果Bot上线更改游戏状态
	s.UpdateGameStatus(0, "LOL~")
}

func RoleMatch(nftNums map[int]int64, DiscordUserID string) error {
	AddRoleIDs := make(map[string]string, 5)
	RemoveRoleIDs := make(map[string]string, 5)
	if len(DiscordUserID) == 0 {
		return errors.New("DiscordUserID is null")
	}
	if err := GetAllRoles(); err != nil {
		return err
	}

	for k, v := range Match {
		for _, value := range v {
			if int64(value[1]) <= nftNums[value[0]] && Roles[k] != nil {
				fmt.Println("need nft num", value[1])
				fmt.Println("nft num", nftNums[value[0]])
				fmt.Println("need ", Roles[k].Name)
				AddRoleIDs[Roles[k].ID] = Roles[k].Name
			}
		}
		if _, ok := AddRoleIDs[Roles[k].ID]; !ok {
			fmt.Println("don't need ", Roles[k].Name)
			RemoveRoleIDs[Roles[k].ID] = Roles[k].Name
		}
	}
	fmt.Println(len(AddRoleIDs))
	fmt.Println(len(RemoveRoleIDs))

	for k, _ := range AddRoleIDs {
		if err := Discord.GuildMemberRoleAdd(global.Config.Discord.GuildID, "699183924684783000", k); err != nil {
			fmt.Println(err.Error())
		}
	}
	for k, _ := range RemoveRoleIDs {
		if err := Discord.GuildMemberRoleRemove(global.Config.Discord.GuildID, DiscordUserID, k); err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}
