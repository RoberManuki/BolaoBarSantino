package repository

import (
	"database/sql"
	//"fmt"
	"bolao/src/app/model"

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
	// Definir a consulta SQL com quebra de linha para melhorar a legibilidade
	query := `
        SELECT "Id", "Rodada", "Time Casa", "Casa Gols", "Time Fora", "Fora Gols", "Data", "Vencedor"
        FROM "schema-bolao-24"."Partida"
        WHERE "Rodada" = $1
		ORDER BY "Id"
    `

	// Executar a consulta passando o filtro.Rodada como parâmetro
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
			&p.Rodada,
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

	// Verificar se houve um erro durante a iteração das linhas
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return partidas, nil
}

func CreatePartida(partida model.Partida) error {
	_, err := dbPartidas.Exec(`INSERT INTO "schema-bolao-24"."Partida" ("Rodada", "Time Casa", "Casa Gols", "Time Fora", "Fora Gols", "Data", "Vencedor") VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		partida.Rodada, partida.TimeCasa, partida.CasaGols, partida.TimeFora, partida.ForaGols, partida.Data, partida.Vencedor)
	return err
}

func UpdatePartida(id string, partida model.Partida) error {
	_, err := dbPartidas.Exec(`UPDATE "schema-bolao-24"."Partida" SET "Rodada" = $1, "Time Casa" = $2, "Casa Gols" = $3, "Time Fora" = $4, "Fora Gols" = $5, "Data" = $6, "Vencedor" = $7 WHERE "Id" = $8`,
		partida.Rodada, partida.TimeCasa, partida.CasaGols, partida.TimeFora, partida.ForaGols, partida.Data, partida.Vencedor, id)
	return err
}
