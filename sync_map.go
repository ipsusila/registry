package registry

type syncMapRegistry[K comparable, E any] struct {
	immutableSyncMapRegistry[K, E]
}

// NewSyncMapRegistry creates registry which are safe for concurrent access
func NewSyncMapRegistry[K comparable, E any]() Registry[K, E] {
	return &syncMapRegistry[K, E]{
		immutableSyncMapRegistry: immutableSyncMapRegistry[K, E]{
			entriesMap: make(map[K]E),
		},
	}
}

// Set value with given key
func (s *syncMapRegistry[K, E]) Set(key K, entry E) E {
	s.Lock()
	defer s.Unlock()

	return s.entriesMap.set(key, entry)
}

// Replace value with given key
func (s *syncMapRegistry[K, E]) Replace(key K, entry E) (E, error) {
	s.Lock()
	defer s.Unlock()

	return s.entriesMap.replace(key, entry)
}

// Remove value with given key
func (s *syncMapRegistry[K, E]) Remove(key K) (E, bool) {
	s.Lock()
	defer s.Unlock()

	return s.entriesMap.remove(key)
}

// Clear value with given key
func (s *syncMapRegistry[K, E]) Clear() error {
	s.Lock()
	defer s.Unlock()
	s.entriesMap = make(map[K]E)

	return nil
}
