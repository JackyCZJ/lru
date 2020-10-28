package lru

import (
	"strconv"
	"testing"
)

func Test_lru(t *testing.T)  {
	c := New(5)
	c.OnEvicted = func(key Key, value interface{}) {
		t.Log("remove key:",key, ", value:",value)
	}
	const (
		s1 = iota
		s2
		s3
		s4
		s5
		s6
	)
	c.Set("s1",s1)
	c.Set("s2",s2)
	c.Set("s3",s3)
	c.Set("s4",s4)
	c.Set("s5",s5)
	c.Set("s6",s6)
	_ ,ok := c.Get("s1")
	if ok{
		t.Fatal("lru false")
	}
	c.Set("s1",s1)
	_ ,ok = c.Get("s2")
	if ok{
		t.Fatal("lru false")
	}
	c.Get("s6")

	if c.ll.Front().Value.(*entry).key != "s6"{
		t.Fatal("lru false")
	}

}

func TestNewSafeCache(t *testing.T) {
	c := New(5)
	sc := NewSafeCache(c)
	sc.Set("s1",1)
	if 1 != sc.Get("s1").(int) {
		t.Fatal("error")
	}
	sc.Get("s2")
	if sc.Stat().NGet != 2{
		t.Fatal("error")
	}
	if sc.Stat().NHit != 1{
		t.Fatal("error")
	}
}

func TestNewFastCache(t *testing.T) {
	sc := NewFastCache(5,10)
	sc.Set("s1",1)
	if 1 != sc.Get("s1").(int) {
		t.Fatal("error")
	}
	if sc.Get("s2") != nil{
		t.Fatal("error")
	}

}

func BenchmarkFastCache(b *testing.B) {
	sc := NewFastCache(100,10)
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		sc.Set(strconv.Itoa(i),i)
	}
	for i:=0 ; i <b.N;i++{
		sc.Get(strconv.Itoa(i))
	}
	for i := 0; i < b.N;i++{
		sc.Del(strconv.Itoa(i))
	}
}

func BenchmarkSafeCache(b *testing.B) {
	sc := NewSafeCache(New(100))
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		sc.Set(strconv.Itoa(i),i)
	}
	for i:=0 ; i <b.N;i++{
		sc.Get(strconv.Itoa(i))
	}
	for i := 0; i < b.N;i++{
		sc.Del(strconv.Itoa(i))
	}
}