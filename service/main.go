package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"dgf.io/inter/service/resources"
)

func main() {
	fmt.Println("Starting service")
	router := gin.Default()

	channels := resources.Channels{}

	router.GET("/channels", channels.Get)
	router.GET("/channels/", channels.Get)
	router.GET("/channels/:id", channels.Get)
	router.POST("/channels", channels.Post)

	router.Run("localhost:8080")
}
