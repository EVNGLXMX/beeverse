package models

type GameObject struct {
	GameId   string `json:"game_id"`
	Score    int64  `json:"score"`
	GameName string `json:"game_name"`
}
