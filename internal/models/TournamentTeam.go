package models

type TournamentTeam struct {
	ID           int `json:"id"`
	TournamentID int `json:"tournament_id"`
	TeamID       int `json:"team_id"`
}
