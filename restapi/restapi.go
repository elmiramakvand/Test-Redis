package restapi

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func RunApi(conn redis.Conn) *gin.Engine {
	r := gin.Default()
	RunApiOnRouter(r, conn)
	return r
}

func RunApiOnRouter(r *gin.Engine, conn redis.Conn) {

	ManageUserRoutes(r, conn)
}

func ManageUserRoutes(r *gin.Engine, conn redis.Conn) {
	Handler := NewUserModel(conn)
	userGroup := r.Group("/api/User")
	{
		userGroup.GET("GetAllUsers", Handler.GetAllUsers)
		userGroup.POST(":operation", Handler.Operation)
		userGroup.DELETE("delete/:id", Handler.Delete)
	}
}
