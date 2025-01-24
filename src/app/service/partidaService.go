package service

import (
	"bolao/src/app/model"
	repository "bolao/src/app/resource"
	"log"
)

func GetPartidas(filtro model.FiltroPartida) ([]model.Partida, error) {
	return repository.GetPartidas(filtro)
}

func UpdatePartida(id int, partida model.Partida) error {
	return repository.UpdatePartida(id, partida)
}

func CreatePartida(partidaCreate model.PartidaCreate) error {

	err := repository.CreatePartida(partidaCreate)
	if err != nil {
		log.Printf("Erro ao criar partida: %v", err)
		return err
	}

	return nil
}

func GetPartidaByID(id int) (model.Partida, error) {
	return repository.GetPartidaByID(id)
}

func DeletePartida(id int) error {
	err := repository.DeletePartida(id)
	if err != nil {
		log.Printf("Erro ao excluir partida na camada de serviço: %v", err)
		return err
	}
	return nil
}
