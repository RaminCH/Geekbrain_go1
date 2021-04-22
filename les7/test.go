package main

import (
	"fmt"
	"reflect"
)

type Eval interface {
	calculateInt()
	calculateFloat()
	calculatePositive()
}

type ValuesFloat struct {
	Numfl1, Numfl2 float64 
	Choicefl float64
	Resultfl float64
}

type Values struct {
	Num1, Num2 int
	Choice int
	Result int
}

func(v *Values) calculateInt() {
	switch v.Choice{ 
		case 1: 
			v.Result=v.Num1+v.Num2
			fmt.Printf("Addition 'Int' is: %d\n",v.Result) 
		case 2: 
			v.Result=v.Num1-v.Num2
			fmt.Printf("Subtraction 'Int' is: %d\n",v.Result) 
		case 3: 
			v.Result=v.Num1*v.Num2
			fmt.Printf("Multiplication 'Int' is: %d\n",v.Result) 
		case 4: 
			v.Result=v.Num1/v.Num2
			fmt.Printf("Division 'Int' is: %d\n",v.Result) 
		default:  
			fmt.Println("Invalid value") 
	}  
}


func(v *ValuesFloat) calculateFloat() {
	
	switch v.Choicefl{ 
		case 1: 
			v.Resultfl=v.Numfl1+v.Numfl2
			fmt.Printf("Addition 'Float' is: %d\n",v.Resultfl) 
		case 2: 
			v.Resultfl=v.Numfl1-v.Numfl2
			fmt.Printf("Subtraction 'Float' is: %d\n",v.Resultfl) 
		case 3: 
			v.Resultfl=v.Numfl1*v.Numfl2
			fmt.Printf("Multiplication 'Float' is: %d\n",v.Resultfl) 
		case 4: 
			v.Resultfl=v.Numfl1/v.Numfl2
			fmt.Printf("Division 'Float' is: %d\n",v.Resultfl) 
		default:  
			fmt.Println("Invalid value") 
	}  

}


func(v *Values) calculatePositive() {
	if v.Num1 <= 0 || v.Num2 <=0 {
		fmt.Println("Please enter positive value")	
	} else {
		switch v.Choice{ 
			case 1: 
			v.Result=v.Num1+v.Num2
			fmt.Printf("Addition 'Positive' is: %d\n",v.Result) 
		case 2: 
			v.Result=v.Num1-v.Num2
			fmt.Printf("Subtraction 'Positive' is: %d\n",v.Result) 
		case 3: 
			v.Result=v.Num1*v.Num2
			fmt.Printf("Multiplication 'Positive' is: %d\n",v.Result) 
		case 4: 
			v.Result=v.Num1/v.Num2
			fmt.Printf("Division 'Positive' is: %d\n",v.Result) 
		default:  
			fmt.Println("Invalid value") 
		} 
	}
} 

func main() {

	var val = &Values{}	
	var valFl = &ValuesFloat{}

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

	if val.Num1 > 0 && val.Num2 > 0 {
		val.calculatePositive()
		if val.Num1%1==0 && val.Num1%2==0 {
			val.calculateInt()
		} else {
			val.calculateFloat()
		}
	} else {
		val.calculatePositive()
	}		
}