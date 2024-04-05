package storage

import (
	"context"
	"rent-car/api/models"
)

type IStorage interface {
	CloseDB()
	Car() ICarStorage
	Customer() ICustomerStorage
	Order() IOrderStorage
	// GetCusCars() IGetCusCars
}

type ICarStorage interface {
	Create(context.Context, models.Car) (string, error)
	GetByID(context.Context, string) (models.Car, error)
	GetAll(ctx context.Context, request models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	Update(context.Context, models.Car) (string, error)
	Delete(ctx context.Context, id string) error
	GetAllOrdersCars(ctx context.Context, req models.GetAllCarsRequest) ([]models.GetOrder, error)
	GetByIDCustomerCarr(ctx context.Context, id string) (models.GetOrder, error)
	GetAvailableCars(ctx context.Context, req models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
}

type ICustomerStorage interface {
	Create(context.Context, models.Customer) (string, error)
	GetByID(ctx context.Context, id string) (models.Customer, error)
	GetAll(ctx context.Context, request models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
	Update(context.Context, models.Customer) (string, error)
	Delete(context.Context, string) error
	GetCustomerCars(ctx context.Context, req models.GetAllCustomersRequest) ([]models.GetOrder, error)
	GetByIDCustomerCar(ctx context.Context, id string) (models.GetOrder, error)
	ChangePassword(ctx context.Context, customer models.Customer) (string, error)
}

type IOrderStorage interface {
	CreateOrder(context.Context, models.CreateOrder) (string, error)
	UpdateOrder(context.Context, models.GetOrder) (string, error)
	GetOne(ctx context.Context, id string) (models.GetOrder, error)
	GetAll(ctx context.Context, req models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error)
	DeleteOrder(context.Context, string) error
}

// type IGetCusCars interface {
// 	GetCusCars(id string, req models.GetAllCarsRequest) ([]models.GetOrder, error)
// }
