package handler

import (
	"fmt"
	"net/http"
	_ "rent-car/api/docs"
	"rent-car/api/models"
	"rent-car/pkg/check"
	"rent-car/pkg/hash"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Router 		/customer/login [PUT]
// @Summary 	login
// @Description This api is used for logining
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		customer body  models.Login true "login"
// @Success		200  {object}  models.Login
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) Login(c *gin.Context) {
	login := models.Login{}

	if err := c.ShouldBindJSON(&login); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidatePhoneNumber(login.Phone); err != nil {
		handleResponseLog(c, h.Log, "error while validating phoneNumber", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.IsValidPassword(login.Password); err != nil {
		handleResponseLog(c, h.Log, "error while validating password", http.StatusBadRequest, err.Error())
	}

	password, err := h.Services.Customer().GetPassword(c.Request.Context(), login.Phone)
	if err != nil {
		handleResponseLog(c, h.Log, "error while Login Customer", http.StatusBadRequest, err.Error())
		return
	}

	err1 := hash.CompareHashAndPassword(password, login.Password)
	if err1 != nil {
		handleResponseLog(c, h.Log, "Wrong Password", http.StatusBadRequest, err1.Error())
		return
	}

	handleResponseLog(c, h.Log, "Login successfully", http.StatusOK, password)
}

// @Security ApiKeyAuth
// CreateCustomer godoc
// @Router 		/customer [POST]
// @Summary 	create a customer
// @Description This api is creates a new customer and returns its id
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		customer body  models.CreateCustomer true "car"
// @Success		200  {object}  models.Customer
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateCustomer(c *gin.Context) {
	customer := models.Customer{}

	if err := c.ShouldBindJSON(&customer); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateEmailAddress(customer.Gmail); err != nil {
		handleResponseLog(c, h.Log, "error while validating email address", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidatePhoneNumber(customer.Phone); err != nil {
		handleResponseLog(c, h.Log, "error while validating phoneNumber", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.IsValidPassword(customer.Password); err != nil {
		handleResponseLog(c, h.Log, "error while validating password", http.StatusBadRequest, err.Error())
	}

	hashedpass, _ := bcrypt.GenerateFromPassword([]byte(
		customer.Password,
	), bcrypt.DefaultCost)

	customer.Password = string(hashedpass)
	id, err := h.Services.Customer().Create(c.Request.Context(), customer)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating customer", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "Created Succesfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// ChangeCustomerPassword godoc
// @Router                /customer/{phone} [PATCH]
// @Summary 			  change customer password
// @Description:          this api changes customer password
// @Tags 			      customer
// @Accept 			      json
// @Produce 		      json
// @Param 			      phone path string true "Customer phone"
// @Param       		  customer body models.ChangePassword true "customer"
// @Success 		      200 {object} models.ChangePassword
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) ChangePassword(c *gin.Context) {
	customer := models.Customer{}

	if err := c.ShouldBindJSON(&customer); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	customer.Phone = c.Param("phone")
	if err := check.ValidatePhoneNumber(customer.Phone); err != nil {
		handleResponseLog(c, h.Log, "error while validating phoneNumber", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.IsValidPassword(customer.Password); err != nil {
		handleResponseLog(c, h.Log, "error while validating password", http.StatusBadRequest, err.Error())
	}

	hashedpass, _ := bcrypt.GenerateFromPassword([]byte(
		customer.Password,
	), bcrypt.DefaultCost)

	customer.Password = string(hashedpass)

	id, err := h.Services.Customer().ChangePassword(c.Request.Context(), customer)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating customer", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "Created Succesfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// UpdateCustomer godoc
// @Router                /customer/{id} [PUT]
// @Summary 			  update a customer
// @Description:          this api updates customer information
// @Tags 			      customer
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Customer ID"
// @Param       		  customer body models.Customer true "customer"
// @Success 		      200 {object} models.Customer
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateCustomer(c *gin.Context) {
	customer := models.Customer{}
	if err := c.ShouldBindJSON(&customer); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}
	customer.Id = c.Param("id")
	err := uuid.Validate(customer.Id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Services.Customer().Update(c.Request.Context(), customer)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating customer", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "Updated Successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GetAllCustomer godoc
// @Router 			/customer [GET]
// @Summary 		get all customer
// @Description 	This API returns customer list
// @Tags 			customer
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllCustomersResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllCustomer(c *gin.Context) {
	var (
		request = models.GetAllCustomersRequest{}
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

	customers, err := h.Services.Customer().GetAll(c.Request.Context(), request)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting customers,", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, customers)
}

// @Security ApiKeyAuth
// GetByIDCustomer godoc
// @Router       /customer/{id} [GET]
// @Summary      return a customer by ID
// @Description  Retrieves a customer by its ID
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        id path string true "Customer ID"
// @Success      200 {object} models.Customer
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDCustomer(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id", id)

	customer, err := h.Services.Customer().GetByID(c.Request.Context(), id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting customer by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, customer)
}

// @Security ApiKeyAuth
// DeleteCustomer godoc
// @Router          /customer/{id} [DELETE]
// @Summary         delete a customer by ID
// @Description     Deletes a customer by its ID
// @Tags            customer
// @Accept          json
// @Produce         json
// @Param           id path string true "Customer ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}
	err = h.Services.Customer().Delete(c.Request.Context(), id)
	if err != nil {
		handleResponseLog(c, h.Log, "Error while deleting customer", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "Deleted successfully", http.StatusOK, id)
}

func (h Handler) GetCustomerCars(c *gin.Context) {
	var (
		request = models.GetAllCustomersRequest{}
	)

	request.Search = c.Param("search")

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

	customers, err := h.Services.Customer().GetCustomerCars(c.Request.Context(), request)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting customers", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "", http.StatusOK, customers)
}

func (h Handler) GetByIDCustomeCarr(c *gin.Context) {
	id := c.Param("id")

	customer, err := h.Services.Customer().GetByIDCustomerCar(c.Request.Context(), id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting customerCar by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, customer)
}

// func (h Handler) GetCustomerCarss(c *gin.Context){
// 	id := c.Param("id")

// 	customer, err := h.Services.Customer().GetCustomerCarss(c.Request.Context(), id)
// 	if err != nil {
// 		handleResponseLog(c, h.Log, "error while getting customerCar by id", http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	handleResponseLog(c, h.Log, "", http.StatusOK, customer)
// }
