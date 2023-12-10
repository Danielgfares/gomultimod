package channel

import (
	"sync"
)

type IChannel interface {
	Write(id, value string) bool
	Read(id string) (string, bool)
}

type Channel struct {
	m     map[string]string
	mutex sync.RWMutex
}

func NewPipe() *Channel {
	return &Channel{
		m:     make(map[string]string),
		mutex: sync.RWMutex{},
	}
}

func (c *Channel) Write(key, value string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// if key is already in the channel then do nothing
	if _, ok := c.m[key]; ok {
		return false
	}
	c.m[key] = value
	return true
}

func (c *Channel) Read(key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if value, ok := c.m[key]; ok {
		delete(c.m, key)
		return value, true
	}
	return "", false
}
