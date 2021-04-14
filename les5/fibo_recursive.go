package main

import (
	"fmt"
)

func Fibonacci(n int) int {

    if n < 0 {
        print("Incorrect input")
	} else if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	} 
	
	return Fibonacci(n-1) + Fibonacci(n-2)
}


func main() {

fmt.Println("Please enter the number for the Fibonacci sequence: ")
var inp int  
fmt.Scan(&inp)

p := Fibonacci(inp)

fmt.Printf("Fibonacci number is: %v\n", p)
}