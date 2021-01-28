package module01

import (
	"math"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	var runTotal float64
	power := 0
	for i := len(value) - 1; i >= 0; i-- {
		index := findIndex(value[i])

		runTotal += float64(index) * (math.Pow(float64(base), float64(power)))
		power += 1
	}
	return int(runTotal)
}

func findIndex(char byte) int {
	const charSet = "0123456789ABCDEF"
	for i, val := range charSet {
		if char == byte(val) {
			return i
		}
	}
	return -1
}
