package array

import (
	"reflect"
	"testing"
)

/*
TestAddTwoNumbers tests solution(s) with the following signature and problem description:

	AddTwoNumbers(num1, num2 []int) []int

Given two numbers as an array like [2,9] and [9,9,9] return the sum of the numbers they
represent like [1,0,2,8], because 29+999=1028.
*/

// * is dereferencing the pointer so *testing is getting the value of testing.
func TestAddTwoNumbers(t *testing.T) {
	// create an array struct
	//num1, num2, sum are all arrays of ints
	tests := []struct{ num1, num2, sum []int }{
		{[]int{1}, []int{}, []int{1}},
		{[]int{1}, []int{0}, []int{1}},
		{[]int{1}, []int{1}, []int{2}},
		{[]int{1}, []int{9}, []int{1, 0}},
		{[]int{2, 5}, []int{3, 5}, []int{6, 0}},
		{[]int{2, 9}, []int{9, 9, 9}, []int{1, 0, 2, 8}},
		{[]int{9, 9, 9}, []int{9, 9, 9}, []int{1, 9, 9, 8}},
	}

	// for index, test in range tests
	for i, test := range tests {

		// if thing := do something; !value {error bc I didn't get the value}
		// general strucutre is if var := function; test return of var { do
		// something }
		//
		//
		// the if var := function actuall does the thing. then we put a semi
		// colon ; with some logic on the var.
		// based on the return}
		// instead of if thing == value {}
		if got := AddTwoNumbers(test.num1, test.num2); !reflect.DeepEqual(got, test.sum) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sum, got)
		}
	}
}
