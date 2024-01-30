package github

import (
	"testing"
)

func TestListAllRepos(t *testing.T) {
	repos, err := ListAllRepos("eyedeekay", "")
	if err != nil {
		t.Error(err)
	}
	if len(repos) == 0 {
		t.Error("No repos found")
	}
	for i, repo := range repos {
		t.Log(i, ": ", *repo.SSHURL)
	}
}
