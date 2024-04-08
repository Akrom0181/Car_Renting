package service

import (
	"context"

	"rent-car/api/models"
	"rent-car/pkg/logger"
	"rent-car/storage"
)

type carService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewCarService(storage storage.IStorage, log logger.ILogger) carService {
	return carService{
		storage: storage,
		log:   log,
	}
}

func (u carService) Create(ctx context.Context, car models.Car) (string, error) {

	pKey, err := u.storage.Car().Create(ctx, car)
	if err != nil {
		u.log.Error("ERROR in service layer while creating car",  logger.Error(err))
		return "", err
	}
	return pKey, nil
}

func (u carService) Update(ctx context.Context, car models.Car) (string, error) {

	pKey, err := u.storage.Car().Update(ctx, car)
	if err != nil {
		u.log.Error("ERROR in service layer while updating car", logger.Error(err))
		return "", err
	}

	return pKey, nil
}

func (u carService) GetByID(ctx context.Context, id string) (models.Car, error) {

	pKey, err := u.storage.Car().GetByID(ctx, id)
	if err != nil {
		u.log.Error("ERROR in service layer while getbyid car", logger.Error(err))
		return models.Car{}, err
	}

	return pKey, nil
}

func (u carService) GetAll(ctx context.Context, req models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {

	pKey, err := u.storage.Car().GetAll(ctx, req)
	if err != nil {
		u.log.Error("ERROR in service layer while GetAll car", logger.Error(err))
		return models.GetAllCarsResponse{}, err
	}

	return pKey, nil
}

func (u carService) Delete(ctx context.Context, id string) error {

	err := u.storage.Car().Delete(ctx, id)
	if err != nil {
		u.log.Error("ERROR in service layer while deleting car", logger.Error(err))
		return err
	}

	return nil
}

func (u carService) GetAvailableCars(ctx context.Context, car models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	cars, err := u.storage.Car().GetAvailableCars(ctx, car)
	if err != nil {
		u.log.Error("ERROR in service layer while getting available cars", logger.Error(err))
		return cars, err
	}
	return cars, nil
}

func (u carService) GetByIDCustomerCar(ctx context.Context, id string) (models.GetOrder, error) {

	pKey, err := u.storage.Car().GetByIDCustomerCarr(ctx, id)
	if err != nil {
		u.log.Error("ERROR in service layer while getbyid___ car", logger.Error(err))
		return models.GetOrder{}, err
	}

	return pKey, nil
}

func (u carService) GetAllCustomerCars(ctx context.Context, req models.GetAllCarsRequest) (resp []models.GetOrder, err error) {

	pKey, err := u.storage.Car().GetAllOrdersCars(ctx, req)
	if err != nil {
		u.log.Error("ERROR in service layer while GetAll car", logger.Error(err))
		return resp, err
	}
	return pKey, nil
}