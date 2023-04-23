package settings

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Base     *BaseCfg
	Log      *LogCfg
	Mongo    *MongoCfg
	MarketDb *DBCfg
	SsoDb    *DBCfg
	Kafka    *KafkaConf
	Redis    *RedisConf
	Syncer   *SyncerCfg
	Sso      *SSO
	Discord  *Discord
}

type DBCfg struct {
	Host         string
	Port         int
	UserName     string
	Password     string
	DBName       string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  int
}

type BaseCfg struct {
	Port               string
	Model              string
	Env                string
	Feishu             string
	Recommend          string
	SSO                string
	User               string
	Password           string
	MirrorContract     string
	MirrorGen2Contract string
}

type LogCfg struct {
	Path    string
	MaxSize int
	MaxAge  int
	Backups int
}

type MongoCfg struct {
	Host     string
	Port     int
	SetName  string
	UserName string
	Password string
	DBName   string
}

type KafkaConf struct {
	Broker           string
	OwnerChangeTopic string
	EventTopic       string
	NftLevelTopic    string
	SolanaTopic      string
	Group            string
}

type RedisConf struct {
	Host     string
	UserName string
	Password string
	DB       int
}

type SyncerCfg struct {
	URL string
}

type SSO struct {
	URL string
}

type Discord struct {
	GuildID  string
	BotToken string
}

var cfg = Config{}

func InitConfig(path string) *Config {
	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		panic("load config error: " + err.Error() + path)
	}
	return &cfg
}
