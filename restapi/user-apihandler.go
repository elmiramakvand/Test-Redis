package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"test-redis/entities"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

var currentUserId int64

type UserModel struct {
	conn redis.Conn
}

func NewUserModel(conn redis.Conn) *UserModel {
	return &UserModel{
		conn: conn,
	}
}

func (userModel UserModel) GetAllUsers(c *gin.Context) {
	rows, err := userModel.conn.Do("KEYS", "user:*")
	if err != nil {
		panic(err)
	}
	var users []entities.User
	for _, k := range rows.([]interface{}) {
		u := entities.User{}
		// Get the data at each key
		reply, err := userModel.conn.Do("GET", k.([]byte))
		if err != nil {
			panic(err)
		}
		// Marshal JSON data to Post
		if err := json.Unmarshal(reply.([]byte), &u); err != nil {
			panic(err)
		}
		// Append each to posts slice
		users = append(users, u)
	}
	c.JSON(http.StatusOK, users)
	return
}

func (userModel UserModel) Operation(c *gin.Context) {
	operation := c.Params.ByName("operation")
	if operation != "add" && operation != "edit" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	var user entities.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	switch strings.ToLower(operation) {
	case "add":
		currentUserId += 1
		user.ID = currentUserId
		// Marshal Post to JSON blob
		b, err := json.Marshal(&user)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		// Save JSON to Redis
		reply, err := userModel.conn.Do("SET", "user:"+strconv.Itoa(int(user.ID)), b)
		fmt.Println("Insert reply :", reply)
		if err != nil {
			fmt.Println("Insert Error :", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

	case "edit":
		b, err := json.Marshal(&user)
		reply, err := userModel.conn.Do("SET", "user:"+strconv.Itoa(int(user.ID)), b)
		fmt.Println("Edit reply :", reply)
		if err != nil {
			fmt.Println("Edit Error :", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	return

}

func (userModel UserModel) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	result, err := userModel.conn.Do("DEL", "user:"+id)
	if err != nil {
		fmt.Println("Delete Error :", err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	fmt.Println("delete user result :", result)
	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	return
}
