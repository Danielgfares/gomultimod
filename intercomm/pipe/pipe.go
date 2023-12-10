package pipe

import (
	"sync"
)

type IPipe interface {
	Write(id, value string) bool
	Read(id string) (string, bool)
}

type Pipe struct {
	m     map[string]string
	mutex sync.RWMutex
}

func NewPipe() *Pipe {
	return &Pipe{
		m:     make(map[string]string),
		mutex: sync.RWMutex{},
	}
}

func (p *Pipe) Write(key, value string) bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// if key is already in the pipe then do nothing
	if _, ok := p.m[key]; ok {
		return false
	}
	p.m[key] = value
	return true
}

func (p *Pipe) Read(key string) (string, bool) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if value, ok := p.m[key]; ok {
		delete(p.m, key)
		return value, true
	}
	return "", false
}
