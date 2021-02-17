// Package for dealing with sums of the elements
// of slices.
package sum

// Sum takes a slice of int and returns the sum
// of its elements.
func Sum(numbers []int) (result int) {
	for _, n := range numbers {
		result += n
	}

	return
}

// SumAll takes a slice of slices of int and returns
// a slice of int with the sum of each slice of int
// in the collection.
func SumAll(numberCollections ...[]int) (result []int) {
	for _, numbers := range numberCollections {
		result = append(result, Sum(numbers))
	}

	return
}

// SumAllTails takes a variable number of int slices and sums
// their tails, the results are returner inside of a slice of int
func SumAllTails(numberCollections ...[]int) (result []int) {
	for _, numbers := range numberCollections {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			tail := numbers[1:]
			result = append(result, Sum(tail))
		}
	}

	return
}
