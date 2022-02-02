import (
	"container/list"
)

type LRUCache struct {
	cacheList *list.List
	cacheMap  map[int]*list.Element
	capacity  int
}

type KV struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cacheList: list.New(),
		cacheMap:  make(map[int]*list.Element),
		capacity:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.cacheMap[key]
	if !ok {
		return -1
	}

	kv, _ := this.cacheList.Remove(e).(KV)
	this.cacheMap[key] = this.cacheList.PushFront(kv)

	return kv.Value
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.cacheMap[key]
	if ok {
		this.cacheList.Remove(e)
		this.cacheMap[key] = this.cacheList.PushFront(KV{
			Key:   key,
			Value: value,
		})
		return
	}

	if len(this.cacheMap) < this.capacity {
		this.cacheMap[key] = this.cacheList.PushFront(KV{
			Key:   key,
			Value: value,
		})
		return
	}

	kv, _ := this.cacheList.Remove(this.cacheList.Back()).(KV)
	delete(this.cacheMap, kv.Key)
	this.cacheMap[key] = this.cacheList.PushFront(KV{
		Key:   key,
		Value: value,
	})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */