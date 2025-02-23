package tohtml

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

//go:embed I2Plogotoopiebanner.png
var logo embed.FS

func I2PLink(dir string) string {
	i2plogopath := filepath.Join(dir, "i2plogo.png")
	file, err := logo.Open("I2Plogotoopiebanner.png")
	if err != nil {
		fmt.Printf("Error opening logo: %s", err)
		return ""
	}
	defer file.Close()
	// read the logo into a byte array
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading logo: %s", err)
		return ""
	}
	if err := ioutil.WriteFile(i2plogopath, bytes, 0644); err != nil {
		fmt.Printf("Error writing logo: %s", err)
		return ""
	}
	gitAddCmd := exec.Command("git", "add", i2plogopath)
	if err := gitAddCmd.Run(); err != nil {
		fmt.Printf("Git Add Error: %s", err)
	}
	return "	<div><a href=\"https://geti2p.net/\"><img class=\"i2plogo\" src=\"" + filepath.Base(i2plogopath) + "\"></img>I2P</a></div>\n"
}
