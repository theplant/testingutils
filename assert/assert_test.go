package assert

import "testing"

func TestAssertEqual__True(t *testing.T) {
	Equal(t, 1, 1, "message1", "message2")
}

func TestAssertEqualAndFatal__True(t *testing.T) {
	EqualAndFatal(t, 2, 2, "message1", "message2")
}

func TestNotEqualAndFatal__True(t *testing.T) {
	NotEqualAndFatal(t, 1, 2, "message1", "message2")
}

func TestNoError__True(t *testing.T) {
	NoError(t, nil, "message")
}

func TestNoErrorAndFatal(t *testing.T) {
	NoErrorAndFatal(t, nil, "message")
}
