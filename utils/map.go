package utils

import (
	"sync"
)

type CMap[KeyType comparable, ValueType any] struct {
	mu      sync.RWMutex
	storage map[KeyType]ValueType
}

func (e *CMap[KeyType, ValueType]) Store(key KeyType, value ValueType) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.storage == nil {
		e.storage = make(map[KeyType]ValueType)
	}
	e.storage[key] = value
}

func (e *CMap[KeyType, ValueType]) Load(key KeyType) (value ValueType, ok bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	value, ok = e.storage[key]
	return
}

func (e *CMap[KeyType, ValueType]) LoadOrStore(key KeyType, value ValueType) (actual ValueType, loaded bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.storage == nil {
		e.storage = make(map[KeyType]ValueType)
	}
	actual, loaded = e.storage[key]
	if !loaded {
		e.storage[key] = value
		actual = value
	}
	return
}

func (e *CMap[KeyType, ValueType]) Delete(key KeyType) {
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.storage, key)
}

func (e *CMap[KeyType, ValueType]) AllValues() []ValueType {
	e.mu.RLock()
	defer e.mu.RUnlock()

	var values []ValueType
	for _, value := range e.storage {
		values = append(values, value)
	}

	return values

}

func (e *CMap[KeyType, ValueType]) Range(f func(key KeyType, value ValueType) bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	for k, v := range e.storage {
		if !f(k, v) {
			break
		}
	}
}

func (e *CMap[KeyType, ValueType]) EntireMap() map[KeyType]ValueType {
	e.mu.RLock()
	defer e.mu.RUnlock()

	mapCopy := make(map[KeyType]ValueType, len(e.storage))
	for k, v := range e.storage {
		mapCopy[k] = v
	}
	return mapCopy
}
