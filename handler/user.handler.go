package handler

import (
	"fmt"
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

	// ตรวจสอบว่ามีรหัสผ่านหรือไม่
	if Register.User_password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "กรุณากรอกรหัสผ่าน",
		})
	}

	// ตรวจสอบว่าเบอร์โทรซ้ำหรือไม่
	var existingUser entity.Register
	if err := database.MYSQL.Debug().Table("User").Where("user_phone = ?", Register.User_Phone).First(&existingUser).Error; err == nil {
		// ถ้าพบเบอร์โทรซ้ำ ให้แจ้งว่าหมายเลขโทรศัพท์ซ้ำ
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "หมายเลขโทรศัพท์นี้มีการลงทะเบียนแล้ว",
		})
	}

	// แฮชรหัสผ่าน
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Register.User_password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "ไม่สามารถแฮชรหัสผ่านได้",
		})
	}

	// เก็บรหัสผ่านที่แฮชแล้วลงใน struct
	Register.User_password = string(hashedPassword)

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

	// ตรวจสอบว่ามีรหัสผ่านหรือไม่
	if RegisterDriver.Raider_password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "กรุณากรอกรหัสผ่าน",
		})
	}

	// ตรวจสอบว่าเบอร์โทรซ้ำหรือไม่
	var existingUser entity.RegisterDriver
	if err := database.MYSQL.Debug().Table("Raiders").Where("raider_phone = ?", RegisterDriver.Raider_Phone).First(&existingUser).Error; err == nil {
		// ถ้าพบเบอร์โทรซ้ำ ให้แจ้งว่าหมายเลขโทรศัพท์ซ้ำ
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "หมายเลขโทรศัพท์นี้มีการลงทะเบียนแล้ว",
		})
	}

	// แฮชรหัสผ่าน
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(RegisterDriver.Raider_password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "ไม่สามารถแฮชรหัสผ่านได้",
		})
	}

	// เก็บรหัสผ่านที่แฮชแล้วลงใน struct
	RegisterDriver.Raider_password = string(hashedPassword)

	// บันทึกข้อมูลผู้ใช้ใหม่ลงในตาราง Raiders
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

// Login
func Login(ctx *fiber.Ctx) error {
	var user entity.User
	var Loginuser entity.LoginUser

	// รับข้อมูลจาก request body (JSON)
	if err := ctx.BodyParser(&Loginuser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ไม่สามารถรับข้อมูลได้",
		})
	}

	// ตรวจสอบว่ามีผู้ใช้ที่ตรงกับอีเมลในฐานข้อมูลหรือไม่
	err := database.MYSQL.Debug().Table("User").Where("user_email = ?", Loginuser.User_email).First(&user).Error
	if err != nil {
		fmt.Println("Error retrieving user:", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "อีเมลไม่ถูกต้อง",
		})
	}

	fmt.Printf("Retrieved user: %+v\n", user)

	// เปรียบเทียบรหัสผ่าน
	err = bcrypt.CompareHashAndPassword([]byte(user.User_password), []byte(Loginuser.User_password))
	if err != nil {
		fmt.Println("Password comparison failed:", err)
		fmt.Printf("Stored password: %s\n", user.User_password) // แสดงรหัสผ่านที่เก็บในฐานข้อมูล
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "รหัสผ่านไม่ถูกต้อง",
		})
	}

	// หากเข้าสู่ระบบสำเร็จ
	return ctx.JSON(fiber.Map{
		"message": "เข้าสู่ระบบสำเร็จ",
		"user":    user,
	})
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

	// ตรวจสอบว่ามีผู้ขับขี่ที่ตรงกับอีเมลในฐานข้อมูลหรือไม่
	err := database.MYSQL.Debug().Table("Raiders").Where("raider_email = ?", LoginDriver.Raider_email).First(&Driver).Error

	if err != nil || Driver.Raider_email == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "อีเมลผู้ขับขี่ไม่ถูกต้อง",
		})
	}

	// เปรียบเทียบรหัสผ่าน
	err = bcrypt.CompareHashAndPassword([]byte(Driver.Raider_password), []byte(LoginDriver.Raider_password))
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "รหัสผ่านผู้ขับขี่ไม่ถูกต้อง",
		})
	}

	// หากเข้าสู่ระบบสำเร็จ
	return ctx.JSON(fiber.Map{
		"message": "เข้าสู่ระบบสำเร็จ",
		"driver":  Driver,
	})
}

//--------------------------------------------------------------------------------------------------------

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

func InsertOrder(ctx *fiber.Ctx) error {
	var order entity.InsertOrder

	// Parse JSON body
	if err := ctx.BodyParser(&order); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// ทำการ insert ข้อมูลลงฐานข้อมูลโดยใช้ GORM
	if err := database.MYSQL.Debug().Table("Order").Create(&order).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert order",
		})
	}

	// ส่ง response เมื่อ insert สำเร็จ
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order inserted successfully",
		"order":   order,
	})
}

// Delete
func DeleteUserAll(ctx *fiber.Ctx) error {
	// ลบข้อมูลผู้ใช้ทั้งหมดจากตาราง User โดยใช้ SQL ตรง
	database.MYSQL.Debug().Exec("DELETE FROM `User`")

	// ส่งข้อความว่าลบสำเร็จ
	return ctx.JSON(fiber.Map{
		"message": "ลบผู้ใช้ทั้งหมดสำเร็จ",
	})
}

func DeleteRaiderAll(ctx *fiber.Ctx) error {
	// ลบข้อมูลผู้ใช้ทั้งหมดจากตาราง User โดยใช้ SQL ตรง
	database.MYSQL.Debug().Exec("DELETE FROM `Raiders`")

	// ส่งข้อความว่าลบสำเร็จ
	return ctx.JSON(fiber.Map{
		"message": "ลบผู้ใช้ทั้งหมดสำเร็จ",
	})
}

func DeleteOrderAll(ctx *fiber.Ctx) error {
	// ลบข้อมูลผู้ใช้ทั้งหมดจากตาราง User โดยใช้ SQL ตรง
	database.MYSQL.Debug().Exec("DELETE FROM `Order`")

	// ส่งข้อความว่าลบสำเร็จ
	return ctx.JSON(fiber.Map{
		"message": "ลบorderทั้งหมดสำเร็จ",
	})
}
