package service

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"rent-car/storage"
)

type customerService struct {
	storage storage.IStorage
}

func NewCustomerService(storage storage.IStorage) customerService {
	return customerService{
		storage: storage,
	}
}

func (u customerService) Create(ctx context.Context, customer models.Customer) (string, error) {

	pKey, err := u.storage.Customer().Create(ctx, customer)
	if err != nil {
		fmt.Println("ERROR in service layer while creating customer", err.Error())
		return "", err
	}

	return pKey, nil
}

func (u customerService) ChangePassword(ctx context.Context, customer models.Customer) (string, error){

	pKey, err := u.storage.Customer().ChangePassword(ctx, customer)
	if err != nil {
		fmt.Println("ERROR in service layer while changing password", err.Error())
		return "", err
	}
	return pKey, nil
}

func (u customerService) Update(ctx context.Context, customer models.Customer) (string, error) {

	pKey, err := u.storage.Customer().Update(ctx, customer)
	if err != nil {
		fmt.Println("ERROR in service layer while updating customer", err.Error())
		return "", err
	}

	return pKey, nil
}

func (u customerService) GetByID(ctx context.Context, id string) (models.Customer, error) {

	pKey, err := u.storage.Customer().GetByID(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyid customer", err.Error())
		return models.Customer{}, err
	}

	return pKey, nil
}

func (u customerService) GetAll(ctx context.Context, req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {

	pKey, err := u.storage.Customer().GetAll(ctx, req)
	if err != nil {
		fmt.Println("ERROR in service layer while GetAll customer", err.Error())
		return models.GetAllCustomersResponse{}, err
	}

	return pKey, nil
}

func (u customerService) Delete(ctx context.Context, id string) error {

	err := u.storage.Customer().Delete(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting customer", err.Error())
		return err
	}

	return nil
}

func (u customerService) GetCustomerCars(ctx context.Context, req models.GetAllCustomersRequest) ([]models.GetOrder, error) {

	pKey, err := u.storage.Customer().GetCustomerCars(ctx, req)
	if err != nil {
		fmt.Println("ERROR in service layer while GetAll customer", err.Error())
		return []models.GetOrder{}, err
	}

	return pKey, nil
}

func (u customerService) GetByIDCustomerCar(ctx context.Context, id string) (models.GetOrder, error) {

	pKey, err := u.storage.Customer().GetByIDCustomerCar(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyid customerCars", err.Error())
		return models.GetOrder{}, err
	}

	return pKey, nil
}