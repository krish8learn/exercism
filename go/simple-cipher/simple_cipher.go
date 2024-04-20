package cipher

import (
	"regexp"
	"strings"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type shift struct {
	distance int
}

var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func NewCaesar() Cipher {
	// panic("Please implement the NewCaesar function")
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	// panic("Please implement the NewShift function")
	if distance == 0 || distance <= -26 || distance >= 26 {
		return nil
	}
	return &shift{
		distance: distance,
	}
}

func (c shift) Encode(input string) string {
	// panic("Please implement the Encode function")
	onlyAlphabetInput := strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(input, ""))
	output := ""
	for _, runeValue := range onlyAlphabetInput {
		for index, letter := range letters {
			if letter == string(runeValue) {
				if index+c.distance >= len(letters) {
					output += letters[(index+c.distance)-len(letters)]
					break
				} else if index+c.distance < 0 {
					output += letters[len(letters)+(index+c.distance)]
					break
				}
				output += letters[index+c.distance]
			}
		}
	}
	return output
}

func (c shift) Decode(input string) string {
	// panic("Please implement the Decode function")
	output := ""
	for _, runeValue := range input {
		for index, letter := range letters {
			if letter == string(runeValue) {
				if index-c.distance < 0 {
					output += letters[(index-c.distance)+len(letters)]
					break
				} else if index-c.distance >= len(letters) {
					output += letters[(index-c.distance)-len(letters)]
					break
				}
				output += letters[index-c.distance]
			}
		}
	}
	return output
}

type vigenere struct {
	key string
}

var vigenereLetterValue = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
	"i": 8,
	"j": 9,
	"k": 10,
	"l": 11,
	"m": 12,
	"n": 13,
	"o": 14,
	"p": 15,
	"q": 16,
	"r": 17,
	"s": 18,
	"t": 19,
	"u": 20,
	"v": 21,
	"w": 22,
	"x": 23,
	"y": 24,
	"z": 25,
}

func NewVigenere(key string) Cipher {
	// panic("Please implement the NewVigenere function")
	if len(key) == 0 || len(key) < 4 {
		return nil
	}

	if satisfied, err := regexp.MatchString(`[^a-z ]+`, key); satisfied || err != nil {
		return nil
	}

	if strings.Contains(key, " ") {
		return nil
	}
	return vigenere{
		key: key,
	}
}

func (v vigenere) Encode(input string) string {
	// panic("Please implement the Encode function")
	input = strings.ToLower(regexp.MustCompile(`[^a-zA-Z ]+`).ReplaceAllString(input, ""))
	inputSlice := strings.Split(strings.ReplaceAll(input, " ", ""), "")
	keySlice := strings.Split(v.key, "")
	value := 0
	output := ""

	for inputSliceIndex, keySliceIndex := 0, 0; inputSliceIndex < len(inputSlice); inputSliceIndex++ {

		value = vigenereLetterValue[inputSlice[inputSliceIndex]] + vigenereLetterValue[keySlice[keySliceIndex]]
		output += letters[value%26]

		if keySliceIndex == len(keySlice)-1 {
			keySliceIndex = -1
		}
		keySliceIndex++
	}

	return output

}

func (v vigenere) Decode(input string) string {
	// panic("Please implement the Decode function")
	input = strings.ToLower(regexp.MustCompile(`[^a-zA-Z ]+`).ReplaceAllString(input, ""))
	inputSlice := strings.Split(strings.ReplaceAll(input, " ", ""), "")
	keySlice := strings.Split(v.key, "")
	value := 0
	output := ""

	for inputSliceIndex, keySliceIndex := 0, 0; inputSliceIndex < len(inputSlice); inputSliceIndex++ {

		value = vigenereLetterValue[inputSlice[inputSliceIndex]] - vigenereLetterValue[keySlice[keySliceIndex]]
		if value < 0 {
			output += letters[len(letters)+value]
			goto NEXT_IT
		}
		output += letters[value%26]

	NEXT_IT:
		if keySliceIndex == len(keySlice)-1 {
			keySliceIndex = -1
		}
		keySliceIndex++
	}

	return output

}
