package main

import(
	"fmt"
	"common"
)

func main() {
	s, err := common.NewSpider("booktxt")
	if err != nil {
		common.ERROR("new spider error:", err.Error())
	}
	err = s.SpiderUrl("http://www.booktxt.net/2_2219")
	if err != nil {
		common.ERROR("new document error:", err.Error())
	}
	var str string
	fmt.Scan(&str)
}