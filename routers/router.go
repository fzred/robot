package routers

import (
	"github.com/astaxie/beego"
	"robot/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/action/send", &controllers.ActionController{}, "*:Send")
}
