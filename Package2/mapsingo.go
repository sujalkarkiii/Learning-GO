package package2

// map[keyType]valueType

import "fmt"

func Learnmap()  {

	m := map[string]int{"a": 1, "b": 2}	
	fmt.Println("Map:", m)

	// if the key is not present than it addds key if present that it updates/rewites

	m["c"]=3
	m["b"]=1
	fmt.Println("Updated Map:", m) 
	// Delete 
	delete(m, "b") 
	fmt.Println("After Deletion:", m)


	val, ok := m["c"]
// here first one will always return val and the next will return if present or not in map lookup
	fmt.Println(val,ok)


	for k, v := range m {
    fmt.Printf("Key: %s, Value: %d\n", k, v)
}

// 
}