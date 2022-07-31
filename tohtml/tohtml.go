package tohtml

import (
	"fmt"
	"io/ioutil"
	"os"
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
	html = strings.Replace(html, title, "<a href=\"/\">"+title+"</a>", 1)
	return html
}

func License() string {
	license, err := ioutil.ReadFile("LICENSE")
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
	navbar := "<div>"
	navbar += "<a href=\"#shownav\">Show navigation</a>"
	navbar += "<div id=\"shownav\">"
	navbar += "<div id=\"hidenav\">"
	for _, file := range files {
		title := strings.Split(file, ".")[0]
		if title == "README" {
			title = "index"
		}
		navbar += "<a href=\"" + title + ".html" + "\">" + title + "</a><br>"
	}
	navbar += "<a href=\"#hidenav\">Hide Navigation</a>"
	navbar += "</div>"
	navbar += "</div>"
	navbar += "</div>\n"
	return navbar
}

func Snowflake() string {
	return "	<div><iframe src=\"https://snowflake.torproject.org/embed.html\" width=\"320\" height=\"240\" frameborder=\"0\" scrolling=\"no\"></iframe></div>\n"
}

func OutputBodyClose() string {
	return "	</body>" + "\n"
}

func OutputHTMLClose() string {
	return "</html>" + "\n"
}
