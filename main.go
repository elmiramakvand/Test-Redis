package main

import (
	"test-redis/config"
	"test-redis/restapi"
)

func main() {
	conn := config.GetDB()
	r := restapi.RunApi(conn)
	//running
	r.Run(":8080")

}
