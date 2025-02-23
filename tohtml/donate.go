package tohtml

import "strings"

func OutputDonationURLs(donate, donatemessage string) string {
	split := strings.Split(donate, ",")
	ret := "<div id=\"donatediv\">"
	ret += "  <div id=\"donatemessage\">"
	ret += "  <a href=\"#donate\">" + donatemessage + "</a>"
	ret += "  </div>"
	ret += "  <div id=\"donate\">"
	ret += "  <div id=\"hidedonate\">"
	for _, addr := range split {
		ret += "    <div class=\"wallet-addr\">"
		ret += "      <a href=\"" + addr + "\">"
		ret += strings.Split(addr, ":")[0]
		ret += "      </a>"
		ret += "      <span id=\"" + strings.Split(addr, ":")[0] + "\">"
		ret += addr
		ret += "      </span>"
		ret += "    </div>"
	}
	ret += "  <a href=\"#hidedonate\">" + "Close donation panel" + "</a>"
	ret += "  </div>"
	ret += "  </div>"
	ret += "</div>"
	ret += "</br>"
	return ret
}
