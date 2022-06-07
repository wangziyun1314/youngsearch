package data

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	init := GetUrls("https://sina.com.cn")
	for i, s := range init {
		if i == 0 {
			fmt.Printf("标题为%s", s)
		}
		if i == 1 {
			fmt.Printf("内容为%s", s)
		}
		fmt.Println(s)
	}
}

func TestGetUrls(t *testing.T) {
	urls := GetUrls("https://finance.sina.com.cn/tech/2022-05-25/doc-imizirau4679602.shtml")
	for _, url := range urls {
		fmt.Println(url[1])
	}
}

func TestSearchForPage(t *testing.T) {
	SearchForPage([2]string{"https://finance.sina.com.cn/tech/2022-05-25/doc-imizirau4679602.shtml", "新浪首页"})
}

func TestInsertPage(t *testing.T) {
	p := &Page{Id: 0, Title: "abc", Content: "sdg", Url: "http://sina.com.cn"}
	InsertPage(p)
}

//func TestGetUs(t *testing.T) {
//	us := GetUs("https://sina.com.cn")
//	fmt.Println(len(us))
//	for _, u := range us {
//		fmt.Printf("%s\n", u[0])
//		fmt.Printf("%s\n", u[1])
//	}
//}

func TestCollector(t *testing.T) {
	Collector()
}

func TestFindPhotoById(t *testing.T) {
	p := FindPhotoById(10)
	fmt.Printf("%#v", p)
}

// 测试范围查找
func TestPut2Csv(t *testing.T) {
	var pages []Page
	Db.Where("id <= ? AND id >= ?", 4, 2).Find(&pages)
	for _, page := range pages {
		fmt.Printf("%#v", page)
	}
}
