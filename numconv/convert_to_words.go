package numconv

import "strconv"

var ones = [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = [...]string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

// Complete list of names for large numbers can be find here https://en.wikipedia.org/wiki/Names_of_large_numbers
var scales = map[int]string{
	0:  "",
	3:  "thousand",
	6:  "million",
	9:  "billion",
	12: "trillion",
	15: "quadrillion",
	18: "quintillion",
	21: "sextillion",
	24: "septillion",
	27: "octillion",
	30: "nonillion",
	33: "decillion",
	36: "undecillion",
	39: "duodecillion",
	42: "tredecillion",
	45: "quattuordecillion",
	48: "quindecillion",
	51: "sexdecillion",
	54: "septendecillion",
	57: "octodecillion",
	60: "novemdecillion",
	63: "vigintillion",
}

func ConvertToWords(n int) string {
	return ConvertToWordsA(strconv.Itoa(n))
}

func ConvertToWordsA(n string) string {
	if len(n) == 0 {
		return ""
	}
	return convertLarge(n, 0)
}

func convertLarge(n string, scale int) string {
	l := len(n)
	if l == 1 && scale == 0 && n[0] == '0' {
		return ones[0]
	}

	if l <= 3 {
		return convertHundredsWithScale(n, scale)
	}

	next := n[:l-3]
	current := n[l-3:]
	block := convertHundredsWithScale(current, scale)

	if block == "" {
		return convertLarge(next, scale+3)
	}

	return join(convertLarge(next, scale+3), block)
}

func convertHundredsWithScale(n string, s int) string {
	h, _ := strconv.Atoi(n)
	if h == 0 {
		return ""
	}
	return join(convertHundreds(h), scales[s])
}

func convertHundreds(n int) string {
	h, t := div(n, 100)

	if h > 0 && t > 0 {
		return ones[h] + " hundred " + convertTens(t)
	}
	if h > 0 {
		return ones[h] + " hundred"
	}
	return convertTens(t)
}

func convertTens(n int) string {
	if n < 20 {
		return ones[n]
	}

	t, o := div(n, 10)
	if o == 0 {
		return tens[t]
	}
	return tens[t] + "-" + ones[o]
}

func join(a string, b string) string {
	if len(a) > 0 && len(b) > 0 {
		return a + " " + b
	}
	if len(a) > 0 {
		return a
	}
	return b
}

func div(x int, y int) (int, int) {
	return x / y, x % y
}
