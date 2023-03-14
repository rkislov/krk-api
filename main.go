package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rkislov/krk-api/initializers"
	"log"
	"net/http"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Не могу загрузить настройки их файда app.env")
	}
	initializers.ConnectDB(&config)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Не могу загрузить настройки их файда app.env")
	}
	router := server.Group("/api")
	router.GET("/healthcheacker", func(ctx *gin.Context) {
		message := "API work"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	log.Fatal(server.Run(":" + config.ServerPort))
}
