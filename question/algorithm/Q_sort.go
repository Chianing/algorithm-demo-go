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

// 归并排序
func sortMerge(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	mid := length / 2
	leftArr := arr[:mid]
	rightArr := arr[mid:]

	return merge(sortMerge(leftArr), sortMerge(rightArr))

}

func merge(leftArr []int, rightArr []int) []int {
	length := len(leftArr) + len(rightArr)
	ret := make([]int, length)

	for i, leftIndex, rightIndex := 0, 0, 0; i < length; i++ {
		if leftIndex >= len(leftArr) {
			ret[i] = rightArr[rightIndex]
			rightIndex++
		} else if rightIndex >= len(rightArr) {
			ret[i] = leftArr[leftIndex]
			leftIndex++
		} else if leftArr[leftIndex] > rightArr[rightIndex] {
			ret[i] = rightArr[rightIndex]
			rightIndex++
		} else {
			ret[i] = leftArr[leftIndex]
			leftIndex++
		}
	}

	return ret

}

// 快速排序
func sortQuick(arr []int, left, right int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	if left < right {
		index := getPartition(arr, left, right)
		sortQuick(arr, left, index-1)
		sortQuick(arr, index+1, right)
	}

	return arr
}

func getPartition(arr []int, left, right int) int {
	index := left + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[left] {
			swap(arr, i, index)
			index++
		}
	}
	index--

	swap(arr, index, left)
	return index
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
