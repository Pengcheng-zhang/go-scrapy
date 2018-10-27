package commons

import(
	"github.com/PuerkitoBio/goquery"
	"time"
	"errors"
)

type HuxiuSpider struct {
	
}

func(this *HuxiuSpider) SpiderUrl() error {
	url := "https://www.huxiu.com"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		DEBUG("get huxiu url failed:", err)
		return err
	}
	var articleLinks []string
	articleLink1,exist1 := doc.Find(".big-pic a").Attr("href")
	articleLink2,exist2 := doc.Find(".big2-pic-index-top>a").Attr("href")
	if exist1 && exist2 {
		articleLinks = append(articleLinks, articleLink1, articleLink2)
	}else{
		return errors.New("未找到")
	}
	
	var ch = make(chan string, 2)
	titleTag := ".article-wrap .t-h1"
	contentTag := ".article-content-wrap"
	for key := range articleLinks {
		ch <- time.Now().String()
		go SpiderArticle(url+ articleLinks[key], titleTag, contentTag, ch)
	}

	for i:=0;i<2;i++ {
		ch <- time.Now().String()
	}
	close(ch)
	return nil
}