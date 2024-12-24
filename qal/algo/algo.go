package algo

func BinarySearchInt(arr []int, target int) int {
	var min int = 0
	var max int = len(arr) - 1

	for min <= max {
		var mid int = (min + max) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}
