package main

import (
	"fmt"
)

func main() {
    var num1 int
    var num2 int
    var choice int=0
    var result int=0

	fmt.Print("Enter 1st number: ") 
	fmt.Scanf("%d",&num1) 
	fmt.Print("Enter 2nd number: ") 
	fmt.Scanf("%d",&num2)
    
    fmt.Println("1: Addition") 
    fmt.Println("2: Subtraction") 
    fmt.Println("3: Multiplication") 
    fmt.Println("4: Division") 

	
	 
    fmt.Print("Enter choice: ") 
    fmt.Scanf("%d",&choice) 
    
     switch choice{ 
       case 1: 
            result=num1+num2
            fmt.Printf("Addition is: %d\n",result) 
       case 2: 
            result=num1-num2
            fmt.Printf("Subtraction is: %d\n",result) 
       case 3: 
            result=num1*num2
            fmt.Printf("Multiplication is: %d\n",result) 
       case 4: 
            result=num1/num2
            fmt.Printf("Division is: %d\n",result) 
       default:  
            fmt.Println("Invalid value") 
   }  
}