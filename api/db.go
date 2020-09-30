package api

import "workspace-go/github-repo-crawler/model"

type DBConnector interface {
	Repositories(username string) (model.Repositories, error)
	AddRepositories(username string, repos model.Repositories) error
	Commits(username, reponame string) (model.Commits, error)
	AddCommits(username, reponame string, commits model.Commits) error
}
