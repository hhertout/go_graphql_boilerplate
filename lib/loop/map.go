package loop

// Map applies a callback function to each element of an input array and returns a new array with the results.
//
// Parameters:
// - array: A slice of elements of type I.
// - callback: A function that takes an element of type I and returns a value of type R.
//
// Returns:
// - A slice of elements of type R, where each element is the result of applying the callback function to the corresponding element of the input array.
//
// Example:
// input := []int{1, 2, 3}
//
//	result := Map(input, func(value int) string {
//	    return strconv.Itoa(value)
//	})
//
// // result is now []string{"1", "2", "3"}
func Map[I any, R any](array []I, callback func(value I) R) []R {
	result := []R{}

	for _, value := range array {
		result = append(result, callback(value))
	}

	return result
}
