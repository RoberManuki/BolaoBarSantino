package service

import (
	"bolao/src/app/model"
	repository "bolao/src/app/resource"
	"errors"
	"log"
	"strings"
)

func GetPartidas(filtro model.FiltroPartida) ([]model.Partida, error) {
	return repository.GetPartidas(filtro)
}

func UpdatePartida(id string, partida model.Partida) error {
	return repository.UpdatePartida(id, partida)
}

func CreatePartida(partidaCreate model.PartidaCreate) error {
	jogouNaRodada := repository.JogouNaRodada(partidaCreate.TimeCasa, partidaCreate.TimeFora, partidaCreate.Rodada)

	if jogouNaRodada {
		errMsg := "Os times já jogaram nesta rodada."
		log.Printf("Erro: %s", errMsg)
		return errors.New(strings.ToLower(errMsg))
	}

	if partidaCreate.TimeCasa == partidaCreate.TimeFora {
		errMsg := "Os times selecionados são iguais."
		return errors.New(strings.ToLower(errMsg))
	}

	err := repository.CreatePartida(partidaCreate)
	if err != nil {
		log.Printf("Erro ao criar partida: %v", err)
		return err
	}

	return nil
}
