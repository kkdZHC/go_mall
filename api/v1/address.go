package v1

import (
	"go_mall/pkg/util"
	"go_mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAddress(c *gin.Context) {
	listAddressService := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&listAddressService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("Address listService err: ", err)
	}
	res := listAddressService.List(c.Request.Context(), claim.ID)
	c.JSON(http.StatusOK, res)
}

func CreateAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createAddressService := service.AddressService{}
	err := c.ShouldBind(&createAddressService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("address CreateAddress err: ", err)
	}
	res := createAddressService.Create(c.Request.Context(), claim.ID)
	c.JSON(http.StatusOK, res)
}

func UpdateAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	updateAddressService := service.AddressService{}
	err := c.ShouldBind(&updateAddressService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("address CreateAddress err: ", err)
	}
	res := updateAddressService.Update(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func DeleteAddress(c *gin.Context) {
	deleteAddressService := service.AddressService{}
	err := c.ShouldBind(&deleteAddressService)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("Address deleteService err: ", err)
	}
	res := deleteAddressService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func ShowAddress(c *gin.Context) {
	showAddressService := service.AddressService{}
	err := c.ShouldBind(&showAddressService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("address ListAddress err: ", err)
	}
	res := showAddressService.Show(c.Request.Context(), c.Param("id"))
	c.JSON(http.StatusOK, res)
}
