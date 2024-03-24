package v1

import (
	"go_mall/pkg/util"
	"go_mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OrderPay(c *gin.Context) {
	orderPay := service.OrderPay{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&orderPay)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		util.LogrusObj.Infoln("orderpay err: ", err)
	}
	res := orderPay.PayDown(c.Request.Context(), claim.ID)
	c.JSON(http.StatusOK, res)

}
