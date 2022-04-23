package algorithm

// 冒泡排序
func sortBubble(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	var tmpNum int

	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmpNum = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmpNum
			}
		}
	}

	return arr

}

// 直接插入排序
func sortInsert(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	var currentNum, sortedIndex int

	for i := 1; i < length; i++ {
		currentNum = arr[i]

		for sortedIndex = i - 1; sortedIndex >= 0 && currentNum < arr[sortedIndex]; sortedIndex-- {
			arr[sortedIndex+1] = arr[sortedIndex]
		}
		arr[sortedIndex+1] = currentNum
	}

	return arr
}

// 希尔排序
func sortShell(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	var currentNum, sortedIndex int

	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			currentNum = arr[i]
			for sortedIndex = i - gap; sortedIndex >= 0 && currentNum < arr[sortedIndex]; sortedIndex -= gap {
				arr[sortedIndex+gap] = arr[sortedIndex]
			}
			arr[sortedIndex+gap] = currentNum
		}
	}

	return arr

}
