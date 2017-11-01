package lrucache

import "container/list"
import "errors"

type CacheNode struct {
	Key, Val interface{}
}

type LruCache struct {
	Cap      int
	Used     int
	lrulist    *list.List
	CacheMap map[interface{}]*list.Element
}

// 初始化一个 lrucache
func NewLruCache() *LruCache {
	cache := &LruCache{}
	cache.Cap = 10
	cache.lrulist = list.New()
	cache.CacheMap = map[interface{}]*list.Element{}
	return cache
}

// 获取缓存的大小
func (lru *LruCache) GetCap() int {
	return lru.Cap
}

// 已使用的大小
func (lru *LruCache) GetUsed() int {
	return lru.lrulist.Len()
}

// set方法
func (lru *LruCache) Set(k, v interface{}) (bool, error) {
	if lru.lrulist == nil {
		return false, errors.New("LRU 未初始化")
	}

	if element, ok := lru.CacheMap[k]; ok {
		lru.lrulist.MoveToFront(element)
		element.Value.(*CacheNode).Val = v
	}
	lru.CacheMap[k] = lru.lrulist.PushFront(&CacheNode{k, v})
	if lru.lrulist.Len() > lru.Cap {
		last := lru.lrulist.Back()
		lastNode := last.Value.(*CacheNode)
		delete(lru.CacheMap, lastNode.Key)
		lru.lrulist.Remove(last)
	}
	return true, nil
}

// get
func (lru *LruCache) Get(k interface{}) (interface{}, error) {
	if lru.lrulist == nil {
		return nil, errors.New("LRU 未初始化")
	}
	if lru.lrulist.Len() == 0 {
		return nil, nil
	}
	if element, ok := lru.CacheMap[k]; ok {
		value := element.Value.(*CacheNode).Val
		lru.lrulist.MoveToFront(element)
		return value, nil
	}
	return nil, nil
}

// 清除某个元素
func (lru *LruCache) Remove(k interface{}) (bool, error) {
	if lru.lrulist == nil {
		return false, errors.New("LRU 未初始化")
	}
	if lru.lrulist.Len() == 0 {
		return false, nil
	}
	if element, ok := lru.CacheMap[k]; ok {
		key := element.Value.(*CacheNode).Key
		delete(lru.CacheMap, key)
		lru.lrulist.Remove(element)
		return true, nil
	}
	return false, nil
}

// 遍历
func (lru *LruCache) Traversal() ([]interface{},error){
	if lru.lrulist == nil {
		return []interface{}{}, errors.New("LRU 未初始化")
	}
	if lru.lrulist.Len() == 0 {
		return []interface{}{}, nil 
	}

	frontEle := lru.lrulist.Front()
	vals := []interface{}{frontEle.Value.(*CacheNode).Val}
	for nxt := frontEle.Next(); nxt != nil ;nxt = nxt.Next(){
		vals = append(vals,nxt.Value.(*CacheNode).Val)	
	}
	return vals,nil

}
