package main

import (
	"fmt"
)


func main() {

	var a int

	fmt.Scan(&a)

	e := a % 10 
	d := a / 10 % 10  
	s := a / 100


	fmt.Printf("Hundreds: %d, Tens: %d, Units: %d\n",s,d,e)
}


