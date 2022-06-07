package data

import (
	"fmt"
	"testing"
	"youngsearch/utils"
)

// 将10亿数据入库 运行到一半失败了
func TestPut2Db(t *testing.T) {
	fileDir := "G:\\BaiduNetdiskDownload\\wukong_release\\wukong_release"
	index := 0
	for i := 0; i < 256; i++ {
		fileSuffix := fmt.Sprintf("\\wukong_100m_%d.csv", i)
		fmt.Println(fileDir + fileSuffix)
		index = Put2Db(fileDir+fileSuffix, index)
	}

}

// 将100m的数据入库
func TestPut2Db1(t *testing.T) {
	fileDir := "G:\\BaiduNetdiskDownload\\wukong_release\\wukong_release"
	index := 0
	for i := 0; i < 1; i++ {
		fileSuffix := fmt.Sprintf("\\wukong_100m_%d.csv", i)
		fmt.Println(fileDir + fileSuffix)
		index = Put2Db(fileDir+fileSuffix, index)
	}

}

// 将100m的数据进行分词并且放入到csv文件中
func TestPut2Csv1(t *testing.T) {
	utils.InitSeg()
	Put2Csv()
}

func TestPut2TermCsv(t *testing.T) {
	err := Put2TermCsv("../tries.csv")
	fmt.Println(err)
}
