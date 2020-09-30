package main

import (
	"os"
	"workspace-go/github-repo-crawler/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	service := api.Service{}
	r.GET("/repositories/:username", service.Repositories)
	r.GET("/repositories/:username/commits/:reponame", service.Commits)

	if err := r.Run(":8080"); err != nil {
		os.Exit(1)
	}
}
