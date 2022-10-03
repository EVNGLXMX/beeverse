package test

import (
	_ "beeverse/routers"
	"beeverse/tests/testvars"
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
