package test

import (
	_ "beeverse/routers"
	"net/http"
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

var gameID = "game1"

func TestGetId(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/games/", nil)
	res := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(res, req)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", res.Code, res.Body.String())

	Convey("Subject: Test Get By ID Response\n", t, func() {
		Convey("Response Should Include Game ID", func() {
			So(res.Body.String(), ShouldContainSubstring, "Dota2")
		})
	})
}
