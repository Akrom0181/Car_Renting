package handler

import (
	"context"
	"fmt"
	"net/http"
	_ "rent-car/api/docs"
	"rent-car/api/models"
	"rent-car/config"
	"rent-car/pkg/check"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// CreateCar godoc
// @Router 		/car [POST]
// @Summary 	create a car
// @Description This api is creates a new car and returns its id
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		car body models.CreateCar true "car"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateCar(c *gin.Context) {
	car := models.Car{}

	if err := c.ShouldBindJSON(&car); err != nil {
		handleResponseLog(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	if err := check.ValidateCarYear(car.Year); err != nil {
		handleResponseLog(c, h.Log, "error while validating car year, year: "+strconv.Itoa(car.Year), http.StatusBadRequest, err.Error())

		return
	}

	ctx, cancel := context.WithTimeout(c, config.Time)
	defer cancel()

	id, err := h.Services.Car().Create(ctx, car)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating car", http.StatusBadRequest, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GetAllAvailableCars godoc
// @Router 			/avacar [GET]
// @Summary 		get all available cars
// @Description 	This API returns available car list
// @Tags 			car
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllCarsResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAvailableCars(c *gin.Context) {
	var (
		request = models.GetAllCarsRequest{}
	)
	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "Error while parsing limit", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	ctx, cancel := context.WithTimeout(c, config.Time)
	defer cancel()

	cars, err := h.Services.Car().GetAvailableCars(ctx, request)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting cars", http.StatusBadRequest, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "", http.StatusOK, cars)
}

// @Security ApiKeyAuth
// UpdateCar godoc
// @Router                /car/{id} [PUT]
// @Summary 			  update a car
// @Description:          this api updates car information
// @Tags 			      car
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Car ID"
// @Param       		  car body models.Car true "car"
// @Success 		      200 {object} models.Car
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateCar(c *gin.Context) {
	car := models.Car{}

	if err := c.ShouldBindJSON(&car); err != nil {
		handleResponseLog(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	if err := check.ValidateCarYear(car.Year); err != nil {
		handleResponseLog(c, h.Log, "error while validating car year, year: "+strconv.Itoa(car.Year), http.StatusBadRequest, err.Error())
		return
	}
	car.Id = c.Param("id")

	err := uuid.Validate(car.Id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating car id,id: "+car.Id, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Time)
	defer cancel()

	id, err := h.Services.Car().Update(ctx, car)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating car", http.StatusBadRequest, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GetAllCars godoc
// @Router 			/car [GET]
// @Summary 		get all cars
// @Description 	This API returns car list
// @Tags 			car
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllCarsResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllCars(c *gin.Context) {
	var (
		request = models.GetAllCarsRequest{}
	)

	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	ctx, cancel := context.WithTimeout(c, config.Time)
	defer cancel()

	cars, err := h.Services.Car().GetAll(ctx, request)
	if err != nil {
		handleResponseLog(c, h.Log, "error while gettign cars", http.StatusBadRequest, err.Error())

		return
	}

	handleResponseLog(c, h.Log, "", http.StatusOK, cars)
}

// @Security ApiKeyAuth
// GetByIDCar godoc
// @Router       /car/{id} [GET]
// @Summary      return a car by ID
// @Description  Retrieves a car by its ID
// @Tags         car
// @Accept       json
// @Produce      json
// @Param        id path string true "Car ID"
// @Success      200 {object} models.Car
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDCar(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	ctx, cancel := context.WithTimeout(c, config.Time)
	defer cancel()

	car, err := h.Services.Car().GetByID(ctx, id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting car by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, car)
}

// @Security ApiKeyAuth
// DeleteCar godoc
// @Router          /car/{id} [DELETE]
// @Summary         delete a car by ID
// @Description     Deletes a car by its ID
// @Tags            car
// @Accept          json
// @Produce         json
// @Param           id path string true "Car ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteCar(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Time)
	defer cancel()

	err = h.Services.Car().Delete(ctx, id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while deleting car", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GetAllCustomerCars godoc
// @Router 			/carcustomer [GET]
// @Summary 		get all customer cars
// @Description 	This API returns car list
// @Tags 			car
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllCarsResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllOrdersCars(c *gin.Context) {
	var (
		request = models.GetAllCarsRequest{}
	)

	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing limit", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	ctx, cancel := context.WithTimeout(c, config.Time)
	defer cancel()

	cars, err := h.Services.Car().GetAllCustomerCars(ctx, request)
	if err != nil {
		fmt.Println("error while getting carsOrder, err: ", err)
		handleResponseLog(c, h.Log, "", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, cars)
}

// @Security ApiKeyAuth
// GetByIDCustomerCar godoc
// @Router       /getcar/{id} [GET]
// @Summary      return a car by ID
// @Description  Retrieves a car by its ID
// @Tags         car
// @Accept       json
// @Produce      json
// @Param        id path string true "Car ID"
// @Success      200 {object} models.GetOrder
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDCustomeCar(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	ctx, cancel := context.WithTimeout(c, config.Time)
	defer cancel()

	customer, err := h.Services.Car().GetByIDCustomerCar(ctx, id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting customerCar by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, customer)
}
