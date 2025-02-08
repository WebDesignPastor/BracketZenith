package models

type TournamentPlayer struct {
	ID           int `json:"id"`
	TournamentID int `json:"tournament_id"`
	PlayerID     int `json:"player_id"`
}
