package registry

import (
	"fmt"
)

type mapRegistry[K comparable, E any] struct {
	entriesMap map[K]E
}

// NewMapRegistry creates registry which are safe for concurrent access
func NewMapRegistry[K comparable, E any]() Registry[K, E] {
	return &mapRegistry[K, E]{
		entriesMap: make(map[K]E),
	}
}

// Register value with given key
func (s *mapRegistry[K, E]) Register(key K, entry E) error {
	if _, dup := s.entriesMap[key]; dup {
		return fmt.Errorf("entry with key `%v` already registered: %w", key, ErrDuplicateEntry)
	}
	s.entriesMap[key] = entry

	return nil
}

// Exists return true if any entry has been registered with `key`
func (s *mapRegistry[K, E]) Exists(key K) bool {
	_, found := s.entriesMap[key]
	return found
}

// Entry value for given key
func (s *mapRegistry[K, E]) Entry(key K) (E, error) {
	v, ok := s.entriesMap[key]
	if !ok {
		return v, fmt.Errorf("error getting entry %v: %w", key, ErrEntryDoesNotExist)
	}
	return v, nil
}

// Entries return list of registered entries
func (s *mapRegistry[K, E]) Entries() []E {
	entries := []E{}
	for _, entry := range s.entriesMap {
		entries = append(entries, entry)
	}
	return entries
}

// Keys return list of registration keys
func (s *mapRegistry[K, E]) Keys() []K {
	keys := []K{}
	for key := range s.entriesMap {
		keys = append(keys, key)
	}
	return keys
}
