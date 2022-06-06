package handler

import (
	"github.com/gin-gonic/gin"
	"sync"
	"youngsearch/search"
	"youngsearch/utils"
)

// SearchHandler 对输入的字符串进行搜索
func SearchHandler(c *gin.Context) {
	param := c.PostForm("search")
	filter := c.PostForm("filter")
	waitGroup := sync.WaitGroup{}
	notLike := make(map[string]int, 0)
	// 如果filter字段为空
	if "" != filter {
		waitGroup.Add(1)
		go func() {
			strs := search.Search(filter)
			for _, str := range strs {
				notLike[str[0]] = 1
			}
			waitGroup.Done()
		}()
	}
	urls := search.Search(param)
	waitGroup.Wait()
	if len(urls) == 0 {
		c.JSON(200, gin.H{
			"message": "not found in database",
		})
		return
	}
	res := make([][2]string, 0)
	if len(notLike) != 0 {
		for _, url := range urls {
			_, ok := notLike[url[1]]
			if !ok {
				res = append(res, url)
			}
		}
		//c.JSON(200, gin.H{
		//	"message": res,
		//})
	} else {
		//c.JSON(200, gin.H{
		//	"message": urls,
		//})
		res = urls
	}
	res = search.TfIdf(utils.Seg.Cut(param, true), res)
	c.JSON(200, gin.H{
		"message": res,
	})
}
