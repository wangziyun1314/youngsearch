package index

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"youngsearch/utils"
)

// 暂时没有用到的索引结构
var urlMap map[int][]string
var indexMap map[string]map[int]int

// BuildIndex 构建上面的索引，没有用到了
func BuildIndex() error {
	urlMap = make(map[int][]string, 0)
	indexMap = make(map[string]map[int]int, 0)
	file, err := os.Open("../wukong50k_release.csv")
	if err != nil {
		fmt.Println(err)
		return err
	}
	scanner := bufio.NewScanner(file)
	index := 0
	scanner.Scan()
	for scanner.Scan() {
		text := scanner.Text()
		reg := regexp.MustCompile(`".*",`)
		find := reg.Find([]byte(text))
		if find == nil {
			continue
		}
		s := string(find)
		fmt.Println(s)
		s2 := s[1 : len(s)-2]
		//s2为正确的
		//fmt.Println(s2)
		reg1 := regexp.MustCompile("[\\p{Han}]+")
		allString := reg1.FindAllString(text, -1)
		if allString == nil || len(allString) == 0 {
			continue
		}
		list := make([]string, 0)
		list = append(list, s2)
		for _, s := range allString {
			cut := utils.Cut(s)
			for _, s2 := range cut {
				list = append(list, s2)
				//indexmap
				allString = append(allString, s2)
				m, ok := indexMap[s2]
				if !ok {
					i := make(map[int]int, 0)
					i[index] = 1
					indexMap[s2] = i
				} else {
					_, ok := m[index]
					if !ok {
						m[index] = 1
					}
				}
			}
		}
		urlMap[index] = list
		index++
	}
	return nil
}

// IndexMap 内存中的索引结构
var IndexMap map[string][]int

// IndexInit 初始化内存中的索引
func IndexInit() {
	IndexMap = Build()
}

// Build 由分词结果构建索引
func Build() map[string][]int {
	res := make(map[string][]int, 0)
	fileName := utils.IndexFileName
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return res
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ",")
		index, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		termsList := strings.Split(split[1], "#")
		for _, list := range termsList {
			i2 := strings.Index(list, "[")
			term := list[0:i2]
			ints, ok := res[term]
			if !ok {
				i := make([]int, 0)
				i = append(i, index)
				res[term] = i
			} else {
				ints = append(ints, index)
				res[term] = ints
			}
		}
	}
	return res
}

// WordTries 内存中的字典树，用来推荐关联搜索
var WordTries *utils.Trie

// InitTris 由常用的搜索词库来进行初始化字典树
func InitTris() {
	WordTries = utils.NewTries()
	file, err := os.Open("./Tries.csv")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		WordTries.Insert(word)
	}
}

// Recommend 根据输入的字符串进行关联词的推荐
func Recommend(word string) []string {
	return WordTries.Recommend(word)
}

// UpdateTries 根据输入的词扩充常用词库，并且更新字典树
func UpdateTries(word string) {
	WordTries.Insert(word)
	file, err := os.OpenFile("./Tries.csv", os.O_APPEND, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{word})
}

// Insert2TriesCsv 把搜索的词加入到字典库中
func Insert2TriesCsv(word string) {
	file, err := os.OpenFile("./Tries.csv", os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{word})
}
