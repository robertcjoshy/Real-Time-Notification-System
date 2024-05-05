package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/robert/notification/admin/admin_handler"
)

func Admin_routes(r *gin.Engine) {
	AdminGroup := r.Group("/admin")
	{
		AdminGroup.GET("/login", admin_handler.Login)
		AdminGroup.POST("/login", admin_handler.Loginpost)
		AdminGroup.GET("/logout", admin_handler.Logout)
		AdminGroup.POST("/sendmessage", admin_handler.SendNotification)
	}
}
