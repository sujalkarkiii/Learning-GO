// package main

// import (
// 	"fmt"
// 	package1 "practice/Package1"
// )

// func init(){
// 	fmt.Println("first ")
// }

// func init(){
// 	fmt.Println("second")
// }

// func main() {
// 	fmt.Println("This is the entry Point")
// 	fmt.Println(package1.Insidemodule1())
// 	Another()
// }

package main

import (
	"fmt"
	package1 "practice/Package1"
	package2 "practice/Package2"
)

func main()  {
	 max:= package2.Instrucure{Firstname: "sujal" }
	 fmt.Print(max.Firstname)


	package1.Controlsif()
	package1.Controlsfor()
    defer fmt.Println("This runs at last") 

	var operation string
	var number1, number2 int


	fmt.Println("Calculator:")
	fmt.Println("Enter the opetation you want to perform (add,subtract,multiply,divide)")
	fmt.Scanf("%s",&operation)
	fmt.Println("Enter first number:")
	fmt.Scanf("%d",&number1)
	fmt.Println("Enter second number:")
	fmt.Scanf("%d",&number2)


	switch operation{
	case "add":
		fmt.Println(number1+number2)
	case "subtract":
		fmt.Println(number1-number2)

	case "multiply":
		fmt.Println(number1*number2)
	case "divide":
		if number2==0{
			panic("cannt divide by zero")
		}
		fmt.Println(number1/number2)
	
	default:
		panic("Invalid Operation") 
	}

}


