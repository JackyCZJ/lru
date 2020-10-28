package lru

import (
	"container/list"
)

type Lru struct {
	MaxEntries int 	//Max Cap , 0 means unlimited.

	//Callback when delete something.
	OnEvicted func(key Key ,value interface{})

	ll 	*list.List

	cache map[interface{}]*list.Element
}

// Key should be any type can be compare/
type Key interface {
}

type entry struct {
	key   Key
	value interface{}
}

func New(maxEntries int) *Lru{
	return &Lru{
		MaxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

func (c *Lru)Set(key Key,value interface{}) {
	if c == nil{
		c.ll = list.New()
		c.cache = make(map[interface{}]*list.Element)
	}
	if v , ok := c.cache[key];ok{
		c.ll.MoveToFront(v)
		v.Value.(*entry).value = value
		return
	}
	ele := c.ll.PushFront(&entry{key,value})
	c.cache[key] = ele
	if c.MaxEntries != 0 &&c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}

}

func (c *Lru)Get(key Key)(interface{},bool){
	if c == nil{
		return nil , false
	}
	if v , ok := c.cache[key]; ok{
		c.ll.MoveToFront(v)
		return v.Value.(*entry).value,true
	}
	return nil,false
}

func (c *Lru) RemoveOldest(){
	if c == nil{
		return
	}
	ele := c.ll.Back()
	if ele != nil{
		c.RemoveElement(ele)
	}
}

func (c *Lru)RemoveElement(e *list.Element){
	c.ll.Remove(e)
	if v  , ok  := e.Value.(*entry);ok {
		delete(c.cache,v.key)
		if c.OnEvicted != nil{
			c.OnEvicted(v.key,v.value)
		}
	}
}

func(c *Lru)Del(key Key){
	if c == nil{
		return
	}
	if v , ok := c.cache[key];ok{
		c.RemoveElement(v)
	}
}

func (c *Lru)Len()int{
	if c == nil{
		return 0
	}
	return c.ll.Len()
}