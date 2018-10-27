package main

import(
	"commons"
)

func main() {
	s, err := commons.NewSpider("opFriends")
	if err != nil {
		commons.ERROR("new spider error:", err.Error())
		return
	}
	err = s.SpiderUrl()
	if err != nil {
		commons.ERROR("new document error:", err.Error())
		return
	}
	return
}