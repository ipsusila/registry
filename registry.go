package registry

import "errors"

// List of error
var (
	ErrDuplicateEntry    = errors.New("duplicate entry")
	ErrEntryDoesNotExist = errors.New("entry does not exists")
)

// Registry interface
type Registry[K comparable, E any] interface {
	Register(key K, entry E) error
	Set(key K, entry E) E
	Replace(key K, entry E) (E, error)
	Remove(key K) (E, bool)
	Exists(key K) bool
	Entry(key K) (E, error)
	Keys() []K
	Entries() []E
}
