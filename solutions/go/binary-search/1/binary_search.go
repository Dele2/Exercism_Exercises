package binarysearch

func SearchInts(list []int, key int) int {

	// Set low and high boundaries
	low := 0
	high := len(list) - 1
	// Calculate middle index
	mid := (low + high) // 2

	for low <= high {
		// Check if element is found
		if list[mid] == key {
			return mid
			// If target is greater, focus on the right half
		} else if list[mid] > key {
			high = mid - 1
			mid = (low + high) // 2
			// If target is less, focus on the left half
		} else {
			low = mid + 1
			mid = (low + high) // 2
		}
	}
	// If element not found, return -1
	return -1
}