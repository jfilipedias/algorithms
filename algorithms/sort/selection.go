package sort

// complexity O(n**2)
func SelectionSort(arr []int) []int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		sortedArr[i], sortedArr[min] = sortedArr[min], sortedArr[i]
	}
	return sortedArr
}
