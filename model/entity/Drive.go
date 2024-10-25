package entity

type Drive struct {
	Drive_image1 string `json:"drive_image1"`
	Drive_image2 string `json:"drive_image2"`
	Order_id     string `json:"order_id"`
	Drive_status string `json:"drive_status"`
	Raider_id    string `json:"raider_id"` // แก้ไขที่นี่
}

type PutDrive struct {
	Drive_image1 string `json:"drive_image1"`
	Drive_image2 string `json:"drive_image2"`
	Drive_status string `json:"drive_status"`
}
