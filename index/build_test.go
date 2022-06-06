package index

import (
	"fmt"
	"testing"
)

func TestBuildIndex(t *testing.T) {
	err := BuildIndex()
	if err != nil {
		fmt.Println(err)
	}
}

//func TestPut2Db(t *testing.T) {
//	Put2Db()
//}

func TestBuild(t *testing.T) {
	build := Build()
	fmt.Printf("%#v\n", build)
}

func TestInitTris(t *testing.T) {
	Insert2TriesCsv("我爱中国")
	Insert2TriesCsv("我爱美国")
	Insert2TriesCsv("我爱学习")
	InitTris()
	UpdateTries("我爱旅游")
	recommend := Recommend("我爱")
	fmt.Printf("%#v", recommend)
}
