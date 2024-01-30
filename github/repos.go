package github

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v58/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"gopkg.in/src-d/go-git.v4"
)

func ListAllRepos(user, token string) ([]*github.Repository, error) {
	if token == "" {
		log.Println("API token taken from environment")
		token = EnvToken
	}
	if user == "" {
		user = EnvUser
	}
	client := github.NewTokenClient(nil, token)
	ctx := context.Background()
	repos := make([]*github.Repository, 0)
	opt := &github.RepositoryListByUserOptions{Type: "public", ListOptions: github.ListOptions{PerPage: 100}}
	triage, _, err := client.Repositories.ListByUser(ctx, user, opt)
	page := 0
	count := -1
	for len(triage) > 0 {
		log.Println("Gathered:", len(repos), "Repositories")
		opt := &github.RepositoryListByUserOptions{Type: "public", ListOptions: github.ListOptions{PerPage: 100, Page: page}}
		triage, _, err := client.Repositories.ListByUser(ctx, user, opt)
		if err != nil {
			return nil, err
		}
		for _, repo := range triage {
			if repo.GetHasPages() {
				repos = append(repos, repo)
			}
		}
		if count == len(repos) {
			break
		}
		count = len(repos)
		page++
	}
	return repos, err
}

func CloneAllRepos(user, token string) error {
	if token == "" {
		log.Println("API token taken from environment")
		token = EnvToken
	}
	if user == "" {
		user = EnvUser
	}
	allRepos, err := ListAllRepos(user, token)
	if err != nil {
		return err
	}
	cloneurl := false
	_, cloneerr := git.PlainClone(allRepos[0].GetName(), false, &git.CloneOptions{
		URL:      *allRepos[0].SSHURL,
		Progress: os.Stdout,
	})
	if cloneerr != nil {
		cloneurl = true
	}
	for _, repo := range allRepos {
		var err error
		if cloneurl {
			_, err = git.PlainClone(repo.GetName(), false, &git.CloneOptions{
				URL:      *repo.CloneURL,
				Progress: os.Stdout,
			})
		} else {
			_, err = git.PlainClone(repo.GetName(), false, &git.CloneOptions{
				URL:      *repo.SSHURL,
				Progress: os.Stdout,
			})
		}
		if err != nil {
			return err
		}
	}
	return nil
}
