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
	actualJson := marshalIfNotString(actual)
	expectedJson := marshalIfNotString(expected)
	if actualJson != expectedJson {
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(expectedJson),
			B:        difflib.SplitLines(actualJson),
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

func marshalIfNotString(v interface{}) (r string) {
	var ok bool
	if r, ok = v.(string); ok {
		return
	}
	rbytes, _ := json.MarshalIndent(v, "", "\t")
	r = string(rbytes)
	return
}
