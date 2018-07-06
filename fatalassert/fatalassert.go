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

// If not equal then fatal.
//
// expectedList and actualList can be slice, array or nil.
// Compare type and elements with ignoring the order.
// For example,
//     * [1]int{} is not equal to []int{}
//     * []int{1, 1, 2} is equal to []int{2, 1, 1}
//
// The time complexity is O(n^2).
func UnorderedListEqual(
	t *testing.T,
	expectedList interface{},
	actualList interface{},
	messages ...interface{},
) {
	t.Helper()

	assert.UnorderedListEqual(t, assert.FatalHandle, expectedList, actualList, messages...)
}
