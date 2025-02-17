package repository

import (
	"bolao/src/app/model"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var dbPartidas *sql.DB

func init() {
	var err error
	dbPartidas, err = sql.Open("postgres", "user=postgres password=sjtbmix6 dbname=santino-bolao sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPartidas(filtro model.FiltroPartida) ([]model.Partida, error) {
	query := `
        SELECT "Id", "Time Casa", "Casa Gols", "Time Fora", "Fora Gols", "Data", "Vencedor"
        FROM "schema-bolao-24"."Partida"
        WHERE "Rodada" = $1
		ORDER BY "Id"
    `

	rows, err := dbPartidas.Query(query, filtro.Rodada)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var partidas []model.Partida
	for rows.Next() {
		var p model.Partida
		if err := rows.Scan(
			&p.Id,
			&p.TimeCasa,
			&p.CasaGols,
			&p.TimeFora,
			&p.ForaGols,
			&p.Data,
			&p.Vencedor,
		); err != nil {
			return nil, err
		}
		partidas = append(partidas, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return partidas, nil
}

func CreatePartida(partidaCreate model.PartidaCreate) error {
	_, err := dbPartidas.Exec(`
        INSERT INTO "schema-bolao-24"."Partida" 
        ("Rodada", "Time Casa", "Casa Gols", "Time Fora", "Fora Gols", "Data", "Vencedor") 
        VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		partidaCreate.Rodada, partidaCreate.TimeCasa, partidaCreate.CasaGols,
		partidaCreate.TimeFora, partidaCreate.ForaGols, partidaCreate.Data,
		partidaCreate.Vencedor)

	if err != nil {
		log.Printf("Erro ao inserir partida no banco de dados: %v", err)
		return err
	}

	return nil
}

func UpdatePartida(id int, partida model.Partida) error {
	_, err := dbPartidas.Exec(`UPDATE "schema-bolao-24"."Partida" SET "Time Casa" = $1, "Casa Gols" = $2, "Time Fora" = $3, "Fora Gols" = $4, "Data" = $5, "Vencedor" = $6 WHERE "Id" = $7`,
		partida.TimeCasa, partida.CasaGols, partida.TimeFora, partida.ForaGols, partida.Data, partida.Vencedor, id)
	return err
}

func JogaramNaRodada(timeCasaId int, timeForaId int, rodada int) bool {

	query := `
		SELECT 1
		FROM "schema-bolao-24"."Partida"
		WHERE "Rodada" = $1 
		AND (("Time Casa" = $2 OR "Time Fora" = $2)
		OR ("Time Casa" = $3 OR "Time Fora" = $3))
	`

	result, err := dbPartidas.Query(query, rodada, timeCasaId, timeForaId)
	if err != nil {
		log.Printf("Erro ao executar a consulta: %v", err)
		return false
	}
	defer result.Close()

	return result.Next()
}

func GetPartidaByID(id int) (model.Partida, error) {
	query := `
        SELECT "Id", "Time Casa", "Casa Gols", "Time Fora", "Fora Gols", "Data", "Vencedor", "Rodada"
        FROM "schema-bolao-24"."Partida"
        WHERE "Id" = $1
    `
	var p model.Partida
	row := dbPartidas.QueryRow(query, id)
	err := row.Scan(
		&p.Id,
		&p.TimeCasa,
		&p.CasaGols,
		&p.TimeFora,
		&p.ForaGols,
		&p.Data,
		&p.Vencedor,
		&p.Rodada,
	)
	if err != nil {
		log.Printf("Erro ao consultar partida: %v", err)
		return p, err
	}
	return p, nil
}

func DeletePartida(id int) error {
	_, err := dbPartidas.Exec(`DELETE FROM "schema-bolao-24"."Partida" WHERE "Id" = $1`, id)
	if err != nil {
		log.Printf("Erro ao excluir a partida: %v", err)
		return err
	}
	return nil
}
