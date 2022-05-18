package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(r Record) bool) []Record {
	// panic("Please implement the Filter function")
	var returnSlice []Record

	for _, value := range in {
		if predicate(value) == true {
			returnSlice = append(returnSlice, value)
		}
	}

	return returnSlice
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(r Record) bool {
	// panic("Please implement the ByDaysPeriod function")
	return func(r Record) bool {
		if p.From <= r.Day && p.To >= r.Day {
			return true
		}
		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(r Record) bool {
	// panic("Please implement the ByCategory function")
	return func(r Record) bool {
		if r.Category == c {
			return true
		}
		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	// panic("Please implement the TotalByPeriod function")
	var totalSpending float64
	for _, value := range Filter(in, ByDaysPeriod(p)) {
		totalSpending += value.Amount
	}
	return totalSpending
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	// panic("Please implement the CategoryExpenses function")

	//category check
	matchedCategoriesRecord := Filter(in, ByCategory(c))

	if matchedCategoriesRecord == nil {
		return 0, errors.New("unknown category")
	}

	//day period check
	totalAmount := TotalByPeriod(matchedCategoriesRecord, p)

	return totalAmount, nil
}
