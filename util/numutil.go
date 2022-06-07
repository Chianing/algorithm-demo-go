package util

func GetMaxInt(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func GetMinInt(i1, i2 int) int {
	if i1 > i2 {
		return i2
	}
	return i1
}

func GetAbsInt(i int) int {
	if i >= 0 {
		return i
	}

	return -i
}
