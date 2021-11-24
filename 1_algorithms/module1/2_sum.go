package module01

// Sum will sum up all of the numbers passed
// in and return the result
// First Solution
// func Sum(numbers []int) int {
// 	res := 0
// 	for _, i := range numbers {
// 		res += i
// 	}

// 	return res
// }

// Second Solution - Recursion
func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + Sum(numbers[1:])
}
