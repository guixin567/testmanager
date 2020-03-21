package datasource

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"test1/config"
)

func NewRedis() *redis.Database {
	initConfig := config.InitConfig()
	var database *redis.Database
	if initConfig != nil {
		iris.New().Logger().Info("redis config success")
		redisConfig := initConfig.Redis
		database = redis.New(redis.Config{
			Network:   redisConfig.NetWork,
			Addr:      redisConfig.Host + ":" + redisConfig.Port,
			Password:  redisConfig.Password,
			Database:  "",
			MaxActive: 10,
			Prefix:    redisConfig.Prefix,
		})
	} else {
		iris.New().Logger().Info("redis config failed")
	}

	return database
}
