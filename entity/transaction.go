package entity

import "time"

type Laundry struct {
	Id         int    `json:"id"`
	Unit       string `json:"unit"`
	Amount     int    `json:"amount"`
	DateIn     time.Time
	DateOut    time.Time
	IdCustomer int `json:"id_customer"`
	IdService  int `json:"id_service"`
}
