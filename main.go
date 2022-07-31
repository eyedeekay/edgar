package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/eyedeekay/edgar/tohtml"
)

var (
	filename = flag.String("filename", "README.md", "The markdown file to convert to HTML")
	author   = flag.String("author", authorDefault(), "The author of the HTML file")
	css      = flag.String("css", "style.css", "The CSS file to use, a default will be generated if one doesn't exist")
	script   = flag.String("script", "", "The script file to use")
	title    = flag.String("title", "", "The title of the HTML file")
)

func authorDefault() string {
	cmd := exec.Command("git", "config", "--get", "user.email")
	out, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.Split(string(out), "@")[0]
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
	output += tohtml.OutputCSSTag(*css)
	output += tohtml.OutputScriptTag(*script)
	output += tohtml.OutputHeaderClose()
	output += tohtml.OutputBodyOpen()
	output += tohtml.OutputBodyClose()
	output += tohtml.OutputHTMLClose()
	fmt.Println(output)
}
