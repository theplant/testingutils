package testingutils

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/theplant/testingutils"
)

func sprintMessages(text string, messages []interface{}) string {
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
	messages ...interface{}) bool {

	t.Helper()
	var diff = testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(sprintMessages("\n"+diff, messages))
		return false
	}

	return true
}

func EqualAndFatal(
	t *testing.T,
	expected interface{},
	actual interface{},
	messages ...interface{}) {

	t.Helper()
	var diff = testingutils.PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Fatal(sprintMessages("\n"+diff, messages))
	}
}

func NoError(t *testing.T, err error, messages ...interface{}) {
	t.Helper()
	if err != nil {
		t.Error(sprintMessages(
			fmt.Sprintf("Got an unexpected error:\n%+v", err),
			messages,
		))
	}
}

func NoErrorAndFatal(t *testing.T, err error, messages ...interface{}) {
	t.Helper()
	if err != nil {
		t.Fatal(sprintMessages(
			fmt.Sprintf("Got an unexpected error:\n%+v", err),
			messages,
		))
	}
}
