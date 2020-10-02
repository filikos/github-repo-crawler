package main

import (
	"fmt"
	"os"
	"workspace-go/github-repo-crawler/api"
	"workspace-go/github-repo-crawler/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dbConn, err := db.InitDB("config/dbConfig.env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	service := api.Service{
		Cache:       api.InitCache(),
		DBConnector: *dbConn,
	}

	r.GET("/repositories/:username", service.Repositories)
	r.GET("/repositories/:username/commits/:reponame", service.Commits)
	r.GET("/recentrepositories", service.RepositoriesDB)

	if err := r.Run(":8080"); err != nil {
		os.Exit(1)
	}
}
