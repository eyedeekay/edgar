package tohtml

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func OutputSourceRepos() string {
	gitAddCmd := exec.Command("git", "remote", "-v")
	if output, err := gitAddCmd.Output(); err != nil {
		fmt.Printf("Git Add Error: %s", err)
		os.Exit(1)
	} else {
		base := strings.Split(string(output), "\t")
		if len(base) > 2 {
			final := strings.Split(base[1], " ")[0]
			final = strings.Replace(final, "git@", "slug-", 1)
			final = strings.Replace(final, "http://", "slug-", 1)
			final = strings.Replace(final, "https://", "slugs-", 1)
			final = strings.Replace(final, ":", "/", 1)
			if strings.Contains(final, "127.0.0.1") || strings.Contains(final, "localhost") || strings.Contains(final, ".i2p/") || strings.Contains(final, ".onion/") {
				final = strings.Replace(final, "slug-", "http://", 1)
			} else {
				final = strings.Replace(final, "slug-", "https://", 1)
			}

			final = strings.Replace(final, "slugs-", "https://", 1)
			ret := "<div id=\"sourcecode\">"
			ret += "    <span id=\"sourcehead\">"
			ret += "    <strong>Get the source code:</strong>"
			ret += "    </span>"
			ret += "    <ul>"
			ret += "        <li>"
			ret += "            <a href=\"" + final + "\">Source Repository: (" + final + ")</a>"
			ret += "        </li>"
			ret += "    </ul>"
			ret += "</div>"
			return ret
		}
	}
	return ""
}
