package division

func InitiatePopulatedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}
func Rest(number int, division int) bool {
	return number%division == 0
}
