package tohtml

import (
	"io/ioutil"

	"github.com/russross/blackfriday/v2"
)

func ReadFirstMarkdownHeader(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	node, err := blackfriday.Parse(bytes)
	if err != nil {
		return "", err
	}

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
