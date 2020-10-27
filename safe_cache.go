package lru

import "sync"

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
	sc.cache.Add(key,value)
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

type Stat struct {
	NGet , NHit int
}

func (sc *SafeCache)Stat()*Stat{
	return &Stat{
		NGet: sc.nGet,
		NHit: sc.nHit,
	}
}


