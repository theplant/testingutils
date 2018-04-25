package assert

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/theplant/testingutils"
)

func SprintMessages(text string, messages []interface{}) string {
	var messagesString string
	for _, mess := range messages {
		jsonBytes, err := json.MarshalIndent(mess, "", "\t")
		if err != nil {
			panic(err)
		}
		messagesString = messagesString + string(jsonBytes) + "\n\n"
	}

	if messagesString == "" {
		return text
	}
	// Remove "\n\n"
	messagesString = messagesString[:len(messagesString)-2]

	return text + "\n" + "Messages:\n" + messagesString
}

func Equal(
	t *testing.T,
	expected interface{},
	actual interface{},
	messages ...interface{},
) bool {
	t.Helper()
	var diff = testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(SprintMessages("\n"+diff, messages))
		return false
	}

	return true
}

func EqualAndFatal(
	t *testing.T,
	expected interface{},
	actual interface{},
	messages ...interface{},
) {
	t.Helper()
	var diff = testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Fatal(SprintMessages("\n"+diff, messages))
	}
}

func NotEqualAndFatal(
	t *testing.T,
	expected interface{},
	actual interface{},
	messages ...interface{},
) {
	t.Helper()
	var diff = testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) == 0 {
		j, err := json.MarshalIndent(actual, "", "\t")
		if err != nil {
			t.Fatal(err)
		}
		t.Fatal(SprintMessages("expected is equal to actual\n"+string(j), messages))
	}
}

func NoError(t *testing.T, err error, messages ...interface{}) {
	t.Helper()
	if err != nil {
		t.Error(SprintMessages(
			fmt.Sprintf("Got an unexpected error:\n%+v", err),
			messages,
		))
	}
}

func NoErrorAndFatal(t *testing.T, err error, messages ...interface{}) {
	t.Helper()
	if err != nil {
		t.Fatal(SprintMessages(
			fmt.Sprintf("Got an unexpected error:\n%+v", err),
			messages,
		))
	}
}

func EqualError(
	t *testing.T,
	expectedErr error,
	actualErr error,
	messages ...interface{},
) {
	t.Helper()
	if expectedErr != actualErr {
		t.Error(SprintMessages(
			fmt.Sprintf("Errors are not equal\nexpected: %+v\nactual: %+v", expectedErr, actualErr),
			messages,
		))
	}
}

func EqualErrorAndFatal(
	t *testing.T,
	expectedErr error,
	actualErr error,
	messages ...interface{},
) {
	t.Helper()
	if expectedErr != actualErr {
		t.Fatal(SprintMessages(
			fmt.Sprintf("Errors are not equal\nexpected: %+v\nactual: %+v", expectedErr, actualErr),
			messages,
		))
	}
}
