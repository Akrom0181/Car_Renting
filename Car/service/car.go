package service

import (
	"context"
	"fmt"

	"net/http"
	"rent-car/api/models"
	"rent-car/storage"
)

type carService struct {
	storage storage.IStorage
}

func NewCarService(storage storage.IStorage) carService {
	return carService{
		storage: storage,
	}
}

func (u carService) Create(ctx context.Context, car models.Car) (string, error) {

	pKey, err := u.storage.Car().Create(ctx, car)
	if err != nil {
		handleResponseLog(ctx, "ERROR in service layer while creating car", http.StatusBadRequest, err.Error())
		return "", err
	}
	return pKey, nil
}

func handleResponseLog(ctx context.Context, s1 string, i int, s2 string) {
	panic("unimplemented")
}



func (u carService) Update(ctx context.Context, car models.Car) (string, error) {

	pKey, err := u.storage.Car().Update(ctx, car)
	if err != nil {
		fmt.Println("ERROR in service layer while updating car", err.Error())
		return "", err
	}

	return pKey, nil
}

func (u carService) GetByID(ctx context.Context, id string) (models.Car, error) {

	pKey, err := u.storage.Car().GetByID(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyid car", err.Error())
		return models.Car{}, err
	}

	return pKey, nil
}

func (u carService) GetAll(ctx context.Context, req models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {

	pKey, err := u.storage.Car().GetAll(ctx, req)
	if err != nil {
		fmt.Println("ERROR in service layer while GetAll car", err.Error())
		return models.GetAllCarsResponse{}, err
	}

	return pKey, nil
}

func (u carService) Delete(ctx context.Context, id string) error {

	err := u.storage.Car().Delete(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting car", err.Error())
		return err
	}

	return nil
}

func (u carService) GetAvailableCars(ctx context.Context, car models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	cars, err := u.storage.Car().GetAvailableCars(ctx, car)
	if err != nil {
		fmt.Println("ERROR in service layer while getting available cars", err.Error())
		return cars, err
	}
	return cars, nil
}

func (u carService) GetByIDCustomerCar(ctx context.Context, id string) (models.GetOrder, error) {

	pKey, err := u.storage.Car().GetByIDCustomerCarr(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyid___ car", err.Error())
		return models.GetOrder{}, err
	}

	return pKey, nil
}

func (u carService) GetAllCustomerCars(ctx context.Context, req models.GetAllCarsRequest) (resp []models.GetOrder, err error) {

	pKey, err := u.storage.Car().GetAllOrdersCars(ctx, req)
	if err != nil {
		fmt.Println("ERROR in service layer while GetAll car", err.Error())
		return resp, err
	}
	return pKey, nil
}