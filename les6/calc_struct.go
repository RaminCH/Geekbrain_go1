package main

import (
	"fmt"
)

type Values struct {
	Num1, Num2 int
	Choice int
	Result int
}

func(v *Values) calculate() {
	switch v.Choice{ 
		case 1: 
			v.Result=v.Num1+v.Num2
			fmt.Printf("Addition is: %d\n",v.Result) 
		case 2: 
			v.Result=v.Num1-v.Num2
			fmt.Printf("Subtraction is: %d\n",v.Result) 
		case 3: 
			v.Result=v.Num1*v.Num2
			fmt.Printf("Multiplication is: %d\n",v.Result) 
		case 4: 
			v.Result=v.Num1/v.Num2
			fmt.Printf("Division is: %d\n",v.Result) 
		default:  
			fmt.Println("Invalid value") 
	}  
} 

func main() {

	val := &Values{0, 0, 0, 0}
    
	fmt.Print("Enter 1st number: ") 
	fmt.Scanf("%d",&val.Num1) 			
	fmt.Print("Enter 2nd number: ") 
	fmt.Scanf("%d",&val.Num2)			

    
    fmt.Println("1: Addition") 
    fmt.Println("2: Subtraction") 
    fmt.Println("3: Multiplication") 
    fmt.Println("4: Division") 


    fmt.Print("Enter choice: ") 
    fmt.Scanf("%d",&val.Choice) 		

	
	val.calculate()
}



// ramin@raminhost:~/go/src/Geekbrains/Go1/les6 (les6_branch)$ go run calc_struct.go 
// Enter 1st number: 5
// Enter 2nd number: 6
// 1: Addition
// 2: Subtraction
// 3: Multiplication
// 4: Division
// Enter choice: 3
// Multiplication is: 30