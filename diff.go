package testingutils

import (
	"encoding/json"

	"fmt"

	"github.com/pmezard/go-difflib/difflib"
)

/*
It convert the two objects into pretty json, and diff them, output the result.
*/
func PrettyJsonDiff(expected interface{}, actual interface{}) (r string) {
	actualJson, _ := json.MarshalIndent(actual, "", "\t")
	expectedJson, _ := json.MarshalIndent(expected, "", "\t")
	if string(actualJson) != string(expectedJson) {
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(string(expectedJson)),
			B:        difflib.SplitLines(string(actualJson)),
			FromFile: "Expected",
			ToFile:   "Actual",
			Context:  3,
		}
		r, _ = difflib.GetUnifiedDiffString(diff)
	}
	return
}

func PrintlnJson(vals ...interface{}) {
	var newvals []interface{}
	for _, v := range vals {
		if s, ok := v.(string); ok {
			newvals = append(newvals, s)
			continue
		}
		j, _ := json.MarshalIndent(v, "", "\t")
		newvals = append(newvals, "\n", string(j))
	}
	fmt.Println(newvals...)
}
