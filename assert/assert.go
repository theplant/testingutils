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
	Indent:                  "\t",
	DisableMethods:          true,
	DisablePointerAddresses: true,
	ContinueOnMethod:        true,
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

func getDiff(
	expected interface{},
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
		panic(fmt.Sprintf("difflib.GetUnifiedDiffString failed: %v", err))
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
		printDiff(t, handleType, expected, actual, messages...)
	}
}

func printDiff(
	t *testing.T,
	handleType HandleType,
	expected interface{},
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	expectedJSON := jsonMarshal(expected)
	actualJSON := jsonMarshal(actual)
	diff := testingutils.PrettyJsonDiff(expected, actual)

	if diff == "" || isJSONNullOrEmpty(expectedJSON) || isJSONNullOrEmpty(actualJSON) {
		diff = getDiff(expected, actual)
	}

	logs := SprintMessages("Expected is not equal to actual:\n"+diff, messages)
	if handleType == ErrorHandle {
		t.Error(logs)
	} else {
		t.Fatal(logs)
	}
}

func jsonMarshal(i interface{}) string {
	j, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		panic(fmt.Sprintf("json.MarshalIndent failed: %v", err))
	}
	return string(j)
}

func dump(iface interface{}) string {
	str := jsonMarshal(iface)
	if str == "" {
		str = SpewConfig.Sdump(iface)
	}
	return reflect.TypeOf(iface).String() + "\n" + str + "\n"
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
		str := dump(actual)
		logs := SprintMessages("Expected is equal to actual, actual is:\n"+str, messages)
		if handleType == ErrorHandle {
			t.Error(logs)
		} else {
			t.Fatal(logs)
		}
	}
}

func isNil(iface interface{}) bool {
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

// If has error then handle.
func NoError(
	t *testing.T,
	handleType HandleType,
	err error,
	messages ...interface{},
) {
	t.Helper()

	if !isNil(err) {
		logs := SprintMessages(fmt.Sprintf("Got an unexpected error:\n%+v", err), messages)
		if handleType == ErrorHandle {
			t.Error(logs)
		} else {
			t.Fatal(logs)
		}
	}
}

// If not nil then handle.
func Nil(
	t *testing.T,
	handleType HandleType,
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	if !isNil(actual) {
		str := dump(actual)
		logs := SprintMessages("Actual is not nil:\n"+str, messages)
		if handleType == ErrorHandle {
			t.Error(logs)
		} else {
			t.Fatal(logs)
		}
	}
}

// If nil then handle.
func NotNil(
	t *testing.T,
	handleType HandleType,
	actual interface{},
	messages ...interface{},
) {
	t.Helper()

	if isNil(actual) {
		logs := SprintMessages("Actual is nil", messages)
		if handleType == ErrorHandle {
			t.Error(logs)
		} else {
			t.Fatal(logs)
		}
	}
}

// If not equal then handle.
//
// expectedList and actualList can be slice, array or nil.
// Compare type and elements with ignoring the order.
// For example,
//     * [1]int{} is not equal to []int{}
//     * []int{1, 1, 2} is equal to []int{2, 1, 1}
//
// The time complexity is O(n^2).
func UnorderedListEqual(
	t *testing.T,
	handleType HandleType,
	expectedList interface{},
	actualList interface{},
	messages ...interface{},
) {
	t.Helper()

	if !isListType(expectedList) {
		if handleType == ErrorHandle {
			t.Errorf("expectedList must be slice, array or nil, but its type is: %v", getType(expectedList))
		} else {
			t.Fatalf("expectedList must be slice, array or nil, but its type is: %v", getType(expectedList))
		}
	}

	if !isListType(actualList) {
		if handleType == ErrorHandle {
			t.Errorf("actualList must be slice, array or nil, but its type is: %v", getType(actualList))
		} else {
			t.Fatalf("actualList must be slice, array or nil, but its type is: %v", getType(actualList))
		}
	}

	if !isUnorderedListEqualedWithTypeCheck(expectedList, actualList) {
		printDiff(t, handleType, expectedList, actualList, messages...)
	}
}

func getType(iface interface{}) string {
	rtype := reflect.TypeOf(iface)
	if rtype == nil {
		return "nil"
	}
	return rtype.Kind().String()
}

func isUnorderedListEqualedWithTypeCheck(list1 interface{}, list2 interface{}) bool {
	isTypeEqualed := isListTypeEqualed(list1, list2)
	if !isTypeEqualed {
		return false
	}

	return isUnorderedListEqualed(list1, list2)
}

// If iface's type is array, slice or nil then return true.
func isListType(iface interface{}) bool {
	if iface == nil {
		return true
	}

	rtype := reflect.TypeOf(iface)

	if rtype.Kind() == reflect.Slice {
		return true
	}

	if rtype.Kind() == reflect.Array {
		return true
	}

	return false
}

// Please make sure list1 and list2 kind is array, slice or nil.
func isListTypeEqualed(list1 interface{}, list2 interface{}) bool {
	if list1 == nil || list2 == nil {
		return list1 == list2
	}

	rtype1 := reflect.TypeOf(list1)
	rtype2 := reflect.TypeOf(list2)

	if rtype1.Kind() != rtype2.Kind() {
		return false
	}

	if rtype1.Elem().Kind() != rtype2.Elem().Kind() {
		return false
	}

	if rtype1.Kind() == reflect.Array {
		if rtype1.Len() != rtype2.Len() {
			return false
		}
	}

	return true
}

// Please make sure list1 and list2 kind is slice, array or nil.
func isUnorderedListEqualed(list1 interface{}, list2 interface{}) bool {
	if list1 == nil || list2 == nil {
		return list1 == list2
	}

	value1 := reflect.ValueOf(list1)
	value2 := reflect.ValueOf(list2)
	len1 := value1.Len()
	len2 := value2.Len()

	if len1 != len2 {
		return false
	}

	foundInList2 := make([]bool, len2)
	for i := 0; i < len1; i++ {
		element1 := value1.Index(i).Interface()
		found := false
		for j := 0; j < len2; j++ {
			if foundInList2[j] {
				continue
			}

			element2 := value2.Index(j).Interface()
			if reflect.DeepEqual(element1, element2) {
				foundInList2[j] = true
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
