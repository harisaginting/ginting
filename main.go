package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"os"
	"log"

	database "github.com/harisaginting/ginting/db"
	router "github.com/harisaginting/ginting/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// DB CONNECTION
	db := database.Connection()

	port 	:= os.Getenv("PORT")
	app 	:= gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	// get default url request
	app.UseRawPath 		   = true
	app.UnescapePathValues = true

	// cors configuration
	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization", "x-source")
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"OPTIONS", "PUT", "POST", "GET", "DELETE"}
	app.Use(cors.New(config))

	// route
	app.GET("/ping", ping)
	app.NoRoute(lostInSpce)

	// API
	api := app.Group("api")
	router.RestV1(api, db)


	// handling server gracefully shutdown
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app,
	}
	// Initializing the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("listen: %s\n", err)
		}
	}()
	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force 🔴")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown 🔴: ", err)
	}
	log.Println("Server exiting 🔴")
}

func lostInSpce(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":        404,
		"data":          nil,
		"error_message": "No Route Found",
	})
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"service_name": os.Getenv("APP_NAME"),
	})
	return
}