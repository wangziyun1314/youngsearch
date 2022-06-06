package utils

import "fmt"

/*type bloomFilteror interface {
	Insert(v interface{})
	Check(v interface{})
	Clear()
}*/
// 布隆过滤器
type bloomFilter struct {
	bits []uint64
	hash Hash
}

type Hash func(v interface{}) (h uint32)

// hash 函数
func hash(v interface{}) (h uint32) {
	h = uint32(0)
	s := fmt.Sprintf("131-%v-%v", v, v)
	bs := []byte(s)
	for i := range bs {
		h += uint32(bs[i]) * 131
	}
	return h
}

// NewBloomFilter 新建一个布隆过滤器
func NewBloomFilter(h Hash) (bf *bloomFilter) {
	if h == nil {
		h = hash
	}
	return &bloomFilter{
		bits: make([]uint64, 0, 0),
		hash: h,
	}
}

// Insert 插入到布隆过滤器
func (bf *bloomFilter) Insert(v interface{}) {
	if bf == nil {
		return
	}
	h := bf.hash(v)
	if h/64+1 > uint32(len(bf.bits)) {
		var tmp []uint64
		if h/64+1 < uint32(len(bf.bits)+1024) {
			tmp = make([]uint64, len(bf.bits)+1024)
		} else {
			tmp = make([]uint64, h/64+1)
		}
		copy(tmp, bf.bits)
		bf.bits = tmp
	}
	bf.bits[h/64] ^= 1 << (h % 64)
}

// Check 判断元素是否在布隆过滤器中
func (bf *bloomFilter) Check(v interface{}) (b bool) {
	if bf == nil {
		return false
	}
	h := bf.hash(v)
	if h/64+1 > uint32(len(bf.bits)) {
		return false
	}
	if bf.bits[h/64]&(1<<(h%64)) > 0 {
		return true
	}
	return false
}

// Clear 将布隆过滤器中的元素清空
func (bf *bloomFilter) Clear() {
	if bf == nil {
		return
	}
	bf.bits = make([]uint64, 0, 0)
}
