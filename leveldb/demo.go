package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, _ := leveldb.OpenFile("/tmp/db", nil)
	defer db.Close()

	db.Put([]byte("key"), []byte("value"), nil)
	data, _ := db.Get([]byte("key"), nil)
	fmt.Printf("%s", data)
	db.Delete([]byte("key"), nil)

}
