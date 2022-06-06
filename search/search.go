package search

import (
	"fmt"
	"youngsearch/data"
	"youngsearch/index"
	"youngsearch/utils"
)

// Search 由给定的字符串来进行搜索
func Search(s string) [][2]string {
	res := make([][2]string, 0)
	seg := utils.Seg
	strings := seg.CutForSearch(s, true)
	if len(strings) == 0 {
		fmt.Println("分词后的切片为空")
		return res
	}
	for _, str := range strings {
		ints, ok := index.IndexMap[str]
		if !ok {
			continue
		} else {
			for _, i := range ints {
				photo := data.FindPhotoById(i)
				res = append(res, [2]string{photo.Url, photo.Content})
			}
		}
	}
	return res
}
