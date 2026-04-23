package bafflingbirthdays

import (
	"math/rand/v2"
	"time"
)

func SharedBirthday(dates []time.Time) bool {
	seen := make(map[string]string)

	for _, date := range dates {
		key := date.Format("01-02")
		if _, exists := seen[key]; exists {
			return true
		}
		seen[key] = ""
	}
	return false
}

func RandomBirthdates(size int) []time.Time {
	dates := make([]time.Time, size)

	// We pick a non-leap year, e.g., 2001
	year := 2001

	for i := 0; i < size; i++ {
		// Pick a random month (1-12)
		month := time.Month(rand.IntN(12) + 1)

		// Determine the max days in that month
		// time.Date with day 0 gives the last day of the previous month.
		// So, passing (year, month+1, 0) gives the last day of 'month'.
		lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()

		// Pick a random day (1 to lastDay)
		day := rand.IntN(lastDay) + 1

		dates[i] = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	}

	return dates
}

func EstimatedProbability(size int) float64 {
	if size <= 1 {
		return 0
	}
	if size > 365 {
		return 1
	}

	// Probability that all birthdays are different
	prob := 1.0
	for i := 0; i < size; i++ {
		prob *= float64(365-i) / 365.0
	}
	// Return probability that at least 2 people share a birthday
	return (1 - prob) * 100
}
