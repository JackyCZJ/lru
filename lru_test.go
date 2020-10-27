package lru

import (
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
	c.Add("s1",s1)
	c.Add("s2",s2)
	c.Add("s3",s3)
	c.Add("s4",s4)
	c.Add("s5",s5)
	c.Add("s6",s6)
	_ ,ok := c.Get("s1")
	if ok{
		t.Fatal("lru false")
	}
	c.Add("s1",s1)
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