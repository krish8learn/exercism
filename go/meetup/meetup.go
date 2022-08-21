package meetup

import "time"

// Define the WeekSchedule type here.

type WeekSchedule int

const (
	First = WeekSchedule(iota)
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	// panic("Please implement the Day function")
	//get a date object
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.FixedZone("UTC", 0))

	match := []int{}

	//compare month
	for date.Month() == month {
		//compare weekday
		if date.Weekday() == weekday {
			//add that week day value into slice
			match = append(match, date.Day())
		}
		//add that week day value into date object
		date = date.AddDate(0, 0, 1)
	}

	switch week {
	case First:
		return match[0]
	case Second:
		return match[1]
	case Third:
		return match[2]
	case Fourth:
		return match[3]
	case Last:
		return match[len(match)-1]
	case Teenth:
		for _, d := range match {
			if d >= 13 {
				return d
			}
		}
	}

	panic("Invalid schedule")
}
