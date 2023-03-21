package resistorcolor

// Colors should return the list of all colors.
var colorsList = []string{
	"black",
	"brown",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"violet",
	"grey",
	"white",
}

func Colors() []string {
	// panic("Please implement the Colors function")
	return colorsList
}

// ColorCode returns the resistance value of the given color.
var colorsMap = map[string]int{
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

func ColorCode(color string) int {
	// panic("Please implement the ColorCode function")
	return colorsMap[color]
}
