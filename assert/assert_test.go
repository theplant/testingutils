package assert_test

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"reflect"
	"runtime"
	"testing"

	"github.com/theplant/testingutils"
	"github.com/theplant/testingutils/assert"
)

type Address struct {
	City string
}

type User struct {
	Name    string
	Age     int
	Address *Address
}

var (
	user1 = User{
		Name: "name1",
		Age:  20,
		Address: &Address{
			City: "city1",
		},
	}

	user2 = User{
		Name: "name2",
		Age:  50,
		Address: &Address{
			City: "city2",
		},
	}
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func RunTest(testName string) (result string, isFatal bool) {
	command := exec.Command("go", "test", "-run", fmt.Sprintf("^%s$", testName))

	output, err := command.CombinedOutput()

	return string(output), err != nil
}

func GetFunctionName(i interface{}) string {
	filename := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return path.Ext(filename)[1:]
}

func getShortTestOutput(fullOutput string) (shortTestOutput string) {
	reader := bufio.NewReader(bytes.NewReader([]byte(fullOutput)))

	reader.ReadString('\n')
	reader.ReadString('\n')
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if line == "FAIL\n" {
			break
		}

		shortTestOutput = shortTestOutput + string(line)
	}

	return "\n" + shortTestOutput
}

func Test(t *testing.T) {
	tests := []struct {
		testFunc       func(*testing.T)
		expectedOutput string
	}{
		{
			testFunc:       TestEqual__Equal,
			expectedOutput: "",
		},

		{
			testFunc: TestEqual__NotEqualAndPrintJSON,
			expectedOutput: `
		--- Expected
		+++ Actual
		@@ -1,7 +1,7 @@
		 {
		-	"Name": "name1",
		-	"Age": 20,
		+	"Name": "name2",
		+	"Age": 50,
		 	"Address": {
		-		"City": "city1"
		+		"City": "city2"
		 	}
		 }
		
`,
		},

		{
			testFunc: TestEqual__NotEqualAndJSONIsNull,
			expectedOutput: `
		--- Expected
		+++ Actual
		@@ -1,4 +1,4 @@
		 (*errors.errorString)({
		-	s: (string) (len=1) "1"
		+	s: (string) (len=1) "2"
		 })
		 
		
`,
		},

		{
			testFunc: TestEqual__NotEqualWithMessages,
			expectedOutput: `
		--- Expected
		+++ Actual
		@@ -1 +1 @@
		-1
		+2
		
		Messages:
		assert_test.User {
			"Name": "name1",
			"Age": 20,
			"Address": {
				"City": "city1"
			}
		}
		
		*assert_test.User {
			"Name": "name1",
			"Age": 20,
			"Address": {
				"City": "city1"
			}
		}
		
		int 1
		
		int 1
		
		bool true
		
		bool false
`,
		},
	}

	for _, test := range tests {
		testName := GetFunctionName(test.testFunc)
		r, isFatal := RunTest(testName)

		if test.expectedOutput == "" {
			if isFatal {
				t.Fatal("Exptected is ok, but actual is not ok.")
			} else {
				continue
			}
		}

		diff := testingutils.PrettyJsonDiff(test.expectedOutput, getShortTestOutput(r))
		if diff != "" {
			t.Fatal("\nName: " + testName + "\n" + diff)
		}
	}
}

func TestEqual__Equal(t *testing.T) {
	assert.Equal(t, assert.FatalHandle, 1, 1)
	assert.Equal(t, assert.FatalHandle, user1, user1)
	assert.Equal(t, assert.FatalHandle, errors.New("1"), errors.New("1"))
}

func TestEqual__NotEqualAndPrintJSON(t *testing.T) {
	assert.Equal(t, assert.FatalHandle, user1, user2)
}

func TestEqual__NotEqualAndJSONIsNull(t *testing.T) {
	assert.Equal(t, assert.FatalHandle, errors.New("1"), errors.New("2"))
}

func TestEqual__NotEqualWithMessages(t *testing.T) {
	assert.Equal(t, assert.FatalHandle, 1, 2, user1, &user1, 1, 01, true, false)
}
