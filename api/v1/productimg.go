package v1

import (
	"go_mall/pkg/util"
	"go_mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProductImg(c *gin.Context) {
	var listProductImg service.ListProductImg
	err := c.ShouldBind(&listProductImg)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("product listProductImg err: ", err)
	}
	res := listProductImg.List(c.Request.Context(), c.Param("id"))
	c.JSON(http.StatusOK, res)
}
