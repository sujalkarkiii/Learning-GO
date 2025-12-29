package package1

import "fmt"

func Controlsif()  {
	age:=20
	if age>=18 {
		fmt.Println("adult")
	} else if age<=10{
		fmt.Println("child")
	} else {
		fmt.Println("teen")
	}
}

func Controlsfor()  {
	for i:=0; i<5;i++{
		fmt.Println(i)
	}
}

// Array are fixed sized collection we can slice array to create sub-section
// maps are key value pair that done support slices
// slices can be reslice append copy itterare and manipulate
