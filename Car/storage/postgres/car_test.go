package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"rent-car/api/models"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-faker/faker/v4"

	"github.com/stretchr/testify/assert"
)

func TestCreateCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.Car{
		Name:  faker.Name(),
		Year:  2011,
		Brand: faker.Word(),
	}

	id, err := carRepo.Create(context.Background(), reqCar)
	if assert.NoError(t, err) {
		createdCar, err := carRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCar.Name, createdCar.Name)
			assert.Equal(t, reqCar.Year, createdCar.Year)
			assert.Equal(t, reqCar.Brand, createdCar.Brand)
		} else {
			return
		}
		fmt.Println("Created car", createdCar)
	}
}

func TestUpdateCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.Car{
		Id:    "f9c1e1b8-7e67-4454-907a-c9d9c8138512",
		Name:  "M8 Competition",
		Brand: "BMW",
		Year:  2010,
	}

	id, err := carRepo.Update(context.Background(), reqCar)
	if assert.NoError(t, err) {
		updatedCar, err := carRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCar.Name, updatedCar.Name)
			assert.Equal(t, reqCar.Brand, updatedCar.Brand)
			assert.Equal(t, reqCar.Year, updatedCar.Year)
		} else {
			return
		}
		fmt.Println("Updated car", updatedCar)
	}
}

func TestGetByIDCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.Car{
		Name:  faker.NAME,
		Model: faker.WORD,
	}

	id, err := carRepo.Create(context.Background(), reqCar)
	assert.NoError(t, err)

	car, err := carRepo.GetByID(context.Background(), id)

	assert.NoError(t, err)

	fmt.Println(car, "passed")
}

func TestGetAllCar(t *testing.T) {
	carRepo := NewCar(db)

	for i := 0; i < 3; i++ {
		reqCar := models.Car{
			Name:        gofakeit.RandomString([]string{"Tesla", "Hundai", "BMW"}),
			Year:        2010 + i,
			Brand:       faker.Word(),
			Model:       faker.Word(),
			HoursePower: 200 + i,
			Colour:      "Yellow",
			EngineCap:   2.0 + float32(i)*0.1,
		}
		_, err := carRepo.Create(context.Background(), reqCar)
		assert.NoError(t, err)
	}

	testCases := []struct {
		name     string
		req      models.GetAllCarsRequest
		expected int
	}{
		{"Get 1st page with limit 3", models.GetAllCarsRequest{Search: "", Page: 1, Limit: 2}, 2},
		{"Get 2nd page with limit 2", models.GetAllCarsRequest{Search: "", Page: 2, Limit: 1}, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cars, err := carRepo.GetAll(context.Background(), tc.req)
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, len(cars.Cars))
		})
	}
}

func TestDeleteCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.Car{
		Name:  faker.NAME,
		Model: faker.WORD,
	}

	id, err := carRepo.Create(context.Background(), reqCar)
	assert.NoError(t, err)

	err = carRepo.Delete(context.Background(), id)

	if assert.NoError(t, err) {
		fmt.Println("passed")
	}
}

func TestGetAllOrdersCars(t *testing.T) {
	DB := NewCar(db)

	request := models.GetAllCarsRequest{
		Page:   1,
		Limit:  10,
		Search: "",
	}

	repo := &carRepo{db: DB.db}

	orders, err := repo.GetAllOrdersCars(context.Background(), request)

	assert.NoError(t, err)

	fmt.Println("pass", orders)
}

func TestGetByIDCustomerCarr(t *testing.T) {

	DB := NewCar(db)
	id := "279c010c-d004-4c6e-9f74-66e3fae81cc9"
	car, err := DB.GetByIDCustomerCarr(context.Background(), id)

	if errors.Is(err, sql.ErrNoRows){
	fmt.Println(car, "passed")
	}

}


