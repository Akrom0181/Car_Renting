package service

import (
	"context"

	"rent-car/api/models"
	"rent-car/pkg/logger"
	"rent-car/storage"
)

type customerService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewCustomerService(storage storage.IStorage, log logger.ILogger) customerService {
	return customerService{
		storage: storage,
		log:     log,
	}
}

func (u customerService) Login(ctx context.Context, login string) (string, error) {

	cus := models.Customer{}
	password, err := u.storage.Customer().GetPassword(ctx, cus.Phone)
	if err != nil {
		u.log.Error("error while getting password in service layer", logger.Error(err))
	}
	return password, nil

}

func (u customerService) Create(ctx context.Context, customer models.Customer) (string, error) {

	pKey, err := u.storage.Customer().Create(ctx, customer)
	if err != nil {
		u.log.Error("ERROR in service layer while creating customer", logger.Error(err))
		return "", err
	}

	return pKey, nil
}

func (u customerService) ChangePassword(ctx context.Context, customer models.Customer) (string, error) {

	pKey, err := u.storage.Customer().ChangePassword(ctx, customer)
	if err != nil {
		u.log.Error("ERROR in service layer while changing password", logger.Error(err))
		return "", err
	}
	return pKey, nil
}

func (u customerService) Update(ctx context.Context, customer models.Customer) (string, error) {

	pKey, err := u.storage.Customer().Update(ctx, customer)
	if err != nil {
		u.log.Error("ERROR in service layer while updating customer", logger.Error(err))
		return "", err
	}

	return pKey, nil
}

func (u customerService) GetByID(ctx context.Context, id string) (models.Customer, error) {

	pKey, err := u.storage.Customer().GetByID(ctx, id)
	if err != nil {
		u.log.Error("ERROR in service layer while getbyid customer", logger.Error(err))
		return models.Customer{}, err
	}

	return pKey, nil
}

func (u customerService) GetAll(ctx context.Context, req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {

	pKey, err := u.storage.Customer().GetAll(ctx, req)
	if err != nil {
		u.log.Error("ERROR in service layer while GetAll customer", logger.Error(err))
		return models.GetAllCustomersResponse{}, err
	}

	return pKey, nil
}

func (u customerService) Delete(ctx context.Context, id string) error {

	err := u.storage.Customer().Delete(ctx, id)
	if err != nil {
		u.log.Error("ERROR in service layer while deleting customer", logger.Error(err))
		return err
	}

	return nil
}

func (u customerService) GetCustomerCars(ctx context.Context, req models.GetAllCustomersRequest) ([]models.GetOrder, error) {

	pKey, err := u.storage.Customer().GetCustomerCars(ctx, req)
	if err != nil {
		u.log.Error("ERROR in service layer while GetAll customer", logger.Error(err))
		return []models.GetOrder{}, err
	}

	return pKey, nil
}

func (u customerService) GetByIDCustomerCar(ctx context.Context, id string) (models.GetOrder, error) {

	pKey, err := u.storage.Customer().GetByIDCustomerCar(ctx, id)
	if err != nil {
		u.log.Error("ERROR in service layer while getbyid customerCars", logger.Error(err))
		return models.GetOrder{}, err
	}

	return pKey, nil
}

func (u customerService) GetPassword(ctx context.Context, phone string) (string, error) {
	pKey, err := u.storage.Customer().GetPassword(ctx, phone)
	if err != nil {
		u.log.Error("ERROR in service layer while getbyID Customer", logger.Error(err))
		return "Error", err
	}

	return pKey, nil
}
