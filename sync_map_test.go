package registry_test

import (
	"sync"
	"testing"

	"github.com/ipsusila/registry"
	"github.com/stretchr/testify/assert"
)

func TestSyncMapRegistry(t *testing.T) {
	reg := registry.NewSyncMapRegistry[string, int]()
	var wg sync.WaitGroup
	items := map[string]int{
		"hundred":  100,
		"ten":      10,
		"thousand": 1000,
	}
	wg.Add(len(items))
	for key, item := range items {
		go func(k string, i int) {
			err := reg.Register(k, i)
			assert.NoError(t, err)
			wg.Done()
		}(key, item)
	}
	wg.Wait()

	// check entry
	v, err := reg.Entry("ten")
	assert.NoError(t, err)
	assert.Equal(t, v, 10)

	assert.True(t, reg.Exists("hundred"))
	assert.False(t, reg.Exists("foo"))

	keys := reg.Keys()
	assert.Equal(t, len(items), len(keys))
	assert.True(t, contains(keys, "hundred"))

	entries := reg.Entries()
	assert.Equal(t, len(entries), len(items))
	assert.True(t, contains(entries, 1000))

	// replace
	v, err = reg.Replace("hundred", 101)
	assert.NoError(t, err)
	assert.Equal(t, v, 100)

	// Set value
	v = reg.Set("foo", -1)
	assert.Equal(t, v, 0)

	v = reg.Set("ten", 11)
	assert.Equal(t, v, 10)

	// remove foo
	v, ok := reg.Remove("foo")
	assert.True(t, ok)
	assert.Equal(t, v, -1)

	// clear
	err = reg.Clear()
	assert.NoError(t, err)
	assert.Zero(t, reg.NumEntries())
}
