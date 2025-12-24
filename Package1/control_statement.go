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


// func Array()  {
// 	var arr[3] int 
// 	arr[0]=10
// 	arr[1]=10
// 	arr[2]=10
// 	fmt.Println(arr)   
// }

// func Slice()  {
// 	slices:= []int{1,2,3}
// 	slices= append(slices, 6)
// 	fmt.Println(slices)
// }


