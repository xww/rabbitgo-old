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
			cleanupInterval: 30*time.Second,
		}

		mutex.Lock()
		cache[table] = t
		mutex.Unlock()
	}
	go t.ExpireItemCheck()
	return t
}
