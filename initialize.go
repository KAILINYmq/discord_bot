package main

import (
	"DiscordRolesBot/global"
	"DiscordRolesBot/pkg/db_engine/mongo"
	"DiscordRolesBot/pkg/db_engine/redis"
	"DiscordRolesBot/pkg/db_engine/sql"
	"DiscordRolesBot/pkg/logging"
	"DiscordRolesBot/pkg/settings"
	"fmt"
	"os"
)

func initConfig() {
	// 根据环境变量读取不同的配置文件
	env := os.Getenv("ENV")
	fmt.Println("current env:", env)
	if env == "dev" {
		global.Config = settings.InitConfig("conf/config_dev.toml")
	} else if env == "staging" {
		global.Config = settings.InitConfig("conf/config_staging.toml")
	} else if env == "prod" {
		global.Config = settings.InitConfig("conf/config_prod.toml")
	} else if env == "uat" {
		global.Config = settings.InitConfig("conf/config_uat.toml")
	} else {
		global.Config = settings.InitConfig("conf/config_local.toml")
	}
}

func initLogger() {
	if global.Config.Base.Env == "dev" {
		global.Logger = logging.InitLogger(logging.WithLogPath(global.Config.Log.Path), logging.WithOutput(logging.ONLY_TERMINAL))
	} else {
		// k8s 直接输出到终端
		global.Logger = logging.InitLogger(logging.WithLogPath(global.Config.Log.Path), logging.WithOutput(logging.ONLY_TERMINAL))
	}
}

func initDB() {
	sql.InitDBEngine(&sql.DbConfig{
		Host:         global.Config.MarketDb.Host,
		Port:         global.Config.MarketDb.Port,
		User:         global.Config.MarketDb.UserName,
		Password:     global.Config.MarketDb.Password,
		DBName:       global.Config.MarketDb.DBName,
		MaxIdleConns: global.Config.MarketDb.MaxIdleConns,
		MaxOpenConns: global.Config.MarketDb.MaxOpenConns,
		MaxLifetime:  global.Config.MarketDb.MaxLifetime,
	})
	sql.InitDBEngine(&sql.DbConfig{
		Host:         global.Config.SsoDb.Host,
		Port:         global.Config.SsoDb.Port,
		User:         global.Config.SsoDb.UserName,
		Password:     global.Config.SsoDb.Password,
		DBName:       global.Config.SsoDb.DBName,
		MaxIdleConns: global.Config.SsoDb.MaxIdleConns,
		MaxOpenConns: global.Config.SsoDb.MaxOpenConns,
		MaxLifetime:  global.Config.SsoDb.MaxLifetime,
	})
}

func initMongo() {
	mongo.InitMongoEngine(&mongo.MongoConfig{
		DB:       global.Config.Mongo.DBName,
		Host:     global.Config.Mongo.Host,
		SetName:  global.Config.Mongo.SetName,
		Port:     global.Config.Mongo.Port,
		User:     global.Config.Mongo.UserName,
		Password: global.Config.Mongo.Password,
	})

}

func initRedis() {
	conf := redis.RedisConfig{
		Addr:     global.Config.Redis.Host,
		Username: global.Config.Redis.UserName,
		Pwd:      global.Config.Redis.Password,
		Db:       global.Config.Redis.DB,
	}
	redis.InitRedis(&conf)
}
