// package postgres

// import (
// 	"context"
// 	"fmt"
// 	"rent-car/api/models"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// // func TestCreateOrder(t *testing.T) {
// // 	orderRepo := NewOrder(db)

// // 	reqOrder := models.CreateOrder{
// // 		CarId: "15c50a82-8206-4e12-95ef-772fe8001223",
// // 		CustomerId: "a6824eb8-0d5c-4072-a79f-34f09062c717",
// // 		FromDate: "2024-04-2 12:12:21",
// // 		ToDate: "2024-05-18 12:12:21",
// // 		Status: "success",
// // 		Payment_status: true,
// // 		Amount: 123123123,
// // 	}

// //		id, err := orderRepo.CreateOrder(context.Background(), reqOrder)
// //		if assert.NoError(t, err) {
// //			createdOrder, err := orderRepo.GetOne(context.Background(), id)
// //			if assert.NoError(t, err) {
// //				assert.Equal(t, reqOrder.FromDate, createdOrder.FromDate)
// //			} else {
// //				return
// //			}
// //			fmt.Println("Created customer", createdOrder)
// //		}
// //	}
// func TestCreateOrder(t *testing.T) {

// 	repo := NewOrder(db)

// 	testOrder := models.CreateOrder{
// 		CarId:          "15c50a82-8206-4e12-95ef-772fe8001223",
// 		CustomerId:     "a6824eb8-0d5c-4072-a79f-34f09062c717",
// 		FromDate:       "2024-04-05",
// 		ToDate:         "2024-04-10",
// 		Status:         "success",
// 		Payment_status: true,
// 		Amount:         100.00,
// 	}

// 	id, err := repo.CreateOrder(context.Background(), testOrder)

// 	if err != nil {
// 		t.Errorf("CreateOrder failed with error: %v", err)
// 	}

// 	if id == "" {
// 		t.Errorf("Expected non-empty ID returned from CreateOrder")
// 	}

// }

// func TestUpdateOrder(t *testing.T) {
// 	customerRepo := NewOrder(db)

// 	reqCustomer := models.GetOrder{
// 		Id:             "bf767d11-5cc8-43cc-85d1-54ed587fc32e",
// 		FromDate:       "2024-03-25 12:12:12",
// 		ToDate:         "2024-04-25 12:12:12",
// 		Status:         "in process",
// 		Payment_status: true,
// 		Amount:         100.00,
// 	}
// 	id, err := customerRepo.UpdateOrder(context.Background(), reqCustomer)

// 	if err != nil {
// 		t.Errorf("UpdateOrder failed with error: %v", err)
// 	}

// 	if id != reqCustomer.Id {
// 		t.Errorf("Expected ID %s, but got %s", reqCustomer.Id, id)
// 	}
// }

// func TestGetByIDOrder(t *testing.T) {
// 	customerRepo := NewOrder(db)

// 	orderID := "bf767d11-5cc8-43cc-85d1-54ed587fc32e"

// 	car, err := customerRepo.GetOne(context.Background(), orderID)

// 	if err != nil {
// 		t.Fatalf("error retrieving customer with ID %s: %v", orderID, err)
// 	}

// 	if car != (models.GetOrder{}) {
// 		t.Errorf("expected nil car but got %+v when retrieving customer with ID %s", car, orderID)
// 	}
// }

// func TestDeleteOrder(t *testing.T) {
// 	orderRepo := NewOrder(db)

// 	orderID := "bf767d11-5cc8-43cc-85d1-54ed587fc32e"

// 	err := orderRepo.DeleteOrder(context.Background(), orderID)

// 	if err == nil {
// 		t.Errorf("order with ID %s", orderID)
// 	}

// }

package postgres

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {

	repo := NewOrder(db)

	testOrder := models.CreateOrder{
		CarId:      "2e195d78-3691-44e2-941a-53d7d5643630",
		CustomerId: "a6824eb8-0d5c-4072-a79f-34f09062c717",
		FromDate:   "2024-04-05",
		ToDate:     "2024-04-10",
		Status:     "success",
		Payment_status:   true,
		Amount:     100.00,
	}

	id, err := repo.CreateOrder(context.Background(), testOrder)
	assert.NoError(t, err)

	fmt.Println("pass", id)

}


func TestUpdateOrder(t *testing.T) {
	repo := NewOrder(db)

	testUpdate := models.GetOrder{
		Id:       "3bf82f8a-0138-4f1e-8ec1-2a68ce7978d1",      
		FromDate: "2024-04-06",   
		ToDate:   "2024-04-12",   
		Status:   "in process",    
		Payment_status:     true,           
		Amount:   100.00,         
	}

	id, err := repo.UpdateOrder(context.Background(), testUpdate)


	if err != nil {
		t.Errorf("UpdateOrder failed with error: %v", err)
	}

	if id != testUpdate.Id {
		t.Errorf("Expected ID %s, but got %s", testUpdate.Id, id)
	}
}


func TestGetAllOrders(t *testing.T) {

	repo := NewOrder(db) 

	testRequest := models.GetAllOrdersRequest{
		Page:   1,
		Limit:  10,
		Search: "in process",
	}

    response, err := repo.GetAll(context.Background(), testRequest)

	if err != nil {
		t.Errorf("GetAllOrders failed with error: %v", err)
	}

	if len(response.Orders) == 0 {
		t.Errorf("Expected non-empty order list, but got an empty list")
	}

}


func TestGetOrderByID(t *testing.T) {
	repo := NewOrder(db)

	testID := "64c9f418-a0b9-4f18-a8b8-3e0212f0a9e0" 

	order, err := repo.GetOne(context.Background(), testID)
	
	if err != nil {
		t.Errorf("GetOrderByID failed with error: %v", err)
	}
	
	if order.Id != testID {
		t.Errorf("Expected order ID %s, but got %s", testID, order.Id)
	}
}


func TestDeleteOrder(t *testing.T) {
	repo := NewOrder(db)

	orderid := "3bf82f8a-0138-4f1e-8ec1-2a68ce7978d1" 

	err := repo.DeleteOrder(context.Background(), orderid)

	if err != nil {
		fmt.Println("Error occurred while deleting order:", err)
		t.Errorf("Failed to delete order with ID %s: %v", orderid, err)
	}
}

