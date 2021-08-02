package redisconnection

import (

	dotenv "cacheisking.com/utils"
	redis "github.com/go-redis/redis/v8"
)

var RDB = redis.NewClient(&redis.Options {
	Addr: dotenv.GoDotEnvVariable("REDISHOST"),
	Password: dotenv.GoDotEnvVariable("REDISPASSWORD"),
	DB: 0,
})

