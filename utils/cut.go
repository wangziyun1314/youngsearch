package utils

import "github.com/yanyiwu/gojieba"

// Cut 根据jieba分词来对字符串进行切分
func Cut(s string) []string {
	seg := Seg
	stirs := seg.CutForSearch(s, true)
	return stirs
}

// Seg 在内存中定义一个jieba分词
var Seg *gojieba.Jieba

// 初始化Seg
func InitSeg() {
	Seg = gojieba.NewJieba()
}
