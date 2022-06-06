package search

import (
	"math"
	"sort"
	"youngsearch/index"
	"youngsearch/utils"
)

type scoreNode struct {
	index int
	score float64
}

// TfIdf 根据tfidf算法对结果进行一个排序
func TfIdf(segs []string, res [][2]string) [][2]string {
	seg := utils.Seg
	m := make(map[int]map[string]int, 0) // 对结果中的每一个content进行分词并且进行统计
	// 统计的过程
	for i, re := range res {
		cut := seg.Cut(re[1], true)
		m2 := make(map[string]int, 0)
		for _, s := range cut {
			i2, ok := m2[s]
			if !ok {
				m2[s] = 1
			} else {
				m2[s] = i2 + 1
			}
		}
		m[i] = m2
	}
	// 对最后的网页进行一个打分，每个单词进行一次计算，并且求出最后的总和，由总和进行一个排序
	scores := make([]scoreNode, len(res))
	for _, term := range segs {
		for i, m2 := range m {
			i2 := m2[term]
			i3 := len(seg.Cut(res[i][1], true))
			tf := float64(i2*1.0) / float64(i3*1.0)
			var ratio float64
			if len(index.IndexMap[term]) != 0 {
				ratio = float64((utils.TotalDocument * 1.0) / (len(index.IndexMap[term]) * 1.0))
			}
			idf := math.Log10(ratio)
			s := scores[i]
			temp := s.score + tf*idf
			scores[i] = scoreNode{index: i, score: temp}
		}
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})
	sortRes := make([][2]string, len(res))
	for i, score := range scores {
		sortRes[i] = res[score.index]
	}
	//fmt.Println("hello")
	return sortRes
}
