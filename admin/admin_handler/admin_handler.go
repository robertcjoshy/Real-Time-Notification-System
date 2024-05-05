package admin_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robert/notification/admin/admin_model"
)

type Message struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "adminlogin.html", gin.H{})
}

func Loginpost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	isadmin_valid, err := admin_model.Getadmin(username, password)
	if !isadmin_valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		c.Abort()
		return
	}

	c.HTML(http.StatusOK, "admin_homepage.html", gin.H{"username": username})
}

func Logout(C *gin.Context) {
	C.HTML(http.StatusOK, "adminlogin.html", gin.H{})
}

func SendNotification(c *gin.Context) {
	var message Message

	err := c.BindJSON(&message)
	//fmt.Println(message)
	//fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"success": true})

}
