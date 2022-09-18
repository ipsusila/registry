package registry

import "errors"

// List of error
var (
	ErrDuplicateEntry    = errors.New("duplicate entry")
	ErrEntryDoesNotExist = errors.New("entry does not exists")
)

type ImmutableRegistry[K comparable, E any] interface {
	Register(key K, entry E) error
	Exists(key K) bool
	Entry(key K) (E, error)
	Keys() []K
	Entries() []E
	NumEntries() int
}

// Registry interface
type Registry[K comparable, E any] interface {
	ImmutableRegistry[K, E]
	Set(key K, entry E) E
	Replace(key K, entry E) (E, error)
	Remove(key K) (E, bool)
	Clear() error
}
