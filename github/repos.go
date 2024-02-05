package github

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/google/go-github/v58/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
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
	ctx := context.TODO()
	repos := make([]*github.Repository, 0)
	opt := &github.RepositoryListByUserOptions{Type: "public", ListOptions: github.ListOptions{PerPage: 100}}
	triage, _, err := client.Repositories.ListByUser(ctx, user, opt)
	page := 0
	count := -1
	timeout := 0 
	for len(triage) > 0 {
		log.Println("Gathered:", len(repos), "Repositories", page)
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
			if timeout > 10 {
				break
			}
			timeout++
		}
		count = len(repos)
		page++
	}
	return repos, err
}

func CloneAllRepos(user, token string, mirroring bool) ([]*github.Repository, error) {
	if token == "" {
		log.Println("API token taken from environment")
		token = EnvToken
	}
	if user == "" {
		user = EnvUser
	}
	allRepos, err := ListAllRepos(user, token)
	if err != nil {
		return nil, err
	}
	cloneurl := false
	_, cloneerr := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:      *allRepos[0].SSHURL,
		Progress: os.Stdout,
	})
	if cloneerr != nil {
		cloneurl = true
	}
	cloned = 0
	count = 0
	for _, repo := range allRepos {
		for count > 15 {
			// don't clone more than 15 at a time.
			// After all the small ones are cloned the last queue will pretty much only be really big ones.
			log.Println("Sleeping until the queue opens up")
			time.Sleep(time.Second * 10)
		}
		go clone(repo, cloneurl, mirroring)
	}
	msg := "Cloned: " + strconv.Itoa(cloned) + " Repositories"
	for cloned < len(allRepos) {
		tmp := "Cloned: " + strconv.Itoa(cloned) + " Repositories"
		if tmp != msg {
			msg = tmp
			log.Println("\n\t", msg)
		}
	}
	return allRepos, err
}

var cloned = 0
var count = 0

func clone(repo *github.Repository, cloneurl, mirroring bool) {
	count++
	if cloneurl {
		log.Println("Cloning", repo.GetName(), *repo.CloneURL)
		_, err := git.PlainClone(repo.GetName(), false, &git.CloneOptions{
			URL:      *repo.CloneURL,
			Progress: os.Stdout,
		})
		if err != nil {
			log.Println(err)
		}
		if mirroring {
			err := mirror(repo)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		log.Println("Cloning", repo.GetName(), *repo.SSHURL)
		_, err := git.PlainClone(repo.GetName(), false, &git.CloneOptions{
			URL:      *repo.SSHURL,
			Progress: os.Stdout,
		})
		if err != nil {
			log.Println(err)
		}
		if mirroring {
			err := mirror(repo)
			if err != nil {
				log.Println(err)
			}
		}
	}
	cloned++
	count--
}

func Mirror(user, token string) error {
	_, err := CloneAllRepos(user, token, true)
	if err != nil {
		return err
	}
	return err
}

func mirror(repo *github.Repository) error {
	self, err := os.Executable()
	if err != nil {
		return err
	}
	log.Println("Mirroring:", repo.GetName(), self)
	Command := exec.Command(self)
	Command.Dir = repo.GetName()
	Command.Env = os.Environ()
	out, err := Command.CombinedOutput()
	if err != nil {
		return err
	}
	log.Println(string(out))
	return nil
}
