package service

import (
	"rent-car/storage"
)

type IServiceManager interface {
	Car() carService
	Customer() customerService
	Order() orderService
}

type Service struct {
	carService      carService
	customerService customerService
	orderService    orderService
}

func New(storage storage.IStorage) Service {
	services := Service{}
	services.carService = NewCarService(storage)
	services.customerService = NewCustomerService(storage)
	services.orderService = NewOrderService(storage)
	return services
}

func (s Service) Car() carService {
	return s.carService
}

func (s Service) Customer() customerService {
	return s.customerService
}

func (s Service) Order() orderService {
	return s.orderService
}


