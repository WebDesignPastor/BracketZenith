package models

type TeamPlayer struct {
	ID       int `json:"id"`
	TeamID   int `json:"team_id"`
	PlayerID int `json:"player_id"`
}
