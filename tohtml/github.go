package tohtml

import (
	"context"
	"fmt"
	"os"

	github "github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func enableGithubPage() error {
	token := os.Getenv("GITHUB_TOKEN")
	if len(token) == 0 {
		return fmt.Errorf("GITHUB_TOKEN not set")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repoName := findGithubRepoName()
	if repoName == "" {
		return fmt.Errorf("could not find github repo name")
	}
	branch := getCurrentBranch()
	path := "/"
	_, _, err := client.Repositories.EnablePages(ctx, FindGithubUsername(), repoName, &github.Pages{
		Source: &github.PagesSource{
			Branch: github.String(branch),
			Path:   github.String(path),
		},
		Public: github.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("could not enable github pages: %v", err)
	}

	return nil
}
