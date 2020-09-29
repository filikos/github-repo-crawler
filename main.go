package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/repositories/:username", test)
	r.GET("/repositories/:username/commits/:reponame", test)
	r.GET("/repositories/:username/commits/:reponame/search/:query", test)

	if err := r.Run(":8080"); err != nil {
		os.Exit(1)
	}
}

func test(c *gin.Context) {
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}
