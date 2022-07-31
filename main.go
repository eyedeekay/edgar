package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/eyedeekay/edgar/tohtml"
	"github.com/yosssi/gohtml"
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
	output += tohtml.OutputCSSTag(*css)
	output += tohtml.OutputScriptTag(*script)
	output += tohtml.OutputHeaderClose()
	output += tohtml.OutputBodyOpen()
	output += tohtml.OutputHTMLFromMarkdown(*filename, *title)
	output += tohtml.Snowflake()
	output += tohtml.OutputBodyClose()
	output += tohtml.OutputHTMLClose()
	output = gohtml.Format(output)
	if *outfile != "" && *outfile != "-" {
		err := ioutil.WriteFile(*outfile, []byte(output), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println(output)
	}
}
