package github

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/google/go-github/v58/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
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
	// de-duplicate the list of repos
	seen := make(map[string]bool)
	j := 0
	for i, repo := range repos {
		if !seen[repo.GetName()] {
			seen[repo.GetName()] = true
			repos[j] = repos[i]
			j++
		}
	}
	repos = repos[:j]
	log.Println("Gathered:", len(repos), "Repositories")
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
		for count > 5 {
			// don't clone more than 5 at a time.
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

func gitClone(repo *github.Repository, cloneurl bool) (err error) {
	// check if git is install on the PATH
	commandPath, err := exec.LookPath("git")
	if cloneurl {
		if err != nil {
			_, err := git.PlainClone(repo.GetName(), false, &git.CloneOptions{
				URL:          *repo.CloneURL,
				Progress:     os.Stdout,
				SingleBranch: true,
				Mirror:       true,
				NoCheckout:   false,
			})
			if err != nil {
				log.Println(err)
			}
		} else {
			command := exec.Command(commandPath, "clone", *repo.CloneURL)
			out, err := command.CombinedOutput()
			if err != nil {
				log.Println(err)
			}
			log.Println(string(out))
		}
	} else {
		if err != nil {
			_, err := git.PlainClone(repo.GetName(), false, &git.CloneOptions{
				URL:          *repo.SSHURL,
				Progress:     os.Stdout,
				SingleBranch: true,
				Mirror:       true,
				NoCheckout:   false,
			})
			if err != nil {
				log.Println(err)
			}
		} else {
			command := exec.Command(commandPath, "clone", *repo.SSHURL)
			out, err := command.CombinedOutput()
			if err != nil {
				log.Println(err)
			}
			log.Println(string(out))
		}
	}
	return err
}

func clone(repo *github.Repository, cloneurl, mirroring bool) {
	count++
	if cloneurl {
		log.Println("Cloning", repo.GetName(), *repo.CloneURL)
		err := gitClone(repo, cloneurl)
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
		err := gitClone(repo, cloneurl)
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
	log.Println("Mirroring:", repo.GetName(), "with", self)
	Command := exec.Command(self)
	Command.Dir = repo.GetName()
	Command.Env = os.Environ()
	Command.Env = append(Command.Env, "GITHUB_USER="+EnvUser)
	Command.Env = append(Command.Env, "GITHUB_TOKEN="+EnvToken)
	out, err := Command.CombinedOutput()
	if err != nil {
		return err
	}
	log.Println(string(out))
	return nil
}
