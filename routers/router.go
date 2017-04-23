package routers

import (
	"pusher/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/scm/", &controllers.ScmController{}, "get,post:BindScm")

	beego.Router("/oauth/", &controllers.OauthController{}, "get:Oauth")
	beego.Router("/oauth/echo", &controllers.OauthController{}, "get:OauthEcho")
}
