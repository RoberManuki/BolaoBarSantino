package model

type Partida struct {
	Id       int    `json:"id"`
	Rodada   int    `json:"rodada"`
	TimeCasa string `json:"time_casa"`
	CasaGols int    `json:"casa_gols"`
	TimeFora string `json:"time_fora"`
	ForaGols int    `json:"fora_gols"`
	Data     string `json:"data"`
	Vencedor string `json:"vencedor"`
}
