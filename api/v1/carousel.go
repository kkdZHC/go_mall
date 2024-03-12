package v1

import (
	"go_mall/pkg/util"
	"go_mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.CarouselService
	err := c.ShouldBind(&listCarousel)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("ListCarousel err: ", err)
	}
	res := listCarousel.List(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
