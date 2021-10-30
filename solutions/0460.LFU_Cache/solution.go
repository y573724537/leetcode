package main

import "container/list"

type LFUCache struct {
	capacity int
	keyMap   map[int]*KVF
	freqMap  map[int]*LinkedHashSet
	minFreq  int
}

type LinkedHashSet struct {
	m map[int]*list.Element
	l *list.List
}

func NewLinkedHashSet() *LinkedHashSet {
	return &LinkedHashSet{
		m: make(map[int]*list.Element),
		l: list.New(),
	}
}

func (lhm *LinkedHashSet) Get(key int) bool {
	_, ok := lhm.m[key]
	return ok
}

func (lhm *LinkedHashSet) Put(key int) {
	_, ok := lhm.m[key]
	if ok {
		return
	}
	lhm.m[key] = lhm.l.PushBack(key)
}

func (lhm *LinkedHashSet) DelOldest(key int) {
	key := lhm.l.Remove(lhm.l.Front()).(int)
	delete(lhm.m, key)
}

func (lhm *LinkedHashSet) Del(key int) {
	e, ok := lhm.m[key]
	if !ok {
		return
	}
	lhm.l.Remove(e)
	delete(lhm.m, key)
}

type KVF struct {
	key   int
	value int
	freq  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		keyMap:   make(map[int]*KVF, capacity),
		freqMap:  make(map[int]*LinkedHashSet),
		minFreq:  0,
	}

}

func (this *LFUCache) Get(key int) int {
	kvf, ok := this.keyMap[key]
	if !ok {
		return -1
	}

}

func (this *LFUCache) Put(key int, value int) {
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
