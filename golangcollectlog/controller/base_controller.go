package controller

import (
	"encoding/json"
	"net/http"
)

func Base(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err:=encoder.Encode("welcome to golang collect log service")
	if err != nil {
		panic(err)
	}

}