package models

import (
	"errors"
)

var (
	Games map[string]*GameObject
)

type GameObject struct {
	GameId   string `json:"game_id"`
	Score    int64  `json:"score"`
	GameName string `json:"game_name"`
}

func init() {
	Games = make(map[string]*GameObject)
	Games["game1"] = &GameObject{"game1", 100, "mario"}
	Games["game2"] = &GameObject{"game2", 101, "crash"}
}

func InsertOne(object GameObject) (game GameObject) {
	// object.GameId = object.GameId
	Games[object.GameId] = &object
	return object
}

func GetOne(GameId string) (object *GameObject, err error) {
	for i, v := range Games {
		if Games[i].GameId == GameId {
			return v, nil
		}
	}
	return nil, errors.New("GameId Does Not Exist")
}

func GetAll() map[string]*GameObject {

	return Games
}

func Update(GameId string, Score int64, GameName string) (gameOb *GameObject, err error) {
	if v, ok := Games[GameId]; ok {
		v.Score = Score
		v.GameName = GameName
		return v, nil
	}
	return nil, errors.New("GameId Does Not Exist")
}

func Delete(GameId string) {
	delete(Games, GameId)
}
