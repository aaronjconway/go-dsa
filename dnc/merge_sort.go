package dnc

// MergeSort solves the problem in O(n*Log n) time and O(n) space.
// take an array of ints
func MergeSort(list []int) []int {
	// initial checking
	if len(list) <= 1 {
		return list
	}

	//what is a left and right?
	left, right := divide(list)
	//do left and right then merge them?
	return merge(left, right)
}

// takes an array of ints and returns two arrays of ints
func divide(list []int) ([]int, []int) {
	//mid point - remember that go lang doesn't do implicit type conversion so
	//the mid point will round down.
	mid := len(list) / 2
	left := MergeSort(list[:mid])
	right := MergeSort(list[mid:])
	return left, right
}

// takes two arrays of ints and returns one array of ints.
func merge(left, right []int) []int {

	output := []int{}

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			output = append(output, left[i])
			i++
		} else {
			output = append(output, right[j])
			j++
		}
	}

	if i >= len(left) {
		output = append(output, right[j:]...)
	} else {
		output = append(output, left[i:]...)
	}
	return output
}
