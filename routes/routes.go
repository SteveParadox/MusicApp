package routes

import (
	"github.com/gin-gonic/gin",
	"./handlers"
)

func RegisterRoutes(r *gin.Engine){
	api := r.Group("/api")
	{
		api.POST("/register", handlers.RegisterUser)
		api.GET("/user/likes", handlers.GetLikes)
	}
}