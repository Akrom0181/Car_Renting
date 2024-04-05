package postgres

import (
	"context"
	"fmt"
	"rent-car/api/models"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	customerRepo := NewCustomer(db)

	reqCustomer := models.Customer{
		FirstName:  "Imronfsdsdf12345",
		LastName:   "Hojiyevdfsdf",
		Gmail:      "imronhosfdjiyev12345@gmail.com",
		Phone:      "+998000000000",
		Is_Blocked: false,
	}

	id, err := customerRepo.Create(context.Background(), reqCustomer)
	if assert.NoError(t, err) {
		createdCustomer, err := customerRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCustomer.FirstName, createdCustomer.FirstName)
			assert.Equal(t, reqCustomer.LastName, createdCustomer.LastName)
			assert.Equal(t, reqCustomer.Gmail, createdCustomer.Gmail)
			assert.Equal(t, reqCustomer.Phone, createdCustomer.Phone)
			assert.Equal(t, reqCustomer.Is_Blocked, createdCustomer.Is_Blocked)

		} else {
			return
		}
		fmt.Println("Created customer", createdCustomer)
	}
}



func TestUpdateCustomer(t *testing.T) {
	customerRepo := NewCustomer(db)

	reqCustomer := models.Customer{
		Id: "01a1515a-916c-41e6-aa23-1828eacce850",
		FirstName:  "Yekgorinji",
		LastName:   "Gashrov",
		Gmail:      "yekgorinji@gmail.com",
		Phone:      "+998953456510",
		Is_Blocked: false,
	}

	id, err := customerRepo.Update(context.Background(), reqCustomer)
	if assert.NoError(t, err) {
		createdCustomer, err := customerRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCustomer.FirstName, createdCustomer.FirstName)
			assert.Equal(t, reqCustomer.LastName, createdCustomer.LastName)
			assert.Equal(t, reqCustomer.Gmail, createdCustomer.Gmail)
			assert.Equal(t, reqCustomer.Phone, createdCustomer.Phone)
			assert.Equal(t, reqCustomer.Is_Blocked, createdCustomer.Is_Blocked)

		} else {
			return
		}
		fmt.Println("Updated customer", createdCustomer)
	}
}

func TestGetByIDCustomer(t *testing.T) {
	customerRepo := NewCustomer(db)

	customerID := "a6824eb8-0d5c-4072-a79f-34f09062c717"

	customer, err := customerRepo.GetByID(context.Background(), customerID)

	assert.NoError(t, err)
	fmt.Println("pass", customer)

	
}

func TestDeleteCustomer(t *testing.T) {
	customerRepo := NewCustomer(db)

	reqCustomer := models.Customer{
		FirstName:  "Imronfsdsdf12345",
		LastName:   "Hojiyevdfsdf",
		Gmail:      "imronhosfdjiyev12345@gmail.com",
		Phone:      "+998953456540",
		Is_Blocked: false,
	}

	id, err := customerRepo.Create(context.Background(), reqCustomer)
	assert.NoError(t, err)
	err = customerRepo.Delete(context.Background(), id)

	if assert.NoError(t, err) {
		fmt.Println("passed")
	}
	}
