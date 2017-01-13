

package main

import (
	"time"
	"fmt"
	
)


)

var (
	k = "testkey"
	v = "testvalue"
)

func main() {
	table := NewCache("table1")
	table.AddItem(k,5,v)
	item, err:= table.GetItem(k)
	if err != nil{
		fmt.Println("add item error:", err)
	}
	fmt.Println(item)
	time.Sleep(50 * time.Second)
}