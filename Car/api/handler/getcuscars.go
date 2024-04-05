// package handler

// import (
// 	"fmt"
// 	"net/http"
// 	"rent-car/api/models"

// 	"github.com/gin-gonic/gin"
// )

// func (h Handler) GetCusCars(c *gin.Context) {
// 	var (
// 		request = models.GetAllCarsRequest{}
// 	)
// 	id := c.Param("id")
// 	fmt.Println("id: ", id)

// 	request.Search = c.Query("search")

// 	page, err := ParsePageQueryParam(c)
// 	if err != nil {
// 		handleResponse(c, "error while parsing page: ", http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	limit, err := ParseLimitQueryParam(c)
// 	if err != nil {
// 		handleResponse(c, "error while parsing limit", http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	fmt.Println("page: ", page)
// 	fmt.Println("limit: ", limit)

// 	request.Page = page
// 	request.Limit = limit
// 	orders, err := h.Store.GetCusCars().GetCusCars(id, request)
// 	if err != nil {
// 		handleResponse(c, "error while getting cuSorders", http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	handleResponse(c, "", http.StatusOK, orders)
// }
package handler