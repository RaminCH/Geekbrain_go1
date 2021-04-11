package main 

import (
	"fmt"
)

func InsertionSort(array []int) []int {
	for i := range array {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j] 
		}
	}
	return array
}

func main() {

	arr := []int{8, 5, 2, 9, 5, 6, 3} 
	
	insrt := InsertionSort(arr)
	fmt.Println(insrt)
}

// Output:
// ramin@raminhost:~/go/src/Geekbrains/Go1/les4 (les4_branch)$ go run insertion_sort.go 
// [2 3 5 5 6 8 9]
