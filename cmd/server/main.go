package main

import (
	"log"

	"anki-project/config"
	"anki-project/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	db := config.InitDB(c.Database.DSN)

	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	r := gin.Default()

	routers.SetupRoutes(r, db)

	if err := r.Run(c.Server.Address); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
