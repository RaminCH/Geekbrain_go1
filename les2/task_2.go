package main

import (
	"fmt"
	"math"
)

func circleDiameter(a float64) float64 {
	
	var diameter = math.Sqrt(a / math.Pi) * 2
	
	return diameter
}

func circleLength(d float64) float64 {
	
	var length = math.Pi * d 

	return length
}

func main() {

	var sq float64

	fmt.Scan(&sq)

	di := math.RoundToEven(circleDiameter(sq))

	le := math.RoundToEven(circleLength(di))

	fmt.Println("Circle diameter is equal: ", di)
	fmt.Println("Circle length is equal: ", le)
}


// ramin@raminhost:~/go/src/Geekbrains/Go1 (les2_branch)$ go run task_2.go 
// 12.56
// Circle diameter is equal:  4
// Circle length is equal:  13
