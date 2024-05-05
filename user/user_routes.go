package user

import (
	"github.com/gin-gonic/gin"
	"github.com/robert/notification/user/user_handler"
)

func User_routes(r *gin.Engine) {
	UserGroup := r.Group("/user")
	{
		UserGroup.GET("/", user_handler.User_homepage)
		UserGroup.GET("/login", user_handler.Login)
		UserGroup.POST("/login", user_handler.User_login)
		UserGroup.GET("/logout", user_handler.User_logout)
	}
}
