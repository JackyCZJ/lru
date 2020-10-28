package lru

import (
	"sync"
)

type SafeCache struct {
	m sync.Mutex
	cache Cache

	nHit , nGet int
}

func NewSafeCache(cache Cache) *SafeCache{
	return &SafeCache{
		cache: cache,
	}
}

func (sc *SafeCache)Set(key string,value interface{}){
	sc.m.Lock()
	defer sc.m.Unlock()
	sc.cache.Set(key,value)
}

func (sc *SafeCache)Get(key string)interface{}{
	if sc == nil{
		return nil
	}
	sc.nGet ++
	sc.m.Lock()
	defer sc.m.Unlock()
	if value , ok := sc.cache.Get(key);ok{
		sc.nHit++
		return value
	}
	return nil
}

func (sc *SafeCache)Del(key Key){
	if sc == nil{
		return
	}
	sc.m.Lock()
	defer sc.m.Unlock()
	sc.cache.Del(key)
}

func (sc *SafeCache)Len()int{
	return sc.cache.Len()
}

type Stat struct {
	NGet , NHit int
}

func (sc *SafeCache)Stat()*Stat{
	return &Stat{
		NGet: sc.nGet,
		NHit: sc.nHit,
	}
}


type FastCache struct {
	shards []*SafeCache
	shardMask uint64
	hash hasher
}


//NewFastCache return a fast cache core with shard
func NewFastCache(maxEntries , shardsNum int)*FastCache{
	fc :=  &FastCache{
		shards:    make([]*SafeCache,shardsNum),
		shardMask: uint64(shardsNum - 1),
		hash: newDefaultHasher(),
	}
	for i := range fc.shards{
		fc.shards[i] = NewSafeCache(New(maxEntries))
	}
	return fc
}

func (c *FastCache)getShard(key Key)*SafeCache{
	hashKey := c.hash.Sum64(key.(string))
	return c.shards[hashKey&c.shardMask]
}

func (c *FastCache)Set(key Key,value string){
	c.getShard(key).Set(key.(string),value)
}

func (c *FastCache)Get(key Key)interface{}{
	return c.getShard(key).Get(key.(string))
}


func (c *FastCache)Del(key Key){
	 c.getShard(key).Del(key.(string))
}

func (c *FastCache)Len()int{
	length := 0
	for _ , shard :=range c.shards{
		length+=shard.Len()
	}
	return length
}