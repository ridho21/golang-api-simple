package entity

type Laundry struct {
	Id         int    `json:"id"`
	Unit       string `json:"unit"`
	Amount     int    `json:"amount"`
	DateIn     string `json:"date_in"`
	DateOut    string `json:"date_out"`
	IdCustomer int    `json:"id_customer"`
	IdService  int    `json:"id_service"`
}
