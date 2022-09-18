package registry

type mapRegistry[K comparable, E any] struct {
	immutableMapRegistry[K, E]
}

// NewMapRegistry creates registry which are safe for concurrent access
func NewMapRegistry[K comparable, E any]() Registry[K, E] {
	return &mapRegistry[K, E]{
		immutableMapRegistry: immutableMapRegistry[K, E]{
			entriesMap: make(map[K]E),
		},
	}
}

// Set value with given key
func (s *mapRegistry[K, E]) Set(key K, entry E) E {
	return s.entriesMap.set(key, entry)
}

// Replace value with given key
func (s *mapRegistry[K, E]) Replace(key K, entry E) (E, error) {
	return s.entriesMap.replace(key, entry)
}

// Remove value with given key
func (s *mapRegistry[K, E]) Remove(key K) (E, bool) {
	return s.entriesMap.remove(key)
}

// Clear entries
func (s *mapRegistry[K, E]) Clear() error {
	s.entriesMap = make(map[K]E)
	return nil
}
