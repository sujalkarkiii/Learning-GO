package main

import (
	"fmt"
	"net/http"
	"sujal/go/main/api"

)

func handlehello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go program!!!!!"))
}



func main()  {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handlehello)
	server.HandleFunc("/api/exibitions",api.Get)
	server.HandleFunc("/api/exibitions/new",api.Post)
	// statically serving a folder
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)
	if err == nil {
		fmt.Println("Error while opeingin the server")
	}
}
