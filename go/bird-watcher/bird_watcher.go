package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	//panic("Please implement the TotalBirdCount() function")
    var total int
    for _, val := range birdsPerDay{
        total += val
    }
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var total int
	var weekCount int
	for i := 0; i < len(birdsPerDay); i++ {
		if i!= 0 && i%7 == 0 {
			weekCount++
			if week == weekCount {
				return total
			}
			total = 0
		}
		total += birdsPerDay[i]
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	//panic("Please implement the FixBirdCountLog() function")
    for i:=0; i<len(birdsPerDay); i=i+2{
        birdsPerDay[i]++
    }
	return birdsPerDay
}
