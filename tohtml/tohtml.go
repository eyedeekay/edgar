package tohtml

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
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
			fmt.Printf("CSS file already contains default CSS, updating it\n")
			err := ioutil.WriteFile(cssFile, []byte(DefaultCSS), 0644)
			if err != nil {
				fmt.Printf("Error writing default CSS file: %s", err)
			}
		}
	}
	return "		<link rel=\"stylesheet\" type=\"text/css\" href=\"" + cssFile + "\" />" + "\n"
}

func OutputScriptTag(scriptFile string) string {
	return "		<script type=\"text/javascript\" src=\"" + scriptFile + "\"></script>" + "\n"
}

func OutputHeaderClose() string {
	return "	</head>" + "\n"
}

func OutputBodyOpen() string {
	return "<body>" + "\n"
}

func outputHTMLFromMarkdown(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s", err)
		panic(err)
	}
	unsafe := blackfriday.Run(bytes)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return string(html) + "\n"
}

func OutputHTMLFromMarkdown(filename, title string) string {
	html := outputHTMLFromMarkdown(filename)
	title = strings.TrimRight(strings.TrimLeft(title, "'"), "'")
	if filename == "README.md" {
		html = strings.Replace(html, title, "<a href=\"/\">"+title+"</a>", 1)
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
	licensehtml += "</div>\n"
	return licensehtml
}

func NavigationBar(files []string) string {
	navbar := "<div id=\"navbar\">"
	navbar += "<a href=\"#shownav\">Show navigation</a>"
	navbar += "<div id=\"shownav\">"
	navbar += "<div id=\"hidenav\">"
	navbar += "<ul>"
	for _, file := range files {
		title := strings.Split(file, ".")[0]
		if title == "README" {
			title = "index"
		}
		navbar += "<li><a href=\"" + title + ".html" + "\">" + title + "</a></li>"
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

func I2PLink() string {
	i2plogopath := "i2plogo.png"
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
	return "	<div><a href=\"https://geti2p.net/\"><img src=\"" + i2plogopath + "\"></img>I2P</a></div>\n"
}

func OutputBodyClose() string {
	return "	</body>" + "\n"
}

func OutputHTMLClose() string {
	return "</html>" + "\n"
}
