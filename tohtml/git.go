package tohtml

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func getCurrentBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.Replace(string(out), "\n", "", -1)
}

func CommitMessage() {
	gitCommitCmd := exec.Command("git", "commit", "-am", "page generation update for: "+time.Now().String())
	if out, err := gitCommitCmd.Output(); err != nil {
		fmt.Printf("Git Commit Error: %s %s", out, err)
	}
	if err := enableGithubPage(); err != nil {
		if strings.Contains(err.Error(), "409") {
			fmt.Println("Page already exists, skipping")
		} else if strings.Contains(err.Error(), "GITHUB_TOKEN not set") {
			fmt.Println("GITHUB_TOKEN not set, skipping")
		} else {
			fmt.Printf("Github Pages Error: %s", err)
		}
	}
}

func findGithubRemote() string {
	// list all the remotes
	cmdRemotes := exec.Command("git", "remote", "-v")
	rout, err := cmdRemotes.Output()
	if err != nil {
		return ""
	}
	var remote string
	// find the github remote
	for _, line := range strings.Split(string(rout), "\n") {
		if strings.Contains(line, "github.com") {
			// store the remote name
			remote = strings.Split(line, "\t")[0]
			break
		}
	}
	return remote
}

func FindGithubUsername() string {
	remote := findGithubRemote()
	cmd := exec.Command("git", "remote", "get-url", remote)
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	var user string
	// trim away the scheme, domain
	user = strings.SplitN(string(out), "github.com", 2)[1]

	// trim away the leading separators
	user = strings.TrimLeft(string(user), "/")
	user = strings.TrimLeft(string(user), ":")

	// split at the /
	user = strings.Split(string(user), "/")[0]
	return strings.Replace(user, "\n", "", -1)
}

func findGithubRepoName() string {
	remote := findGithubRemote()
	fmt.Printf("Remote: '%s'\n", remote)
	cmd := exec.Command("git", "remote", "get-url", remote)
	out, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return ""
	}

	// trim away the scheme, domain
	repo := strings.SplitN(string(out), "github.com", 2)[1]

	// trim away the leading separators
	repo = strings.TrimLeft(string(repo), "/")
	repo = strings.TrimLeft(string(repo), ":")

	// split at the /
	repo = strings.Split(string(repo), "/")[1]
	repo = strings.Replace(string(repo), "\n", "", -1)
	if strings.HasSuffix(string(repo), ".git") {
		repo = strings.TrimRight(string(repo), ".git")
	}
	return repo
}
