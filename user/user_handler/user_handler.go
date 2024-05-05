package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func User_homepage(c *gin.Context) {
	c.HTML(200, "companyhomepage.html", gin.H{})
}
func Login(c *gin.Context) {
	c.HTML(200, "userlogin.html", gin.H{})
}
func User_login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "robertcjoshy" || password == "12345678" {
		c.HTML(200, "userhomepage.html", gin.H{})
		c.Abort()
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
}
func User_logout(c *gin.Context) {
	c.Redirect(http.StatusFound, "/user")
}
