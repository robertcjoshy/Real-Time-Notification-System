package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
	"github.com/robert/notification/admin"
	"github.com/robert/notification/admin/admin_handler"
	"github.com/robert/notification/app"
	m "github.com/robert/notification/middlewares"
	"github.com/robert/notification/user"
	"github.com/robert/notification/user/user_handler"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	bolbolinstance := app.Build()
	admin_bolbol := &admin_handler.Server{Bolbol: bolbolinstance}
	user_bolbol := &user_handler.Server{Bolbol: bolbolinstance}
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")
	user.User_routes(router, user_bolbol)
	admin.Admin_routes(router, admin_bolbol)

	// ratelimiter
	limiter := rate.NewLimiter(1, 5) // Example: Allow 5 requests per second with burst of 1
	router.Use(m.RateLimitMiddleware(limiter))

	// Create the server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Run the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
