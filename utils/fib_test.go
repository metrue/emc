package utils

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	type Case struct {
		Input  int
		Output int
	}

	cases := []Case{
		Case{
			Input:  0,
			Output: 0,
		},
		Case{
			Input:  1,
			Output: 1,
		},
		Case{
			Input:  2,
			Output: 1,
		},
		Case{
			Input:  3,
			Output: 2,
		},
		Case{
			Input:  4,
			Output: 3,
		},
		Case{
			Input:  100000,
			Output: 2754320626097736315,
		},
	}
	for _, c := range cases {
		res := Fib(c.Input)
		if res != c.Output {
			t.Fatalf("should get %d, but got %d", c.Output, res)
		}
	}
}

func TestFibNumbers(t *testing.T) {
	type Case struct {
		Input  int
		Output []int
	}

	cases := []Case{
		Case{
			Input:  0,
			Output: []int{},
		},
		Case{
			Input:  1,
			Output: []int{0},
		},
		Case{
			Input:  2,
			Output: []int{0, 1},
		},
		Case{
			Input:  3,
			Output: []int{0, 1, 1},
		},
		Case{
			Input:  5,
			Output: []int{0, 1, 1, 2, 3},
		},
	}
	for _, c := range cases {
		res := FibNumbers(c.Input)
		if !reflect.DeepEqual(res, c.Output) {
			t.Fatalf("should get %v, but got %v", c.Output, res)
		}
	}
}
