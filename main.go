package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strings"
	"workspace-go/github-repo-crawler/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/repositories/:username", repositories)
	r.GET("/repositories/:username/commits/:reponame", commits)

	if err := r.Run(":8080"); err != nil {
		os.Exit(1)
	}
}

func repositories(c *gin.Context) {
	username, _ := c.Params.Get("username")

	if len(username) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username can't be empty",
		})
		return
	}

	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, nil)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var respData model.Repositories
	err = json.Unmarshal(body, &respData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	repoNames := make([]string, 0)
	for _, v := range respData {
		repoNames = append(repoNames, v.Name)
	}

	c.JSON(http.StatusOK, repoNames)
}

func commits(c *gin.Context) {
	username, _ := c.Params.Get("username")
	reponame, _ := c.Params.Get("reponame")

	if len(username) == 0 || len(reponame) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or reponame can't be empty",
		})
		return
	}

	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", username, reponame))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, nil)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var respData []model.CommitMeta
	err = json.Unmarshal(body, &respData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	search := c.Query("search")
	if len(search) == 0 {
		commits := make(model.Commits, 0)
		max := math.Max(20, float64(len(respData)))
		for i := 0; i < int(max); i++ {
			commits = append(commits, respData[i].Commit)
		}

		c.JSON(http.StatusOK, commits)
		return
	}

	commits := make(model.Commits, 0)
	max := math.Max(20, float64(len(respData)))
	for i := 0; i < int(max); i++ {
		// check if search is in commit msg
		if strings.Contains(respData[i].Commit.Message, search) {
			commits = append(commits, respData[i].Commit)
		}
	}

	c.JSON(http.StatusOK, commits)
}