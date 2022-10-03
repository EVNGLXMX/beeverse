package test

import (
	"beeverse/models/mymongo"
	_ "beeverse/routers"
	"beeverse/tests/testvars"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestGetGame(t *testing.T) {
	r := httptest.NewRequest("GET", "/v1/games/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGET", "Code[%d]\n%s", w.Code, w.Body.String())
	// var responseBody mymongo.Game
	// json.Unmarshal(w.Body.Bytes(), &responseBody)

	Convey("Subject: Test Get All Games Response\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Response Should Include All Games", func() {
			So(w.Body.String(), ShouldContainSubstring, "dota2")
			So(w.Body.String(), ShouldContainSubstring, "god of war")
			//So(w.Body.String(), ShouldContainSubstring, "testgame")
		})
	})
}

func TestPostGame(t *testing.T) {
	var testGame = testvars.TestGame
	byteArray, _ := json.Marshal(testGame)
	r := httptest.NewRequest("POST", "/v1/games/", bytes.NewBuffer(byteArray))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestPOST", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Game Post Response\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
		Convey("The Result Should Contain Success Message", func() {
			So(w.Body.String(), ShouldContainSubstring, "success")
		})
	})
}

func TestPostDuplicateGame(t *testing.T) {
	var testGame = testvars.TestGame
	byteArray, _ := json.Marshal(testGame)
	r := httptest.NewRequest("POST", "/v1/games/", bytes.NewBuffer(byteArray))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestPost", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Game Post Response 2\n", t, func() {
		Convey("Status Code Should Be 409", func() {
			So(w.Code, ShouldEqual, 409)
		})
		Convey("The Result Should Contain Error Message", func() {
			So(w.Body.String(), ShouldContainSubstring, "error")
		})
	})
}

// TODO : FIX UpdatedOn NOT MATCHING

// func TestGetGameByID(t *testing.T) {
// 	var testGame = testvars.TestGame
// 	r := httptest.NewRequest("GET", "/v1/games/"+testGame.GameID, nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)
// 	beego.Trace("testing", "TestGET", "Code[%d]\n%s", w.Code, w.Body.String())
// 	var responseBody mymongo.Game
// 	json.Unmarshal(w.Body.Bytes(), &responseBody)

// 	Convey("Subject: Test Get All Games Response\n", t, func() {
// 		Convey("Status Code Should Be 200", func() {
// 			So(w.Code, ShouldEqual, 200)
// 		})
// 		Convey("Response Should Include TestGame", func() {
// 			So(responseBody, ShouldResemble, testGame)
// 		})
// 	})
// }

func TestPutGame(t *testing.T) {
	var updateGame = testvars.TestGame
	var updateData = mymongo.MetaData{Updated: true, ServerStatus: true}
	updateGame.MetaData = updateData
	updateGame.UpdatedOn = time.Now()
	byteArray, _ := json.Marshal(updateGame)
	r := httptest.NewRequest("PUT", "/v1/games/"+updateGame.GameID, bytes.NewBuffer(byteArray))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestPut", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Game Put/Update Response\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestDeleteGame(t *testing.T) {
	var deleteGame = testvars.TestGame
	r := httptest.NewRequest("DELETE", "/v1/games/"+deleteGame.GameID, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestDelete", "Code[%d]\n%s", w.Code, w.Body.String())
	Convey("Subject: Test Delete\n", t, func() {
		Convey("Response Should Include Delete Success Message", func() {
			So(w.Body.String(), ShouldContainSubstring, "delete success")
		})
	})
}
