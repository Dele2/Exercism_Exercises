package binarysearch

func SearchInts(list []int, key int) int {

	// Set low and high boundaries
	low := 0
	high := len(list) - 1
	// Calculate middle index
	mid := (low + high) // 2

	return recurs(list, key, low, mid, high)
}

func recurs(list []int, key int, low int, mid int, high int) int {

	if low > high {
		return -1
	} else if list[mid] == key {
		return mid
	} else if list[mid] > key {
		high = mid - 1
		mid = (low + high) / 2
		return recurs(list, key, low, mid, high)
	} else {
		low = mid + 1
		mid = (low + high) / 2
		return recurs(list, key, low, mid, high)
	}
}