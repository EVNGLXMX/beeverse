package mymongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoHandler struct {
	Client     *mongo.Client
	Database   string
	Collection string
}

func NewHandler() *MongoHandler {
	credential := options.Credential{
		Username: beego.AppConfig.String("mongodb::username"),
		Password: beego.AppConfig.String("mongodb::password"),
	}
	url := beego.AppConfig.String("mongodb::url")
	timeout := 300 * time.Second
	opts := options.ClientOptions{ConnectTimeout: &timeout}
	opts.SetDirect(true)
	opts.ApplyURI(url)
	opts.SetAuth(credential)
	timeOut, _ := beego.AppConfig.Int64("mongodb::context_timeout")
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	client, err := mongo.Connect(ctx, &opts)

	if err != nil {
		log.Println(err)
	}
	mh := &MongoHandler{
		Client:     client,
		Database:   beego.AppConfig.String("mongodb::database"),
		Collection: beego.AppConfig.String("mongodb::game_collection"),
	}
	fmt.Print(beego.AppConfig.String("mongodb::database"))
	return mh
}

func (mh *MongoHandler) Disconnect() {
	er := mh.Client.Disconnect(context.TODO()) //close the handler
	if er != nil {
		log.Println(er)
	}
}

// func ping(client *mongo.Client, ctx context.Context) error{

//     // mongo.Client has Ping to ping mongoDB, deadline of
//     // the Ping method will be determined by cxt
//     // Ping method return error if any occurred, then
//     // the error can be handled.
//     mh:= NewHandler()
//     fmt.Println("connected successfully")
//     return nil
// }
