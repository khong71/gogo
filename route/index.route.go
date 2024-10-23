package route

import (
	"gogo/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Get("GetUsers", handler.GetUsers)

	r.Get("GetUserid", handler.GetUser_id)

	r.Delete("DeleteUserAll", handler.DeleteUserAll)
	r.Delete("DeleteRaiderAll", handler.DeleteRaiderAll)
	r.Delete("DeleteOrderAll", handler.DeleteOrderAll)

	//Register
	r.Post("Register", handler.Register)
	r.Post("RegisterDriver", handler.RegisterDriver)


	r.Post("insertOrder", handler.InsertOrder)
	
	//Login
	r.Post("LoginUser", handler.Login)
	r.Post("LoginDriver", handler.LoginDriver)
	//get
	r.Get("GetDriver", handler.GetDriver)

	


	

}
