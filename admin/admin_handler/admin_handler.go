package admin_handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robert/notification/admin/admin_model"
	"github.com/robert/notification/app"
	"github.com/robert/notification/app/entity"
)

type Message struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type Server struct {
	Bolbol *app.Bobol
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

func (s *Server) SendNotification(c *gin.Context) {
	var message Message

	err := c.BindJSON(&message)
	log.Println("message = ", message)
	//var actual_notification entity.Messagenotification
	if err != nil {
		log.Println("error in binding = ", err)
		c.JSON(http.StatusOK, gin.H{"success": false})
		c.Abort()
		return
	}

	id, _ := strconv.Atoi(message.Subject)
	log.Println("subject to string = ", id)
	errr := s.Bolbol.Notify(c, id, entity.Messagenotification{Basenotification: entity.Basenotification{Createdat: time.Now().Unix()}, Noty: message.Message})
	if errr != nil {
		log.Println("error in notify function = ", errr)
		c.JSON(http.StatusOK, gin.H{"success": false})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"success": true})
}
