package tohtml

import (
	"log"
	"strings"
)

func NavigationBar(files []string, out string) string {
	count := len(strings.Split(out, "/"))
	navbar := "<div id=\"navbar\">"
	navbar += "<a href=\"#shownav\">Show navigation</a>"
	navbar += "<div id=\"shownav\">"
	navbar += "<div id=\"hidenav\">"
	navbar += "<ul>"
	navbar += "<li><a href=\"" + ".." + "\">" + "Up one level ^" + "</a></li>"
	for _, file := range files {
		log.Println("checking file for navbar inclusion ", file)
		if count > 1 {
			log.Println("count, file", count, file)
			if count-1 < len(strings.Split(file, "/")) {
				spl := strings.Split(file, "/")[count-1:]
				file = strings.Join(spl, "/")
			} else {
				spl := strings.Split(file, "/")[:]
				file = strings.Join(spl, "/")
			}

		}
		if strings.HasSuffix(file, ".md") {
			title := strings.Split(strings.ReplaceAll(file, "README", "index"), ".")[0]
			if title == "README" {
				title = "index"
			}
			navbar += "<li><a href=\"" + title + ".html" + "\">" + title + "</a></li>"
		} else {
			//if file == "index.html" {
			title := strings.ReplaceAll(file, "README", "index")
			if title == "README" {
				title = "index"
			}
			navbar += "<li><a href=\"" + title + "\">" + title + "</a></li>"
			//}
		}
	}
	navbar += "</ul>"
	navbar += "<br>"
	navbar += "<a href=\"#hidenav\">Hide Navigation</a>"
	navbar += "</div>"
	navbar += "</div>"
	navbar += "</div>\n"
	return navbar
}
