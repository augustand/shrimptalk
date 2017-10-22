package routers

import (
	"github.com/lifeisgo/shrimptalk/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.UserController{}, "get:Register")
	beego.Router("/all", &controllers.MainController{}, "get:All")
	//beego.Router("/:id/talk/:talk", &controllers.TalkController{}, "get:Talk")
	//	beego.Router("/talk/:id", &controllers.TalkController{}, "get:Detail") //s
	beego.Router("/new", &controllers.TalkController{}, "get:New")
	beego.Router("/:id/new", &controllers.TalkController{}, "post:PostNew")
	beego.Router("/login/:id", &controllers.LoginController{})
	beego.Router("/", &controllers.LoginController{}, "get:List")
	beego.Router("/success", &controllers.LoginController{}, "get:Success")
	ns := beego.NewNamespace("/talk",
		//		beego.NSRouter("/user/:user_id([0-9]+)/info", &controllers.MainController{}), //注意逗号不要忘记了
		// 查找所有 list
		beego.NSRouter("/findall", &controllers.TalkController{}, "get:FindMyTalk"),
		beego.NSRouter("/mytalk", &controllers.TalkController{}, "get:FindNowTalk"),
		beego.NSRouter("/:id", &controllers.TalkController{}, "get:Detail"),
		beego.NSRouter("/:id/answer", &controllers.TalkController{}, "get:Answer"),
		beego.NSRouter("/:talkhex/:id/postanswer", &controllers.TalkController{}, "post:PostAnswer"),
	)
	//注册 namespace
	beego.AddNamespace(ns)
}
