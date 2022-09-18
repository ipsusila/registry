package registry

import (
	"sync"
)

type syncMapRegistry[K comparable, E any] struct {
	sync.RWMutex
	mapRegistry[K, E]
}

// NewSyncMapRegistry creates registry which are safe for concurrent access
func NewSyncMapRegistry[K comparable, E any]() Registry[K, E] {
	return &syncMapRegistry[K, E]{
		mapRegistry: mapRegistry[K, E]{
			entriesMap: make(map[K]E),
		},
	}
}

// Register value with given key
func (s *syncMapRegistry[K, E]) Register(key K, entry E) error {
	s.Lock()
	defer s.Unlock()

	return s.mapRegistry.Register(key, entry)
}

// Set value with given key
func (s *syncMapRegistry[K, E]) Set(key K, entry E) E {
	s.Lock()
	defer s.Unlock()

	return s.mapRegistry.Set(key, entry)
}

// Replace value with given key
func (s *syncMapRegistry[K, E]) Replace(key K, entry E) (E, error) {
	s.Lock()
	defer s.Unlock()

	return s.mapRegistry.Replace(key, entry)
}

// Remove value with given key
func (s *syncMapRegistry[K, E]) Remove(key K) (E, bool) {
	s.Lock()
	defer s.Unlock()

	return s.mapRegistry.Remove(key)
}

// Exists return true if any entry has been registered with `key`
func (s *syncMapRegistry[K, E]) Exists(key K) bool {
	s.RLock()
	defer s.RUnlock()

	return s.mapRegistry.Exists(key)
}

// Entry value for given key
func (s *syncMapRegistry[K, E]) Entry(key K) (E, error) {
	s.RLock()
	defer s.RUnlock()

	return s.mapRegistry.Entry(key)
}

// Entries return list of registered entries
func (s *syncMapRegistry[K, E]) Entries() []E {
	s.RLock()
	defer s.RUnlock()

	return s.mapRegistry.Entries()
}

// Keys return list of registration keys
func (s *syncMapRegistry[K, E]) Keys() []K {
	s.RLock()
	defer s.RUnlock()

	return s.mapRegistry.Keys()
}
