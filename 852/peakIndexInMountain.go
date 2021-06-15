package peakIndexInMountain

func peakIndexInMountainArray(arr []int) int {
	var index int
	var arrLen int
	preIndex := 0

	for {
		arrLen = len(arr)
		index = arrLen / 2
		if arr[index] < arr[index+1] {
			arr = arr[index-1:]
			preIndex += index - 1
		} else if arr[index] < arr[index-1] {
			arr = arr[:index+1]
		} else {
			return index + preIndex
		}
	}
	return 0
}