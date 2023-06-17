package tohtml

import (
	"bufio"
	"context"
	"crypto/sha256"
	"embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	github "github.com/google/go-github/v45/github"
	"github.com/russross/blackfriday/v2"
	"github.com/yosssi/gohtml"
	"golang.org/x/oauth2"
)

func ReadFirstMarkdownHeader(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	markdown := blackfriday.New()
	node := markdown.Parse(bytes)
	if err != nil {
		return "", err
	}
	var title string
	// walk the nodes until we find an h1, then return the text
	node.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if node.Type == blackfriday.Heading && node.Level == 1 {
			title = strings.Split(node.FirstChild.String(), ": ")[1]
			return blackfriday.Terminate
		}
		return blackfriday.GoToNext
	})
	return title, nil
}

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

func OutputCSSTag(cssFile string) string {
	if _, err := os.Stat(cssFile); err != nil {
		err := ioutil.WriteFile(cssFile, []byte(DefaultCSS), 0644)
		if err != nil {
			fmt.Printf("Error writing default CSS file: %s", err)
		}
	} else {
		// check if the phrase: /* edgar default CSS file */ is in the file
		bytes, err := ioutil.ReadFile(cssFile)
		if err != nil {
			fmt.Printf("Error reading CSS file: %s", err)
		}
		if strings.Contains(string(bytes), "/* edgar default CSS file */") {
			err := ioutil.WriteFile(cssFile, []byte(DefaultCSS), 0644)
			if err != nil {
				fmt.Printf("Error writing default CSS file: %s", err)
			}
		}
	}
	return "		<link rel=\"stylesheet\" type=\"text/css\" href=\"" + filepath.Base(cssFile) + "\" />" + "\n"
}

func OutputShowHiderCSSTag(cssFile string) string {
	//cssFile := "showhider.css"
	if _, err := os.Stat(cssFile); err != nil {
		err := ioutil.WriteFile(cssFile, []byte(ShowHiderCSS), 0644)
		if err != nil {
			fmt.Printf("Error writing default CSS file: %s", err)
		}
	} else {
		// check if the phrase: /* edgar default CSS file */ is in the file
		bytes, err := ioutil.ReadFile(cssFile)
		if err != nil {
			fmt.Printf("Error reading CSS file: %s", err)
		}
		if strings.Contains(string(bytes), "/* edgar showhider CSS file */") {
			err := ioutil.WriteFile(cssFile, []byte(ShowHiderCSS), 0644)
			if err != nil {
				fmt.Printf("Error writing default CSS file: %s", err)
			}
		}
	}
	return "		<link rel=\"stylesheet\" type=\"text/css\" href=\"" + filepath.Base(cssFile) + "\" />" + "\n"
}

func OutputDarkLightCSSTag(cssFile string) string {
	// cssFile := "darklight.css"
	if _, err := os.Stat(cssFile); err != nil {
		err := ioutil.WriteFile(cssFile, []byte(DarkLightCSS), 0644)
		if err != nil {
			fmt.Printf("Error writing default CSS file: %s", err)
		}
	} else {
		// check if the phrase: /* edgar default CSS file */ is in the file
		bytes, err := ioutil.ReadFile(cssFile)
		if err != nil {
			fmt.Printf("Error reading CSS file: %s", err)
		}
		if strings.Contains(string(bytes), "/* edgar darklight CSS file */") {
			err := ioutil.WriteFile(cssFile, []byte(DarkLightCSS), 0644)
			if err != nil {
				fmt.Printf("Error writing default CSS file: %s", err)
			}
		}
	}
	return "		<link rel=\"stylesheet\" type=\"text/css\" href=\"" + filepath.Base(cssFile) + "\" />" + "\n"
}

func OutputScriptTag(scriptFile string) string {
	if strings.Contains(scriptFile, ",") {
		output := ""
		for _, script := range strings.Split(scriptFile, ",") {
			output += outputScriptTag(script)
		}
		return output
	}
	return outputScriptTag(scriptFile)
}

func outputScriptTag(scriptFile string) string {
	scriptType := moduleType(scriptFile)
	return "		<script type=\"" + scriptType + "\" src=\"" + scriptFile + "\"></script>" + "\n"
}

func moduleType(scriptFile string) string {
	scriptType := "text/javascript"
	readFile, err := os.Open(scriptFile)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		if strings.HasPrefix(strings.TrimSpace(fileScanner.Text()), "import") {
			scriptType := "module"
			return scriptType
		}
	}
	readFile.Close()
	return scriptType
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

func License() string {
	licensePath := "LICENSE"
	if _, err := os.Stat(licensePath); err != nil {
		licensePath = "LICENSE.md"
		if _, err := os.Stat(licensePath); err != nil {
			return ""
		}
	}
	license, err := ioutil.ReadFile(licensePath)
	if err != nil {
		fmt.Printf("Error reading license: %s", err)
		return ""
	}
	licensehtml := "<div>"
	licensehtml += "<a href=\"#show\">Show license</a>"
	licensehtml += "<div id=\"show\">"
	licensehtml += "<div id=\"hide\"><pre><code>" + string(license) + "</code></pre>"
	licensehtml += "<a href=\"#hide\">Hide license</a>"
	licensehtml += "</div>"
	licensehtml += "</div>"
	licensehtml += "</div></div>\n"
	return licensehtml
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

func NavigationBar(files []string, out string) string {
	count := len(strings.Split(out, "/"))
	navbar := "<div id=\"navbar\">"
	navbar += "<a href=\"#shownav\">Show navigation</a>"
	navbar += "<div id=\"shownav\">"
	navbar += "<div id=\"hidenav\">"
	navbar += "<ul>"
	navbar += "<li><a href=\"" + ".." + "\">" + "Up one level ^" + "</a></li>"
	for _, file := range files {
		log.Println("checking file for navbar inclusion ", file)
		if count > 1 {
			log.Println("count, file", count, file)
			if count-1 < len(strings.Split(file, "/")) {
				spl := strings.Split(file, "/")[count-1:]
				file = strings.Join(spl, "/")
			} else {
				spl := strings.Split(file, "/")[:]
				file = strings.Join(spl, "/")
			}

		}
		if strings.HasSuffix(file, ".md") {
			title := strings.Split(strings.ReplaceAll(file, "README", "index"), ".")[0]
			if title == "README" {
				title = "index"
			}
			navbar += "<li><a href=\"" + title + ".html" + "\">" + title + "</a></li>"
		} else {
			//if file == "index.html" {
			title := strings.ReplaceAll(file, "README", "index")
			if title == "README" {
				title = "index"
			}
			navbar += "<li><a href=\"" + title + "\">" + title + "</a></li>"
			//}
		}
	}
	navbar += "</ul>"
	navbar += "<br>"
	navbar += "<a href=\"#hidenav\">Hide Navigation</a>"
	navbar += "</div>"
	navbar += "</div>"
	navbar += "</div>\n"
	return navbar
}

func Snowflake() string {
	return "	<div><iframe src=\"https://snowflake.torproject.org/embed.html\" width=\"320\" height=\"240\" frameborder=\"0\" scrolling=\"no\"></iframe></div>\n"
}

//go:embed I2Plogotoopiebanner.png
var logo embed.FS

func I2PLink(dir string) string {
	i2plogopath := filepath.Join(dir, "i2plogo.png")
	file, err := logo.Open("I2Plogotoopiebanner.png")
	if err != nil {
		fmt.Printf("Error opening logo: %s", err)
		return ""
	}
	defer file.Close()
	// read the logo into a byte array
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading logo: %s", err)
		return ""
	}
	if err := ioutil.WriteFile(i2plogopath, bytes, 0644); err != nil {
		fmt.Printf("Error writing logo: %s", err)
		return ""
	}
	gitAddCmd := exec.Command("git", "add", i2plogopath)
	if err := gitAddCmd.Run(); err != nil {
		fmt.Printf("Git Add Error: %s", err)
	}
	return "	<div><a href=\"https://geti2p.net/\"><img class=\"i2plogo\" src=\"" + i2plogopath + "\"></img>I2P</a></div>\n"
}

func OutputBodyClose() string {
	return "	</body>" + "\n"
}

func OutputHTMLClose() string {
	return "</html>" + "\n"
}

func OutputSourceRepos() string {
	gitAddCmd := exec.Command("git", "remote", "-v")
	if output, err := gitAddCmd.Output(); err != nil {
		fmt.Printf("Git Add Error: %s", err)
		os.Exit(1)
	} else {
		base := strings.Split(string(output), "\t")
		if len(base) > 2 {
			final := strings.Split(base[1], " ")[0]
			final = strings.Replace(final, "git@", "slug-", 1)
			final = strings.Replace(final, "http://", "slug-", 1)
			final = strings.Replace(final, "https://", "slugs-", 1)
			final = strings.Replace(final, ":", "/", 1)
			if strings.Contains(final, "127.0.0.1") || strings.Contains(final, "localhost") || strings.Contains(final, ".i2p/") || strings.Contains(final, ".onion/") {
				final = strings.Replace(final, "slug-", "http://", 1)
			} else {
				final = strings.Replace(final, "slug-", "https://", 1)
			}

			final = strings.Replace(final, "slugs-", "https://", 1)
			ret := "<div id=\"sourcecode\">"
			ret += "    <span id=\"sourcehead\">"
			ret += "    <strong>Get the source code:</strong>"
			ret += "    </span>"
			ret += "    <ul>"
			ret += "        <li>"
			ret += "            <a href=\"" + final + "\">Source Repository: (" + final + ")</a>"
			ret += "        </li>"
			ret += "    </ul>"
			ret += "</div>"
			return ret
		}
	}
	return ""
}

func OutputDonationURLs(donate, donatemessage string) string {
	split := strings.Split(donate, ",")
	ret := "<div id=\"donatediv\">"
	ret += "  <div id=\"donatemessage\">"
	ret += "  <a href=\"#donate\">" + donatemessage + "</a>"
	ret += "  </div>"
	ret += "  <div id=\"donate\">"
	ret += "  <div id=\"hidedonate\">"
	for _, addr := range split {
		ret += "    <div class=\"wallet-addr\">"
		ret += "      <a href=\"" + addr + "\">"
		ret += strings.Split(addr, ":")[0]
		ret += "      </a>"
		ret += "      <span id=\"" + strings.Split(addr, ":")[0] + "\">"
		ret += addr
		ret += "      </span>"
		ret += "    </div>"
	}
	ret += "  <a href=\"#hidedonate\">" + "Close donation panel" + "</a>"
	ret += "  </div>"
	ret += "  </div>"
	ret += "</div>"
	ret += "</br>"
	return ret
}

func head(num int) string {
	var r string
	for i := 0; i < num; i++ {
		r += "="
	}
	return r
}

func IsOpenDirectory(path string) bool {
	if bytes, err := ioutil.ReadFile(path); err != nil {
		return true
	} else {
		if strings.Contains(string(bytes), "**Directory Listing:**") {
			return true
		}
	}
	return false
}
func OpenDirectory() string {
	wd, _ := os.Getwd()
	files, err := ioutil.ReadDir(wd)
	if err != nil {
		log.Fatal(err)
	}
	var readme string
	readme += fmt.Sprintf("%s\n", filepath.Base(wd))
	readme += fmt.Sprintf("%s\n", head(len(filepath.Base(wd))))
	readme += fmt.Sprintf("%s\n", "")
	readme += fmt.Sprintf("%s\n", "**Directory Listing:**")
	readme += fmt.Sprintf("%s\n", "")
	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file.Name(), file.IsDir())
			bytes, err := ioutil.ReadFile(file.Name())
			if err != nil {
				panic(err)
			}
			sum := fmt.Sprintf("%x", sha256.Sum256(bytes))
			readme += fmt.Sprintf(" - [%s](%s) : `%d` : `%s` - `%s`\n", file.Name(), file.Name(), file.Size(), file.Mode(), sum)
		} else {
			fmt.Println(file.Name(), file.IsDir())
			readme += fmt.Sprintf(" - [%s](%s) : `%d` : `%s`\n", file.Name(), file.Name(), file.Size(), file.Mode())
		}
	}
	return readme
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
			gitAddCmd := exec.Command("git", "add", "-v", out, ".nojekyll", css, filepath.Join(dir, "darklight.css"), filepath.Join(dir, "showhider.css"))
			if err := gitAddCmd.Run(); err != nil {
				fmt.Printf("Git Add Error: %s", err)
				os.Exit(1)
			}
		}
	} else {
		fmt.Println(final)
	}
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

func getCurrentBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.Replace(string(out), "\n", "", -1)
}
