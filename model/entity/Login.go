package entity

type LoginUser struct {
	User_email string `json:"user_email"`	
	User_password string `json:"user_password"`
}

type LoginDriver struct {
	Raider_email string `json:"raider_email"`
	Raider_password string `json:"raider_password"`
}