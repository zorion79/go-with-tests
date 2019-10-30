package sum

// Sum returns the sum of integers
func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}

// All returns new array contains sum of arrays
func All(arr ...[]int) []int {
	var result []int

	for _, a := range arr {
		result = append(result, Sum(a))
	}
	return result
}

// AllTails returns new array contains sum of tail elements
func AllTails(arr ...[]int) []int {
	var result []int
	for _, a := range arr {
		if len(a) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(a[1:]))
		}
	}
	return result
}
