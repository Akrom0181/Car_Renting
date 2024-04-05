package handler

import (
	"context"
	"fmt"
	"net/http"
	_ "rent-car/api/docs"
	"rent-car/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// CreateOrder godoc
// @Router 		/order [POST]
// @Summary 	create a order
// @Description This api is creates a new order and returns its id
// @Tags 	    order
// @Accept		json
// @Produce		json
// @Param		order body  models.CreateOrderr true "car"
// @Success		200  {object}  models.GetOrder
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateOrder(c *gin.Context) {
	order := models.CreateOrder{}

	if err := c.ShouldBindJSON(&order); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.Order().Create(context.Background(), order)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating order", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Order is created successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// UpdateOrder godoc
// @Router                /order/{id} [PUT]
// @Summary 			  update a order
// @Description:          this api updates order information
// @Tags 			      order
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Order ID"
// @Param       		  order body models.GetOrder true "car"
// @Success 		      200 {object} models.GetOrder
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateOrder(c *gin.Context) {
	order := models.GetOrder{}
	if err := c.ShouldBindJSON(&order); err != nil {
		handleResponseLog(c, h.Log, "error while decoding body", http.StatusBadRequest, err.Error())
		return
	}

	order.Id = c.Param("id")
	err := uuid.Validate(order.Id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating"+order.Id, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Services.Order().Update(context.Background(), order)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating order", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// GetAllOrder godoc
// @Router 			/order [GET]
// @Summary 		get all order
// @Description 	This API returns order list
// @Tags 			order
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllOrdersResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllOrder(c *gin.Context) {
	var (
		request = models.GetAllOrdersRequest{}
	)

	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing page: ", http.StatusInternalServerError, err.Error())
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
	orders, err := h.Services.Order().GetAll(context.Background(), request)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting orders", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "", http.StatusOK, orders)
}

// @Security ApiKeyAuth
// GetByIDOrder godoc
// @Router       /order/{id} [GET]
// @Summary      return a order by ID
// @Description  Retrieves a order by its ID
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        id path string true "Order ID"
// @Success      200 {object} models.GetOrder
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetOne(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	customer, err := h.Services.Order().GetByID(context.Background(), id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting order by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, customer)
}

// func (c Controller) DeleteOrder(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	fmt.Println("id", id)

// 	err := uuid.Validate(id)
// 	if err != nil {
// 		fmt.Println("error while validating id,err:", err.Error())
// 		handleResponseLog(w, h.Log, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	err = c.Store.Order().Delete(id)
// 	if err != nil {
// 		fmt.Println("error while deleting order, err:", err)
// 		handleResponseLog(w, h.Log, http.StatusInternalServerError, err)
// 		return
// 	}
// 	handleResponseLog(w, h.Log, http.StatusOK, id)
// }
