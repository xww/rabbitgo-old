package cache

import (
	"sync"
	"time"
)

var (
	cache = make(map[string]*Table)
	mutex sync.RWMutex
)

func NewCache(table string) *Table {
	mutex.RLock()
	t, ok := cache[table]
	mutex.RUnlock()

	if !ok {
		t = &Table{
			name:  table,
			items: make(map[interface{}]*Item),
			cleanupInterval: 3*time.Second,
			isLock: false,
			reveiveitems:make(chan *Item),
		}

		mutex.Lock()
		cache[table] = t
		mutex.Unlock()
	}
	go t.ExpireItemCheck()
	go t.ProcessItem()
	return t
}
