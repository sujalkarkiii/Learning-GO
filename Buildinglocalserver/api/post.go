package api

import (
	"encoding/json"
	"net/http"
	"sujal/go/main/data"
)

func Post(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var exhibitions data.Exibition
		err:=json.NewDecoder(r.Body).Decode(&exhibitions)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusBadRequest)
		}
		data.Add(exhibitions)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Ok"))
		} else { //we want all
		http.Error(w, "unsuported Method", http.StatusMethodNotAllowed)
	}

}
