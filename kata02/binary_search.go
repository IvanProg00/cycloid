package kata02

func IterativeSearch(num int, arr []int) int {
	for i, v := range arr {
		if v == num {
			return i
		}
	}

	return -1
}

func RecursiveSearch(num int, arr []int) int {
	return recursiveSearch(num, arr, 0)
}

func recursiveSearch(num int, arr []int, pos int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		if arr[0] == num {
			return pos
		} else {
			return -1
		}
	}

	center := len(arr) / 2

	if num == arr[center] {
		return pos + center
	}
	if num < arr[center] {
		return recursiveSearch(num, arr[:center], pos)
	}
	if num > arr[center] {
		return recursiveSearch(num, arr[center+1:], pos+center+1)
	}

	return -1
}
