package mymongo

import (
	"context"
	"log"
	"time"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type metaData struct {
	ServerStatus bool  `json:"server_status" bson:"server_status"`
	PlayerCount  int64 `json:"player_count" bson:"player_count"`
	Updated      bool  `json:"updated" bson:"updated"`
}

type Game struct {
	GameID      string    `json:"game_id" bson:"game_id"`
	Genres      []string  `json:"genres" bson:"genres"`
	ReleaseDate time.Time `json:"release_date" bson:"release_date"`
	CreatedOn   time.Time `json:"created_on" bson:"created_on"`
	UpdatedOn   time.Time `json:"updated_on" bson:"updated_on"`
	MetaData    metaData  `json:"metadata" bson:"metadata"`
}

func (game *Game) GetAllGames(filter interface{}) (*[]Game, error) {
	timeOut, _ := beego.AppConfig.Int64("mongodb::context_timeout")
	var games []Game
	mh := NewHandler()
	collection := mh.Client.Database(mh.Database).Collection(beego.AppConfig.String("mongodb::game_collection"))
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return &games, err
	}
	for cur.Next(ctx) {
		game := &Game{}
		er := cur.Decode(game)
		if er != nil {
			log.Print(er)
		}
		games = append(games, *game)
	}

	mh.Disconnect()
	return &games, err
}

func (game *Game) GetGameByFilter(filter interface{}) error {
	timeOut, _ := beego.AppConfig.Int64("mongodb::context_timeout")
	mh := NewHandler()
	collection := mh.Client.Database(mh.Database).Collection(beego.AppConfig.String("mongodb::game_collection"))
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	err := collection.FindOne(ctx, filter).Decode(game)
	mh.Disconnect()
	return err
}

func (game *Game) InsertOne() (*mongo.InsertOneResult, error) {
	game.CreatedOn = time.Now()
	game.UpdatedOn = time.Now()

	timeOut, _ := beego.AppConfig.Int64("mongodb::context_timeout")
	mh := NewHandler()
	collection := mh.Client.Database(mh.Database).Collection(beego.AppConfig.String("mongodb::game_collection"))
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	result, err := collection.InsertOne(ctx, game)
	mh.Disconnect()
	return result, err
}

func (game *Game) UpdateOne(filter interface{}) (*mongo.UpdateResult, error) {
	timeOut, _ := beego.AppConfig.Int64("mongodb::context_timeout")
	mh := NewHandler()
	collection := mh.Client.Database(mh.Database).Collection(beego.AppConfig.String("mongodb::game_collection"))
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": game})
	mh.Disconnect()
	return result, err
}

func (game *Game) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {
	timeOut, _ := beego.AppConfig.Int64("mongodb::context_timeout")
	mh := NewHandler()
	collection := mh.Client.Database(mh.Database).Collection(beego.AppConfig.String("mongodb::game_collection"))
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	result, err := collection.DeleteOne(ctx, filter)
	mh.Disconnect()
	return result, err
}
