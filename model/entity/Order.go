package entity

type InsertOrder struct {
	OrderImage      string `json:"order_image"`
	OrderInfo       string `json:"order_info"`
	OrderSenderID   string `json:"order_sender_id"`
	OrderReceiverID string `json:"order_receiver_id"`
	OrderRaiderID   string `json:"order_raider_id"`
}
