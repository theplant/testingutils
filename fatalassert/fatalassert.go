package fatalassert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/theplant/testingutils"
	"github.com/theplant/testingutils/assert"
)

func prettyPrintDiff(
	t *testing.T,
	expected,
	actual interface{},
) (
	diff string,
) {
	expectedString := spew.Sdump(expected)
	actualString := spew.Sdump(actual)
	diff, err := difflib.GetUnifiedDiffString(
		difflib.UnifiedDiff{
			A:        difflib.SplitLines(expectedString),
			B:        difflib.SplitLines(actualString),
			FromFile: "Expected",
			FromDate: "",
			ToFile:   "Actual",
			ToDate:   "",
			Context:  3,
		})
	if err != nil {
		t.Fatal("difflib.GetUnifiedDiffString failed", err)
	}
	return diff
}

func isJSONNullOrEmpty(str string) bool {
	if str == "null" || str == "{}" {
		return true
	} else {
		return false
	}
}

// If not equal then fatal.
func Equal(
	t *testing.T,
	expected interface{},
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		expectedJSON := jsonMarshal(t, expected)
		actualJSON := jsonMarshal(t, actual)
		diff := testingutils.PrettyJsonDiff(expected, actual)

		if diff == "" || isJSONNullOrEmpty(expectedJSON) || isJSONNullOrEmpty(actualJSON) {
			diff = prettyPrintDiff(t, expected, actual)
		}

		t.Fatal(assert.SprintMessages("\n"+diff, messages))
	}
}

func jsonMarshal(t *testing.T, i interface{}) string {
	j, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Fatal("json.MarshalIndent failed", err)
	}
	return string(j)
}

// If equal then fatal.
func NotEqual(
	t *testing.T,
	expected interface{},
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	if reflect.DeepEqual(expected, actual) {
		str := jsonMarshal(t, actual)
		if str == "" {
			str = spew.Sdump(actual)
		}

		t.Fatal(assert.SprintMessages("Expected is equal to actual, actual is:\n"+str, messages))
	}
}

func IsNil(iface interface{}) bool {
	if iface == nil {
		return true
	}
	return reflect.ValueOf(iface).IsNil()
}

// If has error, then fatal.
func NoError(
	t *testing.T,
	err error,
	messages ...interface{},
) {
	t.Helper()

	if !IsNil(err) {
		t.Fatal(assert.SprintMessages(fmt.Sprintf("Got an unexpected error:\n%+v", err), messages))
	}
}
