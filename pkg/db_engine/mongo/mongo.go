package mongo

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoConfig struct {
	DB       string
	Host     string
	SetName  string
	Port     int
	User     string
	Password string
}

func InitMongoEngine(config *MongoConfig) {
	var uri string
	if config.User != "" {
		uri = fmt.Sprintf("mongodb://%v:%v@%v/?replicaSet=%v&authSource=admin", config.User, config.Password, config.Host, config.SetName)
	} else {
		uri = fmt.Sprintf("mongodb://%v/?replicaSet=%v&authSource=admin", config.Host, config.SetName)
	}
	fmt.Println("uri", uri)
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: time.Second * 2}, config.DB, options.Client().ApplyURI(uri).SetMaxPoolSize(30).SetMaxConnIdleTime(time.Second*10))
	if err != nil {
		errMsg := "mongo connect error:" + err.Error()
		panic(errMsg)
	}
}
