package cache

import (
	"testing"
	"time"
	"fmt"
)

var (
	k = "testkey"
	v = "testvalue"
)

func TestCache(t *testing.T) {
	table := NewCache("table1")
	table.AddItem(k,5,v)
	item, err:= table.GetItem(k)
	if err != nil{
		t.Error("add item error:", err)
	}
	fmt.Println(item)
	time.Sleep(50 * time.Second)
}
