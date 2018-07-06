package assert

import (
	"testing"
)

type User struct {
	Name string
}

func TestIsUnorderedListEqualedWithTypeCheck(t *testing.T) {
	tests := []struct {
		name              string
		list1             interface{}
		list2             interface{}
		expectedIsEqualed bool
	}{
		{
			name:              "1",
			list1:             nil,
			list2:             nil,
			expectedIsEqualed: true,
		},

		{
			name:              "2",
			list1:             nil,
			list2:             []string{},
			expectedIsEqualed: false,
		},

		{
			name:              "3",
			list1:             []int{},
			list2:             []string{},
			expectedIsEqualed: false,
		},

		{
			name:              "4",
			list1:             [1]string{},
			list2:             []string{},
			expectedIsEqualed: false,
		},

		{
			name:              "5",
			list1:             []User{},
			list2:             []string{},
			expectedIsEqualed: false,
		},

		{
			name:              "6",
			list1:             []int{1, 2},
			list2:             []int{1, 1, 2},
			expectedIsEqualed: false,
		},

		{
			name:              "7",
			list1:             []interface{}{1, 2},
			list2:             []int{1, 2},
			expectedIsEqualed: false,
		},

		{
			name:              "8",
			list1:             [3]int{1, 2, 3},
			list2:             [3]int{1, 2, 3},
			expectedIsEqualed: true,
		},

		{
			name:              "9",
			list1:             []int{1, 2, 2, 3, 3, 3},
			list2:             []int{3, 3, 3, 2, 2, 1},
			expectedIsEqualed: true,
		},

		{
			name:              "10",
			list1:             []int{},
			list2:             []int{},
			expectedIsEqualed: true,
		},

		{
			name:              "11",
			list1:             [1]int{1},
			list2:             []int{1},
			expectedIsEqualed: false,
		},

		{
			name:              "12",
			list1:             []User{{Name: "A"}, {Name: "B"}},
			list2:             []User{{Name: "B"}, {Name: "A"}},
			expectedIsEqualed: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotIsEqualed := isUnorderedListEqualedWithTypeCheck(test.list1, test.list2)
			if gotIsEqualed != test.expectedIsEqualed {
				t.Errorf("Exptected is not equal to actual\nExptected: %v\nActual: %v", test.expectedIsEqualed, gotIsEqualed)
			}
		})
	}
}
