package handler

import (
	"bolao/src/app/service"
	"encoding/json"

	//"io/ioutil"
	"net/http"
	//"strings"
)

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetTimes(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func GetTimes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	times, err := service.GetTimes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(times)
}
