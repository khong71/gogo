package entity

type Register struct {
	User_name     string `json:"user_name"`
	User_email    string `json:"user_email"`
	User_Phone    int    `json:"user_phone"`
	User_image    string `json:"user_image"`
	User_address  string `json:"user_address"`
	User_password string `json:"user_password"`
	User_location string `json:"user_location"`
}

type RegisterDriver struct {
	Raider_name  string `json:"raider_name"`
	Raider_email string `json:"raider_email"`
	Raider_Phone int    `json:"raider_phone"`
	Raider_image string `json:"raider_image"`
	// Raider_numder   string `json:"raider_numder-plate"`
	Raider_password string `json:"raider_password"`
	Raider_location string `json:"raider_location"`
}
