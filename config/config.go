package config

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func GetDB() (conn redis.Conn) {

	// dsn := "host=localhost user=postgres password=8elmira8 dbname=zapood port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// client = redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379", // host:port of the redis server
	// 	Password: "",               // no password set
	// 	DB:       0,                // use default DB
	// })

	// return
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	return conn
}
