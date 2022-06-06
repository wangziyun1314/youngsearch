package utils

import (
	"fmt"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	tries := NewTries()
	tries.Insert("我爱中国")
	tries.Insert("我爱美国")
	word := tries.SearchWord("我爱中国")
	lastNode := tries.FindLastNode("我爱")
	fmt.Println(word)
	fmt.Printf("%#+v", lastNode)
	recommend := tries.Recommend("我爱")
	fmt.Println(recommend)
}
