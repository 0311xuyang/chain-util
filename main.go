package main

import (
	"fmt"

	"github.com/0311xuyang/chain-util/route"
	"github.com/gin-gonic/gin"
)

// LoadRoute defines the routes for the server.
func LoadRoute(server *gin.Engine) {
	dbController := server.Group("/db")
	dbController.GET("/get/:id", route.GetUser)
	dbController.POST("/set", route.SetUser)
}

// StartServer a gin server
func StartServer() {
	gin.SetMode(gin.DebugMode)
	server := gin.Default()
	LoadRoute(server)
	if err := server.Run(":8080"); err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	StartServer()
}
