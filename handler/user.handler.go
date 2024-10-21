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

func Login(ctx *fiber.Ctx) error {
	// Check Available Username
	var user entity.User
	

	var Loginuser entity.LoginUser
	// รับข้อมูลจาก request body (JSON)
	if err := ctx.BodyParser(&Loginuser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ไม่สามารถรับข้อมูลได้",
		})
	}

	err := database.MYSQL.Debug().Table("User").Find(&user, "user_email = ?", Loginuser.User_email).Error
	if err != nil || user.User_email == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "อีเมลไม่ถูกต้อง",
		})
	}

	if err != nil || user.User_password != "" {
		if user.User_password != Loginuser.User_password {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "รหัสผ่านไม่ถูกต้อง",
			})
		} else {
			return ctx.Status(fiber.StatusUnauthorized).JSON(user)
		}
	}

	return nil
}

func LoginDriver(ctx *fiber.Ctx) error {
	// Check Available Username

	var LoginDriver entity.LoginDriver
	// รับข้อมูลจาก request body (JSON)
	if err := ctx.BodyParser(&LoginDriver); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ไม่สามารถรับข้อมูลได้",
		})
	}


	var Driver entity.Driver

	err := database.MYSQL.Debug().Table("Raiders").Find(&Driver, "raider_email = ?", LoginDriver.Raider_email).Error

	if err != nil || Driver.Raider_email == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong raider_email ",
		})
	}

	if err != nil || Driver.Raider_password != "" {

		if Driver.Raider_password != LoginDriver.Raider_password {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "wrong raider_password",
			})
		} else {
			return ctx.Status(fiber.StatusUnauthorized).JSON(Driver)
		}
	}

	return nil
}

// get
func GetUsers(ctx *fiber.Ctx) error {
	var user []entity.User

	database.MYSQL.Debug().Table("User").Find(&user)
	ctx.JSON(user)

	return ctx.JSON(user)
}

func GetDriver(ctx *fiber.Ctx) error {
	var Driver []entity.Driver

	database.MYSQL.Debug().Table("Raiders").Find(&Driver)
	ctx.JSON(Driver)

	return ctx.JSON(Driver)
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
