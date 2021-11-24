package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
// First Solution - Faster
// func Reverse(word string) string {
// 	var res strings.Builder
// 	for i := len(word) - 1; i >= 0; i-- {
// 		res.WriteByte(word[i])
// 	}

// 	return res.String()
// }

// Second Solution
func Reverse(word string) string {
	var res string
	for _, s := range word {
		res = string(s) + res
	}

	return res
}
