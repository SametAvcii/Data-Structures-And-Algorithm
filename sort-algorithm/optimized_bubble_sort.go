package sort_algorithm

func bubbleSort(array []int, size int) {

	var swapped bool = false
	for i := 0; i < size; i++ {
		swapped = false

		for j := 0; j < size-i-1; j++ {
			if array[j+1] < array[j] {
				var temp int
				temp = array[j]
				array[j] = array[j+1]
				array[j+1] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}