package package2

import "fmt"

func Learnslices()  {
	var arr [5]int
	arr=[5]int{1,2,3,4,5}
	slice:=arr[1:4]
	
	fmt.Println("Slice:", slice)

	slice=append(slice, 60)
	fmt.Println("Appended Slice:", slice)


	// fmt.Println(arr)
	// fmt.Println("Element at index 2:", arr[2])

	for i ,v:= range slice{
		fmt.Printf("Index %d: value %d\n",i,v)
	}
	
sub := slice[1:4]
 fmt.Println("Resliced:", sub)
}