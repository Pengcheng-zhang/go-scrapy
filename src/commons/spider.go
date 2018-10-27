package commons

import (
	"github.com/PuerkitoBio/goquery"
	"models"
	"strings"
	"time"
	"errors"
)

type Spider interface {
	SpiderUrl() error
}

func NewSpider(from string) (Spider, error) {
	DEBUG("new spider xxx")
	switch from {
	case "booktxt":
		return new(BookTextSpider), nil
	case "huxiu":
		return new(HuxiuSpider), nil
	case "36kr":
		return new(ThreeSixKrSpider), nil
	case "opFriends":
		return new(OpFriends), nil	
	default:
		return nil, errors.New("系统暂未处理该类型的配置文件")
	}
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

func SpiderArticle(url, titleTag, contentTag string, ch chan string) {
	defer func(){
		<- ch
	}()
	DEBUG("spider:",url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		DEBUG("get article url failed:", err)
		return
	}
	article := models.ArticleModel{}
	articleTitle := doc.Find(titleTag).Text()
	articleContent, err := doc.Find(contentTag).Html()
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