package entity

type Register struct {
	User_name string `json:"user_name"`
	User_email string `json:"user_email"`
	User_Phone string `json:"user_phone"`
	User_image string `json:"user_image"`
	User_address string `json:"user_address"`	
	User_password string `json:"user_password"`
	User_location string `json:"user_location"`
}