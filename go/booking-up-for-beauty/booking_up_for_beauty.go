package booking
import (
	"fmt"
	"time"
)

func Schedule(date string) time.Time {
	// panic("Please implement the Schedule function")
	// Mon Jan 2 15:04:05 -0700 MST 2006
	form := "1/02/2006 15:04:05"
	result, _ := time.Parse(form, date)

	return result
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	// panic("Please implement the HasPassed function")
	// Mon Jan 2 15:04:05 -0700 MST 2006
	form := "January 2, 2006 15:04:05"
	result, _ := time.Parse(form, date)

	return result.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	// panic("Please implement the IsAfternoonAppointment function")
	// Mon Jan 2 15:04:05 -0700 MST 2006
	form := "Monday, January 2, 2006 15:04:05"
	result, _ := time.Parse(form, date)

	inputHour := result.Hour()
	if inputHour <= 18 && inputHour >= 12 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	// panic("Please implement the Description function")
	// Mon Jan 2 15:04:05 -0700 MST 2006
	// Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."
	form := "1/2/2006 15:04:05"
	result, _:= time.Parse(form, date)
	layout := "Monday, January 2, 2006, at 15:04"
	res := result.Format(layout)
	return fmt.Sprintf("You have an appointment on %v.", res)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	// panic("Please implement the AnniversaryDate function")
	// Mon Jan 2 15:04:05 -0700 MST 2006
	year := time.Now().Year()
	layout := "2006-01-02"
	str := fmt.Sprintf("%v-09-15", year)
	result, _ := time.Parse(layout, str)
	// log.Println("-->", result)
	return result
}

