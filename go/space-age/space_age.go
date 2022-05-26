package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	// panic("Please implement the Age function")

	//convert value to earth year
	earthyear := seconds / 31557600

	switch planet {
	case "Earth":
		return earthyear
	case "Mercury":
		return earthyear / 0.2408467
	case "Venus":
		return earthyear / 0.61519726
	case "Mars":
		return earthyear / 1.8808158
	case "Jupiter":
		return earthyear / 11.862615
	case "Saturn":
		return earthyear / 29.447498
	case "Uranus":
		return earthyear / 84.016846
	case "Neptune":
		return earthyear / 164.79132
	}

	return 0.00
}
