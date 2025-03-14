package utils

import (
	"sync"
)

type ConcurrentMap[KeyType comparable, ValueType any] struct {
	shards     []*ConcurrentMapShared[KeyType, ValueType]
	shardingFn func(key KeyType) uint32
	shardCount int
}

type ConcurrentMapShared[KeyType comparable, ValueType any] struct {
	items map[KeyType]ValueType
	sync.RWMutex
}

func NewConcurrentMap[K comparable, V any](sharding func(key K) uint32, shardCount int) ConcurrentMap[K, V] {
	m := ConcurrentMap[K, V]{
		shardingFn: sharding,
		shards:     make([]*ConcurrentMapShared[K, V], shardCount),
		shardCount: shardCount,
	}
	for i := 0; i < shardCount; i++ {
		m.shards[i] = &ConcurrentMapShared[K, V]{items: make(map[K]V)}
	}
	return m
}

func (m ConcurrentMap[K, V]) GetShard(key K) *ConcurrentMapShared[K, V] {
	return m.shards[uint(m.shardingFn(key))%uint(m.shardCount)]
}

// MStore stores the given data in the map.
func (m ConcurrentMap[K, V]) MStore(data map[K]V) {
	for key, value := range data {
		shard := m.GetShard(key)
		shard.Lock()
		shard.items[key] = value
		shard.Unlock()
	}
}

// Store the given value under the specified key.
func (m ConcurrentMap[K, V]) Store(key K, value V) {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

// Load retrieves an element from map under given key.
func (m ConcurrentMap[K, V]) Load(key K) (V, bool) {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	// Get item from shard.
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

// Delete removes an element from the map.
func (m ConcurrentMap[K, V]) Delete(key K) {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// Count returns the number of elements within the map.
func (m ConcurrentMap[K, V]) Count() int {
	count := 0
	for i := 0; i < m.shardCount; i++ {
		shard := m.shards[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

func (m ConcurrentMap[K, V]) Keys() []K {
	count := m.Count()
	ch := make(chan K, count)
	go func() {
		// Foreach shard.
		wg := sync.WaitGroup{}
		wg.Add(m.shardCount)
		for _, shard := range m.shards {
			go func(shard *ConcurrentMapShared[K, V]) {
				// Foreach key, value pair.
				shard.RLock()
				for key := range shard.items {
					ch <- key
				}
				shard.RUnlock()
				wg.Done()
			}(shard)
		}
		wg.Wait()
		close(ch)
	}()

	// Generate keys
	keys := make([]K, 0, count)
	for k := range ch {
		keys = append(keys, k)
	}
	return keys
}

func (m ConcurrentMap[KeyType, ValueType]) Values() []ValueType {
	count := m.Count()
	ch := make(chan ValueType, count)
	go func() {
		// Foreach shard.
		wg := sync.WaitGroup{}
		wg.Add(m.shardCount)
		for _, shard := range m.shards {
			go func(shard *ConcurrentMapShared[KeyType, ValueType]) {
				// Foreach key, value pair.
				shard.RLock()
				for _, value := range shard.items {
					ch <- value
				}
				shard.RUnlock()
				wg.Done()
			}(shard)
		}
		wg.Wait()
		close(ch)
	}()

	// Generate keys
	values := make([]ValueType, 0, count)
	for v := range ch {
		values = append(values, v)
	}
	return values
}

func Fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
