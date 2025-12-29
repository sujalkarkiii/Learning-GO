package main

import (
	"fmt"
	"net/http"
)

func handlehello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go program!!!!!"))
}





func main()  {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handlehello)

	// statically serving a folder
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)
	if err == nil {
		fmt.Println("Error while opeingin the server")
	}
}
