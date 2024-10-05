package entity

type User struct {
	User_id   int    `json:"user_id" gorm:"primaryKey"`
	User_name string `json:"user_name"`
	User_email string `json:"user_email"`
	User_password string `json:"user_password"`
	User_location string `json:"user_location"`
	User_image string `json:"user_image"`
	User_Phone string `json:"user_phone"`
	User_address string `json:"user_address"`	
}

type PostUser struct {
	User_name string `json:"user_name"`
	User_email string `json:"user_email"`
	User_password string `json:"user_password"`
	User_location string `json:"user_location"`
	User_image string `json:"user_image"`
	User_Phone string `json:"user_phone"`
	User_address string `json:"user_address"`	
}
