package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Chetan7595/task-manager/internal/routes"
	"github.com/Chetan7595/task-manager/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to Connect to Database: %v", err)
	}

	routes.SetupRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server {
		Addr: ":" + port,
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("🚀 Server running on port %s", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}