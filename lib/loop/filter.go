package loop

// Filter applies a callback function to each element of an input array and returns a new array with the elements that satisfy the callback condition.
//
// Parameters:
// - array: A slice of elements of type I.
// - callback: A function that takes an element of type I and returns a boolean. If the callback returns true, the element is included in the result.
//
// Returns:
// - A slice of elements of type I, where each element satisfies the condition specified by the callback function.
//
// Example:
// input := []int{1, 2, 3, 4, 5}
//
//	result := Filter(input, func(value int) bool {
//	    return value%2 == 0
//	})
//
// // result is now []int{2, 4}
func Filter[I any](array []I, callback func(value I) bool) []I {
	result := []I{}

	for _, value := range array {
		if callback(value) {
			result = append(result, value)
		}
	}

	return result
}
