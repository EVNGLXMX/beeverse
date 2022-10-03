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

var game = models.GameObject{GameId: "game4", Score: 100, GameName: "portal"}

func TestGamePost(t *testing.T) {
	byteArray, _ := json.Marshal(game)
	r := httptest.NewRequest("POST", "/v1/games/", bytes.NewBuffer(byteArray))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGamePost", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Game Post Response\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestGetGamebyID(t *testing.T) {
	r := httptest.NewRequest("GET", "/v1/games/game4", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
	var responseBody models.GameObject
	json.Unmarshal(w.Body.Bytes(), &responseBody)
	// rBytes, _ := json.Marshal(w.Body.String())
	// fmt.Println(ob)
	Convey("Subject: Test Get By ID Response\n", t, func() {
		Convey("Response Should Resemble Post Game", func() {
			So(responseBody, ShouldResemble, game)
		})
	})
}
