package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robert/notification/app"
	"github.com/robert/notification/app/entity"
)

type Server struct {
	Bolbol *app.Bobol
}

func User_homepage(c *gin.Context) {
	c.HTML(200, "companyhomepage.html", gin.H{})
}
func Login(c *gin.Context) {
	c.HTML(200, "userlogin.html", gin.H{})
}

/*
	func (s *Server) User_login(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "robertcjoshy" || password == "1001" {
			id := 1001
			c.HTML(200, "userhomepage.html", gin.H{})
			fmt.Println("after rendring userhomepage")
			nots, err := s.Notifymessage(c, id)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
				c.Abort()
				return
			}
			n := len(nots)
			fmt.Println("PRINTING NOTS")
			fmt.Println(nots)
			if n != 0 {
				c.JSON(200, gin.H{"message": "n != zero"})
				c.Abort()
				return
			}
		}
		c.JSON(400, gin.H{"error": "invalid credentials"})
	}
*/
func (s *Server) User_login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Dummy check for username and password
	if username == "robertcjoshy" && password == "1001" {
		// Render the user homepage first
		c.HTML(200, "userhomepage.html", gin.H{})
		return
	}

	// If credentials are invalid, return an error response
	c.JSON(400, gin.H{"error": "invalid credentials"})
}

func User_logout(c *gin.Context) {
	c.Redirect(http.StatusFound, "/user")
}

func (s *Server) Notifymessage(c *gin.Context, id int) ([]entity.Notification, error) {
	//id := c.Param("id")
	/*
		//lastupdate, err := strconv.ParseInt(time, 10, 54)
		ids, _ := strconv.Atoi(id)
		if len(id) == 0 {
			c.JSON(400, gin.H{"error": "no timestamp"})
			c.Abort()
			return
		}
	*/
	//var s app.Bobol
	nots, err := s.Bolbol.Getnotifications(c, id)
	if err != nil {
		return nil, err
	}
	return nots, err
	//c.HTML(200, "userhomepage.html", gin.H{"message": nots})

}

func (s *Server) FetchNotifications(c *gin.Context) {
	id := 1001 // This should be dynamically determined based on the logged-in user
	nots, err := s.Notifymessage(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching notifications"})
		return
	}
	n := len(nots)
	specific := nots[0]
	data := specific.(entity.Messagenotification).Noty
	//fmt.Println(specific.Isnotification())
	/*for _, ele := range nots {
		if not, ok := ele.(entity.Messagenotification); ok {
			fmt.Println(not.Noty)
		}
	}*/
	if n != 0 {
		c.JSON(200, gin.H{"message": data})
		return
	}
	c.JSON(200, gin.H{"message": "no notifications"})
}
