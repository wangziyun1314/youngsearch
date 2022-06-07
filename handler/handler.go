package handler

import (
	"github.com/gin-gonic/gin"
	"sync"
	"youngsearch/index"
	"youngsearch/search"
	"youngsearch/utils"
)

// SearchHandler 对输入的字符串进行搜索
func SearchHandler(c *gin.Context) {
	param := c.DefaultQuery("content", "")
	if param == "" {
		c.JSON(-1, gin.H{
			"message": "content is empty",
		})
		return
	}
	filter := c.DefaultQuery("filter", "")
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
	res := make(map[[2]string]int, 0)
	if len(notLike) != 0 {
		for _, url := range urls {
			_, ok := notLike[url[0]]
			if !ok {
				res[url] = 1
			}
		}
	} else {
		for _, url := range urls {
			res[url] = 1
		}
	}
	msg := make([][2]string, 0)
	for k, _ := range res {
		msg = append(msg, k)
	}
	msg = search.TfIdf(utils.Seg.Cut(param, true), msg)
	if !utils.BloomFilter.Check(param) {
		index.UpdateTries(param)
	}
	recommend := index.WordTries.Recommend(param)
	index.WordTries.Insert(param)
	pages := utils.GetPages(msg)
	c.JSON(200, gin.H{
		"pages":     pages,
		"recommend": recommend,
	})
}
