package handler

import (
	"bolao/src/app/model"
	partidaRepository "bolao/src/app/resource"
	partidaService "bolao/src/app/service"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// CRUD partidas
func PartidaHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		Get(w, r)
	case http.MethodPost:
		Create(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
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

	partidas, err := partidaService.GetPartidas(model.FiltroPartida{Rodada: rodada})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(partidas)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var partidaCreate model.PartidaCreate

	err := json.NewDecoder(r.Body).Decode(&partidaCreate)
	if err != nil {
		http.Error(w, "Erro ao processar dados", http.StatusBadRequest)
		return
	}

	err = partidaService.CreatePartida(partidaCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Partida criada com sucesso!"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := extrairID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	err = partidaService.UpdatePartida(id, partida)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Partida atualizada com sucesso!"))
}

// Função auxiliar para extrair e validar o id da URL
func extrairID(r *http.Request) (int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/partidas/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("ID inválido")
	}
	return id, nil
}

func PartidaByID(w http.ResponseWriter, r *http.Request) {
	id, err := extrairID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet: // GET
		if id == 0 {
			http.Error(w, "ID da partida não informado", http.StatusBadRequest)
			return
		}

		log.Println("api/GET")
		partida, err := partidaService.GetPartidaByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(partida)

	case http.MethodPut: // PUT
		var partida model.Partida
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("api/PUT")
		err = json.Unmarshal(body, &partida)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = partidaService.UpdatePartida(id, partida)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Partida atualizada com sucesso!"))

	case http.MethodDelete:
		err := partidaService.DeletePartida(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Partida excluída com sucesso!"))

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func ValidarPartidaHandler(w http.ResponseWriter, r *http.Request) {
	rodadaStr := r.URL.Query().Get("rodada")
	timeCasaStr := r.URL.Query().Get("timeCasa")
	timeForaStr := r.URL.Query().Get("timeFora")

	rodada, _ := strconv.Atoi(rodadaStr)
	timeCasa, _ := strconv.Atoi(timeCasaStr)
	timeFora, _ := strconv.Atoi(timeForaStr)

	jogaram := partidaRepository.JogaramNaRodada(rodada, timeCasa, timeFora)

	response := map[string]bool{"jaJogaram": jogaram}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
