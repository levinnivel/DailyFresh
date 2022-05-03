package model

type Shipment struct {
	ID          int    `form:"id" json:"id"`
	ArrivalDate string `form:"arrival_date" json:"arrival_date"`
	ShipDate    string `form:"ship_date" json:"ship_date"`
	OrderID     int    `form:"order_id" json:"order_id"`
}
