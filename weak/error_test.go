package weak

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeakError(t *testing.T) {
	assert.Error(t, WeakError{I: io.EOF})
	assert.NotErrorIs(t, WeakError{I: io.EOF}, io.EOF)
}

func TestCompressError(t *testing.T) {
	ce := CompressError{}
	assert.Error(t, ce)
	ce = ce.Shrink()
	assert.Nil(t, ce)

	assert.NotErrorIs(t, ce, io.EOF)
	ce = CompressError{io.EOF}
	assert.ErrorIs(t, ce, io.EOF)
}
