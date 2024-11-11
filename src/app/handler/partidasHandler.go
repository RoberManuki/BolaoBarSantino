package handler

import (
	"bolao/src/app/model"
	repository "bolao/src/app/resource"
	"bolao/src/app/service"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

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

	rodadaStr := r.URL.Query().Get("rodada")
	rodada := 1
	if rodadaStr != "" {
		var err error
		rodada, err = strconv.Atoi(rodadaStr)
		if err != nil {
			http.Error(w, "Rodada inválida", http.StatusBadRequest)
			return
		}
	}

	partidas, err := service.GetPartidas(model.FiltroPartida{Rodada: rodada})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(partidas)
}

func CreatePartida(w http.ResponseWriter, r *http.Request) {
	var partidaCreate model.PartidaCreate

	err := json.NewDecoder(r.Body).Decode(&partidaCreate)
	if err != nil {
		http.Error(w, "Erro ao processar dados", http.StatusBadRequest)
		return
	}

	err = service.CreatePartida(partidaCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Partida criada com sucesso!"))
}

func UpdatePartida(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/partidas/")
	var partida model.Partida

	body, err := io.ReadAll(r.Body)
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

func ValidarPartidaHandler(w http.ResponseWriter, r *http.Request) {
	rodadaStr := r.URL.Query().Get("rodada")
	timeCasaStr := r.URL.Query().Get("timeCasa")
	timeForaStr := r.URL.Query().Get("timeFora")

	rodada, _ := strconv.Atoi(rodadaStr)
	timeCasa, _ := strconv.Atoi(timeCasaStr)
	timeFora, _ := strconv.Atoi(timeForaStr)

	// Verificar se os times já jogaram
	jogaram := repository.JogaramNaRodada(timeCasa, timeFora, rodada)

	response := map[string]bool{"jaJogaram": jogaram}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
