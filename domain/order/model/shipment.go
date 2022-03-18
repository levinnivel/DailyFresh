package model

type Shipment struct {
	Arrival_Date int `form:"arrival_date" json:"arrival_date"`
	Order
	Ship_Date string `form:"ship_date" json:"ship_date"`
}