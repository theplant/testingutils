package testingutils

import "testing"

func TestAssertEqual__True(t *testing.T) {
	AssertEqual(t, 1, 1, "message1", "message2")
}

func TestAssertEqualAndFatal__True(t *testing.T) {
	AssertEqualAndFatal(t, 2, 2, "message1", "message2")
}
