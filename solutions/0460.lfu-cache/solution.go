package main

import (
	"container/list"
	"fmt"
)

type LFUCache struct {
	capacity int
	len      int
	minFreq  int
	keyMap   map[int]*KVF
	freqMap  map[int]*LinkedHashMap
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		len:      0,
		minFreq:  0,
		keyMap:   make(map[int]*KVF),
		freqMap:  make(map[int]*LinkedHashMap),
	}
}

func (this *LFUCache) Get(key int) int {
	kvf, ok := this.keyMap[key]
	if !ok {
		return -1
	}

	lhm0 := this.freqMap[kvf.F]
	lhm0.Del(key)
	if lhm0.Len() == 0 {
		delete(this.freqMap, kvf.F)
	}

	if kvf.F == this.minFreq && lhm0.Len() == 0 {
		this.minFreq = kvf.F + 1
	}

	kvf.F++
	lhm1, ok := this.freqMap[kvf.F]
	if !ok {
		lhm1 = NewLinkedHashMap()
		this.freqMap[kvf.F] = lhm1
	}
	lhm1.Put(kvf)

	return kvf.V
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	kvf, ok := this.keyMap[key]
	if ok {
		kvf.V = value

		lhm0 := this.freqMap[kvf.F]
		lhm0.Del(key)
		if lhm0.Len() == 0 {
			delete(this.freqMap, kvf.F)
		}
		if kvf.F == this.minFreq && lhm0.Len() == 0 {
			this.minFreq = kvf.F + 1
		}

		kvf.F++
		lhm1, ok := this.freqMap[kvf.F]
		if !ok {
			lhm1 = NewLinkedHashMap()
			this.freqMap[kvf.F] = lhm1
		}
		lhm1.Put(kvf)

		return
	}

	if this.len == this.capacity {
		minLhm := this.freqMap[this.minFreq]
		last := minLhm.DelLast()
		delete(this.keyMap, last.K)
		if minLhm.Len() == 0 {
			delete(this.freqMap, this.minFreq)
			this.minFreq = 0
		}
	} else {
		this.len++
	}

	lhm1, ok := this.freqMap[1]
	if !ok {
		lhm1 = NewLinkedHashMap()
	}
	this.minFreq = 1
	this.freqMap[1] = lhm1

	kvf = &KVF{
		K: key,
		V: value,
		F: 1,
	}
	lhm1.Put(kvf)
	this.keyMap[key] = kvf
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Printf("%v\n", cache.Get(1))
	cache.Put(3, 3)
	fmt.Printf("%v\n", cache.Get(2))
	fmt.Printf("%v\n", cache.Get(3))
	cache.Put(4, 4)
	fmt.Printf("%v\n", cache.Get(1))
	fmt.Printf("%v\n", cache.Get(3))
	fmt.Printf("%v\n", cache.Get(4))
}

type KVF struct {
	K int
	V int
	F int
}

type LinkedHashMap struct {
	l *list.List
	m map[int]*list.Element
}

func NewLinkedHashMap() *LinkedHashMap {
	return &LinkedHashMap{
		l: list.New(),
		m: make(map[int]*list.Element),
	}
}

func (lhm *LinkedHashMap) Put(kvf *KVF) {
	ele, ok := lhm.m[kvf.K]
	if ok {
		lhm.l.Remove(ele)
	}
	lhm.m[kvf.K] = lhm.l.PushFront(kvf)
}

func (lhm *LinkedHashMap) Get(key int) (*KVF, bool) {
	ele, ok := lhm.m[key]
	if !ok {
		return nil, false
	}

	kvf, _ := lhm.l.Remove(ele).(*KVF)
	lhm.m[key] = lhm.l.PushFront(kvf)

	return kvf, true
}

func (lhm *LinkedHashMap) Del(key int) {
	ele, ok := lhm.m[key]
	if !ok {
		return
	}

	lhm.l.Remove(ele)
	delete(lhm.m, key)
}

func (lhm *LinkedHashMap) Len() int {
	return lhm.l.Len()
}

func (lhm *LinkedHashMap) DelLast() *KVF {
	kvf, _ := lhm.l.Remove(lhm.l.Back()).(*KVF)
	delete(lhm.m, kvf.K)

	return kvf
}
