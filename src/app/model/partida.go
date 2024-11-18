package model

type Partida struct {
	Id       int    `json:"id"`
	Rodada   int    `json:"rodada"`
	TimeCasa int    `json:"time_casa"`
	CasaGols *int   `json:"casa_gols"`
	TimeFora int    `json:"time_fora"`
	ForaGols *int   `json:"fora_gols"`
	Data     string `json:"data"`
	Vencedor string `json:"vencedor"`
}
