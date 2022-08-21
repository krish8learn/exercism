package allergies

var allergicItemsScore = map[uint]string{
	1: "eggs",
	2: "peanuts",
	4: "shellfish",
	8: "strawberries",
	16: "tomatoes",
	32: "chocolate",
	64: "pollen",
	128: "cats",
}

func Allergies(allergies uint) []string {
	// panic("Please implement the Allergies function")
	var allergicItems []string

	for key, value := range allergicItemsScore{
		if key&allergies == key{
			allergicItems = append(allergicItems, value)
		}
	}
	return allergicItems
}

func AllergicTo(allergies uint, allergen string) bool {
	// panic("Please implement the AllergicTo function")

	allergicItems := Allergies(allergies)
	for _, item := range allergicItems {
		if item == allergen {
			return true
		}
	}
	return false
}
