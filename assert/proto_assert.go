package assert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/pmezard/go-difflib/difflib"
)

var ProtoJSONMarshaler = &jsonpb.Marshaler{
	EnumsAsInts:  false,
	EmitDefaults: false,
	Indent:       "\t",
	OrigName:     false,
}

func protoJSONMarshal(message proto.Message) string {
	j, err := ProtoJSONMarshaler.MarshalToString(message)
	if err != nil {
		panic(fmt.Sprintf("ProtoJSONMarshaler.MarshalToString failed: %v", err))
	}
	return j
}

func getProtoDiff(
	expected proto.Message,
	actual proto.Message,
) (
	diff string,
) {
	expectedString := protoJSONMarshal(expected)
	actualString := protoJSONMarshal(actual)
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

func printProtoDiff(
	t *testing.T,
	handleType HandleType,
	expected proto.Message,
	actual proto.Message,
	messages ...interface{},
) {
	t.Helper()

	expectedJSON := protoJSONMarshal(expected)
	actualJSON := protoJSONMarshal(actual)
	diff := getProtoDiff(expected, actual)

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

// If not equal then handle.
func ProtoEqual(
	t *testing.T,
	handleType HandleType,
	expected proto.Message,
	actual proto.Message,
	messages ...interface{},
) {
	t.Helper()

	if !proto.Equal(expected, actual) {
		printProtoDiff(t, handleType, expected, actual, messages...)
	}
}

func protoDump(message proto.Message) string {
	str := protoJSONMarshal(message)
	if str == "" {
		str = SpewConfig.Sdump(message)
	}
	return reflect.TypeOf(message).String() + "\n" + str + "\n"
}

// If equal then handle.
func ProtoNotEqual(
	t *testing.T,
	handleType HandleType,
	expected proto.Message,
	actual proto.Message,
	messages ...interface{},
) {
	t.Helper()

	if proto.Equal(expected, actual) {
		str := protoDump(actual)
		logs := SprintMessages("Expected is equal to actual, actual is:\n"+str, messages)
		if handleType == ErrorHandle {
			t.Error(logs)
		} else {
			t.Fatal(logs)
		}
	}
}
