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
	Repositories(username string) (interface{}, error)
	AddRepositories(username string, repos model.Repositories) error
	Commits(username, reponame string) (interface{}, error)
	AddCommits(username, reponame string, commits model.Commits) error
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

func (c *Cache) AddRepositories(username string, repos model.Repositories) error {

	return c.cache.Replace(username, repos, cache.DefaultExpiration)
}

func (c *Cache) AddCommits(username, reponame string, commits model.Commits) error {

	c.cache.Set(fmt.Sprintf("%s%s", username, reponame), commits, cache.DefaultExpiration)
	return nil
}

func (c *Cache) Commits(username, reponame string) (interface{}, error) {

	repos, found := c.cache.Get(fmt.Sprintf("%s%s", username, reponame))
	if !found {
		return nil, fmt.Errorf("cache unable to find entries for %s", username)
	}

	return repos, nil
}
