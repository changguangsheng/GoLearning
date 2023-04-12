package lru

import "container/list"

// Cache is a LRU cache. It is not safe for concurrent access.
// 包含字典和双向链表的结构体类型 Cache，方便实现后续的增删查改操作。
type Cache struct {
	maxBytes int64                    //允许使用的最大内存
	nbytes   int64                    //当前已使用的内存
	ll       *list.List               //双向链表
	cache    map[string]*list.Element //键是字符串，值是双向链表中对应节点的指针。
	//optional and executed when an entry is purged.
	OnEvicted func(key string, value Value) //某条记录被移除时的回调函数，可以为 nil。
}

// 双向链表节点的数据类型
// 在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射。
type entry struct {
	key   string
	value Value
}

// 实现了 Value 接口的任意类型，该接口只包含了一个方法 Len() int，用于返回值所占用的内存大小。
type Value interface {
	Len() int
}

// 实例化 Cache，实现 New() 函数：
func New(maxBytes int64, onEvicated func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicated,
	}
}

// 查找功能
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		//将链表中的节点 ele 移动到队尾（双向链表作为队列，队首队尾是相对的，在这里约定 front 为队尾）
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return nil, false
}

// 删除功能
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back() //c.ll.Back() 取到队首节点，从链表中删除。
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)                                //delete(c.cache, kv.key)，从字典中 c.cache 删除该节点的映射关系。
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len()) //更新当前所用的内存 c.nbytes。
		if c.OnEvicted != nil {                                //如果回调函数 OnEvicted 不为 nil，则调用回调函数。如果回调函数 OnEvicted 不为 nil，则调用回调函数。
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

//新增/修改

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		//如果键存在，则更新对应节点的值，并将该节点移到队尾。
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		//首先队尾添加新节点 &entry{key, value}, 并字典中添加 key 和节点的映射关系。
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	//更新 c.nbytes，如果超过了设定的最大值 c.maxBytes，则移除最少访问的节点
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
