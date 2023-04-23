package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"reflect"
	"time"
)

var cli *redis.Client

type RedisConfig struct {
	Addr     string
	Username string
	Pwd      string
	Db       int
}

func InitRedis(conf *RedisConfig) {
	cli = redis.NewClient(&redis.Options{
		Addr:        conf.Addr,
		Username:    conf.Username,
		Password:    conf.Pwd,
		DB:          conf.Db,
		PoolSize:    10,
		IdleTimeout: time.Minute * 1,
	})

	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		panic("redis connect error: " + err.Error())
	}
}

func Get(key string) (interface{}, error) {
	result, err := cli.Get(context.Background(), key).Result()
	return result, err
}

func Set(key string, value interface{}, timeout time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return cli.Set(context.Background(), key, data, timeout).Err()
}

func Delete(key string) error {
	return cli.Del(context.Background(), key).Err()
}

func SAdd(key string, values interface{}) error {
	mv := make([]interface{}, 0)

	v := reflect.ValueOf(values)
	if v.Kind() != reflect.Slice {
		return errors.New("not a interface slice")
	}

	l := v.Len()

	for i := 0; i < l; i++ {
		_v := v.Index(i).Interface()
		vv, err := json.Marshal(&_v)
		if err != nil {
			return err
		}
		mv = append(mv, vv)
	}

	return cli.SAdd(context.Background(), key, mv...).Err()
}

func SRandMember(key string, number int64) ([]string, error) {
	return cli.SRandMemberN(context.Background(), key, number).Result()
}

func HMSet(key string, values interface{}) error {
	v := reflect.ValueOf(values)
	if v.Kind() != reflect.Slice {
		return errors.New("not a slice")
	}
	if v.Len()%2 != 0 {
		return errors.New("slice member error")
	}
	return cli.HMSet(context.Background(), key, values).Err()
}

func HMGet(key string, field ...string) ([]interface{}, error) {
	return cli.HMGet(context.Background(), key, field...).Result()
}

func ZAdd(key string, members ...*redis.Z) *redis.IntCmd {
	return cli.ZAdd(context.Background(), key, members...)
}

func ZRank(key string, member string) *redis.IntCmd {
	return cli.ZRank(context.Background(), key, member)
}

func ZRevRank(key string, member string) *redis.IntCmd {
	return cli.ZRevRank(context.Background(), key, member)
}

func CloseRedisConnection() {
	cli.Close()
}
