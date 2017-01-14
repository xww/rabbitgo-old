package cache

import (
	"sync"
	"time"
	"log"
	"fmt"
)

// CacheTable is a table within the cache
type Table struct {
	sync.RWMutex

	// The table's name.
	name string
	// All cached items.
	items map[interface{}]*Item

	// Timer responsible for triggering cleanup.
	//cleanupTimer *time.Timer
	// Current timer duration.
	cleanupInterval time.Duration

	// The logger used for this table.
	logger *log.Logger


	// Callback method triggered when trying to load a non-existing key.
	//loadData func(key interface{}, args ...interface{}) *Item
	// Callback method triggered when adding a new item to the cache.
	//addedItem func(item *Item)
	// Callback method triggered before deleting an item from the cache.
	//aboutToDeleteItem func(item *Item)

	isLock bool
}



func (table * Table) GetCount() int64 {
	table.RLock()
	defer table.RUnlock()
	return int64(len(table.items))
}

func (table * Table) GetItem(key interface{}) (*Item, error){
	table.RLock()
	defer table.RUnlock()
	item, ok := table.items[key]
	if ok {
		item.UpdateAccess()
		return item, nil
	}
	return nil, ErrKeyNotFound
}

func (table * Table) AddItem(key interface{}, lifespan time.Duration, value interface{}) *Item{
	item := NewItem(key, lifespan, value)
	table.Lock()
	defer table.Unlock()
	table.items[key] = item
	return item
}

func (table * Table) DeleteItem(key interface{}) error{
	table.Lock()
	defer table.Unlock()
	_, ok := table.items[key]
	if !ok{
		return ErrKeyNotFound
	}
	delete(table.items, key)
	return nil
}

func (table * Table) KeyExist(key interface{}) bool{
	table.Lock()
	defer table.Unlock()
	_, ok := table.items[key]
	if ok {
		return  true
	}
	return false
}

func (table * Table) Flush() {
	table.Lock()
	defer table.Unlock()

}



func (table *Table) ExpireItemCheck() {
	tc := time.Tick(table.cleanupInterval)
	for {
		<- tc
		fmt.Println("begin clean up table ", table.name)
		table.Lock()
		items := table.items
		table.Unlock()
		now := time.Now()
		for key, item := range items {
			item.RLock()
			lifeSpan := item.lifeSpan
			accessedOn := item.accessedOn
			item.RUnlock()
			if lifeSpan == 0 {
				continue
			}
			if now.Sub(accessedOn) >= lifeSpan {
				// Item has excessed its lifespan.
				table.DeleteItem(key)
			}
		}
		//table.Unlock()
	}
}



