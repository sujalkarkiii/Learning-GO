package package1
import "fmt"
type Address struct {
    City string
}

type User struct {
    Name    string
    Addr    Address
}


func Runningthis()  {
		u1 := User{ Name: "Sujal", Addr: Address{City: "Kathmandu"} }
		// also can be done as
		u2:=User{"Sujal",Address{"biratnagar"}}
		fmt.Println(u1)
		fmt.Println(u2)

		
}