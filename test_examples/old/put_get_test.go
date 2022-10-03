package test

import (
	"beeverse/models"
	_ "beeverse/routers"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

var game1 = models.GameObject{GameId: "game1", Score: 105, GameName: "putgame"}

func TestPut(t *testing.T) {
	byteArray, _ := json.Marshal(game1)
	r := httptest.NewRequest("PUT", "/v1/games/game1", bytes.NewBuffer(byteArray))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGamePut", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Game Put Response\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestGetUpdatedGame(t *testing.T) {
	r := httptest.NewRequest("GET", "/v1/games/game1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGetUpdatedGame", "Code[%d]\n%s", w.Code, w.Body.String())
	var responseBody models.GameObject
	json.Unmarshal(w.Body.Bytes(), &responseBody)

	Convey("Subject: Test Get Put Game Response\n", t, func() {
		Convey("Response Should Resemble Updated Game JSON", func() {
			So(responseBody, ShouldResemble, game1)
			// So(responseBody, ShouldNotResemble, game1)
		})
	})
}
