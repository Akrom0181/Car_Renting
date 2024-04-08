package api

import (
	"fmt"
	"rent-car/api/handler"
	"rent-car/pkg/logger"
	"rent-car/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func New(services service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.NewStrg(services, log)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.Use(authMiddleware)

	r.Use(getHeaders)

	r.POST("/car", h.CreateCar)
	// r.GET("/car/:id", h.GetAllCars)
	r.GET("/car", h.GetAllCars)
	r.GET("/car/:id", h.GetByIDCar)
	r.GET("/avacar", h.GetAvailableCars)
	r.GET("/getcar/:id", h.GetByIDCustomeCarr)
	r.GET("/carcustomer", h.GetAllOrdersCars)
	r.PUT("/car/:id", h.UpdateCar)
	r.DELETE("/car/:id", h.DeleteCar)
	// r.PATCH("/car/:id", h.UpdateUserPassword)

	r.GET("/customer/:id", h.GetByIDCustomer)
	r.GET("/customer", h.GetAllCustomer)
	r.POST("/customer", h.CreateCustomer)
	r.PUT("/customer/:id", h.UpdateCustomer)
	r.DELETE("/customer/:id", h.DeleteCustomer)
	r.PATCH("/customer/:phone", h.ChangePassword)
	r.PUT("customer/login", h.Login)


	r.GET("/order", h.GetAllOrder)
	r.GET("/order/:id", h.GetOne)
	r.POST("/order", h.CreateOrder)
	r.PUT("/order/:id", h.UpdateOrder)

	// r.GET("/getcus/:id", h.GetCusCars)

	return r
}

// func authMiddleware(c *gin.Context) {
// 	auth := c.GetHeader("Authorization")
// 	if auth == "" {
// 		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
// 	}
// 	c.Next()
// }

func getHeaders(c *gin.Context) {
	headers := c.Request.Header

	for key, values := range headers {
		fmt.Printf("key %v, value %v\n", key, values)
	}

}
