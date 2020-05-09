package numconv

import "testing"

func TestConvertToWordsForOnes(t *testing.T) {
	tables := []struct {
		n string
		w string
	}{
		{"0", "zero"},
		{"1", "one"},
		{"2", "two"},
		{"3", "three"},
		{"4", "four"},
		{"10", "ten"},
		{"15", "fifteen"},
		{"19", "nineteen"},
	}

	for _, table := range tables {
		w := ConvertToWordsA(table.n)
		if w != table.w {
			t.Errorf("convertToWords (%s) was incorrect, got: %s, want: %s.", table.n, w, table.w)
		}
	}
}

func TestConvertToWordsForTens(t *testing.T) {
	tables := []struct {
		n string
		w string
	}{
		{"20", "twenty"},
		{"22", "twenty-two"},
		{"35", "thirty-five"},
		{"45", "forty-five"},
		{"60", "sixty"},
		{"68", "sixty-eight"},
		{"99", "ninety-nine"},
	}

	for _, table := range tables {
		w := ConvertToWordsA(table.n)
		if w != table.w {
			t.Errorf("ConvertToWords (%s) was incorrect, got: %s, want: %s.", table.n, w, table.w)
		}
	}
}

func TestConvertToWordsForHundreds(t *testing.T) {
	tables := []struct {
		n string
		w string
	}{
		{"100", "one hundred"},
		{"101", "one hundred one"},
		{"115", "one hundred fifteen"},
		{"200", "two hundred"},
		{"850", "eight hundred fifty"},
		{"999", "nine hundred ninety-nine"},
	}

	for _, table := range tables {
		w := ConvertToWordsA(table.n)
		if w != table.w {
			t.Errorf("ConvertToWords (%s) was incorrect, got: %s, want: %s.", table.n, w, table.w)
		}
	}
}

func TestConvertToWordsForLargeNumbers(t *testing.T) {
	tables := []struct {
		n string
		w string
	}{
		{"1000", "one thousand"},
		{"1001", "one thousand one"},
		{"20115", "twenty thousand one hundred fifteen"},
		{"1000000", "one million"},
		{"1000010", "one million ten"},
		{"1540100991245", "one trillion five hundred forty billion one hundred million nine hundred ninety-one thousand two hundred forty-five"},
		{"1540100991245154015", "one quintillion five hundred forty quadrillion one hundred trillion nine hundred ninety-one billion two hundred forty-five million one hundred fifty-four thousand fifteen"},
	}

	for _, table := range tables {
		w := ConvertToWordsA(table.n)
		if w != table.w {
			t.Errorf("ConvertToWords (%s) was incorrect, got: %s, want: %s.", table.n, w, table.w)
		}
	}
}
