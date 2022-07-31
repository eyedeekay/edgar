package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/eyedeekay/edgar/tohtml"
	github "github.com/google/go-github/v45/github"
	"github.com/yosssi/gohtml"
	"golang.org/x/oauth2"
)

var (
	filename  = flag.String("filename", "README.md", "The markdown file to convert to HTML")
	author    = flag.String("author", authorDefault(), "The author of the HTML file")
	css       = flag.String("css", "style.css", "The CSS file to use, a default will be generated if one doesn't exist")
	script    = flag.String("script", hasScript(), "The script file to use")
	title     = flag.String("title", "", "The title of the HTML file")
	outfile   = flag.String("out", "index.html", "The name of the output file")
	snowflake = flag.Bool("snowflake", true, "add a snowflake to the page footer")
)

func authorDefault() string {
	user := findGithubUsername()
	if user != "" {
		return user
	}
	cmd := exec.Command("git", "config", "--get", "user.email")
	out, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.Split(string(out), "@")[0]
}

func hasScript() string {
	_, err := os.Stat("script.js")
	if err != nil {
		return ""
	}
	return "script.js"
}

func main() {
	flag.Parse()
	if *title == "" {
		var err error
		*title, err = tohtml.ReadFirstMarkdownHeader("README.md")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	output := tohtml.OutputHTMLOpen()
	output += tohtml.OutputHeaderOpen()
	output += tohtml.OutputTitleTag(*title)
	output += tohtml.OutputMetaTag("author", *author)
	output += tohtml.OutputMetaTag("description", findGithubRepoName())
	output += tohtml.OutputMetaTag("keywords", getCurrentBranch())
	output += tohtml.OutputCSSTag(*css)
	output += tohtml.OutputScriptTag(*script)
	output += tohtml.OutputHeaderClose()
	output += tohtml.OutputBodyOpen()
	output += tohtml.OutputHTMLFromMarkdown(*filename, *title)
	output += tohtml.Snowflake()
	output += tohtml.License()
	output += tohtml.OutputBodyClose()
	output += tohtml.OutputHTMLClose()
	output = gohtml.Format(output)
	err := ioutil.WriteFile(".nojekyll", []byte{}, 0644)
	if err != nil {
		fmt.Printf("78 \n %s", err)
		os.Exit(1)
	}
	if *outfile != "" && *outfile != "-" {
		if err := ioutil.WriteFile(*outfile, []byte(output), 0644); err != nil {
			fmt.Printf("83\n %s", err)
			os.Exit(1)
		}
		gitAddCmd := exec.Command("git", "add", *outfile, ".nojekyll")
		if err := gitAddCmd.Run(); err != nil {
			fmt.Printf("88\n %s", err)
			os.Exit(1)
		}
		gitCommitCmd := exec.Command("git", "commit", "-am", "update "+*outfile)
		if out, err := gitCommitCmd.Output(); err != nil {
			fmt.Printf("93\n %s\n %s", out, err)
			os.Exit(1)
		}
		if err := enableGithubPage(); err != nil {
			if strings.Contains(err.Error(), "409") {
				fmt.Println("Page already exists, skipping")
				os.Exit(0)
			}
			fmt.Printf("102\n %s", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(output)
	}
}

func findGithubUsername() string {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	// trim away the scheme, domain
	user := strings.SplitN(string(out), "github.com", 2)[1]
	// trim away the leading separators
	user = strings.TrimLeft(string(user), "/")
	user = strings.TrimLeft(string(user), ":")
	// split at the /
	user = strings.Split(string(user), "/")[0]
	return strings.Replace(user, "\n", "", -1)
}

func findGithubRepoName() string {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	// trim away the scheme, domain
	repo := strings.SplitN(string(out), "github.com", 2)[1]
	// trim away the leading separators
	repo = strings.TrimLeft(string(repo), "/")
	repo = strings.TrimLeft(string(repo), ":")
	// split at the /
	repo = strings.Split(string(repo), "/")[1]
	repo = strings.TrimRight(string(repo), ".git")
	return strings.Replace(repo, "\n", "", -1)
}

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
	_, _, err := client.Repositories.EnablePages(ctx, findGithubUsername(), repoName, &github.Pages{
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

func getCurrentBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.Replace(string(out), "\n", "", -1)
}
