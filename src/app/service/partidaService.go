package service

import (
	"bolao/src/app/model"
	repository "bolao/src/app/resource"
)

func GetPartidas(filtro model.FiltroPartida) ([]model.Partida, error) {
	return repository.GetPartidas(filtro)
}

func CreatePartida(partidaCreate model.PartidaCreate) error {
	return repository.CreatePartida(partidaCreate)
}

func UpdatePartida(id string, partida model.Partida) error {
	return repository.UpdatePartida(id, partida)
}
