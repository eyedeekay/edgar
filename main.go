package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/eyedeekay/edgar/tohtml"
	github "github.com/google/go-github/v45/github"
	"github.com/yosssi/gohtml"
	"golang.org/x/oauth2"
)

var (
	filename  = flag.String("filename", listAllMarkdownFiles(), "The markdown file to convert to HTML, or a comma separated list of files")
	author    = flag.String("author", authorDefault(), "The author of the HTML file")
	css       = flag.String("css", "style.css", "The CSS file to use, a default will be generated if one doesn't exist")
	script    = flag.String("script", hasScript(), "The script file to use.")
	title     = flag.String("title", "", "The title of the HTML file, if blank it will be generated from the first h1 in the markdown file.")
	outfile   = flag.String("out", "index.html", "The name of the output file(Only used for the first file, others will be named `inputfile.html`)")
	snowflake = flag.Bool("snowflake", true, "add a snowflake to the page footer")
	i2plink   = flag.Bool("i2plink", true, "add an i2p link to the page footer. Logo courtesy of @Shoalsteed and @mark22k")
	nodonate  = flag.Bool("nodonate", true, "disable the donate section(change the -donate wallet addresses before setting this to true)")
	donate    = flag.String(
		"donate",
		"monero:4A2BwLabGUiU65C5JRfwXqFTwWPYNSmuZRjbTDjsu9wT6wV6kMFyXn83ydnVjVcR7BCsWh8B5b4Z9b6cmqjfZiFd9sBUpWT,bitcoin:1D1sDmyZAs5q2Lb29q8TBnGhEJK7vfp5PJ,ethereum:0x539a4356bb0566a39376CaC3F50B558F77E84eC9",
		"add donation section to cryptocurrency wallets. Use the address URL schemes, separated by commas(no spaces). Change them before running unless you want the money to go to me.",
	)
	donatemessage = flag.String("support", "Support independent development"+myDirectory(), "change message/CTA for donations section.")
)

func myDirectory() string {
	override := os.Getenv("PROJECT_NAME")
	if override == "" {
		wd, err := os.Getwd()
		if err != nil {
			return ""
		}
		return " of " + filepath.Base(wd)
	}
	return override
}

func listAllMarkdownFiles() string {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var fileList []string

	if _, err := os.Stat("README.md"); err == nil {
		fileList = append(fileList, "README.md")
	}

	for _, file := range files {
		if !file.IsDir() {
			if strings.HasSuffix(file.Name(), ".md") {
				if file.Name() != "README.md" {
					fileList = append(fileList, file.Name())
				}
			} else if strings.HasSuffix(file.Name(), ".html") {
				mdExtension := strings.ReplaceAll(file.Name(), ".html", ".md")
				if _, err := os.Stat(mdExtension); err != nil {
					fileList = append(fileList, file.Name())
				}
			}
		}
	}
	if _, err := os.Stat("docs"); err == nil {
		docs, err := ioutil.ReadDir("docs")
		if err != nil {
			log.Fatal(err)
		}
		tohtml.OutputCSSTag("docs/styles.css")
		tohtml.OutputShowHiderCSSTag("docs/showhider.css")
		gitAddCmd := exec.Command("git", "add", "docs/styles.css", "docs/showhider.css")
		if err := gitAddCmd.Run(); err != nil {
			fmt.Printf("Git Add Error: %s", err)
			os.Exit(1)
		}
		//var fileList []string
		for _, file := range docs {
			if !file.IsDir() {
				if strings.HasSuffix(filepath.Join("docs", file.Name()), ".md") {
					if filepath.Join("docs", file.Name()) != "README.md" {
						fileList = append(fileList, filepath.Join("docs", file.Name()))
					}
				} else if strings.HasSuffix(filepath.Join("docs", file.Name()), ".html") {
					mdExtension := strings.ReplaceAll(filepath.Join("docs", file.Name()), ".html", ".md")
					if _, err := os.Stat(mdExtension); err != nil {
						fileList = append(fileList, filepath.Join("docs", file.Name()))
					}
				}
			}
		}
	}
	if _, err := os.Stat("./doc"); err == nil {
		docs, err := ioutil.ReadDir("doc")
		if err != nil {
			log.Fatal(err)
		}
		tohtml.OutputCSSTag("doc/styles.css")
		tohtml.OutputShowHiderCSSTag("doc/showhider.css")
		gitAddCmd := exec.Command("git", "add", "doc/styles.css", "doc/showhider.css")
		if err := gitAddCmd.Run(); err != nil {
			fmt.Printf("Git Add Error: %s", err)
			os.Exit(1)
		}
		//var fileList []string
		for _, file := range docs {
			if !file.IsDir() {
				if strings.HasSuffix(filepath.Join("doc", file.Name()), ".md") {
					if filepath.Join("doc", file.Name()) != "README.md" {
						fileList = append(fileList, filepath.Join("doc", file.Name()))
					}
				} else if strings.HasSuffix(filepath.Join("doc", file.Name()), ".html") {
					mdExtension := strings.ReplaceAll(filepath.Join("doc", file.Name()), ".html", ".md")
					if _, err := os.Stat(mdExtension); err != nil {
						fileList = append(fileList, filepath.Join("doc", file.Name()))
					}
				}
			}
		}
	}
	return strings.Join(fileList, ",")
}

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
	filesList := strings.Split(*filename, ",")
	if *title == "" {
		var err error
		*title, err = tohtml.ReadFirstMarkdownHeader(filesList[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	for index, file := range filesList {
		if index == 0 {
			runGenerator(file, *outfile)
		} else {
			out := strings.Split(strings.ReplaceAll(strings.TrimLeft(file, "."), "README", "index"), ".")[0] + ".html"
			runGenerator(file, out)
		}
	}
}

func runGenerator(file, out string) {
	if strings.HasSuffix(file, ".html") {
		return
	}
	fmt.Println("Converting", file, "to", out)
	filesList := strings.Split(*filename, ",")
	output := tohtml.OutputHTMLOpen()
	output += tohtml.OutputHeaderOpen()
	output += tohtml.OutputTitleTag(*title)
	output += tohtml.OutputMetaTag("author", *author)
	output += tohtml.OutputMetaTag("description", findGithubRepoName())
	output += tohtml.OutputMetaTag("keywords", getCurrentBranch())
	output += tohtml.OutputCSSTag(*css)
	output += tohtml.OutputShowHiderCSSTag("showhider.css")
	if *script != "" {
		output += tohtml.OutputScriptTag(*script)
	}
	output += tohtml.OutputHeaderClose()
	output += tohtml.OutputBodyOpen()
	output += tohtml.NavigationBar(filesList)
	output += tohtml.OutputHTMLFromMarkdown(file, *title)
	output += tohtml.OutputSourceRepos()
	if !*nodonate || *donate == "" {
		output += tohtml.OutputDonationURLs(*donate, *donatemessage)
	}
	output += tohtml.License()
	if *snowflake {
		output += tohtml.Snowflake()
	}
	if *i2plink {
		output += tohtml.I2PLink()
	}
	output += tohtml.OutputBodyClose()
	output += tohtml.OutputHTMLClose()
	final := gohtml.Format(output)
	err := ioutil.WriteFile(".nojekyll", []byte{}, 0644)
	if err != nil {
		fmt.Printf("No Jekyll Error \n %s", err)
		os.Exit(1)
	}
	if out != "" && out != "-" {
		if err := ioutil.WriteFile(out, []byte(final), 0644); err != nil {
			fmt.Printf("Output Error %s", err)
			os.Exit(1)
		}
		// determine if git is installled
		cmd := exec.Command("git", "--version")
		_, err := cmd.Output()
		if err != nil {
			fmt.Printf("Git not installed \n %s\n", err)
		} else {
			gitAddCmd := exec.Command("git", "add", out, ".nojekyll", *css, "showhider.css")
			if err := gitAddCmd.Run(); err != nil {
				fmt.Printf("Git Add Error: %s", err)
				os.Exit(1)
			}
			gitCommitCmd := exec.Command("git", "commit", "-am", "update "+out)
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
	} else {
		fmt.Println(final)
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
			fmt.Printf("Found remote %s\n", remote)
			break
		}
	}
	fmt.Printf("Looked up Github Remote: %s\n", remote)
	return remote
}

func findGithubUsername() string {
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
	repo = strings.TrimRight(string(repo), ".git")
	return repo
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
