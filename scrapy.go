package main

import(
	"commons"
)

func main() {
	s, err := commons.NewSpider("huxiu")
	if err != nil {
		commons.ERROR("new spider error:", err.Error())
		return
	}
	err = s.SpiderUrl("https://www.huxiu.com")
	if err != nil {
		commons.ERROR("new document error:", err.Error())
		return
	}
}