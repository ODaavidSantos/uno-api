package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize router
	r := gin.Default()

	// Initialize routes
	initilizeRoutes(r)

	// Run the server
	r.Run(":8080")
}