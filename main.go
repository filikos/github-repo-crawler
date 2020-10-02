package main

import (
	"fmt"
	"os"
	"workspace-go/github-repo-crawler/api"
	"workspace-go/github-repo-crawler/db"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:    "Github-Repo-Crawler",
		Version: "v1.0.0",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Usage: "Port the Rest-API will listen on.",
				Value: 8080,
			},
			&cli.PathFlag{
				Name:        "configPath",
				Usage:       "Path to *.env postgres config file.",
				Value:       "./config/dbConfig.env",
				DefaultText: "./config/dbConfig.env",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		r := gin.Default()

		dbConn, err := db.InitDB(c.String("configPath"))
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

		if err := r.Run(fmt.Sprintf(":%v", c.Int("port"))); err != nil {
			os.Exit(1)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
