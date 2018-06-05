package commons

import(
	"github.com/PuerkitoBio/goquery"
	"models"
	"strings"
	"time"
	"errors"
)

type HuxiuSpider struct {
	
}

func getArticleByName(title string) bool{
	article := models.ArticleModel{}
	GetDbInstance().Where("title = ?", title).First(&article)
	if article.Id > 0 {
		return true
	}
	return false
}

func addArticle(article *models.ArticleModel) bool{
	result := getArticleByName(article.Title)
	if result == true {
		DEBUG("add article failed:", article.Title)
		return false
	}
	err := GetDbInstance().Create(&article).Error
	if err == nil {
		return true
	}
	DEBUG("add article fail:", err.Error())
	return false
}

func spiderArticle(url string, ch chan string) {
	defer func(){
		<- ch
	}()
	DEBUG("spider:",url)
	doc, err := goquery.NewDocument("https://www.huxiu.com"+url)
	if err != nil {
		DEBUG("get article url failed:", err)
		return
	}
	article := models.ArticleModel{}
	articleTitle := doc.Find(".article-wrap .t-h1").Text()
	articleContent, err := doc.Find(".article-content-wrap").Html()
	articleTitle = strings.Replace(articleTitle, "\n","", -1)
	articleTitle = strings.Replace(articleTitle, " ", "", -1)
	//时讯
	article.Type = 23 
	article.Title = articleTitle
	article.Content = articleContent
	article.CreatorId = 2
	article.Status = "P"
	article.LastReplyTime = time.Now().Format("2006-01-02 15:04:05")
	addArticle(&article)
}

func(this *HuxiuSpider) SpiderUrl(url string) error {
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
	for key := range articleLinks {
		ch <- time.Now().String()
		go spiderArticle(articleLinks[key], ch)
	}

	for i:=0;i<2;i++ {
		ch <- time.Now().String()
	}
	close(ch)
	return nil
}