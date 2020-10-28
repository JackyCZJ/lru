package lru


type Cache interface {
	Set(key Key,value interface{})
	Get(key Key)(interface{},bool)
	Len()int
	Del(key Key)
}