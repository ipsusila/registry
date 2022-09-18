package registry_test

import (
	"testing"

	"github.com/ipsusila/registry"
	"github.com/stretchr/testify/assert"
)

func TestMapRegistry(t *testing.T) {
	reg := registry.NewMapRegistry[string, int]()
	err := reg.Register("hundred", 100)
	assert.NoError(t, err)
	err = reg.Register("ten", 10)
	assert.NoError(t, err)

	v, err := reg.Entry("ten")
	assert.NoError(t, err)
	assert.Equal(t, v, 10)
}
