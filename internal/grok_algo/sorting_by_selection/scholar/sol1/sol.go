package sol1

import "fmt"

func Solution() {
	arr := []int{5, 3, 6, 2, 10}
	res := selectionSort(arr)
	fmt.Println(res)
}

func selectionSort(arr []int) []int {
	arrLen := len(arr)
	newArr := make([]int, arrLen)
	copy(newArr, arr)

	for i := 0; i < arrLen; i++ {
		minIndex := i
		for j := i + 1; j < arrLen; j++ {
			if newArr[j] < newArr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			newArr[i], newArr[minIndex] = newArr[minIndex], newArr[i]
		}
	}
	return newArr
}
