package models

type GetOrder struct {
	Id             string   `json:"id"`
	Car            Car      `json:"car"`
	Customer       Customer `json:"customer"`
	FromDate       string   `json:"from_date"`
	ToDate         string   `json:"to_date"`
	Status         string   `json:"status"`
	Payment_status bool     `json:"payment_status"`
	Amount         float64  `json:"amount"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
}

type CreateOrder struct {
	Id             string  `json:"id"`
	CarId          string  `json:"car_id"`
	CustomerId     string  `json:"customer_id"`
	FromDate       string  `json:"from_date"`
	ToDate         string  `json:"to_date"`
	Status         string  `json:"status"`
	Payment_status bool    `json:"payment_status"`
	Amount         float64 `json:"amount"`
	Created_at     string  `json:"created_at"`
	Updated_at     string  `json:"updated_at"`
}

type CreateOrderr struct {
	CarId          string  `json:"car_id"`
	CustomerId     string  `json:"customer_id"`
	FromDate       string  `json:"from_date"`
	ToDate         string  `json:"to_date"`
	Status         string  `json:"status"`
	Payment_status bool    `json:"payment_status"`
	Amount         float64 `json:"Amount"`
}

type GetAllOrders struct {
	Orders []GetOrder `json:"orders"`
	Count  int        `json:"count"`
}

type GetAllOrdersRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllOrdersResponse struct {
	Orders []GetOrder `json:"orders"`
	Count  int16      `json:"count"`
}
