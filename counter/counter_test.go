package counter

import "testing"

// TestCountVowelsAndConsonants verifies vowel/consonant logic
func TestCountVowelsAndConsonants(t *testing.T) {

	tests := []struct {
		input      string
		vowels     int
		consonants int
	}{
		{"hello", 2, 3},
		{"GoLang", 2, 4},
		{"I love coding", 5, 6},
		{"123!!!", 0, 0},
		{"", 0, 0},
	}

	for _, test := range tests {
		v, c := CountVowelsAndConsonants(test.input)

		if v != test.vowels || c != test.consonants {
			t.Errorf(
				"Input: %s | Expected (v:%d, c:%d) but got (v:%d, c:%d)",
				test.input,
				test.vowels,
				test.consonants,
				v,
				c,
			)
		}
	}
}