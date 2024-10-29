package repository

import (
	"database/sql"
	//"fmt"
	"bolao/src/app/model"

	_ "github.com/lib/pq"
)

var dbTimes *sql.DB

func init() {
	var err error
	dbTimes, err = sql.Open("postgres", "user=postgres password=sjtbmix6 dbname=santino-bolao sslmode=disable")

	if err != nil {
		panic(err)
	}
}

func GetTimes() ([]model.Time, error) {
	query := `
        SELECT "Id", "Nome", "Estádio", "Cidade", "Estado"
        FROM "schema-bolao-24"."Time"
		ORDER BY "Id"
    `

	rows, err := dbTimes.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var times []model.Time
	for rows.Next() {
		var p model.Time

		if err := rows.Scan(
			&p.Id,
			&p.Nome,
			&p.Estadio,
			&p.Cidade,
			&p.Estado,
		); err != nil {
			return nil, err
		}

		times = append(times, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return times, nil
}
