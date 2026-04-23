package lineup

import "fmt"

func Format(name string, number int) string {

	suffix := ""
	switch number % 10 {
	case 1:
		suffix = "st"
	case 2:
		suffix = "nd"
	case 3:
		suffix = "rd"
	default:
		suffix = "th"
	}
	if number%100 >= 11 && number%100 <= 13 {
		suffix = "th"
	}

	str := fmt.Sprintf("%s, you are the %d%s customer we serve today. Thank you!", name, number, suffix)

	return str
}
