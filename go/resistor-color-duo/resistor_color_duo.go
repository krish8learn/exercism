package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	// panic("Implement the Value function")
	resistorValue := 0

	for index, value := range colors {
		if index == 1 {
			resistorValue += resistorColor[value]
			break
		} else if index == 2 {
			break
		}
		resistorValue += resistorColor[value] * 10
	}

	return resistorValue
}

var resistorColor = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}
