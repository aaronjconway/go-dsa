// touched
package array

import "math"

// AddTwoNumbers solves the problem in O(n) time and O(1) space.

// Takes two int arrays and outputs an array which is the sum.
// better explanation in the test file
// [1,2]  + [2] =[1,4]

func AddTwoNumbers(num1, num2 []int) []int {

	// equalizeLengths makes arrays same length by appending 0's
	// I believe the point of this is so that we can move things in place
	// instead of appending to the list.
	num1, num2 = equalizeLengths(num1, num2)

	carry := false

	// start i at the last index available
	// while i is greater than -1 aka when it is >= 0
	for i := len(num1) - 1; i > -1; i-- {

		//take the right most number and add the right most number of the other
		//array
		num1[i] += num2[i]

		// if the last addition was greater than 10 we will have to carry the 10
		// over to the next digit in the form of incrementing by one because
		// it's already int the one's place.

		if carry {
			//add one
			num1[i]++
			//we're done witht this so set back to false
			carry = false
		}

		if num1[i] >= 10 {
			//make sure to remove 10 from the num1[i] so 18 = 8
			num1[i] -= 10
			//set carry to true so that we add that 1 to the next value
			carry = true

		}
	}

	// if we do the addition and carry is still true then we've exceeded the
	// lengths of both of the arrays and need to add a one. [5] + [5] = [1,0]
	if carry {
		num1 = append([]int{1}, num1...)
	}

	return num1
}

// takes two nums int arrays and outputs two arrays
// the point of this is to make the two arrays the same length by appending
// zero's to the array that is shorter.
func equalizeLengths(num1, num2 []int) ([]int, []int) {

	//the len of the diff between the two arrays
	diff := int(math.Abs(float64(len(num2) - len(num1))))

	// init an empty array
	zeros := []int{}

	//create an array of 0's
	// if diff is 3 then i will do 0,1,2 getting all three digits.
	// this is why i is 0 based. then it's easy to loop through on i< len of
	// something
	for i := 0; i < diff; i++ {
		zeros = append(zeros, 0)
	}
	//  have to append to the left of the array
	// like we would in addition.
	// example: 1000 + 1
	//[1,0,0,0] + [1] = [1,0,0,0] + [0,0,0,1]

	if len(num2) > len(num1) {
		num1 = append(zeros, num1...)

	} else if len(num1) > len(num2) {
		num2 = append(zeros, num2...)
	}
	return num1, num2
}
