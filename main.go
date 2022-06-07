package main

import (
	"github.com/gin-gonic/gin"
	"youngsearch/handler"
	"youngsearch/index"
	"youngsearch/utils"
)

func main() {
	Init()
	r := gin.Default()
	r.GET("/search", func(c *gin.Context) {
		handler.SearchHandler(c)
	})
	r.Run()
}

// Init 对一些必要的数据结构的一些初始化操作
func Init() {
	index.IndexInit() // 初始化索引
	utils.InitSeg()   // 初始化jieba分词
	index.InitTris()  // 初始化字典树
}
