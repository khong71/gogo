package entity

type InsertOrder struct {
	OrderImage      *string `json:"order_image"`       // อนุญาตให้เป็น nil
	OrderInfo       *string `json:"order_info"`        // อนุญาตให้เป็น nil
	OrderSenderID   *string `json:"order_sender_id"`   // ต้องมีค่า
	OrderReceiverID *string `json:"order_receiver_id"` // อนุญาตให้เป็น nil
}

type GetOrder struct {
	Orderid         int    `json:"order_id" gorm:"column:order_id;primaryKey"`
	OrderImage      string `json:"order_image"`
	OrderInfo       string `json:"order_info"`
	OrderSenderID   string `json:"order_sender_id"`
	OrderReceiverID string `json:"order_receiver_id"`
	Status          string `json:"status"`
}

type PutOrder struct {
	Status string `json:"status"`
}
