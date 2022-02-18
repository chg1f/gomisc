package weak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	assert.Equal(t, true, IsEmpty(nil))
	assert.Equal(t, true, IsEmpty(""))
	assert.Equal(t, true, IsEmpty(0))
	var m map[interface{}]interface{}
	assert.Equal(t, true, IsEmpty(m))
	assert.Equal(t, true, IsEmpty(map[interface{}]interface{}{}))
	var a [1]interface{}
	assert.Equal(t, false, IsEmpty(a))
	assert.Equal(t, true, IsEmpty([...]interface{}{}))
	var s []interface{}
	assert.Equal(t, true, IsEmpty(s))
	assert.Equal(t, true, IsEmpty([]interface{}{}))
	var f func()
	assert.Equal(t, true, IsEmpty(f))
}
