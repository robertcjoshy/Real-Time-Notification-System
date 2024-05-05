package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robert/notification/admin"
	"github.com/robert/notification/user"
	"github.com/robert/notification/utils"
)

func main() {
	utils.Opendatabaseconnection()
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")
	user.User_routes(router)
	admin.Admin_routes(router)
	router.Run()
}
