package main

import (
	"flag"
	"fmt"
	"os"
)

func showHelp() {
	top := fmt.Sprintf(`Edgar(Everything does get a README): Static Site Generator for the Paradoxically Prolific
=========================================================================================
This is a static site generator which is intended to generate a page for a piece of software from markdown documents stored inside its VCS.
It will generate a page from any directory containing markdown files, but it's especially useful for Github Pages. If you use the .nojekyll option, it will bypass Jekyll processing, allowing you to use directories and files that Jekyll would normally ignore, such as those starting with an underscore.

`)
	block := fmt.Sprintf("\n```\n")
	os.Stderr = os.Stdout
	help := fmt.Sprintf("%s%s%s%s", top, block, flag.Lookup("help").Usage, block)
	fmt.Println(help)
}
