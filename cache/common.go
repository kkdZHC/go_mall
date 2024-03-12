package cache

import (
	"strconv"

	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
)

var RedisClient *redis.Client
var (
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		panic("读取配置文件出错")
	}
	LoadRedis(file)
	Redis()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		//Password: RedisPw,
		DB: int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
