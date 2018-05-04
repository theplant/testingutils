//go:generate stringer -output stringers.go -type "HandleType"
package assert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/theplant/testingutils"
)

type HandleType int

const (
	ErrorHandle HandleType = iota
	FatalHandle
)

var SpewConfig = spew.ConfigState{
	Indent:           "\t",
	DisableMethods:   true,
	ContinueOnMethod: true,
}

func SprintMessages(text string, messages []interface{}) string {
	var messagesString string
	for _, mess := range messages {
		jsonBytes, err := json.MarshalIndent(mess, "", "\t")
		if err != nil {
			panic(err)
		}
		messagesString = messagesString + reflect.TypeOf(mess).String() + " " + string(jsonBytes) + "\n\n"
	}

	if messagesString == "" {
		return text
	}
	// Remove "\n\n"
	messagesString = messagesString[:len(messagesString)-2]

	return text + "\n" + "Messages:\n" + messagesString
}

func prettyPrintDiff(
	t *testing.T,
	expected,
	actual interface{},
) (
	diff string,
) {
	expectedString := SpewConfig.Sdump(expected)
	actualString := SpewConfig.Sdump(actual)
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
		t.Fatalf("difflib.GetUnifiedDiffString failed: %v", err)
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

// If not equal then handle.
func Equal(
	t *testing.T,
	handleType HandleType,
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

		logs := SprintMessages("Expected is not equal to actual:\n"+diff, messages)
		if handleType == ErrorHandle {
			t.Error(logs)
		} else {
			t.Fatal(logs)
		}
	}
}

func jsonMarshal(t *testing.T, i interface{}) string {
	j, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Fatalf("json.MarshalIndent failed: %v", err)
	}
	return string(j)
}

// If equal then handle.
func NotEqual(
	t *testing.T,
	handleType HandleType,
	expected interface{},
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	if reflect.DeepEqual(expected, actual) {
		str := jsonMarshal(t, actual)
		if str == "" {
			str = SpewConfig.Sdump(actual)
		}

		logs := SprintMessages("Expected is equal to actual, actual is:\n"+str, messages)
		if handleType == ErrorHandle {
			t.Error(logs)
		} else {
			t.Fatal(logs)
		}
	}
}

func IsNil(iface interface{}) bool {
	if iface == nil {
		return true
	}
	return reflect.ValueOf(iface).IsNil()
}

// If equal then handle.
func EqualError(
	t *testing.T,
	handleType HandleType,
	expected error,
	actual error,
	messages ...interface{},
) {
	t.Helper()

	Equal(t, handleType, expected, actual, messages)
}

// If has error, then handle.
func NoError(
	t *testing.T,
	handleType HandleType,
	err error,
	messages ...interface{},
) {
	t.Helper()

	if !IsNil(err) {
		logs := SprintMessages(fmt.Sprintf("Got an unexpected error:\n%+v", err), messages)
		if handleType == ErrorHandle {
			t.Error(logs)
		} else {
			t.Fatal(logs)
		}
	}
}
