// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beeverse/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/games",
			beego.NSInclude(
				&controllers.GamesController{},
			),
			beego.NSRouter("/", &controllers.GamesController{}, "get:GetAll"),
			beego.NSRouter("/", &controllers.GamesController{}, "post:Post"),
			beego.NSRouter("/:gameId", &controllers.GamesController{}, "get:GetGameByID"),
			beego.NSRouter("/:gameId", &controllers.GamesController{}, "put:Put"),
			beego.NSRouter("/:gameId", &controllers.GamesController{}, "delete:Delete"),
		),
		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.TestController{},
			),
			beego.NSRouter("/hello", &controllers.TestController{}, "get:Test"),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
			beego.NSRouter("/", &controllers.UserController{}, "get:GetAll"),
			beego.NSRouter("/", &controllers.UserController{}, "post:Post"),
			beego.NSRouter("/login", &controllers.UserController{}, "get:Login"),
			beego.NSRouter("/logout", &controllers.UserController{}, "get:Logout"),
			beego.NSRouter("/:uid", &controllers.UserController{}, "get:Get"),
			beego.NSRouter("/:uid", &controllers.UserController{}, "put:Put"),
			beego.NSRouter("/:uid", &controllers.UserController{}, "delete:Put"),
		),
	)
	beego.AddNamespace(ns)
}
