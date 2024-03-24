package array

import "math"

// AddTwoNumbers solves the problem in O(n) time and O(1) space.
// a func atha ttkaes in a number and an array of ints
// outputs an array of ints

func AddTwoNumbers(num1, num2 []int) []int {

	// make arrays same length by appending 0's
	// I believe the point of this is so that we can move things in place
	// instead of appending to the list.
	num1, num2 = equalizeLengths(num1, num2)

	carry := false
	// start i at the last index
	// while I is less than one
	// while I is greater than -1 aka when it is >= 0

	for i := len(num1) - 1; i > -1; i-- {
		//increment by i
		num1[i] += num2[i]

		if carry {
			//inc the num1
			num1[i]++
			carry = false
		}

		if num1[i] >= 10 {
			num1[i] -= 10
			carry = true
		}
	}

	if carry {
		num1 = append([]int{1}, num1...)
	}
	return num1
}

// takes two nums int arrays and outputs two arrays
// the point of this is to make the two arrays the same length by appending
// zero's to the array that is shorter.
func equalizeLengths(num1, num2 []int) ([]int, []int) {

	// get the diff in length of the two arrays

	diff := int(math.Abs(float64(len(num2) - len(num1))))

	// inti an empty array
	zeros := []int{}

	//create an array of 0's
	for i := 0; i < diff; i++ {
		zeros = append(zeros, 0)
	}

	if len(num2) > len(num1) {
		num1 = append(zeros, num1...)
	} else if len(num1) > len(num2) {
		num2 = append(zeros, num2...)
	}
	return num1, num2
}
