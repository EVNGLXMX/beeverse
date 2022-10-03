package test

import (
	_ "beeverse/routers"
	"beeverse/tests/testvars"
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

var testGame = testvars.TestGame

func TestPost(t *testing.T) {
	byteArray, _ := json.Marshal(testGame)
	r := httptest.NewRequest("POST", "/v1/games/", bytes.NewBuffer(byteArray))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestPost", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Game Post Response\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
		Convey("The Result Should Contain Success Message", func() {
			So(w.Body.String(), ShouldContainSubstring, "success")
		})
	})
}

func TestPostDuplicate(t *testing.T) {
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
