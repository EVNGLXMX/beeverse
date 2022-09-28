package test

import (
	_ "beeverse/routers"
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
func TestDelete(t *testing.T) {
	r := httptest.NewRequest("DELETE", "/v1/games/game1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestDelete", "Code[%d]\n%s", w.Code, w.Body.String())
	Convey("Subject: Test Delete\n", t, func() {
		Convey("Response Should Include Delete Success Message", func() {
			So(w.Body.String(), ShouldContainSubstring, "delete success!")
		})
	})
}
func TestGetDeletedGame(t *testing.T) {
	r := httptest.NewRequest("GET", "/v1/games/game1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestGetDeletedGame", "Code[%d]\n%s", w.Code, w.Body.String())
	Convey("Subject: Test Get Deleted Game\n", t, func() {
		Convey("Statue Code Should Be 404 ", func() {
			// So(w.Code, ShouldEqual, 200)
			So(w.Code, ShouldEqual, 404)
		})
	})
}
