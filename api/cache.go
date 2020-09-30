package api

import (
	"fmt"
	"time"
	"workspace-go/github-repo-crawler/model"

	"github.com/patrickmn/go-cache"
)

const (
	expirationMinutes = 1
	purgeMinutes      = 5
)

type CacheConnector interface {
	AddRepositories(username string, repos model.Repositories) 
	AddCommits(username, reponame string, commits model.Commits) 
	Repositories(username string) (interface{}, error)
	Commits(username, reponame string) (interface{}, error)
}

type Cache struct {
	cache cache.Cache
}

func InitCache() Cache {

	c := cache.New(expirationMinutes*time.Minute, purgeMinutes*time.Minute)

	return Cache{
		cache: *c,
	}
}

func (c *Cache) Repositories(username string) (interface{}, error) {

	repos, found := c.cache.Get(username)
	if !found {
		return nil, fmt.Errorf("cache unable to find entries for %s", username)
	}

	return repos, nil
}

func (c *Cache) AddRepositories(username string, repos model.Repositories)  {

	c.cache.Set(username, repos, cache.DefaultExpiration)
}

func (c *Cache) AddCommits(username, reponame string, commits model.Commits) {

	c.cache.Set(fmt.Sprintf("%s%s", username, reponame), commits, cache.DefaultExpiration)
}

func (c *Cache) Commits(username, reponame string) (interface{}, error) {

	repos, found := c.cache.Get(fmt.Sprintf("%s%s", username, reponame))
	if !found {
		return nil, fmt.Errorf("cache unable to find entries for %s", username)
	}

	return repos, nil
}
