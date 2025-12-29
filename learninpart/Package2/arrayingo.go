package package2

import "fmt"

func Learnarray()  {
	var arr [5]int
	arr=[5]int{1,2,3,4,5}
	fmt.Println(arr)
	fmt.Println("Element at index 2:", arr[2])

	for i ,v:= range arr{
		fmt.Printf("Index %d: value %d\n",i,v)
	}
	
	arr[1] = 99 
	fmt.Println("Modified Array:", arr)
}