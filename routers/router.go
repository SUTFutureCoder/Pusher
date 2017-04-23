package routers

import (
	"pusher/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/scm/", &controllers.ScmController{}, "get,post:BindScm")
}
