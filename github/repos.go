package github

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/google/go-github/v58/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
)

func ListAllRepos(user, token string) ([]*github.Repository, error) {
	if token == "" {
		if EnvToken != "" {
			log.Println("API token taken from environment")
			token = EnvToken
		}
	}
	if user == "" {
		user = EnvUser
	}
	ctx := context.TODO()
	client := github.NewTokenClient(ctx, token)
	repos := make([]*github.Repository, 0)
	opt := &github.RepositoryListByUserOptions{Type: "public", ListOptions: github.ListOptions{PerPage: 100}}
	triage, _, err := client.Repositories.ListByUser(ctx, user, opt)
	if err != nil {
		return nil, err
	}
	page := 0
	count := -1
	timeout := 0
	for len(triage) > 0 {
		log.Println("Gathered:", len(repos), "Repositories", page)
		opt := &github.RepositoryListByUserOptions{Type: "public", ListOptions: github.ListOptions{PerPage: 100, Page: page}}
		triage, resp, err := client.Repositories.ListByUser(ctx, user, opt)
		if len(triage) == 0 || resp.NextPage == 0 || err != nil {
			break
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
		if EnvToken != "" {
			log.Println("API token taken from environment")
			token = EnvToken
		}
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
	cloned.Store(0)
	count.Store(0)
	for _, repo := range allRepos {
		for count.Load() > 5 {
			// don't clone more than 5 at a time.
			// After all the small ones are cloned the last queue will pretty much only be really big ones.
			log.Println("Sleeping until the queue opens up")
			time.Sleep(time.Second * 10)
		}
		go clone(repo, cloneurl, mirroring)
	}
	msg := "Cloned: " + strconv.Itoa(int(cloned.Load())) + " Repositories"
	for cloned.Load() < int32(len(allRepos)) {
		tmp := "Cloned: " + strconv.Itoa(int(cloned.Load())) + " Repositories"
		if tmp != msg {
			msg = tmp
			log.Println("\n\t", msg)
		}
	}
	return allRepos, err
}

var (
	cloned atomic.Int32
	count  atomic.Int32
)

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
	count.Add(1)
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
	cloned.Add(1)
	count.Add(-1)
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
