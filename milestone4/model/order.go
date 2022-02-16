package model

import "time"

type Item struct {
	Quantity int `json:"quantity"`
	ItemID string `json:"item_id"`
}

type Order struct {
	Id string `json:"id"`
	Name string     `json:"name"`
	Timestamp time.Time `json:"timestamp"`
	CustomerID string `json:"customer_id"`
	Items []Item `json:"items"`
}



