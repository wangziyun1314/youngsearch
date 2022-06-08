package utils

import (
	"fmt"
	"strings"
)

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

var nodeCnt int

type AvlNode struct {
	val   int
	left  *AvlNode
	right *AvlNode
	h     int
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func (node *AvlNode) height() int {
	if node == nil {
		return 0
	}
	return node.h
}

// 对节点y进行向右旋转操作，返回旋转后新的根节点x
//        y                              x
//       / \                           /   \
//      x   T4     向右旋转 (y)        z     y
//     / \       - - - - - - - ->    / \   / \
//    z   T3                       T1  T2 T3 T4
//   / \
// T1   T2

func adjustLeft(y *AvlNode) *AvlNode {
	if y == nil {
		panic("invalid node")
	}
	var x = y.left
	y.left = x.right
	x.right = y
	y.h = max(y.left.height(), y.right.height()) + 1
	x.h = max(x.left.height(), x.right.height()) + 1
	return x
}

// 对节点y进行向左旋转操作，返回旋转后的新的根节点x
//            y                                            x
//          /   \                                        /    \
//         T4    x             向左旋转                  y        z
//              /  \          ------------->          /   \    /   \
//             T1     z                              T4  T1   T2    T3
//                 /  \
//               T2   T3

func adjustRight(y *AvlNode) *AvlNode {
	if y == nil {
		panic("invalid node")
	}
	var x = y.right
	y.right = x.left
	x.left = y
	y.h = max(y.left.height(), y.right.height()) + 1
	x.h = max(x.left.height(), x.right.height()) + 1
	return x
}

// left 重并且left.right 重
//
//        y                                               Y
//       / \                                            /   \
//      x   T4       先调整x,把x向左旋转                 T3     T4          和上面的一样
//     / \       - - - - - - -------- ->             /   \             ---------------->
//    z   T3                                        X    T2
//       / \                                      /   \
//     T1   T2                                   Z     T1
//

func adjustLeftRight(node *AvlNode) *AvlNode {
	left := adjustRight(node.left)
	node.left = left
	return adjustLeft(node)
}

//right 重  并且 right.left重
//            y                                             y
//          /   \                                         /   \
//         T4    x               先调整x                 T4     Z              和上面一样
//              /  \          ------------->                 /   \      -------------------->
//             z    T1                                      T2    X
//           /  \                                                / \
//         T2   T3                                             T3   T1
//

func adjustRightLeft(node *AvlNode) *AvlNode {
	right := adjustLeft(node.right)
	node.right = right
	return adjustRight(node)
}
func Find(node *AvlNode, val int) *AvlNode {
	if node == nil {
		return nil
	}
	if node.val == val {
		return node
	} else if node.val < val {
		return Find(node.right, val)
	} else {
		return Find(node.left, val)
	}
}

func InsertNode(node *AvlNode, val int) *AvlNode {
	if node == nil {
		nodeCnt++
		return &AvlNode{
			val: val,
			h:   1,
		}
	}
	switch {
	case node.val < val:
		node.right = InsertNode(node.right, val)
		if node.right.height()-node.left.height() == 2 {
			//right 重
			if val > node.right.val {
				//right 重 并且 right.right 重
				node = adjustRight(node)
			} else {
				// right 重 right.left 重
				node = adjustRightLeft(node)
			}
		}
	case val < node.val:
		node.left = InsertNode(node.left, val)
		if node.left.height()-node.right.height() == 2 {
			// left 重
			if node.left.val < val {
				// left 重 并且 left.right重
				node = adjustLeftRight(node)
			} else {
				// left 重 并且 left.left 重
				node = adjustLeft(node)
			}
		}
	default:

	}
	node.h = max(node.left.height(), node.right.height()) + 1
	return node
}

func findMinNode(node *AvlNode) *AvlNode {
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

func findMaxNode(node *AvlNode) *AvlNode {
	if node == nil {
		return nil
	}
	for node.right != nil {
		node = node.right
	}
	return node
}

// 找到小于等于val的最大值
func findFloorNode(node *AvlNode, val int) *AvlNode {
Loop:
	for node != nil {
		if node.val == val {
			break Loop
		} else if node.val < val {
			node = node.right
		} else {
			left := findFloorNode(node.left, val)
			if left != nil && left.val < node.val {
				node = left
			}
			break Loop
		}
	}
	return node
}

// 找到大于等于val的最小值
func findCeilNode(node *AvlNode, val int) *AvlNode {
	if node == nil {
		return nil
	}
find:
	for node != nil {
		if node.val == val {
			break find
		} else if node.val > val {
			node = node.left
		} else {
			var s = findCeilNode(node.right, val)
			if s != nil && s.val > node.val {
				node = s
			}
			break find
		}
	}
	return node
}

func adjustNode(node *AvlNode) *AvlNode {
	switch node.left.height() - node.right.height() {
	case 2:
		// left 重
		if node.left.height() > node.right.height() {
			// left 重 left.left 重
			node = adjustLeft(node)
		} else {
			// left 重 left.right 重
			node = adjustLeftRight(node)
		}
	case -2:
		// right 重
		if node.left.height() > node.right.height() {
			// right 重 并且 right.left 重
			node = adjustRightLeft(node)
		} else {
			// right 重 并且 right.right 重
			node = adjustRight(node)
		}
	}
	return node
}

// 删除节点
func deleteNode(node *AvlNode, val int) *AvlNode {
	if node == nil {
		return nil
	}
	if node.val < val {
		node.right = deleteNode(node.right, val)
	} else if node.val > val {
		node.left = deleteNode(node.left, val)
	} else {
		nodeCnt--
		if node.left == nil {
			if node.right == nil {
				node = nil
			} else {
				node = node.right
			}
		} else {
			if node.right == nil {
				node = node.left
			} else {
				var nextNode = findMinNode(node.right)
				node.val = nextNode.val
				node.right = deleteNode(node.right, nextNode.val)
			}
		}
	}
	if node != nil {
		node.h = max(node.left.height(), node.right.height()) + 1
		node = adjustNode(node)
	}
	return node
}

// 从小到大对输出avl的值
func ascPrint(node *AvlNode) {
	if node == nil {
		return
	}
	ascPrint(node.left)
	fmt.Printf("%v\n", node.val)
	ascPrint(node.right)
}

func (node *AvlNode) String() string {
	return fmt.Sprintf("%v", node.val)
}

//  把avl画出来
func (node *AvlNode) draw(prefix string, isTail bool, str *strings.Builder) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "|   "
		} else {
			newPrefix += "    "
		}
		node.right.draw(newPrefix, false, str)
	}
	str.WriteString(prefix)
	if isTail {
		str.WriteString("└── ")
	} else {
		str.WriteString("┌── ")
	}
	str.WriteString(node.String() + "\n")
	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "|    "
		}
		node.left.draw(newPrefix, true, str)
	}
}
