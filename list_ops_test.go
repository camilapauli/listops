package listops

import (
	"fmt"
	"reflect"
	"testing"
)

var foldTestCases = []struct {
	name     string
	property string
	fn       binFunc
	initial  int
	list     IntList
	want     int
}{
	{
		name:     "empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x * y },
		initial:  2,
		want:     2,
		list:     []int{},
	},
	{
		name:     "direction independent function applied to non-empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x + y },
		initial:  5,
		want:     15,
		list:     []int{1, 2, 3, 4},
	},
	{
		name:     "direction independent function applied to non-empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x - y },
		initial:  0,
		want:     -10,
		list:     []int{1, 2, 3, 4},
	},
	{
		name:     "direction dependent function applied to non-empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x / y },
		initial:  5,
		want:     0,
		list:     []int{2, 5},
	},
	{
		name:     "empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x * y },
		initial:  2,
		want:     2,
		list:     []int{},
	},
	{
		name:     "direction independent function applied to non-empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x + y },
		initial:  5,
		want:     15,
		list:     []int{1, 2, 3, 4},
	},
	{
		name:     "direction independent function applied to non-empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x - y },
		initial:  100,
		want:     -94,
		list:     []int{1, 3, 8},
	},
	{
		name:     "direction dependent function applied to non-empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x / y },
		initial:  5,
		want:     2,
		list:     []int{2, 5},
	},
}

func TestFold(t *testing.T) {
	var got int
	for _, tt := range foldTestCases {
		if tt.property == "foldr" {
			got = tt.list.Foldr(tt.fn, tt.initial)
		} else {
			got = tt.list.Foldl(tt.fn, tt.initial)
		}
		if got != tt.want {
			t.Fatalf("FAIL: %s: %q -- expected: %d, actual: %d", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}

	}

}

func ExampleIntList_Foldl() {
	testCase := foldTestCases[2]
	list := testCase.list
	fn := testCase.fn
	initial := testCase.initial
	fmt.Println(list.Foldl(fn, initial))
	// Output: -10
}

func ExampleIntList_Foldr() {
	testCase := foldTestCases[6]
	list := testCase.list
	fn := testCase.fn
	initial := testCase.initial
	fmt.Println(list.Foldr(fn, initial))
	// Output: -94
}

var filterTestCases = []struct {
	name     string
	property string
	fn       predFunc
	list     []int
	want     []int
}{
	{
		name:     "empty list",
		property: "filter",
		fn:       func(n int) bool { return n%2 == 1 },
		list:     []int{},
		want:     []int{},
	},
	{
		name:     "non-empty list",
		property: "filter",
		fn:       func(n int) bool { return n%2 == 1 },
		list:     []int{1, 2, 3, 4, 5},
		want:     []int{1, 3, 5},
	},
}

func TestFilterMethod(t *testing.T) {
	for _, tt := range filterTestCases {
		in := IntList(tt.list)
		got := in.Filter(tt.fn)
		if !reflect.DeepEqual(IntList(tt.want), got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}

	}
}

func ExampleIntList_Filter() {
	testCase := filterTestCases[1]
	list := IntList(testCase.list)
	fn := testCase.fn
	fmt.Println(list.Filter(fn))
	// Output: [1 3 5]
}

var lengthTestCases = []struct {
	name     string
	property string
	list     IntList
	want     int
}{
	{
		name:     "empty list",
		property: "length",
		list:     []int{},
		want:     0,
	},
	{
		name:     "non-empty list",
		property: "length",
		list:     []int{1, 2, 3, 4},
		want:     4,
	},
}

func TestLengthMethod(t *testing.T) {
	for _, tt := range lengthTestCases {
		got := tt.list.Length()
		if tt.want != got {
			t.Fatalf("FAIL: %s: %q -- expected: %d, actual: %d", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}

	}
}

func ExampleIntList_Length() {
	testCase := lengthTestCases[1]
	list := testCase.list
	fmt.Println(list.Length())
	// Output: 4
}

var mapTestCases = []struct {
	name     string
	property string
	list     IntList
	fn       unaryFunc
	want     IntList
}{
	{
		name:     "empty list",
		property: "map",
		list:     []int{},
		fn:       func(x int) int { return x + 1 },
		want:     []int{},
	},
	{
		name:     "non-empty list",
		property: "map",
		list:     []int{1, 3, 5, 7},
		fn:       func(x int) int { return x + 1 },
		want:     []int{2, 4, 6, 8},
	},
}

func TestMapMethod(t *testing.T) {
	for _, tt := range mapTestCases {
		got := tt.list.Map(tt.fn)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}

	}
}

func ExampleIntList_Map() {
	testCase := mapTestCases[1]
	list := IntList(testCase.list)
	fn := testCase.fn
	fmt.Println(list.Map(fn))
	// Output: [2 4 6 8]
}

var reverseTestCases = []struct {
	name     string
	property string
	list     IntList
	want     IntList
}{
	{
		name:     "empty list",
		property: "reverse",
		list:     []int{},
		want:     []int{},
	},
	{
		name:     "non-empty list",
		property: "reverse",
		list:     []int{1, 3, 5, 7},
		want:     []int{7, 5, 3, 1},
	},
}

func TestReverseMethod(t *testing.T) {
	for _, tt := range reverseTestCases {
		got := tt.list.Reverse()
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}

	}
}

func ExampleIntList_Reverse() {
	testCase := reverseTestCases[1]
	list := IntList(testCase.list)
	fmt.Println(list.Reverse())
	// Output: [7 5 3 1]
}

var appendTestCases = []struct {
	name       string
	property   string
	list       IntList
	appendThis IntList
	want       IntList
}{
	{
		name:       "empty list",
		property:   "append",
		list:       []int{},
		appendThis: []int{},
		want:       []int{},
	},
	{
		name:       "empty list to list",
		property:   "append",
		list:       []int{},
		appendThis: []int{1, 2, 3, 4},
		want:       []int{1, 2, 3, 4},
	},
	{
		name:       "non-empty lists",
		property:   "append",
		list:       []int{1, 2},
		appendThis: []int{2, 3, 4, 5},
		want:       []int{1, 2, 2, 3, 4, 5},
	},
}

func TestAppendMethod(t *testing.T) {
	for _, tt := range appendTestCases {
		got := tt.list.Append(tt.appendThis)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}

	}
}

func ExampleIntList_Append() {
	testCase := appendTestCases[2]
	list := IntList(testCase.list)
	appendList := testCase.appendThis
	fmt.Println(list.Append(appendList))
	// Output: [1 2 2 3 4 5]
}

var concatTestCases = []struct {
	name     string
	property string
	list     IntList
	args     []IntList
	want     IntList
}{
	{
		name:     "empty list",
		property: "concat",
		list:     []int{},
		args:     []IntList{},
		want:     []int{},
	},
	{
		name:     "list of lists",
		property: "concat",
		list:     []int{1, 2},
		args:     []IntList{[]int{3}, []int{}, []int{4, 5, 6}},
		want:     []int{1, 2, 3, 4, 5, 6},
	},
}

func TestConcatMethod(t *testing.T) {
	for _, tt := range concatTestCases {
		got := tt.list.Concat(tt.args)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}

func ExampleIntList_Concat() {
	testCase := concatTestCases[1]
	list := IntList(testCase.list)
	args := testCase.args
	fmt.Println(list.Concat(args))
	// Output: [1 2 3 4 5 6]
}
