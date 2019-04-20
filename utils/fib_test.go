package utils

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	type Case struct {
		Input  int64
		Output *big.Int
	}

	cases := []Case{
		Case{
			Input:  0,
			Output: big.NewInt(0),
		},
		Case{
			Input:  1,
			Output: big.NewInt(1),
		},
		Case{
			Input:  2,
			Output: big.NewInt(1),
		},
		Case{
			Input:  3,
			Output: big.NewInt(2),
		},
		Case{
			Input:  4,
			Output: big.NewInt(3),
		},
		Case{
			Input:  10,
			Output: big.NewInt(55),
		},
	}
	for _, c := range cases {
		res := Fib(c.Input)
		if !reflect.DeepEqual(res, c.Output) {
			t.Fatalf("should get %d, but got %d", c.Output, res)
		}
	}
}

func TestFibNumbers(t *testing.T) {
	type Case struct {
		Input  int64
		Output []*big.Int
	}

	cases := []Case{
		Case{
			Input:  0,
			Output: []*big.Int{},
		},
		Case{
			Input:  1,
			Output: []*big.Int{big.NewInt(0)},
		},
		Case{
			Input: 2,
			Output: []*big.Int{
				big.NewInt(0),
				big.NewInt(1),
			},
		},
		Case{
			Input: 3,
			Output: []*big.Int{
				big.NewInt(0),
				big.NewInt(1),
				big.NewInt(1),
			},
		},
		Case{
			Input: 5,
			Output: []*big.Int{
				big.NewInt(0),
				big.NewInt(1),
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(3),
			},
		},
	}
	for _, c := range cases {
		res := FibNumbers(c.Input)
		if !reflect.DeepEqual(res, c.Output) {
			t.Fatalf("should get %v, but got %v", c.Output, res)
		}
	}
}
