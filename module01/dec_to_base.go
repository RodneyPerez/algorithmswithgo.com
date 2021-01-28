package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	values := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	if dec == 0 {
		return ""
	}
	place := dec % base
	answer := values[place]
	dec = dec / base
	//	var answer string
	//	for {
	//		if dec == 0 {
	//			break
	//		}
	//		place := dec % base
	//		answer = values[place] + answer
	//		dec = dec / base
	//	}
	return DecToBase(dec, base) + answer
}
