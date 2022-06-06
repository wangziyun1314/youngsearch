package search

import (
	"fmt"
	"testing"
	"youngsearch/index"
)

func TestSearch(t *testing.T) {
	index.IndexInit()
	search := Search("苹果")
	fmt.Printf("%#v\n", search)
}
