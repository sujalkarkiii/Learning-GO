package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sujal/go/main/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	//  api/exhibitions?id=34
	id := r.URL.Query()["id"]
	if id != nil { //we will try to server only one
		finalId, err := strconv.Atoi(id[0])
		fmt.Println(finalId)
		fmt.Println(len(data.Getall()))

		if err == nil && finalId < len(data.Getall()) {
			fmt.Println("am i here")

			json.NewEncoder(w).Encode(data.Getall()[finalId])

		} else {
			fmt.Println("am i here in else block")
			http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
		}
	} else { //we want all

		json.NewEncoder(w).Encode(data.Getall())
	}

}
