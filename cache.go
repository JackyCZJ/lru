package lru


type Cache interface {
	Add(key Key,value interface{})
	Get(key Key)(interface{},bool)
}