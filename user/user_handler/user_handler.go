package user_handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robert/notification/app"
	"github.com/robert/notification/app/entity"
)

type Server struct {
	Bolbol *app.Bobol
}

func Login(c *gin.Context) {
	c.HTML(200, "userlogin.html", gin.H{})
}

func (s *Server) User_login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	userid, _ := strconv.Atoi(username)
	// Dummy check for username and password
	if username == "1001" && password == "1001" {
		// Render the user homepage first
		c.SetCookie("userid", strconv.Itoa(userid), 3600, "/", "localhost", false, true)
		c.HTML(200, "usernotificationpage.html", gin.H{})
		return
	}

	// If credentials are invalid, return an error response
	c.JSON(400, gin.H{"error": "invalid credentials"})
}

func User_logout(c *gin.Context) {
	c.SetCookie("session", "userid", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/user/login")
}

func (s *Server) Notifymessage(c *gin.Context, id int, timestamp int64) ([]entity.Notification, error) {

	nots, err := s.Bolbol.Getnotifications(c, id, timestamp)
	if err != nil {
		return nil, err
	}
	return nots, err

}

func (s *Server) FetchNotifications(c *gin.Context) {

	userID, _ := c.Get("userid") // This should be dynamically determined based on the logged-in user

	id, _ := userID.(int)

	timestamp := time.Now().Unix() // this hsuld be dynamically determined

	nots, err := s.Notifymessage(c, id, timestamp)
	if err != nil {
		log.Println("error in fetching notification or no notification")
		c.JSON(200, gin.H{"error": "error fetching notifications"}) // it must show internel server error but changeged to 200
		return
	}
	n := len(nots)

	data := make([]string, 0)
	for _, value := range nots {
		temp := value.(entity.Messagenotification).Noty
		data = append(data, temp)
	}

	if n != 0 {
		c.JSON(200, gin.H{"message": data})
		return
	}
	c.JSON(200, gin.H{"message": "no notifications"})
}
