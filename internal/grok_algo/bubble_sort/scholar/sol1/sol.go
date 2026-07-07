package sol1

import "fmt"

func Solution() {
	arr := []int{5, 3, 6, 2, 10}
	res := bubbleSort(arr)
	fmt.Println(res)
}

func bubbleSort(arr []int) []int {
	arrLen := len(arr)
	newArr := make([]int, arrLen)
	copy(newArr, arr)

	for i := 0; i < arrLen; i++ {
		fmt.Println("I is", arr[i])

		for j := 0; j < arrLen-i; j++ {
			fmt.Println(" J is", arr[j])

		}

		fmt.Println("----------")
	}

	return newArr
}
