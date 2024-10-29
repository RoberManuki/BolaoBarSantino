package handler

import (
	"bolao/src/app/model"
	"bolao/src/app/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// PartidaHandler lida com solicitações para /partidas
func PartidaHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetPartidas(w, r)
	case http.MethodPost:
		CreatePartida(w, r)
	case http.MethodPut:
		UpdatePartida(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func GetPartidas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Substitua este código pelo real acesso ao banco de dados
	partidas, err := service.GetPartidas(model.FiltroPartida{Rodada: 1})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(partidas)
}

func CreatePartida(w http.ResponseWriter, r *http.Request) {
	var partida model.Partida
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &partida)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.CreatePartida(partida)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func UpdatePartida(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/partidas/")
	var partida model.Partida
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &partida)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.UpdatePartida(id, partida)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
