package models

type Customer struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Gmail      string `json:"gmail"`
	Phone      string `json:"phone"`
	Is_Blocked bool   `json:"isblocked"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	// CarsCount   int    `json:"carscount"`
	// OrdersCount int    `json:"orderscount"`
	// Orders     []GetOrder `json:"orders"`
}

type CreateCustomer struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Gmail      string `json:"gmail"`
	Phone      string `json:"phone"`
	Is_Blocked bool   `json:"isblocked"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type Login struct {
	Phone      string
	Password string
}

type ChangePassword struct {
	Password string `json:"password"`
}

type GetPassword struct {
	Password string `json:"password"`
}

type GetAllCustomersResponse struct {
	Customers []Customer `json:"customers"`
	Count     int16      `json:"count"`
}
type GetAllCustomersRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

// type Cars struct {
// 	Car []Car
// }

// type CustomerCarsResponse struct {
// 	Cars []Cars `json:"cars"`
// }

type OrderCount struct {
	Count []Customer
}
