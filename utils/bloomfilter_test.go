package utils

import (
	"fmt"
	"testing"
)

func TestBloomFilter_Insert(t *testing.T) {
	filter := NewBloomFilter(nil)
	strings := []string{"中国", "上海", "北京", "深圳", "广州", "南京", "成都", "河南", "浙江", "杭州"}
	strs := []string{"中国", "上海", "北京", "深圳", "广州", "南京", "成都", "河南", "浙江", "杭州", "郑州", "南阳", "邓州", "中国"}
	for _, s := range strings {
		filter.Insert(s)
	}
	for _, str := range strs {
		fmt.Println(filter.Check(str))
	}

}
