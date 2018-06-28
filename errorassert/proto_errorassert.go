package errorassert

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/theplant/testingutils/assert"
)

// If not equal then error.
func ProtoEqual(
	t *testing.T,
	expected proto.Message,
	actual proto.Message,
	messages ...interface{},
) {
	t.Helper()

	assert.ProtoEqual(t, assert.ErrorHandle, expected, actual, messages...)
}

// If equal then error.
func ProtoNotEqual(
	t *testing.T,
	expected proto.Message,
	actual proto.Message,
	messages ...interface{},
) {
	t.Helper()

	assert.ProtoNotEqual(t, assert.ErrorHandle, expected, actual, messages...)
}
