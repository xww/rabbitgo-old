package cache

import (
	"sync"
	"time"

)

type Item struct {
	sync.RWMutex

	// The item's key.
	key interface{}
	// The item's data.
	value interface{}
	// How long will the item live in the cache when not being accessed/kept alive.
	lifeSpan time.Duration

	// Creation timestamp.
	createdOn time.Time
	// Last access timestamp.
	accessedOn time.Time
	// How often the item was accessed.
	accessCount int64

	// Callback method triggered
	extendFunction func(key interface{})
}

func NewItem(key interface{}, lifeSpan time.Duration, value interface{}) *Item {
	t := time.Now()
	return &Item{
		key:           key,
		lifeSpan:      lifeSpan,
		createdOn:     t,
		accessedOn:    t,
		accessCount:   0,
		extendFunction: nil,
		value:          value,
	}
}

func (item *Item) UpdateLifeSpan(t time.Duration){
	item.Lock()
	defer item.Unlock()
	item.lifeSpan = t

}
func (item *Item) GetLifeSpan() time.Duration{
	return item.lifeSpan
}

func (item *Item) UpdateAccess() {
	item.Lock()
	defer item.Unlock()
	item.accessedOn = time.Now()
	item.accessCount++
}

func (item *Item) GetKey() interface{} {
	return item.key
}

func (item *Item) GetValue() interface{}{
	return item.value
}

func (item *Item) GetAccessCount() int64{
	return item.accessCount
}

func (item *Item) GetCreateTime() time.Time{
	return item.createdOn
}

func (item *Item) SetExtendFunction(f func(interface{})){
	item.extendFunction = f
}



