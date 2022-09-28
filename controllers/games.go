package controllers

import (
	"beeverse/models/helper"
	"beeverse/models/mymongo"
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
)

// Operations about games
type GamesController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.GameObject
// @Failure 403 :gameId is empty
// @router / [get]
func (controller *GamesController) GetAll() {
	result, err := helper.GetAllGames()
	if err != nil {
		controller.Ctx.ResponseWriter.WriteHeader(500)
		controller.Data["json"] = map[string]string{"message": string(err.Error())}
	} else {
		if result != nil {
			controller.Data["json"] = result
		} else {
			controller.Ctx.ResponseWriter.WriteHeader(404)
			controller.Data["json"] = map[string]string{"message": "data not found"}
		}
	}
	controller.ServeJSON()
}

// @Title Get
// @Description find games by gameId
// @Param	gameId		path 	string	true		"the gameId you want to get"
// @Success 200  {object} models.GameObject
// @Failure 403 :gameId is empty
// @Failure 404 :game does not exist
// @router /:gameId [get]
func (controller *GamesController) GetGameByID() {
	gameId := controller.Ctx.Input.Param(":gameId")
	if gameId != "" {
		result, err := helper.GetGameByID(gameId)
		if err != nil {
			controller.Data["json"] = map[string]string{"error": err.Error()}
			controller.Ctx.ResponseWriter.WriteHeader(404)
			// controller.CustomAbort(404, "Game does not exist")
		} else {
			controller.Data["json"] = result
		}
	}
	controller.ServeJSON()
}

// @Title Create
// @Description create games
// @Param	body		body 	models.GameObject	true		"The games content"
// @Success 200 {string} models.GameObject.gameId
// @Failure 403 body is empty
// @router / [post]
func (controller *GamesController) Post() {
	var game mymongo.Game
	json.Unmarshal(controller.Ctx.Input.RequestBody, &game)

	if game.GameID != "" {
		_, err := helper.GetGameByID(game.GameID)
		if err == nil {
			controller.Data["json"] = map[string]string{"error": "game already exists"}
			controller.Ctx.ResponseWriter.WriteHeader(409)
			controller.ServeJSON()
			controller.Abort("")
		}
	}

	_, err := helper.InsertGame(game)
	if err != nil {
		log.Println(err)
	}
	controller.Data["json"] = map[string]string{"message": "create success"}
	controller.Ctx.ResponseWriter.WriteHeader(201)
	controller.ServeJSON()
}

// @Title Update
// @Description update the games
// @Param	gameId		path 	string	true		"The gameId you want to update"
// @Param	body		body 	models.GameObject	true		"The body"
// @Success 201  {object} models.GameObject
// @Failure 403 :gameId is empty
// @router /:gameId [put]
func (controller *GamesController) Put() {
	gameId := controller.Ctx.Input.Param(":gameId")
	if gameId != "" {
		_, err := helper.GetGameByID(gameId)
		if err != nil {
			controller.Data["json"] = map[string]string{"error": err.Error()}
			controller.Ctx.ResponseWriter.WriteHeader(404)
			controller.ServeJSON()
			controller.Abort("")
		}
	}
	var game mymongo.Game
	json.Unmarshal(controller.Ctx.Input.RequestBody, &game)

	_, err := helper.UpdateGame(gameId, game)
	if err != nil {
		controller.Data["json"] = err.Error()
	} else {
		controller.Ctx.ResponseWriter.WriteHeader(201)
		controller.Data["json"] = map[string]string{"message": "update success!"}
	}
	controller.ServeJSON()
}

// @Title Delete
// @Description delete the games
// @Param	gameId		path 	string	true		"The gameId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 gameId is empty
// @router /:gameId [delete]
func (controller *GamesController) Delete() {
	var result *mongo.DeleteResult
	var err error
	gameId := controller.Ctx.Input.Param(":gameId")
	result, err = helper.DeleteGame(gameId)

	if err != nil {
		controller.Ctx.ResponseWriter.WriteHeader(404)
		controller.Data["json"] = map[string]string{"message": "game does not exist"}
	} else if result.DeletedCount == 0 {
		controller.Ctx.ResponseWriter.WriteHeader(404)
		controller.Data["json"] = result
	} else {
		controller.Data["json"] = map[string]string{"message": "delete success"}
	}
	controller.ServeJSON()
}
