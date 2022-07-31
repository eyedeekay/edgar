package tohtml

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
	return "<html>"
}

func OutputHeaderOpen() string {
	return "<head>"
}

func OutputTitleTag(title string) string {
	return "<title>" + title + "</title>"
}

func OutputMetaTag(name, content string) string {
	return "<meta name=\"" + name + "\" content=\"" + content + "\" />"
}

func OutputCSSTag(cssFile string) string {
	if _, err := os.Stat(cssFile); err != nil {
		err := ioutil.WriteFile(cssFile, []byte(DefaultCSS), 0644)
		if err != nil {
			fmt.Errorf("Error writing default CSS file: %s", err)
		}
	}
	return "<link rel=\"stylesheet\" type=\"text/css\" href=\"" + cssFile + "\" />"
}

func OutputScriptTag(scriptFile string) string {
	return "<script type=\"text/javascript\" src=\"" + scriptFile + "\"></script>"
}

func OutputHeaderClose() string {
	return "</head>"
}

func OutputBodyOpen() string {
	return "<body>"
}

func OutputBodyClose() string {
	return "</body>"
}

func OutputHTMLClose() string {
	return "</html>"
}
