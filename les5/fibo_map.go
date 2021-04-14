// package main

// import (
// 	"fmt"
// )


	

// func fibonacci() func() int {
	
// 	first, second := 0, 1
// 	return func() int {
// 		ret := first
// 		first, second = second, first+second
// 		return ret
// 	}
	
// }


// func main() {

// fmt.Println("Please enter the number for the Fibonacci sequence: ")
// var inp int  
// fmt.Scan(&inp)

// f := fibonacci()
// 	for i := 0; i <= inp; i++ {
// 		fmt.Println(f())
// 	}
// }


package main

import (
	"fmt"
)


	

func fibonacci() func() int {
	m := make(map[int]int)
	m = append(m, ret)
	
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
	
}


func main() {

fmt.Println("Please enter the number for the Fibonacci sequence: ")
var inp int  
fmt.Scan(&inp)

f := fibonacci()
	for i := 0; i <= inp; i++ {
		fmt.Println(f())
	}
}

