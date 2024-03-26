// touched
package array

// BubbleSort solves the problem in O(n^2) time and O(1) space.
// takes an array of int's and returns the ints sorted
// modifies the original array
func BubbleSort(input []int) {
	//loop over the input
	for i := range input {
		//loop over the input again this is why it's n^2. we have to move each
		//value to the end of the array if it's larger.
		//best case n^2/2
		for j := range input {
			if input[i] < input[j] {
				//swap if so
				input[i], input[j] = input[j], input[i]
			}
		}
	}
}
