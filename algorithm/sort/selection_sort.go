package sort

// complexity O(n**2)
func SelectionSort(arr []int) []int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	for i := 0; i < len(sortedArr); i++ {
		min := i
		for j := i + 1; j < len(sortedArr); j++ {
			if sortedArr[j] < sortedArr[min] {
				min = j
			}
		}
		sortedArr[i], sortedArr[min] = sortedArr[min], sortedArr[i]
	}
	return sortedArr
}
