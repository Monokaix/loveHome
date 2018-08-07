package routers

import (
	"myproject/loveHome/controllers"
	"github.com/astaxie/beego"
)

func init(){
	beego.Router("/",&controllers.AreaController{})
	beego.Router("api/v1.0/areas",&controllers.AreaController{},"get:GetArea")
	///api/v1.0/users 注册
	beego.Router("/api/v1.0/users", &controllers.UserController{},"post:Reg")
	//api/v1.0/session
	beego.Router("/api/v1.0/session", &controllers.SessionController{},"get:GetSessionData;delete:DeleteSessionData")
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{},"get:GetHouseIndex")
}
