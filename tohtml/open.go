package tohtml

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
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

func OpenDirectory() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	files, err := os.ReadDir(wd)
	if err != nil {
		return "", err
	}
	readme := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", filepath.Base(wd), head(len(filepath.Base(wd))), "", "**Directory Listing:**", "")
	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file.Name(), file.IsDir())
			bytes, err := os.ReadFile(file.Name())
			if err != nil {
				return "", err
			}
			sum := fmt.Sprintf("%x", sha256.Sum256(bytes))
			readme += fmt.Sprintf(" - [%s %s](%s)\n", file.Name(), sum, file.Name())
		} else {
			fmt.Println(file.Name(), file.IsDir())
			readme += fmt.Sprintf(" - [%s](%s)\n", file.Name(), file.Name())
		}
	}
	return readme, err
}
