package assert

import (
	"errors"
	"testing"
)

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

func TestNoErrorAndFatal__true(t *testing.T) {
	NoErrorAndFatal(t, nil, "message")
}

func TestEqualError__True(t *testing.T) {
	err := errors.New("err")
	EqualError(t, err, err, "message1", "message2")
}

func TestEqualErrorAndFatal__True(t *testing.T) {
	err := errors.New("err")
	EqualErrorAndFatal(t, err, err, "message1", "message2")
}
