package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var reversedWord string
	for _, r := range word {
		reversedWord = string(r) + reversedWord
	}
	return reversedWord
}
