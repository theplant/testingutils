package fatalassert

import (
	"testing"

	"github.com/theplant/testingutils/assert"
)

// If not equal then fatal.
func Equal(
	t *testing.T,
	expected interface{},
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	assert.Equal(t, assert.FatalHandle, expected, actual, messages...)
}

// If equal then fatal.
func NotEqual(
	t *testing.T,
	expected interface{},
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	assert.NotEqual(t, assert.FatalHandle, expected, actual, messages...)
}

// If not equal then fatal.
func EqualError(
	t *testing.T,
	expected error,
	actual error,
	messages ...interface{},
) {
	t.Helper()

	assert.EqualError(t, assert.FatalHandle, expected, actual, messages...)
}

// If has error then fatal.
func NoError(
	t *testing.T,
	err error,
	messages ...interface{},
) {
	t.Helper()

	assert.NoError(t, assert.FatalHandle, err, messages...)
}

// If not nil then fatal.
func Nil(
	t *testing.T,
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	assert.Nil(t, assert.FatalHandle, actual, messages...)
}

// If nil then fatal.
func NotNil(
	t *testing.T,
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	assert.NotNil(t, assert.FatalHandle, actual, messages...)
}
