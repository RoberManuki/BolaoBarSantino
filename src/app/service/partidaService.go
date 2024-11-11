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

	if err := ValidacoesFormulario(partidaCreate); err != nil {
		return err
	}

	err := repository.CreatePartida(partidaCreate)
	if err != nil {
		log.Printf("Erro ao criar partida: %v", err)
		return err
	}

	return nil
}

func ValidacoesFormulario(partidaCreate model.PartidaCreate) error {

	if partidaCreate.TimeCasa == partidaCreate.TimeFora {
		errMsg := "Os times selecionados são iguais."
		return errors.New(strings.ToLower(errMsg))
	}

	jogouNaRodada := repository.JogaramNaRodada(partidaCreate.TimeCasa, partidaCreate.TimeFora, partidaCreate.Rodada)
	if jogouNaRodada {
		errMsg := "Um dos times já jogou nesta rodada."
		log.Printf("Erro: %s", errMsg)
		return errors.New(strings.ToLower(errMsg))
	}

	return nil
}
