package v1

import (
	"go_mall/pkg/util"
	"go_mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createOrderService := service.OrderService{}
	err := c.ShouldBind(&createOrderService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("order CreateOrder err: ", err)
	}
	res := createOrderService.Create(c.Request.Context(), claim.ID)
	c.JSON(http.StatusOK, res)
}

func DeleteOrder(c *gin.Context) {
	deleteOrderService := service.OrderService{}
	err := c.ShouldBind(&deleteOrderService)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("Order deleteService err: ", err)
	}
	res := deleteOrderService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func ShowOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showOrderService := service.OrderService{}
	err := c.ShouldBind(&showOrderService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("order ListOrder err: ", err)
	}
	res := showOrderService.Show(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func ListOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showOrderService := service.OrderService{}
	err := c.ShouldBind(&showOrderService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("order ListOrder err: ", err)
	}
	res := showOrderService.List(c.Request.Context(), claim.ID)
	c.JSON(http.StatusOK, res)
}
