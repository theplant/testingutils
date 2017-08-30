package testingutils

import (
	"encoding/json"
	"testing"
)

func sprint(diff string, messages []interface{}) string {
	var messagesString string
	for _, mess := range messages {
		jsonBytes, err := json.MarshalIndent(mess, "", "\t")
		if err != nil {
			panic(err)
		}
		messagesString = messagesString + string(jsonBytes) + "\n\n"
	}

	var printDiff = "\n" + diff + "\n"

	if messagesString == "" {
		return printDiff
	}
	// Remove "\n\n"
	messagesString = messagesString[:len(messagesString)-2]

	return printDiff + "Message:\n" + messagesString
}

func AssertEqual(
	t *testing.T,
	expected interface{},
	actual interface{},
	messages ...interface{}) bool {

	t.Helper()
	var diff = PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Error(sprint(diff, messages))
		return false
	}

	return true
}

func AssertEqualAndFatal(
	t *testing.T,
	expected interface{},
	actual interface{},
	messages ...interface{}) {

	t.Helper()
	var diff = PrettyJsonDiff(expected, actual)
	if len(diff) > 0 {
		t.Fatal(sprint(diff, messages))
	}
}
