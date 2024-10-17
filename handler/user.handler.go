package handler

import (
	"gogo/database"
	"gogo/model/entity"

	"github.com/gofiber/fiber/v2"
)

// Register

func Register(ctx *fiber.Ctx) error {
	var Register entity.Register

	// รับข้อมูลจาก request body
	if err := ctx.BodyParser(&Register); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ไม่สามารถรับข้อมูลได้",
		})
	}

	// บันทึกข้อมูลผู้ใช้ใหม่ลงในตาราง User
	if result := database.MYSQL.Debug().Table("User").Create(&Register); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "ไม่สามารถเพิ่มข้อมูลได้",
		})
	}

	// ส่งข้อมูลผู้ใช้ที่ถูกเพิ่มกลับไป
	return ctx.JSON(fiber.Map{
		"message": "เพิ่มผู้ใช้สำเร็จ",
	})
}

func RegisterDriver(ctx *fiber.Ctx) error {
	var RegisterDriver entity.RegisterDriver

	// รับข้อมูลจาก request body
	if err := ctx.BodyParser(&RegisterDriver); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ไม่สามารถรับข้อมูลได้",
		})
	}

	// บันทึกข้อมูลผู้ใช้ใหม่ลงในตาราง User
	if result := database.MYSQL.Debug().Table("Raiders").Create(&RegisterDriver); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "ไม่สามารถเพิ่มข้อมูลได้",
		})
	}

	// ส่งข้อมูลผู้ใช้ที่ถูกเพิ่มกลับไป
	return ctx.JSON(fiber.Map{
		"message": "เพิ่มผู้ใช้สำเร็จ",
	})
}

//-------------------------------------------------------------------------------------------------------

func GetUsers(ctx *fiber.Ctx) error {
	var user []entity.User

	database.MYSQL.Debug().Table("User").Find(&user)
	ctx.JSON(user)

	return ctx.JSON(user)
}

func GetUser_id(ctx *fiber.Ctx) error {
	var idx = ctx.Query("id")
	var user []entity.User

	database.MYSQL.Debug().Table("User").Where(idx).Find(&user)
	ctx.JSON(user)

	return ctx.JSON(user)
}

func DeleteUserAll(ctx *fiber.Ctx) error {
	// ลบข้อมูลผู้ใช้ทั้งหมดจากตาราง User โดยใช้ SQL ตรง
	database.MYSQL.Debug().Exec("DELETE FROM `User`")

	// ส่งข้อความว่าลบสำเร็จ
	return ctx.JSON(fiber.Map{
		"message": "ลบผู้ใช้ทั้งหมดสำเร็จ",
	})
}
