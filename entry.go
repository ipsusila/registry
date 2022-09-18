package registry

import "fmt"

type entryMap[K comparable, E any] map[K]E

func (e entryMap[K, E]) set(key K, entry E) E {
	old := e[key]
	e[key] = entry

	return old
}

// Replace value with given key
func (e entryMap[K, E]) replace(key K, entry E) (E, error) {
	old, ok := e[key]
	if !ok {
		return old, fmt.Errorf("replace entry %v: %w", key, ErrEntryDoesNotExist)
	}
	e[key] = entry

	return old, nil
}

// Remove value with given key
func (e entryMap[K, E]) remove(key K) (E, bool) {
	old, ok := e[key]
	if ok {
		delete(e, key)
	}

	return old, ok
}

func (e entryMap[K, E]) numEntries() int {
	return len(e)
}

func (e entryMap[K, E]) register(key K, entry E) error {
	if _, dup := e[key]; dup {
		return fmt.Errorf("entry with key `%v` already registered: %w", key, ErrDuplicateEntry)
	}
	e[key] = entry

	return nil
}

func (e entryMap[K, E]) exists(key K) bool {
	_, found := e[key]
	return found
}

// Entry value for given key
func (e entryMap[K, E]) entry(key K) (E, error) {
	v, ok := e[key]
	if !ok {
		return v, fmt.Errorf("error getting entry %v: %w", key, ErrEntryDoesNotExist)
	}
	return v, nil
}

// Entries return list of registered entries
func (e entryMap[K, E]) entries() []E {
	entries := []E{}
	for _, entry := range e {
		entries = append(entries, entry)
	}
	return entries
}

// Keys return list of registration keys
func (e entryMap[K, E]) keys() []K {
	keys := []K{}
	for key := range e {
		keys = append(keys, key)
	}
	return keys
}
