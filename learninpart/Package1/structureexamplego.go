package package1

import "fmt"

type Person struct {
	Id   int
	Name string
}


// func(u Person) Greet(){
// 	u.Name="dharan"
// 	fmt.Println("Hello, my name is", u.Name, "and I am", u.Age, "years old.")
// }


// method with pointer reciver to modify struct 


// func (u *Person) UpdateName(){
// 		u.Name = "dharan"
// 		fmt.Println("Hello, my name is", u.Name, "and I am", u.Age, "years old.")
// 	}


func Run() {
	// u1 := Person{1, "sujal", 25}
	// fmt.Println(u1)

	// update
	// u1.Name = "karki"
	// fmt.Println(u1)
	// pointer to struct

	// u3 := Person{1, "Ram", 25}
	// p := &u3
	// p.Name = "hari"
	// fmt.Println(p)
	// u1.Greet()
	// u1.UpdateName()
	// fmt.Println(u1)

	// result on not using function as ponter
// {1 sujal 25}
// {1 karki 25}
// &{1 hari 25}
// Hello, my name is dharan and I am 25 years old.
// {1 karki 25}

people := []Person{ {1, "Ram"},
					{2, "Sita"}}
fmt.Println("People:", people)


// Iterate over slice 
for _, p := range people { 
	fmt.Printf("ID: %d, Name: %s\n", p.Id, p.Name) 
}
// Append new person 
people = append(people, Person{3, "Hari"}) 
fmt.Println("After append:", people)



}
