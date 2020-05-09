package numconv

var ones = [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = [...]string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

// Complete list of names for large numbers can be find here https://en.wikipedia.org/wiki/Names_of_large_numbers
var scalesForLargeNumbers = map[int]string{
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
	if n < 100 {
		return convertTens(n)
	}
	return convertLargeNumber(n)
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

func convertHundreds(n int) string {
	if n == 0 {
		return ""
	}

	h, t := div(n, 100)

	if h > 0 && t > 0 {
		return ones[h] + " hundred " + convertTens(t)
	}
	if h > 0 {
		return ones[h] + " hundred"
	}
	return convertTens(t)
}

func convertLargeNumber(n int) string {
	s := 3
	words := convertHundreds(n % 1000)
	n /= 1000

	for ok := true; ok; ok = n > 0 {
		block := n % 1000
		if block > 0 {
			scale := convertHundreds(block) + " " + scalesForLargeNumbers[s]
			words = join(scale, words)
		}
		n /= 1000
		s += 3
	}

	return words
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
