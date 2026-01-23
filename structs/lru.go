package structs

import "container/list"

// LRU LRU缓存 / LRU cache
type LRU[K comparable, V any] struct {
	capacity int
	cache    map[K]*list.Element
	list     *list.List
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

// NewLRU 创建LRU缓存 / Create LRU cache
func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	if capacity <= 0 {
		capacity = 100
	}
	return &LRU[K, V]{
		capacity: capacity,
		cache:    make(map[K]*list.Element),
		list:     list.New(),
	}
}

// Get 获取值 / Get value
func (l *LRU[K, V]) Get(key K) (V, bool) {
	var zero V
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(*entry[K, V]).value, true
	}
	return zero, false
}

// Set 设置值 / Set value
func (l *LRU[K, V]) Set(key K, value V) {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		elem.Value.(*entry[K, V]).value = value
		return
	}
	if l.list.Len() >= l.capacity {
		back := l.list.Back()
		if back != nil {
			l.list.Remove(back)
			delete(l.cache, back.Value.(*entry[K, V]).key)
		}
	}
	elem := l.list.PushFront(&entry[K, V]{key: key, value: value})
	l.cache[key] = elem
}

// Remove 删除键值对 / Remove key-value pair
func (l *LRU[K, V]) Remove(key K) bool {
	if elem, ok := l.cache[key]; ok {
		l.list.Remove(elem)
		delete(l.cache, key)
		return true
	}
	return false
}

// Clear 清空缓存 / Clear cache
func (l *LRU[K, V]) Clear() {
	l.cache = make(map[K]*list.Element)
	l.list = list.New()
}

// Size 返回当前大小 / Return current size
func (l *LRU[K, V]) Size() int {
	return l.list.Len()
}

// Cap 返回容量 / Return capacity
func (l *LRU[K, V]) Cap() int {
	return l.capacity
}
