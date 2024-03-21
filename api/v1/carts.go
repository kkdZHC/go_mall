package v1

import (
	"go_mall/pkg/util"
	"go_mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createCartService := service.CartService{}
	err := c.ShouldBind(&createCartService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("cart CreateCart err: ", err)
	}
	res := createCartService.Create(c.Request.Context(), claim.ID)
	c.JSON(http.StatusOK, res)
}

func UpdateCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	updateCartService := service.CartService{}
	err := c.ShouldBind(&updateCartService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("cart CreateCart err: ", err)
	}
	res := updateCartService.Update(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func DeleteCart(c *gin.Context) {
	deleteCartService := service.CartService{}
	err := c.ShouldBind(&deleteCartService)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("Cart deleteService err: ", err)
	}
	res := deleteCartService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func ListCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showCartService := service.CartService{}
	err := c.ShouldBind(&showCartService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("cart ListCart err: ", err)
	}
	res := showCartService.List(c.Request.Context(), claim.ID)
	c.JSON(http.StatusOK, res)
}
