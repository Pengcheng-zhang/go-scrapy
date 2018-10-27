package commons
//面向对象数据爬取
import(
	"fmt"
	"models"
	"github.com/PuerkitoBio/goquery"
	"time"
	// "strings"
	"regexp"
	"strconv"
)

type OpFriends struct {
	
}

func(this *OpFriends) SpiderUrl() error {
	url := "http://date.jobbole.com/page/"
	page := 1
	result := queryPageInfo(url + strconv.Itoa(page))
	if result == nil {
		for page=2;result == nil&&page <5;page++ {
			fmt.Println(page)
			result = queryPageInfo(url + strconv.Itoa(page))
			fmt.Println(result)
		}
	}
	fmt.Println(result)
	return result
}

func queryPageInfo(url string) error{
	doc, err := goquery.NewDocument(url)
	if err != nil {
		DEBUG("get jobbole url failed:", err)
		return err
	}
	var jobboleLinks []string
	items := doc.Find(".media-body")
	items.Each(func(index int, sel *goquery.Selection) {
		href,exist1 := sel.Find(".p-tit>a").Attr("href")
		if exist1 {
			jobboleLinks = append(jobboleLinks, href)
		}
	})
	
	var ch = make(chan string, 5)
	for _,linkUrl := range jobboleLinks {
		ch <- time.Now().String()
		go getFriendsInfo(linkUrl, ch)
	}

	for i:=0;i<5;i++ {
		ch <- time.Now().String()
	}
	close(ch)
	return nil
}
func getFriendsInfo(url string, ch chan string) {
	defer func(){
		<- ch
	}()
	DEBUG("spider:",url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		DEBUG("get article url failed:", err)
		return
	}
	friends := models.OpfriendsModel{}
	articleTitle := doc.Find(".p-single .p-tit-single").Text()
	if articleTitle == "" {
		return
	}
	articleContent, err := doc.Find(".p-single .p-entry>p").Html()
	if articleContent == "" {
		return
	}
	reg := regexp.MustCompile(`\n`)
	var subMatchString [][]string
	var matchStrings []string
	var matchString string
	matchLength := 0
	regString := []string{`出生(.*?)：(.*?)<br/>`,`身高(.*?)：(.*?)<br/>`,
	`体重(.*?)：(.*?)<br/>`,`学历(.*?)：(.*?)<br/>`,`(所在地|城市)(.*?)：(.*?)<br/>`,
	`户籍(.*?)：(.*?)<br/>`,`籍贯(.*?)：(.*?)<br/>`,`职业(.*?)：(.*?)<br/>`,
	`收入(.*?)：(.*?)<br/>`,`兴趣(.*?)：(.*?)<br/>`, `异地(.*?)：(.*?)<br/>`,
	`结婚(.*?)：(.*?)<br/>`, `脱颖(.*?)：(.*?)<br/>`, `(最低|另一半)(.*?)：(.*?)<br/>`,
	`自我(.*?)：(.*?)<br/>`, `离异(.*?)：(.*?)<br/>`, `几个小孩(.*?)：(.*?)<br/>`,
	`特(殊|别)要求(.*?)：(.*?)<br/>`, `父母(.*?)：(.*?)<br/>`, `(兄弟|姐妹)(.*?)：(.*?)<br/>`}
	for regIndex,regType := range regString {
		reg = regexp.MustCompile(regType)
		subMatchString = reg.FindAllStringSubmatch(articleContent, -1)
		if subMatchString == nil {
			continue
		}
		matchStrings = subMatchString[0]
		matchLength = len(matchStrings)
		matchString = matchStrings[matchLength-1]
		switch regIndex {
		case 0:
			friends.BirthDay = matchString
		case 1:
			friends.Height = matchString
		case 2:
			friends.Weight = matchString
		case 3:
			friends.Education = matchString
		case 4:
			friends.CurrentCity = matchString	
		case 5:
			friends.RegistCity = matchString
		case 6:
			friends.BornCity = matchString	
		case 7:
			friends.Profession = matchString
		case 8:
			friends.InCome = matchString
		case 9:
			friends.Interest = matchString
		case 10:
			friends.PlaceOther = matchString
		case 11:
			friends.MarryYears = matchString
		case 12:
			friends.ShowMeSpecial = matchString
		case 13:
			friends.RequestBase = matchString
		case 14:
			friends.SelfRecommend = matchString
		case 15:
			friends.Married = matchString
		case 16:
			friends.ChildNum = matchString
		case 17:
			friends.RequestOther = matchString
		case 18:
			friends.Parents = matchString	 	
		}
	}

	// catch person images
	reg = regexp.MustCompile(`src="(.*?)"`)
	subMatchString = reg.FindAllStringSubmatch(articleContent, -1)
	if subMatchString == nil {
		return
	}
	for key,value := range subMatchString {
		switch key {
		case 0:
			friends.ImageUrlOne = value[1]
		case 1:
			friends.ImageUrlTwo = value[1]
		case 2:
			friends.ImageUrlThree = value[1]
		case 3:
			friends.ImageUrlFour = value[1]
		default:
			break;
		}
	}
	friends.SourceDest = url
	addFriends(&friends)
}

func addFriends(opFriends *models.OpfriendsModel) bool{
	err := GetDbInstance().Create(&opFriends).Error
	if err == nil {
		return true
	}
	DEBUG("add opfriends fail:", err.Error())
	return false
}
