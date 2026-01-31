package global

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

var Db *sql.DB
var RedisClient *redis.Client

func GetDb() *sql.DB {
	return Db
}

// func CloseDb(){}

func GetRedis() *redis.Client {
	return RedisClient
}
