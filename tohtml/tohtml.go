package tohtml

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	github "github.com/google/go-github/v45/github"
	"github.com/russross/blackfriday/v2"
	"github.com/yosssi/gohtml"
	"golang.org/x/oauth2"
)

func OutputHTMLOpen() string {
	return "<html>" + "\n"
}

func OutputHeaderOpen() string {
	return "	<head>" + "\n"
}

func OutputTitleTag(title string) string {
	return "		<title>" + strings.TrimRight(strings.TrimLeft(title, "'"), "'") + "</title>" + "\n"
}

func OutputMetaTag(name, content string) string {
	return "		<meta name=\"" + name + "\" content=\"" + content + "\" />" + "\n"
}

func OutputMetaEquiv(name, content string) string {
	return "		<meta http-equiv=\"" + name + "\" content=\"" + content + "\" />" + "\n"
}

func OutputHeaderClose() string {
	return "	</head>" + "\n"
}

func OutputBodyOpen() string {
	return "<body>    <input type=\"checkbox\" id=\"checkboxDarkLight\"><div class=\"container\">" + "\n"
}

func outputHTMLFromMarkdown(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s", err)
		panic(err)
	}

	html := blackfriday.Run(bytes, blackfriday.WithExtensions(blackfriday.CommonExtensions))
	//unsafe := github_flavored_markdown.Markdown(bytes)
	//html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return "<a id=\"returnhome\" href=\"/\">/</a>" + string(html) + "\n"
}

func OutputHTMLFromMarkdown(filename, title string) string {
	html := outputHTMLFromMarkdown(filename)
	title = strings.TrimRight(strings.TrimLeft(title, "'"), "'")
	if filename == "README.md" {
		//html = strings.Replace(html, title, "<a href=\".\">"+title+"</a>", 1)
	}
	return html
}

func folderList() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func OutputBodyClose() string {
	return "	</body>" + "\n"
}

func OutputHTMLClose() string {
	return "</html>" + "\n"
}

func head(num int) string {
	var r string
	for i := 0; i < num; i++ {
		r += "="
	}
	return r
}

func RunGenerator(file, out, filename, title, author, css, script, donate, donatemessage string, nodonate, snowflake, i2plink bool, i2pequiv string) {
	if strings.HasSuffix(file, ".html") {
		return
	}
	fmt.Println("Converting", file, "to", out)
	dir := filepath.Dir(file)
	filesList := strings.Split(filename, ",")
	output := OutputHTMLOpen()
	output += OutputHeaderOpen()
	output += OutputTitleTag(title)
	output += OutputMetaTag("author", author)
	output += OutputMetaTag("description", findGithubRepoName())
	output += OutputMetaTag("keywords", getCurrentBranch())
	if i2pequiv != "" {
		output += OutputMetaEquiv("i2p-location", i2pequiv)
	}
	output += OutputCSSTag(filepath.Join(dir, css))
	output += OutputShowHiderCSSTag(filepath.Join(dir, "showhider.css"))
	output += OutputDarkLightCSSTag(filepath.Join(dir, "darklight.css"))
	if script != "" {
		output += OutputScriptTag(script)
	}
	output += OutputHeaderClose()
	output += OutputBodyOpen()
	output += NavigationBar(filesList, out)
	output += OutputHTMLFromMarkdown(file, title)
	output += OutputSourceRepos()
	if !nodonate || donate == "" {
		output += OutputDonationURLs(donate, donatemessage)
	}
	output += License()
	if snowflake {
		output += Snowflake()
	}
	if i2plink {
		output += I2PLink(dir)
	}
	output += OutputBodyClose()
	output += OutputHTMLClose()
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
			gitAddCmd := exec.Command("git", "add", "-v", out, ".nojekyll", filepath.Join(dir, css), filepath.Join(dir, "darklight.css"), filepath.Join(dir, "showhider.css"))
			if err := gitAddCmd.Run(); err != nil {
				fmt.Printf("Git Add Error: %s", err)
				os.Exit(1)
			}
		}
	} else {
		fmt.Println(final)
	}
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
