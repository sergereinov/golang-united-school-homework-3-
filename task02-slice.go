package homework

func reverse(input []int64) (result []int64) {
	size := len(input)
	result = make([]int64, size)
	for i, v := range input {
		result[size-1-i] = v
	}
	return
}
