package tohtml

import (
	"fmt"
	"io/ioutil"
	"os"
)

func License() string {
	licensePath := "LICENSE"
	if _, err := os.Stat(licensePath); err != nil {
		licensePath = "LICENSE.md"
		if _, err := os.Stat(licensePath); err != nil {
			return ""
		}
	}
	license, err := ioutil.ReadFile(licensePath)
	if err != nil {
		fmt.Printf("Error reading license: %s", err)
		return ""
	}
	licensehtml := "<div>"
	licensehtml += "<a href=\"#show\">Show license</a>"
	licensehtml += "<div id=\"show\">"
	licensehtml += "<div id=\"hide\"><pre><code>" + string(license) + "</code></pre>"
	licensehtml += "<a href=\"#hide\">Hide license</a>"
	licensehtml += "</div>"
	licensehtml += "</div>"
	licensehtml += "</div></div>\n"
	return licensehtml
}
