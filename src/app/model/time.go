package model

type Time struct {
	Id      int    `json:"id"`
	Nome    string `json:"nome"`
	Estadio string `json:"estadio"`
	Cidade  string `json:"cidade"`
	Estado  string `json:"estado"`
}
