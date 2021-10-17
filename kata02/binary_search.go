package kata02

func IterativeSearch(num int, arr []int) int {
	for i, v := range arr {
		if v == num {
			return i
		}
	}

	return -1
}
