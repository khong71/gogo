package handler

import (
	"gogo/database"
	"gogo/model/entity"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
//Login
// ฟังก์ชันสำหรับ Login ผู้ใช้ (User)
// LoginUser ใช้สำหรับเข้าสู่ระบบผู้ใช้
func LoginUser(ctx *fiber.Ctx) error {
    var loginRequest entity.LoginUser

    // รับข้อมูลจาก request body
    if err := ctx.BodyParser(&loginRequest); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "ไม่สามารถรับข้อมูลได้",
        })
    }

    // ค้นหาผู้ใช้ในฐานข้อมูล
    var user entity.User
    if result := database.MYSQL.Debug().Table("User").Where("user_email = ?", loginRequest.User_email).First(&user); result.Error != nil {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "อีเมลหรือรหัสผ่านไม่ถูกต้อง",
        })
    }

    // ตรวจสอบรหัสผ่าน
    if err := bcrypt.CompareHashAndPassword([]byte(user.User_password), []byte(loginRequest.User_password)); err != nil {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "อีเมลหรือรหัสผ่านไม่ถูกต้อง",
        })
    }

    // ส่งข้อความสำเร็จเมื่อ login ผ่าน
    return ctx.JSON(fiber.Map{
        "message": "เข้าสู่ระบบสำเร็จ",
        "user": fiber.Map{
            "user_email": user.User_email, // ส่งกลับอีเมลผู้ใช้ถ้าจำเป็น
        },
    })
}

func LoginDriver(ctx *fiber.Ctx) error {
	var loginRequest entity.LoginDriver

	// รับข้อมูลจาก request body
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ไม่สามารถรับข้อมูลได้",
		})
	}

	// ค้นหาผู้ขับขี่ในฐานข้อมูล
	var driver entity.Driver
	if result := database.MYSQL.Debug().Table("Driver").Where("raider_email = ?", loginRequest.Raider_email).First(&driver); result.Error != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "อีเมลหรือรหัสผ่านไม่ถูกต้อง",
		})
	}

	// ตรวจสอบรหัสผ่าน
	if err := bcrypt.CompareHashAndPassword([]byte(driver.Raider_password), []byte(loginRequest.Raider_password)); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "อีเมลหรือรหัสผ่านไม่ถูกต้อง",
		})
	}

	// ส่งข้อความสำเร็จเมื่อ login ผ่าน
	return ctx.JSON(fiber.Map{
		"message": "เข้าสู่ระบบสำเร็จ (Driver)",
	})
}
//-------------------------------------------------------------------------------------------------------


//get
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
