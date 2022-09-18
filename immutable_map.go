package registry

type immutableMapRegistry[K comparable, E any] struct {
	entriesMap entryMap[K, E]
}

// NewImmutableMapRegistry creates registry which are safe for concurrent access
func NewImmutableMapRegistry[K comparable, E any]() ImmutableRegistry[K, E] {
	return &immutableMapRegistry[K, E]{
		entriesMap: make(map[K]E),
	}
}

// NumEntries return number of entries
func (s *immutableMapRegistry[K, E]) NumEntries() int {
	return s.entriesMap.numEntries()
}

// Register value with given key
func (s *immutableMapRegistry[K, E]) Register(key K, entry E) error {
	return s.entriesMap.register(key, entry)
}

// Exists return true if any entry has been registered with `key`
func (s *immutableMapRegistry[K, E]) Exists(key K) bool {
	return s.entriesMap.exists(key)
}

// Entry value for given key
func (s *immutableMapRegistry[K, E]) Entry(key K) (E, error) {
	return s.entriesMap.entry(key)
}

// Entries return list of registered entries
func (s *immutableMapRegistry[K, E]) Entries() []E {
	return s.entriesMap.entries()
}

// Keys return list of registration keys
func (s *immutableMapRegistry[K, E]) Keys() []K {
	return s.entriesMap.keys()
}
