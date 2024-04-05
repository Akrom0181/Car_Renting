package service

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"rent-car/storage"
)

type orderService struct {
	storage storage.IStorage
}

func NewOrderService(storage storage.IStorage) orderService {
	return orderService{
		storage: storage,
	}
}

func (u orderService) Create(ctx context.Context, order models.CreateOrder) (string, error) {

	pKey, err := u.storage.Order().CreateOrder(ctx, order)
	if err != nil {
		fmt.Println("ERROR in service layer while creating order", err.Error())
		return "", err
	}

	return pKey, nil
}

func (u orderService) Update(ctx context.Context, order models.GetOrder) (string, error) {

	pKey, err := u.storage.Order().UpdateOrder(ctx, order)
	if err != nil {
		fmt.Println("ERROR in service layer while updating order", err.Error())
		return "", err
	}

	return pKey, nil
}

func (u orderService) GetByID(ctx context.Context, id string) (models.GetOrder, error) {

	pKey, err := u.storage.Order().GetOne(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyid order", err.Error())
		return models.GetOrder{}, err
	}

	return pKey, nil
}

func (u orderService) GetAll(ctx context.Context, req models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error) {

	pKey, err := u.storage.Order().GetAll(ctx, req)
	if err != nil {
		fmt.Println("ERROR in service layer while GetAll order", err.Error())
		return models.GetAllOrdersResponse{}, err
	}

	return pKey, nil
}

// func (u orderService) Delete(ctx context.Context, id string) error {

// 	err := u.storage.Order().Delete(ctx, id)
// 	if err != nil {
// 		fmt.Println("ERROR in service layer while deleting order", err.Error())
// 		return err
// 	}

// 	return nil
// }
