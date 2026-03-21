package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    
    birdCount := 0
    
	for idx := 0; idx < len(birdsPerDay); idx++ {
        birdCount += birdsPerDay[idx]
    }
    return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdCount := 0
	start := (week - 1) * 7
	end := start + 7

	for i := start; i < end; i++ {
		birdCount += birdsPerDay[i]
	}
	return birdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i] += 1
    }
    return birdsPerDay
}
