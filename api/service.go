package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"workspace-go/github-repo-crawler/db"
	"workspace-go/github-repo-crawler/model"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Cache       Cache
	DBConnector db.Database
}

func (s *Service) Repositories(c *gin.Context) {

	username, _ := c.Params.Get("username")
	if len(username) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username can't be empty",
		})
		return
	}

	repos, err := s.Cache.Repositories(username)
	if err == nil {
		reposTyped, ok := repos.(model.Repositories)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, reposTyped.GetNames())
		return
	}

	fmt.Println(err)

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

	s.Cache.AddRepositories(username, respData)
	c.JSON(http.StatusOK, respData.GetNames())
}

func (s *Service) Commits(c *gin.Context) {

	username, _ := c.Params.Get("username")
	reponame, _ := c.Params.Get("reponame")
	if len(username) == 0 || len(reponame) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or reponame can't be empty",
		})
		return
	}

	search := c.Query("search")
	var respData model.CommitMetas

	commits, err := s.Cache.Commits(username, reponame)
	if err != nil {

		fmt.Print(err)

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

		err = json.Unmarshal(body, &respData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		s.Cache.AddCommits(username, reponame, respData.GetCommits())

		if len(search) == 0 {
			c.JSON(http.StatusOK, respData.GetCommits())
			return
		}

		commits := respData.GetCommits()
		c.JSON(http.StatusOK, commits.GetCommitsBySearch(search))
		return
	}

	commitsTyped, ok := commits.(model.Commits)
	if !ok {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(search) == 0 {
		c.JSON(http.StatusOK, commitsTyped)
		return
	}

	c.JSON(http.StatusOK, commitsTyped.GetCommitsBySearch(search))
}
