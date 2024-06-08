package user

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robert/notification/user/user_handler"
)

func authMiddleware(c *gin.Context) {
	// Skip the middleware for the login endpoint
	if strings.HasPrefix(c.FullPath(), "/user/login") {
		c.Next()
		return
	}

	userids, err := c.Cookie("userid")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is required"})
		c.Abort()
		return
	}
	userid, _ := strconv.Atoi(userids)
	log.Println("AFTER CONVERTING TO INT")
	c.Set("userid", userid)
	c.Next()
}

func User_routes(r *gin.Engine, user_bolbol *user_handler.Server) {
	UserGroup := r.Group("/user")
	UserGroup.Use(authMiddleware)
	{
		UserGroup.GET("/login", user_handler.Login)
		UserGroup.POST("/login", user_bolbol.User_login)
		UserGroup.GET("/notification", user_bolbol.FetchNotifications)
		UserGroup.GET("/logout", user_handler.User_logout)
	}
}
