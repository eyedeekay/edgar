package tohtml

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
