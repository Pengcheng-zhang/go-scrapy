package commons

import(
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"time"
	"strings"
)

type ThreeSixKrSpider struct {
	
}

func(this *ThreeSixKrSpider) SpiderUrl() error {
	url := "https://36kr.com"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		DEBUG("get huxiu url failed:", err)
		return err
	}
	var articleLinks []string
	items := doc.Find(".pc_list li")
	items.Each(func(index int, sel *goquery.Selection) {
		if len(articleLinks) > 2 {
			return
		}
		href,exist1 := sel.Find(".inner_li>a").Attr("href")
		fmt.Println("href:", href)
		if exist1 && !strings.Contains(href, "video") {
			articleLinks = append(articleLinks, href)
		}
	})
	
	var ch = make(chan string, 2)
	titleTag := ".pc_list li:first .am-text-break>a>span"
	contentTag := ""
	for key := range articleLinks {
		ch <- time.Now().String()
		go SpiderArticle(url+articleLinks[key], titleTag, contentTag, ch)
	}

	for i:=0;i<2;i++ {
		ch <- time.Now().String()
	}
	close(ch)
	return nil
}