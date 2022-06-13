package util

func ReverseList(list []int) []int {
	size := len(list)
	if size < 2 {
		return list
	}

	for i := 0; i < size/2; i++ {
		Swap(list, i, size-i-1)
	}

	return list

}

func Swap(list []int, i1, i2 int) {
	size := len(list)
	if i1 < 0 || i1 >= size || i2 < 0 || i2 >= size {
		return
	}

	list[i1], list[i2] = list[i2], list[i1]

}
