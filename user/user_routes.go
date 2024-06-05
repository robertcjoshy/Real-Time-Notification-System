package user

import (
	"github.com/gin-gonic/gin"
	"github.com/robert/notification/user/user_handler"
)

func User_routes(r *gin.Engine, user_bolbol *user_handler.Server) {
	UserGroup := r.Group("/user")
	{
		UserGroup.GET("/", user_handler.User_homepage)
		UserGroup.GET("/login", user_handler.Login)
		UserGroup.POST("/login", user_bolbol.User_login)
		UserGroup.GET("/notification", user_bolbol.FetchNotifications)
		UserGroup.GET("/logout", user_handler.User_logout)
	}
}
