package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	// panic("Please implement the New function")
	//total minutes
	minutes := hour*60 + minute

	//hour
	hours := (minutes / 60) % 24
	minutes = minutes % 60

	// Deal with negatives now
	if minutes < 0 {
		hours -= 1    // Roll back one hour
		minutes += 60 // Make mins positive
	}
	if hours < 0 {
		hours += 24 // Normalize to 24 hours
	}

	return Clock{hours, minutes}
}

func (c Clock) Add(m int) Clock {
	// panic("Please implement the Add function")
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
