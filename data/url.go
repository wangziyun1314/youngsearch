package data

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strings"
	"youngsearch/utils"
)

type Page struct {
	gorm.Model
	Id      int
	Url     string //网页的url
	Title   string //网页的标题
	Content string //网页的内容
}

// Page GetUrls 对输入的url进行解析，得到其中的标题和内容和其中包含的url
//func GetUrls(url string) [][2]string {
//	res := make([][2]string, 0)
//	page, err := goquery.ParseUrl(url)
//	if err != nil {
//		fmt.Println(url)
//		fmt.Println(err)
//		fmt.Println("打开网页失败")
//		return res
//	}
//	text := page.Find("title")
//	builder := strings.Builder{}
//	for i := 0; i < text.Length(); i++ {
//		fmt.Fprint(&builder, text.Eq(i).Text())
//	}
//	array1 := [2]string{url, builder.String()}
//	res = append(res, array1)
//	//content := page.Find("p")
//	//sb := strings.Builder{}
//	//for i := 0; i < content.Length(); i++ {
//	//	fmt.Fprint(&sb, content.Eq(i).Text())
//	////}
//	//res = append(res, sb.String())
//	t := page.Find("a")
//	for i := 0; i < t.Length(); i++ {
//		d := t.Eq(i).Attr("href")
//		strings.Replace(d, "\t", "", -1)
//		strings.Replace(d, "\n", "", -1)
//		if strings.HasPrefix(d, "http") {
//			text := t.Eq(i).Text()
//			array := [2]string{d, text}
//			res = append(res, array)
//		}
//
//	}
//	//fmt.Println(len(res))
//	return res
//}

//Page 存储爬取到的网页信息

//SearchForPage 根据给定的数组来进行层序遍历爬取网页的信息，url[0] 是网页的url
// url[1] 是网页的title
func SearchForPage(url [2]string) {
	res := make([]*Page, 0)
	bloomFilter := utils.NewBloomFilter(nil)
	index := 0
	queue := utils.NewQueue()
	queue.Offer(url)
	for !queue.IsEmpty() {
		stirs := queue.Poll()
		u := stirs.([2]string)
		if bloomFilter.Check(u[0]) {
			continue
		}
		bloomFilter.Insert(u[0])
		urls := GetUrls(u[0])
		if len(urls) == 0 {
			continue
		}
		p := &Page{Id: index, Title: urls[0][1], Content: u[1], Url: u[0]}
		res = append(res, p)
		for i := 1; i < len(urls); i++ {
			queue.Offer(urls[i])
		}
		index++
		if index == 10 {
			InsertPages(res)
			res = make([]*Page, 0)
		}
	}
}

// GetUrls 将给定的url里面的内容解析出来，生成一个数组切片，其中数组的第一个元素是url，第二个元素是url的内容
// 暂时也没有用到，因为不支持网页
func GetUrls(url string) [][2]string {
	res := make([][2]string, 0)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return res
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		fmt.Println("code error")
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	head := doc.Find("head")
	attr, exists := head.Find("meta").Attr("charset")
	if !exists {

	}
	fmt.Printf(attr)
	find := doc.Find("a")
	for i := 0; i < find.Length(); i++ {
		u, exists := find.Eq(i).Attr("href")
		if !exists {
			continue
		}
		if !strings.HasPrefix(u, "http") {
			continue
		}
		replace := strings.Replace(u, "\t", "", -1)
		s := strings.Replace(replace, "\n", "", -1)
		text := find.Eq(i).Text()
		s2 := strings.Replace(text, "\t", "", -1)
		s3 := strings.Replace(s2, "\n", "", -1)
		arr := [2]string{s, s3}
		res = append(res, arr)
	}
	return res
}

// Collector 由colly爬虫框架来进行网页的爬取，暂时不支持，因为网页暂时用不到
func Collector() {

	fName := "urls.csv"
	file, err := os.Create(fName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"url", "title"})

	c := colly.NewCollector(
		colly.MaxDepth(10),
		colly.Debugger(&debug.LogDebugger{}),
	)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
		text := e.Text
		fmt.Println(text)
		c.Visit(e.Request.AbsoluteURL(link))
		writer.Write([]string{link, text})
	})

	c.Visit("http://sina.com.cn")
}

// BuildIndex 把网页的信息建立一个倒排索引，但是这个方法暂时没有用到，因为暂时不支持网页
func BuildIndex(pages []*Page) map[string]int {
	m := make(map[*Page][]string)
	res := make(map[string]int)
	for _, page := range pages {
		cut := utils.Cut(page.Title)
		m[page] = cut
	}
	for p, words := range m {
		for _, word := range words {
			res[word] = p.Id
		}
	}
	return res
}
