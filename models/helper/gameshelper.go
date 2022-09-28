package helper

import (
	"beeverse/models/mymongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllGames() (map[string]interface{}, error) {
	var game mymongo.Game
	var error error
	games, error := game.GetAllGames(bson.M{})
	if len(*games) > 0 {
		result := map[string]interface{}{
			"games": games,
		}
		return result, error
	}
	return nil, error
}

func GetGameByID(id string) (*mymongo.Game, error) {
	var game mymongo.Game
	err := game.GetGameByFilter(bson.M{"game_id": id})
	return &game, err
}

func InsertGame(game mymongo.Game) (*mongo.InsertOneResult, error) {
	result, error := game.InsertOne()
	return result, error
}

func UpdateGame(id string, game mymongo.Game) (*mongo.UpdateResult, error) {
	var updateResult *mongo.UpdateResult
	updateResult, err := game.UpdateOne(bson.M{"game_id": id})
	return updateResult, err
}

func DeleteGame(id string) (*mongo.DeleteResult, error) {
	var game mymongo.Game
	result, err := game.DeleteOne(bson.M{"game_id": id})
	return result, err
}
