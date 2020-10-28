package lru

import "hash"

func newDefaultHasher()hasher{
	return hasher{}
}

type hasher struct {
	hash.Hash
}

const (
	offset64 = 13695981039345556037
	prime64 = 1099511628211
)

func (hasher)Sum64(key string)uint64{
	var h uint64 = offset64
	for i := 0 ; i < (len(key));i++{
		h ^=uint64(key[i])
		h *= prime64
	}
	return h
}
