package registry

import (
	"sync"
)

type immutableSyncMapRegistry[K comparable, E any] struct {
	sync.RWMutex
	entriesMap entryMap[K, E]
}

// NewImmutableSyncMapRegistry creates registry which are safe for concurrent access
func NewImmutableSyncMapRegistry[K comparable, E any]() ImmutableRegistry[K, E] {
	return &immutableSyncMapRegistry[K, E]{
		entriesMap: make(map[K]E),
	}
}

// NumEntries return number of entries
func (s *immutableSyncMapRegistry[K, E]) NumEntries() int {
	s.RLock()
	defer s.RUnlock()

	return s.entriesMap.numEntries()
}

// Register value with given key
func (s *immutableSyncMapRegistry[K, E]) Register(key K, entry E) error {
	s.Lock()
	defer s.Unlock()

	return s.entriesMap.register(key, entry)
}

// Exists return true if any entry has been registered with `key`
func (s *immutableSyncMapRegistry[K, E]) Exists(key K) bool {
	s.RLock()
	defer s.RUnlock()

	return s.entriesMap.exists(key)
}

// Entry value for given key
func (s *immutableSyncMapRegistry[K, E]) Entry(key K) (E, error) {
	s.RLock()
	defer s.RUnlock()

	return s.entriesMap.entry(key)
}

// Entries return list of registered entries
func (s *immutableSyncMapRegistry[K, E]) Entries() []E {
	s.RLock()
	defer s.RUnlock()

	return s.entriesMap.entries()
}

// Keys return list of registration keys
func (s *immutableSyncMapRegistry[K, E]) Keys() []K {
	s.RLock()
	defer s.RUnlock()

	return s.entriesMap.keys()
}
