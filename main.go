package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/eyedeekay/edgar/github"
	"github.com/eyedeekay/edgar/tohtml"
	"github.com/eyedeekay/edgar/torrent"
)

var (
	filename  = flag.String("filename", listAllMarkdownFiles(), "The markdown file to convert to HTML, or a comma separated list of files")
	author    = flag.String("author", authorDefault(), "The author of the HTML file")
	css       = flag.String("css", "style.css", "The CSS file to use, a default will be generated if one doesn't exist")
	script    = flag.String("script", hasScript(), "The script file to use.")
	title     = flag.String("title", "", "The title of the HTML file, if blank it will be generated from the first h1 in the markdown file.")
	outfile   = flag.String("out", "index.html", "The name of the output file(Only used for the first file, others will be named `inputfile.html`)")
	commit    = flag.Bool("checkin", false, "commit page generation changes after running")
	snowflake = flag.Bool("snowflake", true, "add a snowflake to the page footer")
	i2plink   = flag.Bool("i2plink", true, "add an i2p link to the page footer. Logo courtesy of @Shoalsteed and @mark22k")
	mirror    = flag.String("mirror", "", "Use edgar to download all github repos with pages activated to the current directory")
	nodonate  = flag.Bool("nodonate", true, "disable the donate section(change the -donate wallet addresses before setting this to true)")
	donate    = flag.String(
		"donate",
		"monero:4A2BwLabGUiU65C5JRfwXqFTwWPYNSmuZRjbTDjsu9wT6wV6kMFyXn83ydnVjVcR7BCsWh8B5b4Z9b6cmqjfZiFd9sBUpWT,bitcoin:1D1sDmyZAs5q2Lb29q8TBnGhEJK7vfp5PJ,ethereum:0x539a4356bb0566a39376CaC3F50B558F77E84eC9",
		"add donation section to cryptocurrency wallets. Use the address URL schemes, separated by commas(no spaces). Change them before running unless you want the money to go to me.",
	)
	donatemessage = flag.String("support", "Support independent development"+myDirectory(), "change message/CTA for donations section.")
	help          = flag.Bool("help", false, "Show usage.")
	i2pequiv      = flag.String("i2p-location", "", "An i2p-location http-equiv value")
	wsHost        = flag.String("webseed-host", "http://idk.i2p", "The webseed host for torrent-based mirrors")
)

var recursive = os.Getenv("EDGAR_RECURSIVE")

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
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var fileList []string

	if _, err := os.Stat("README.md"); err != nil {
		od, err := tohtml.OpenDirectory()
		if err != nil {
			panic(err)
		}
		err = os.WriteFile("README.md", []byte(od), 0o644)
		if err != nil {
			panic(err)
		}
	}
	if tohtml.IsOpenDirectory("README.md") {
		od, err := tohtml.OpenDirectory()
		if err != nil {
			panic(err)
		}
		err = os.WriteFile("README.md", []byte(od), 0o644)
		if err != nil {
			panic(err)
		}
	}

	fileList = append(fileList, "README.md")

	if recursive != "" {
		generateNonRecursive(fileList)
	} else {
		fileList = generateRecursive(files, fileList)
	}
	return strings.Join(fileList, ",")
}

func generateRecursive(files []os.DirEntry, fileList []string) []string {
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
		} else {
			if _, err := os.Stat(filepath.Join(file.Name(), "index.html")); err == nil {
				fileList = append(fileList, filepath.Join(file.Name(), "index.html"))
			}
		}
	}
	if _, err := os.Stat("docs"); err == nil {
		docs, err := os.ReadDir("docs")
		if err != nil {
			log.Fatal(err)
		}
		tohtml.OutputCSSTag("docs/style.css")
		tohtml.OutputShowHiderCSSTag("docs/showhider.css")
		gitAddCmd := exec.Command("git", "add", "docs/style.css", "docs/showhider.css")
		if err := gitAddCmd.Run(); err != nil {
			fmt.Printf("Git Add Error: %s", err)
			os.Exit(1)
		}
		// var fileList []string
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
	if _, err := os.Stat("doc"); err == nil {
		docs, err := os.ReadDir("doc")
		if err != nil {
			log.Fatal(err)
		}
		tohtml.OutputCSSTag("doc/style.css")
		tohtml.OutputShowHiderCSSTag("doc/showhider.css")
		gitAddCmd := exec.Command("git", "add", "doc/style.css", "doc/showhider.css")
		if err := gitAddCmd.Run(); err != nil {
			fmt.Printf("Git Add Error: %s", err)
			os.Exit(1)
		}
		// var fileList []string
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
	return fileList
}

func generateNonRecursive(fileList []string) {
	log.Println("walking dir recursively...")
	err := filepath.Walk(".",
		func(path string, file os.FileInfo, err error) error {
			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}
				if err == filepath.SkipDir {
					return nil
				}
				return err
			}
			if !file.IsDir() {
				log.Println("recursing to path", path)
				if strings.HasSuffix(path, ".md") {
					if path != "README.md" {
						fileList = append(fileList, filepath.Join(path))
					}
				} else if strings.HasSuffix(file.Name(), ".html") {
					mdExtension := strings.ReplaceAll(file.Name(), ".html", ".md")
					if _, err := os.Stat(mdExtension); err != nil {
						fileList = append(fileList, filepath.Join(path))
					}
				}
			} else {
				if _, err := os.Stat(filepath.Join(path, "index.html")); err == nil {
					fileList = append(fileList, filepath.Join(path, "index.html"))
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func authorDefault() string {
	user := tohtml.FindGithubUsername()
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
	if *help {
		showHelp()
		return
	}
	if *mirror != "" {
		github.Mirror(*mirror, "")
	}
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
	// get the name of the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	currentDir = filepath.Base(currentDir)
	files, err := github.DownloadLatestReleaseAssets(*author, currentDir, *author, "")
	if err != nil {
		log.Println(err)
	}
	for _, file := range files {
		log.Println("Generating torrent for:", file)
		err := torrent.CreateTorrent(torrent.CreateTorrentConfig{
			PieceLength: 256,
			Name:        file,
			Output:      file + ".torrent",
			Comment:     "",
			Announces:   []string{"http://w7tpbzncbcocrqtwwm3nezhnnsw4ozadvi2hmvzdhrqzfxfum7wa.b32.i2p/a"},
			WebSeeds:    []string{*wsHost + "/" + currentDir + "/" + file},
			NoDate:      true,
		})
		if err != nil {
			log.Println(err)
		}
	}
	if *commit {
		tohtml.CommitMessage()
	}
}

func runGenerator(file, out string) {
	tohtml.RunGenerator(file, out, *filename, *title, *author, *css, *script, *donate, *donatemessage, *nodonate, *snowflake, *i2plink, *i2pequiv)
}
