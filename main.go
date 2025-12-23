package main

import (
	"fmt"
	package1 "practice/Package1"
)

func init(){
	fmt.Println("first ")
}

func init(){
	fmt.Println("second")
}

func main() {
	fmt.Println("This is the entry Point")
	fmt.Println(package1.Insidemodule1())
	Another()
}
