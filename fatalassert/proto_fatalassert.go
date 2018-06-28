package fatalassert

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/theplant/testingutils/assert"
)

// If not equal then fatal.
func ProtoEqual(
	t *testing.T,
	expected proto.Message,
	actual proto.Message,
	messages ...interface{},
) {
	t.Helper()

	assert.ProtoEqual(t, assert.FatalHandle, expected, actual, messages...)
}

// If equal then fatal.
func ProtoNotEqual(
	t *testing.T,
	expected proto.Message,
	actual proto.Message,
	messages ...interface{},
) {
	t.Helper()

	assert.ProtoNotEqual(t, assert.FatalHandle, expected, actual, messages...)
}
