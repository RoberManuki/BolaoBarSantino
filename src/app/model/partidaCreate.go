package model

// PartidaCreate representa os dados necessários para criar uma partida
type PartidaCreate struct {
	Rodada   int    `json:"rodada"`
	TimeCasa int    `json:"time_casa"`
	CasaGols int    `json:"casa_gols"`
	TimeFora int    `json:"time_fora"`
	ForaGols int    `json:"fora_gols"`
	Data     string `json:"data"`
	Vencedor string `json:"vencedor"`
}
