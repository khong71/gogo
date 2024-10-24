package handler

import (
	"fmt"
	"gogo/database"
	"gogo/model/entity"
	"log"

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

func GetRaider_id(ctx *fiber.Ctx) error {
	var idx = ctx.Query("id")
	var Driver []entity.Driver

	database.MYSQL.Debug().Table("Raiders").Where(idx).Find(&Driver)
	ctx.JSON(Driver)

	return ctx.JSON(Driver)
}

func GetOrders(ctx *fiber.Ctx) error {
	var orders []entity.GetOrder

	// ควรใช้ slice เนื่องจากดึงข้อมูลหลายแถว
	database.MYSQL.Debug().Table("Order").Find(&orders)
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	// ส่งออกข้อมูลในรูปแบบ JSON
	return ctx.JSON(orders)
}

func GetOrdersreceiverList(ctx *fiber.Ctx) error {
	// รับค่า id จาก query parameter
	id := ctx.Query("id")
	var orders []struct {
		OrderID         uint    `json:"order_id"`          // ID ของ Order
		OrderInfo       *string `json:"order_info"`        // ข้อมูลเกี่ยวกับ Order
		OrderImage      *string `json:"order_image"`       // รูปภาพของ Order
		OrderSenderID   uint    `json:"order_sender_id"`   // ID ของผู้ส่ง
		OrderReceiverID uint    `json:"order_receiver_id"` // ID ของผู้รับ
		User_name       string  `json:"user_name"`         // ชื่อผู้ใช้
		User_Phone      string  `json:"user_phone"`
		User_image      string  `json:"user_image"`
	}

	// ตรวจสอบว่าค่า id มีอยู่หรือไม่
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing 'id' query parameter",
		})
	}

	// ใช้ GORM เพื่อทำการ JOIN ตาราง Order และ User
	result := database.MYSQL.Debug().
		Table("Order").
		Select("*").
		Joins("JOIN User ON Order.order_receiver_id = User.user_id"). // กำหนดเงื่อนไขการ JOIN
		Where("Order.order_receiver_id = ?", id).                     // กำหนดเงื่อนไขการค้นหาข้อมูล
		Find(&orders)

	// ตรวจสอบว่ามีข้อผิดพลาดหรือไม่
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve orders",
		})
	}

	// ส่งข้อมูลในรูปแบบ JSON
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	return ctx.JSON(orders)
}

func GetOrdersSendList(ctx *fiber.Ctx) error {
	// รับค่า id จาก query parameter
	id := ctx.Query("id")
	var orders []struct {
		OrderID         uint    `json:"order_id"`          // ID ของ Order
		OrderInfo       *string `json:"order_info"`        // ข้อมูลเกี่ยวกับ Order
		OrderImage      *string `json:"order_image"`       // รูปภาพของ Order
		OrderSenderID   uint    `json:"order_sender_id"`   // ID ของผู้ส่ง
		OrderReceiverID uint    `json:"order_receiver_id"` // ID ของผู้รับ
		UserName        string  `json:"user_name"`         // ชื่อผู้ใช้ (ผู้รับ)
		UserPhone       string  `json:"user_phone"`        // เบอร์โทรศัพท์ของผู้ใช้ (ผู้รับ)
		UserImage       string  `json:"user_image"`        // รูปภาพของผู้ใช้ (ผู้รับ)
	}

	// ตรวจสอบว่าค่า id มีอยู่หรือไม่
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing 'id' query parameter",
		})
	}

	// ใช้ GORM เพื่อทำการ JOIN ตาราง Order และ User
	result := database.MYSQL.Debug().
		Table("Order").
		Select("*").
		Joins("JOIN User ON Order.order_receiver_id = User.user_id"). // กำหนดเงื่อนไขการ JOIN
		Where("Order.order_sender_id = ?", id).                       // กำหนดเงื่อนไขการค้นหาข้อมูลจาก sender_id
		Find(&orders)

	// ตรวจสอบว่ามีข้อผิดพลาดหรือไม่
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve orders",
		})
	}

	// ส่งข้อมูลในรูปแบบ JSON
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	return ctx.JSON(orders)
}

func GetInfoOrder(ctx *fiber.Ctx) error {
	// รับค่า id จาก query parameter
	id := ctx.Query("id")

	type OrderResponse struct {
		OrderID          string `json:"order_id"`
		OrderSenderID    string `json:"order_sender_id"`
		OrderReceiverID  string `json:"order_receiver_id"`
		OrderImage       string `json:"order_image"`
		OrderInfo        string `json:"order_info"`
		UserSenderName   string `json:"user_sender_name"`
		UserReceiverName string `json:"user_receiver_name"`
		UserLocation     string `json:"user_location"`
		UserImage        string `json:"user_image"`
		UserPhone        string `json:"user_phone"`
	}

	var orders []OrderResponse

	// ตรวจสอบว่าค่า id มีอยู่หรือไม่
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing 'id' query parameter",
		})
	}

	// Query to join Order with UserSender and UserReceiver only
	result := database.MYSQL.Debug().
		Table("Order").
		Select("Order.order_id, Order.order_sender_id, Order.order_receiver_id, Order.order_image, Order.order_info, UserSender.user_name AS user_sender_name, UserReceiver.user_name AS user_receiver_name, UserSender.user_location AS user_location, UserSender.user_image AS user_image, UserSender.user_phone AS user_phone").
		Joins("JOIN User AS UserSender ON UserSender.user_id = Order.order_sender_id").
		Joins("JOIN User AS UserReceiver ON UserReceiver.user_id = Order.order_receiver_id").
		Where("Order.order_receiver_id = ?", id).
		Find(&orders)

	// ตรวจสอบว่ามีข้อผิดพลาดหรือไม่
	if result.Error != nil {
		log.Printf("Failed to retrieve orders: %v", result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve orders",
		})
	}

	// ส่งข้อมูลในรูปแบบ JSON
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	return ctx.JSON(orders)
}

func GetInfoDriver(ctx *fiber.Ctx) error {
	// รับค่า id จาก query parameter
	id := ctx.Query("id")

	type OrderResponse struct {
		OrderID          string `json:"order_id"`
		OrderSenderID    string `json:"order_sender_id"`
		OrderReceiverID  string `json:"order_receiver_id"`
		OrderImage       string `json:"order_image"`
		OrderInfo        string `json:"order_info"`
		UserSenderName   string `json:"user_sender_name"`
		UserReceiverName string `json:"user_receiver_name"`
		UserLocation     string `json:"user_location"`
		UserImage        string `json:"user_image"`
		UserPhone        string `json:"user_phone"`
		RaiderName       string `json:"raider_name"`  // Added for raider's name
		RaiderPhone      string `json:"raider_phone"` // Added for raider's phone
	}

	var orders []OrderResponse

	// ตรวจสอบว่าค่า id มีอยู่หรือไม่
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing 'id' query parameter",
		})
	}

	// Query to join Drive with Order and Raiders
	result := database.MYSQL.Debug().
		Table("drive").
		Select("drive.order_id, `Order`.order_sender_id, `Order`.order_receiver_id, `Order`.order_image, `Order`.order_info, UserSender.user_name AS user_sender_name, UserReceiver.user_name AS user_receiver_name, UserSender.user_location, UserSender.user_image, UserSender.user_phone, Raiders.raider_name, Raiders.raider_phone").
		Joins("JOIN `Order` ON `Order`.order_id = drive.order_id").
		Joins("JOIN User AS UserSender ON UserSender.user_id = `Order`.order_sender_id").
		Joins("JOIN User AS UserReceiver ON UserReceiver.user_id = `Order`.order_receiver_id").
		Joins("JOIN Raiders ON Raiders.raider_id = drive.raider_id").
		Where("`Order`.order_receiver_id = ?", id).
		Find(&orders)

	// ตรวจสอบว่ามีข้อผิดพลาดหรือไม่
	if result.Error != nil {
		log.Printf("Failed to retrieve orders: %v", result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve orders",
		})
	}

	// ส่งข้อมูลในรูปแบบ JSON
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	return ctx.JSON(orders)
}

// post
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

func InsertDrive(ctx *fiber.Ctx) error {
	var drives entity.Drive

	// Parse JSON body
	if err := ctx.BodyParser(&drives); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// ทำการ insert ข้อมูลลงฐานข้อมูลโดยใช้ GORM
	if err := database.MYSQL.Debug().Table("drive").Create(&drives).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert order",
		})
	}

	// ส่ง response เมื่อ insert สำเร็จ
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order inserted successfully",
		"order":   drives,
	})
}

// func InsertDrive(ctx *fiber.Ctx) error {
//     // สร้างตัวแปรเพื่อเก็บข้อมูลที่รับจาก body
//     var drives entity.Drive

//     // อ่านข้อมูล JSON จาก request body
//     if err := ctx.BodyParser(&drives); err != nil {
//         return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//             "error": "Invalid input",
//         })
//     }

//     // เชื่อมต่อกับฐานข้อมูลและบันทึกข้อมูล
//     if err := database.MYSQL.Create(&drives).Error; err != nil {
//         return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//             "error": "Failed to insert drive",
//         })
//     }

//     // ส่งกลับผลลัพธ์เมื่อสำเร็จ
//     return ctx.Status(fiber.StatusCreated).JSON(drives)
// }

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

func Putstatus(ctx *fiber.Ctx) error {
	id := ctx.Query("id")
    var put entity.PutDrive

    // Parse the request body into the PutDrive struct
    if err := ctx.BodyParser(&put); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // Update the drive record in the database
    if err := database.MYSQL.Debug().Table("drive").Where("order_id = ?", id).Updates(put).Error; err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update order",
        })
    }

    // Send a success response
    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Status updated successfully",
        "data":    put, // Return the updated data if needed
    })
}

func PutstatusOrder(ctx *fiber.Ctx) error {
    id := ctx.Query("id") // Get the order ID from the query parameter
    var put entity.PutOrder

    // Parse the request body into the PutOrder struct
    if err := ctx.BodyParser(&put); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // Update the order record in the database
    if err := database.MYSQL.Debug().Table("Order").Where("order_id = ?", id).Updates(put).Error; err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update order",
        })
    }

    // Send a success response
    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Order updated successfully",
        "data":    put, // Return the updated data if needed
    })
}

