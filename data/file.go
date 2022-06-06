package data

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

// Put2Csv 将数据库中的数据进行分词预处理，然后将结果放到csv文件中
// 再使用之前需要进行Seg的初始化
func Put2Csv() error {
	fileName := "../details_37w.csv"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"id", "detail"})
	var photos []Photo
	Db.Find(&photos)
	for _, p := range photos {
		content := p.Content
		cut := utils.Cut(content)
		index := p.Index
		m := make(map[string][]int)
		// 进行分词，记录每一个分词出现的次数、位置
		for i, s := range cut {
			ints, ok := m[s]
			if !ok {
				i2 := make([]int, 0)
				i2 = append(i2, 1)
				i2 = append(i2, i)
				m[s] = i2
			} else {
				ints = append(ints, i)
				ints[0]++
				m[s] = ints
			}
		}
		builder := strings.Builder{}
		for k, v := range m {
			builder.WriteString(k)
			fmt.Fprint(&builder, v)
			builder.WriteString("#")
		}
		s := builder.String()
		s2 := s[0 : len(s)-1]
		writer.Write([]string{strconv.Itoa(index), s2})
	}
	return nil
}

func Put2Db(fileName string, index int) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		text := scanner.Text()
		reg := regexp.MustCompile(`".*",`)
		find := reg.Find([]byte(text))
		if find == nil {
			continue
		}
		s := string(find)
		s2 := s[1 : len(s)-2]
		//s2为正确的
		//fmt.Println(s2)
		reg1 := regexp.MustCompile("[\\p{Han}]+")
		allString := reg1.FindAllString(text, -1)
		if allString == nil || len(allString) == 0 {
			continue
		}
		sb := strings.Builder{}
		for _, s3 := range allString {
			sb.WriteString(s3)
		}
		photo := &Photo{Index: index, Url: s2, Content: sb.String()}
		Db.Create(photo)
		index++
	}
	return index
}

func Put2Csvs(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"id", "detail"})
	var photos []Photo
	Db.Find(&photos)
	for _, p := range photos {
		content := p.Content
		cut := utils.Cut(content)
		index := p.Index
		m := make(map[string][]int)
		// 进行分词，记录每一个分词出现的次数、位置
		for i, s := range cut {
			ints, ok := m[s]
			if !ok {
				i2 := make([]int, 0)
				i2 = append(i2, 1)
				i2 = append(i2, i)
				m[s] = i2
			} else {
				ints = append(ints, i)
				ints[0]++
				m[s] = ints
			}
		}
		builder := strings.Builder{}
		for k, v := range m {
			builder.WriteString(k)
			fmt.Fprint(&builder, v)
			builder.WriteString("#")
		}
		s := builder.String()
		s2 := s[0 : len(s)-1]
		writer.Write([]string{strconv.Itoa(index), s2})
	}
	return nil
}
