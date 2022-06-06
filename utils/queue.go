package utils

// Queue 使用双向链表实现一个队列
type Queue struct {
	head *node //队头
	tail *node // 队尾
	size int64 // 队列中的元素个数
}

// 队列中的节点
type node struct {
	data interface{}
	pre  *node
	next *node
}

// 新建一个节点
func newNode(v interface{}) *node {
	return &node{data: v}
}

// NewQueue create a queue
func NewQueue() *Queue {
	return &Queue{nil, nil, 0}
}

// GetSize 获取队列中元素的个数
func (q *Queue) GetSize() int64 {
	return q.size
}

// IsEmpty 判断队列中是否有元素
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Peek 得到队头的节点的值，但是不出队列
func (q *Queue) Peek() interface{} {
	if q.head == nil {
		return nil
	}
	return q.head.data
}

// Offer 将队尾入队
func (q *Queue) Offer(v interface{}) {
	node := newNode(v)
	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		node.pre = q.tail
		q.tail.next = node
		q.tail = node
	}
	q.size++
}

// Poll 队头出队
func (q *Queue) Poll() interface{} {
	if q.size == 0 {
		return nil
	}
	node := q.head
	if q.head.next == nil {
		q.head = nil
	} else {
		q.head = q.head.next
		q.head.pre.next = nil
		q.head.pre = nil
	}
	q.size--
	return node.data
}
