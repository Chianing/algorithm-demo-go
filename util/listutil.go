package util

func ReverseList(list []any) []any {
	size := len(list)
	if size < 2 {
		return list
	}

	for i := 0; i < size/2; i++ {
		Swap(list, i, size-i-1)
	}

	return list

}

func Swap(list []any, i1, i2 int) {
	size := len(list)
	if i1 < 0 || i1 >= size || i2 < 0 || i2 >= size {
		return
	}

	tmp := list[i1]
	list[i1] = list[i2]
	list[i2] = tmp

}
