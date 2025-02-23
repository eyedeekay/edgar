package tohtml

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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
			err := ioutil.WriteFile(cssFile, []byte(DefaultCSS), 0644)
			if err != nil {
				fmt.Printf("Error writing default CSS file: %s", err)
			}
		}
	}
	return "		<link rel=\"stylesheet\" type=\"text/css\" href=\"" + filepath.Base(cssFile) + "\" />" + "\n"
}

func OutputShowHiderCSSTag(cssFile string) string {
	//cssFile := "showhider.css"
	if _, err := os.Stat(cssFile); err != nil {
		err := ioutil.WriteFile(cssFile, []byte(ShowHiderCSS), 0644)
		if err != nil {
			fmt.Printf("Error writing default CSS file: %s", err)
		}
	} else {
		// check if the phrase: /* edgar default CSS file */ is in the file
		bytes, err := ioutil.ReadFile(cssFile)
		if err != nil {
			fmt.Printf("Error reading CSS file: %s", err)
		}
		if strings.Contains(string(bytes), "/* edgar showhider CSS file */") {
			err := ioutil.WriteFile(cssFile, []byte(ShowHiderCSS), 0644)
			if err != nil {
				fmt.Printf("Error writing default CSS file: %s", err)
			}
		}
	}
	return "		<link rel=\"stylesheet\" type=\"text/css\" href=\"" + filepath.Base(cssFile) + "\" />" + "\n"
}

func OutputDarkLightCSSTag(cssFile string) string {
	// cssFile := "darklight.css"
	if _, err := os.Stat(cssFile); err != nil {
		err := ioutil.WriteFile(cssFile, []byte(DarkLightCSS), 0644)
		if err != nil {
			fmt.Printf("Error writing default CSS file: %s", err)
		}
	} else {
		// check if the phrase: /* edgar default CSS file */ is in the file
		bytes, err := ioutil.ReadFile(cssFile)
		if err != nil {
			fmt.Printf("Error reading CSS file: %s", err)
		}
		if strings.Contains(string(bytes), "/* edgar darklight CSS file */") {
			err := ioutil.WriteFile(cssFile, []byte(DarkLightCSS), 0644)
			if err != nil {
				fmt.Printf("Error writing default CSS file: %s", err)
			}
		}
	}
	return "		<link rel=\"stylesheet\" type=\"text/css\" href=\"" + filepath.Base(cssFile) + "\" />" + "\n"
}
