package sql

import (
	"DiscordRolesBot/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

var (
	marketDbEngine *gorm.DB
	ssoDbEngine    *gorm.DB
)

type DbConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBName       string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  int
}

func InitDBEngine(config *DbConfig) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
	}), &gorm.Config{})

	fmt.Println("db uri", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName))
	if err != nil {
		panic("connect db error: " + err.Error())
	}
	sqlDb, _ := db.DB()
	if err := sqlDb.Ping(); err != nil {
		panic("DB ping error: " + err.Error())
	}

	sqlDb.SetConnMaxLifetime(time.Minute * time.Duration(config.MaxLifetime))
	sqlDb.SetMaxIdleConns(config.MaxIdleConns)
	sqlDb.SetMaxOpenConns(config.MaxOpenConns)

	if global.Config.Base.Env != "prod" {
		db = db.Debug()
	}
	if strings.HasPrefix(config.DBName, "market") {
		marketDbEngine = db
	} else {
		ssoDbEngine = db
	}
}

func GetMarketDbEngine() *gorm.DB {
	return marketDbEngine
}

func GetSsoDbEngine() *gorm.DB {
	return ssoDbEngine
}
