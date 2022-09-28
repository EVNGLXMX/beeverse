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

var newGame = models.GameObject{GameId: "game4", Score: 100, GameName: "postGame"}
var updateGame = models.GameObject{GameId: "game4", Score: 500, GameName: "updateGame"}

// POST
func TestGamePostCRUD(t *testing.T) {
	byteArray, _ := json.Marshal(newGame)
	r := httptest.NewRequest("POST", "/v1/games/", bytes.NewBuffer(byteArray))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGamePostCRUD", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Game Post Response\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// GET <POST>
func TestGetPostGameCRUD(t *testing.T) {
	r := httptest.NewRequest("GET", "/v1/games/game4", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGetPostGameCRUD", "Code[%d]\n%s", w.Code, w.Body.String())
	var responseBody models.GameObject
	json.Unmarshal(w.Body.Bytes(), &responseBody)

	Convey("Subject: Test Get By ID Response\n", t, func() {
		Convey("Response Should Resemble Post Game", func() {
			So(responseBody, ShouldResemble, newGame)
		})
	})
}

// PUT
func TestPutCRUD(t *testing.T) {
	byteArray, _ := json.Marshal(updateGame)
	r := httptest.NewRequest("PUT", "/v1/games/game4", bytes.NewBuffer(byteArray))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestPutCRUD", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Game Put Response\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
		Convey("Response Should Include Update Success Message", func() {
			So(w.Body.String(), ShouldContainSubstring, "update success!")
		})
	})
}

// GET <PUT>
func TestGetPutGameCRUD(t *testing.T) {
	r := httptest.NewRequest("GET", "/v1/games/game4", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGetPutGameCRUD", "Code[%d]\n%s", w.Code, w.Body.String())
	var responseBody models.GameObject
	json.Unmarshal(w.Body.Bytes(), &responseBody)

	Convey("Subject: Test Get Put Game Response\n", t, func() {
		Convey("Response Should Resemble Put game", func() {
			So(responseBody, ShouldResemble, updateGame)
			// So(responseBody, ShouldNotResemble, game1)
		})
	})
}

// DELETE
func TestDeleteCRUD(t *testing.T) {
	r := httptest.NewRequest("DELETE", "/v1/games/game4", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestDeleteCRUD", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Delete\n", t, func() {
		Convey("Response Should Include Delete Success Message", func() {
			So(w.Body.String(), ShouldContainSubstring, "delete success!")
		})
	})
}

// GET <DELETE>
func TestGetDeleteCRUD(t *testing.T) {
	r := httptest.NewRequest("GET", "/v1/games/game4", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGetDeleteCRUD", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Get Deleted Game\n", t, func() {
		Convey("Statue Code Should Be 404 ", func() {
			// So(w.Code, ShouldEqual, 200)
			So(w.Code, ShouldEqual, 404)
		})
	})
}
