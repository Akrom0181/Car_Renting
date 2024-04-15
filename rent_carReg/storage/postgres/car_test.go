package postgres

import (
	"clone/rent_car_us/api/models"
	"context"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.Car{
		Name:   faker.Name(),
		Brand:  faker.Word(),
		Colour: faker.Name(),
	}

	id, err := carRepo.Create(context.Background(), reqCar)
	if assert.NoError(t, err) {
		createdCar, err := carRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCar.Name, createdCar.Name)
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
		Id:     "a49f419b-f31a-4a90-8c4d-3663b5ab6da3",
		Name:   "asdasdas",
		Brand:  "asdasd",
		Colour: "asdasda",
	}

	id, err := carRepo.Update(context.Background(), reqCar)
	if assert.NoError(t, err) {
		updateCar, err := carRepo.GetByID(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCar.Name, updateCar.Name)
			assert.Equal(t, reqCar.Brand, updateCar.Brand)
		} else {
			return
		}
		fmt.Println("update car", updateCar)
	}
}

func TestGetAllCar(t *testing.T) {
	carRepo := NewCar(db)
	car, err := carRepo.GetAllCars(context.Background(), models.GetAllCarsRequest{})
	if assert.NoError(t, err) {
		fmt.Println("get all car", car.Count)
	} else {
		return
	}

}
