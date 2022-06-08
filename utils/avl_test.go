package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestAvlNode_String(t *testing.T) {
	var root *AvlNode
	root = InsertNode(root, 1)
	root = InsertNode(root, 1)
	root = InsertNode(root, 1)
	root = InsertNode(root, 1)
	for i := 1; i < 50000000; i++ {
		root = InsertNode(root, i)
	}
	start := time.Now()
	find := Find(root, 255)
	seconds := float64(time.Since(start).Nanoseconds())
	fmt.Printf("节点为%v\n", find)
	fmt.Printf("时间为%f\n", seconds)
}
