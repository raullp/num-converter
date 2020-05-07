package numconv

var ones = [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = [...]string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

// Complete list of names for large numbers could be find here https://en.wikipedia.org/wiki/Names_of_large_numbers
var namesForLargeNumbers = map[int]string{
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
	if n < 1000 {
		return convertHundreds(n)
	}
	return convertLargeNumber(n)
}

func convertTens(n int) string {
	if n < 20 {
		return ones[n]
	}
	t := tens[n/10]
	o := n % 10
	if o == 0 {
		return t
	}
	return t + "-" + ones[o]
}

func convertHundreds(n int) string {
	if n == 0 {
		return ""
	}

	h := n / 100
	t := n % 100

	if h > 0 && t > 0 {
		return ones[h] + " hundred " + convertTens(t)
	}
	if h > 0 {
		return ones[h] + " hundred"
	}
	return convertTens(t)
}

func convertLargeNumber(n int) string {
	scale := 3
	w := convertHundreds(n % 1000)
	n /= 1000

	for ok := true; ok; ok = n > 0 {
		h := n % 1000
		if h > 0 {
			s := convertHundreds(n%1000) + " " + namesForLargeNumbers[scale]
			if len(w) > 0 {
				w = s + " " + w	
			} else {
				w = s
			}
		}
		n /= 1000
		scale += 3
	}

	return w
}
