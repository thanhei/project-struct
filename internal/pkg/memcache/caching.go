package memcache

import (
	"go-training/internal/common"
	"sync"
	"time"
)

type Caching interface {
	Write(k string, value interface{})
	Read(k string) interface{}
}

type caching struct {
	store  map[string]interface{}
	locker *sync.RWMutex
}

func NewCaching() *caching {
	return &caching{
		store:  make(map[string]interface{}),
		locker: new(sync.RWMutex),
	}
}

func (c *caching) Write(k string, value interface{}) {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.store[k] = value
}

func (c *caching) Read(k string) interface{} {
	c.locker.RLock()
	defer c.locker.RUnlock()
	return c.store[k]
}

func (c *caching) WriteTTL(k string, value interface{}, exp int) {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.store[k] = value

	go func() {
		defer common.AppRecovery()
		<-time.NewTimer(time.Second * time.Duration(exp)).C
		c.Write(k, nil)
	}()
}
