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

func TestPut(t *testing.T) {
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
