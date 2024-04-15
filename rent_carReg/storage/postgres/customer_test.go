package postgres

import (
	"clone/rent_car_us/api/models"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	customerRepo := NewCustomer(db)

	reqCustomer := models.CreateCustomer{
		First_name: "ssss",
		Last_name:  "Tsssss",
		Gmail:      "sunnatillo@gmail.com",
		Phone:      "+998939513121",
	}

	id, err := customerRepo.Create(context.Background(), reqCustomer)
	if assert.NoError(t, err) {
		updatedCustomer, err := customerRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCustomer.First_name, updatedCustomer.First_name)
			assert.Equal(t, reqCustomer.Last_name, updatedCustomer.Last_name)
			assert.Equal(t, reqCustomer.Gmail, updatedCustomer.Gmail)
			assert.Equal(t, reqCustomer.Phone, updatedCustomer.Phone)
		} else {
			return
		}
		fmt.Println("Created customer", updatedCustomer)
	}
}

func TestUpdateCustomer(t *testing.T) {
	customerRepo := NewCustomer(db)

	reqCustomer := models.GetCustomer{
		Id:         "01f6e0600-008b-494b-b54d-025ed2a4af0c",
		First_name: "ssss",
		Last_name:  "Tsssss",
		Gmail:      "sunnatillo@gmail.com",
		Phone:      "+998939513121",
	}

	id, err := customerRepo.Update(context.Background(), reqCustomer)
	if assert.NoError(t, err) {
		updatedCustomer, err := customerRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCustomer.First_name, updatedCustomer.First_name)
			assert.Equal(t, reqCustomer.Last_name, updatedCustomer.Last_name)
			assert.Equal(t, reqCustomer.Gmail, updatedCustomer.Gmail)
			assert.Equal(t, reqCustomer.Phone, updatedCustomer.Phone)

		} else {
			return
		}
		fmt.Println("Updated customer", updatedCustomer)
	}
}

func TestGetByIDCustomer(t *testing.T) {
	customerRepo := NewCustomer(db)

	customerID := "1f6e0600-008b-494b-b54d-025ed2a4af0c"

	customer, err := customerRepo.GetByID(context.Background(), customerID)

	if err != nil {
		t.Fatalf("error retrieving customer with ID %s: %v", customerID, err)
	}

	if customer != (models.GetCustomer{}) {
		t.Errorf("expected nil customer but got %+v when retrieving customer with ID %s", customer, customerID)
	}
}

func TestDeleteCustomer(t *testing.T) {
	customerRepo := NewCustomer(db)

	customerID := "1f6e0600-008b-494b-b54d-025ed2a4af0c"

	err := customerRepo.Delete(context.Background(), customerID)

	if err == nil {
		t.Errorf("customer Delete ID %s", customerID)
	}

}
