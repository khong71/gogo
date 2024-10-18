package route

import (
	"gogo/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Get("GetUsers", handler.GetUsers)

	r.Get("GetUserid", handler.GetUser_id)

	r.Delete("DeleteUserAll", handler.DeleteUserAll)

	//Register
	r.Post("Register", handler.Register)
	r.Post("RegisterDriver", handler.RegisterDriver)

	//Login

	//get
	r.Get("GetDriver", handler.GetDriver)

}
