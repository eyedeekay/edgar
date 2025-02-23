package tohtml

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func OutputScriptTag(scriptFile string) string {
	if strings.Contains(scriptFile, ",") {
		output := ""
		for _, script := range strings.Split(scriptFile, ",") {
			output += outputScriptTag(script)
		}
		return output
	}
	return outputScriptTag(scriptFile)
}

func outputScriptTag(scriptFile string) string {
	scriptType := moduleType(scriptFile)
	return "		<script type=\"" + scriptType + "\" src=\"" + scriptFile + "\"></script>" + "\n"
}

func moduleType(scriptFile string) string {
	scriptType := "text/javascript"
	readFile, err := os.Open(scriptFile)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		if strings.HasPrefix(strings.TrimSpace(fileScanner.Text()), "import") {
			scriptType := "module"
			return scriptType
		}
	}
	readFile.Close()
	return scriptType
}
