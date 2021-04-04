package main

import (
	"fmt"
)

func rectSquare(a, b int) int {
	
	var square = a * b
	
	return square
}

func main() {

	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	sq := rectSquare(a, b)

	fmt.Println("Rectangle square is equal: ", sq)
}


// ramin@raminhost:~/go/src/Geekbrains/Go1 (les2_branch)$ go run task_1.go 
// 4
// 5
// Rectangle square is equal:  20