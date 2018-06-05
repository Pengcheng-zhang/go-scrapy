package commons

import(
	"github.com/PuerkitoBio/goquery"
	"models"
	"strings"
)

type BookTextSpider struct {
	
}

func getBookByName(bookname string) bool{
	book := models.BookModel{}
	GetDbInstance().Where("title = ?", bookname).First(&book)
	if book.Id > 0 {
		return true
	}
	return false
}

func addBook(book *models.BookModel) bool{
	DEBUG("add book:", book.Title)
	err := GetDbInstance().Create(&book).Error
	if err == nil {
		return true
	}
	DEBUG("add book fail:", err.Error())
	return false
}

func addChapter(chapter models.ChapterModel) bool {
	err := GetDbInstance().Create(&chapter).Error
	if err == nil {
		return true
	}
	DEBUG("add chapter fail:", err.Error())
	return false
}

func spiderChapter(bookId int, chapter models.ChapterModel) {
	if chapter.Url != "" {
		doc, err := goquery.NewDocument(chapter.Url)
		if err != nil {
			DEBUG("get chapter details error:", bookId, err)
			return
		}
		content := doc.Find("#content").Text()
		content = GbKToUTF8(content)
		content = strings.Replace(content, "聽", " ", -1)
		chapter.Content = content
		addChapter(chapter)
	}
}

func(this *BookTextSpider) SpiderUrl(url string) error {
	book := models.BookModel{}
	book.Url = url
	doc, err := goquery.NewDocument(url)
	if err != nil {
		DEBUG("get booktxt url failed:", err)
	}
	bookname := GbKToUTF8(doc.Find("#info h1").Text())

	bFind := getBookByName(bookname)
	if bFind == false {
		book = models.BookModel{Title: bookname}
		addBook(&book)
	}
	var chapters []models.ChapterModel
	doc.Find("#list dd").Each(func(i int, contentSelection *goquery.Selection){
		if i < 9 {
			return
		}
		pre := i - 9
		next := i - 7
		name := GbKToUTF8(contentSelection.Find("a").Text())

		href, _ := contentSelection.Find("a").Attr("href")
		chapter := models.ChapterModel{BookId: book.Id, Name: name, Url: "http://www.booktxt.net"+ href, Order: i-8, Pre:pre, Next: next}
		chapters = append(chapters, chapter)
	})
	for _, chapter := range chapters {
		go spiderChapter(book.Id, chapter)
	}
	return nil
}