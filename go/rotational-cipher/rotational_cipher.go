package rotationalcipher

const charsSmall = "abcdefghijklmnopqrstuvwxyz"
const charsCapital = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RotationalCipher(plain string, shiftKey int) string {
	// panic("Please implement the RotationalCipher function")
	var output string

	for _, value := range plain {

		if value >= 65 && value <= 90 {
			//capital letter
			for index, char := range charsCapital {
				if value == char {
					if index+shiftKey <= 25 {
						output += string(charsCapital[index+shiftKey])
					} else if index+shiftKey > 25 {
						output += string(charsCapital[(index+shiftKey)-26])
					}
				}
			}

		} else if value >= 97 && value <= 122 {
			//small letter
			for index, char := range charsSmall {
				if value == char {
					if index+shiftKey <= 25 {
						output += string(charsSmall[index+shiftKey])
					} else if index+shiftKey > 25 {
						output += string(charsSmall[(index+shiftKey)-26])
					}
				}
			}

		} else {
			//for other special chars
			output += string(value)
		}

	}

	return output
}

/*
alternate solution
	var output []rune
	for _, r := range input {
		base := 0
		switch {
		case r >= 'A' && r <= 'Z':
			base = int('A')
		case r >= 'a' && r <= 'z':
			base = int('a')
		}
		if base != 0 {
			//this formula is key
			r = rune(((int(r) - base + rot) % 26) + base)
		}
		output = append(output, r)
	}
	return string(output)
*/
