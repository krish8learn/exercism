package resistorcolortrio

import (
	"fmt"
	"math"
	"strconv"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {

	// value determination
	resistorValue := 0
	for index, color := range colors {
		if index > 2 {
			break
		} else if index == 0 {
			resistorValue += resistorColor[color]
		} else if index == 2 {
			resistorValue = resistorValue * int(math.Pow10(resistorColor[color]))
		} else {
			resistorValue = (resistorValue * 10) + resistorColor[color]
		}
	}

	// string conversion
	resistorValueString := ""
	if resistorValue%int(math.Pow10(9)) == 0 && resistorValue != 0 {
		resistorValueString = strconv.Itoa(resistorValue/int(math.Pow10(9))) + " gigaohms"
		fmt.Println(resistorValueString)
	} else if resistorValue%int(math.Pow10(6)) == 0 && resistorValue != 0 {
		resistorValueString = strconv.Itoa(resistorValue/int(math.Pow10(6))) + " megaohms"
		fmt.Println(resistorValueString)
	} else if resistorValue%int(math.Pow10(3)) == 0 && resistorValue != 0 {
		resistorValueString = strconv.Itoa(resistorValue/int(math.Pow10(3))) + " kiloohms"
		fmt.Println(resistorValueString)
	} else {
		resistorValueString = strconv.Itoa(resistorValue) + " ohms"
		fmt.Println(resistorValueString)
	}

	return resistorValueString
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
