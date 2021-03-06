package errorassert

import (
	"testing"

	"github.com/theplant/testingutils/assert"
)

// If not equal then error.
func Equal(
	t *testing.T,
	expected interface{},
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	assert.Equal(t, assert.ErrorHandle, expected, actual, messages...)
}

// If equal then error.
func NotEqual(
	t *testing.T,
	expected interface{},
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	assert.NotEqual(t, assert.ErrorHandle, expected, actual, messages...)
}

// If not equal then error.
func EqualError(
	t *testing.T,
	expected error,
	actual error,
	messages ...interface{},
) {
	t.Helper()

	assert.EqualError(t, assert.ErrorHandle, expected, actual, messages...)
}

// If has error then error.
func NoError(
	t *testing.T,
	err error,
	messages ...interface{},
) {
	t.Helper()

	assert.NoError(t, assert.ErrorHandle, err, messages...)
}

// If not nil then error.
func Nil(
	t *testing.T,
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	assert.Nil(t, assert.ErrorHandle, actual, messages...)
}

// If nil then error.
func NotNil(
	t *testing.T,
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	assert.NotNil(t, assert.ErrorHandle, actual, messages...)
}

// If not equal then error.
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

	assert.UnorderedListEqual(t, assert.ErrorHandle, expectedList, actualList, messages...)
}
