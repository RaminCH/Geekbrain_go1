package main

import (
	"fmt"
)

type Eval interface {
	calculate()
	// calculateFloat()
	calculatePositive()
}

type ValuesFloat struct {
	Num1, Num2 float64
	Choice     float64
	Result     float64
}

type Values struct {
	Num1, Num2 int
	Choice     int
	Result     int
}

func (v *Values) calculate() {
	switch v.Choice {
	case 1:
		v.Result = v.Num1 + v.Num2
		fmt.Printf("Addition 'Int' is: %d\n", v.Result)
	case 2:
		v.Result = v.Num1 - v.Num2
		fmt.Printf("Subtraction 'Int' is: %d\n", v.Result)
	case 3:
		v.Result = v.Num1 * v.Num2
		fmt.Printf("Multiplication 'Int' is: %d\n", v.Result)
	case 4:
		v.Result = v.Num1 / v.Num2
		fmt.Printf("Division 'Int' is: %d\n", v.Result)
	default:
		fmt.Println("Invalid value")
	}
}

func (v *ValuesFloat) calculate() {

	switch v.Choice {
	case 1:
		v.Result = v.Num1 + v.Num2
		fmt.Printf("Addition 'Float' is: %f\n", v.Result)
	case 2:
		v.Result = v.Num1 - v.Num2
		fmt.Printf("Subtraction 'Float' is: %f\n", v.Result)
	case 3:
		v.Result = v.Num1 * v.Num2
		fmt.Printf("Multiplication 'Float' is: %f\n", v.Result)
	case 4:
		v.Result = v.Num1 / v.Num2
		fmt.Printf("Division 'Float' is: %f\n", v.Result)
	default:
		fmt.Println("Invalid value")
	}

}

func (v *Values) calculatePositive() {
	if v.Num1 <= 0 || v.Num2 <= 0 {
		fmt.Println("Please enter positive value")
	} else {
		switch v.Choice {
		case 1:
			v.Result = v.Num1 + v.Num2
			fmt.Printf("Addition 'Positive' is: %d\n", v.Result)
		case 2:
			v.Result = v.Num1 - v.Num2
			fmt.Printf("Subtraction 'Positive' is: %d\n", v.Result)
		case 3:
			v.Result = v.Num1 * v.Num2
			fmt.Printf("Multiplication 'Positive' is: %d\n", v.Result)
		case 4:
			v.Result = v.Num1 / v.Num2
			fmt.Printf("Division 'Positive' is: %d\n", v.Result)
		default:
			fmt.Println("Invalid value")
		}
	}
}

func (v *ValuesFloat) calculatePositive() {
	if v.Num1 <= 0 || v.Num2 <= 0 {
		fmt.Println("Please enter positive value")
	} else {
		switch v.Choice {
		case 1:
			v.Result = v.Num1 + v.Num2
			fmt.Printf("Addition 'Positive' is: %f\n", v.Result)
		case 2:
			v.Result = v.Num1 - v.Num2
			fmt.Printf("Subtraction 'Positive' is: %f\n", v.Result)
		case 3:
			v.Result = v.Num1 * v.Num2
			fmt.Printf("Multiplication 'Positive' is: %f\n", v.Result)
		case 4:
			v.Result = v.Num1 / v.Num2
			fmt.Printf("Division 'Positive' is: %f\n", v.Result)
		default:
			fmt.Println("Invalid value")
		}
	}
}
func measure(e Eval) {
	fmt.Println(e)
	e.calculatePositive()
	// e.calculateFloat()
	e.calculatePositive()
}

func main() {

	val := &Values{Num1: 5, Num2: 6}            //for ints
	valfl := &ValuesFloat{Num1: 5.2, Num2: 6.7} //for float

	measure(val)
	measure(valfl)

	// var i interface{} = val
	// _, ii := i.(Eval)
	// fmt.Println("Result is: ",ii)  // Result is: False

	// var funcs = []Eval{
	// 	&Values{Num1: 5, Num2: 6},
	// 	&ValuesFloat{Num1: 5.2, Num2: 6.7},
	// }

	fmt.Println("1: Addition")
	fmt.Println("2: Subtraction")
	fmt.Println("3: Multiplication")
	fmt.Println("4: Division")

	fmt.Print("Enter choice: ")
	fmt.Scanf("%d", &val.Choice)

	// for _, row := range funcs {
	// 	row.calculateInt()
	// 	row.calculatePositive()
	// 	row.calculateFloat()
	// }

}
