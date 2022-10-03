package test

import (
	//"beeverse/models/mymongo"
	_ "beeverse/routers"
	//"encoding/json"
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

func TestGet(t *testing.T) {
	r := httptest.NewRequest("GET", "/v1/games/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGetPostGameCRUD", "Code[%d]\n%s", w.Code, w.Body.String())
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
