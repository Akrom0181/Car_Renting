package handler

import (
	"backend_course/rent_car/api/models"
	"backend_course/rent_car/pkg/check"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CustomerLogin godoc
// @Router       /customer/login [POST]
// @Summary      Customer login
// @Description  Customer login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.CustomerLoginRequest true "login"
// @Success      201  {object}  models.CustomerLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) CustomerLogin(c *gin.Context) {
	loginReq := models.CustomerLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate login & password

	if err := check.ValidatePhoneNumber(loginReq.Login); err != nil {
		handleResponseLog(c, h.Log, "error while validating phoneNumber", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.IsValidPassword(loginReq.Password); err != nil {
		handleResponseLog(c, h.Log, "error while validating password", http.StatusBadRequest, err.Error())
	}

	loginResp, err := h.Services.Auth().CustomerLogin(c.Request.Context(), loginReq)
	if err != nil {
		handleResponseLog(c, h.Log, "unauthorized", http.StatusUnauthorized, err)
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, loginResp)

}
