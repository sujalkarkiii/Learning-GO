package package1

import (
    "encoding/json"
    "fmt"
)

type Random struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func Jsoning() {
    u := Random{
        Name:  "Sujal",
        Age:   22,
        Email: "sujal@gmail.com",
    }

    jsonData, err := json.Marshal(u)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(jsonData))

	arr := [3]int{10, 20, 30}
    slice := []string{"Go", "Node", "Python"}
    m := map[string]int{
        "apple":  5,
        "banana": 10,
    }

    arrJSON, _ := json.Marshal(arr)
    sliceJSON, _ := json.Marshal(slice)
    mapJSON, _ := json.Marshal(m)

    fmt.Println("Array JSON:", string(arrJSON))
    fmt.Println("Slice JSON:", string(sliceJSON))
    fmt.Println("Map JSON:", string(mapJSON))
}
