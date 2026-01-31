package sql

import (
	"database/sql"
	"demo02/global"

	"github.com/redis/go-redis/v9"
)

func ConnextSQL() {
	var err error
	global.Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/myseconddb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

func ConnectRedis() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
