package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	// panic("Please implement the Convert function")
	sound := ""
	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		sound = strconv.Itoa(number)
		return sound
	}

	for number > 2 {
		if number%3 == 0 {
			if strings.Contains(sound, "Pling"){
				break
			}
			sound += "Pling"
			number = number / 3
			
		} else if number%5 == 0 {
			if strings.Contains(sound, "Plang"){
				break
			}
			sound += "Plang"
			number = number / 5
			
		} else if number%7 == 0 {
			if strings.Contains(sound, "Plong"){
				break
			}
			sound += "Plong"
			number = number / 7
			
		}
	}

	return sound
}
