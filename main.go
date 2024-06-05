package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robert/notification/admin"
	"github.com/robert/notification/admin/admin_handler"
	"github.com/robert/notification/app"
	"github.com/robert/notification/user"
	"github.com/robert/notification/user/user_handler"
	"github.com/robert/notification/utils"
)

func main() {
	utils.Opendatabaseconnection()
	router := gin.Default()
	bolbolinstance := app.Build()
	admin_bolbol := &admin_handler.Server{Bolbol: bolbolinstance}
	user_bolbol := &user_handler.Server{Bolbol: bolbolinstance}
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")
	user.User_routes(router, user_bolbol)
	admin.Admin_routes(router, admin_bolbol)
	router.Run()
}
