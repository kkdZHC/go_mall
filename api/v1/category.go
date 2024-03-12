package v1

import (
	"go_mall/pkg/util"
	"go_mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	var listCategory service.CategoryService
	err := c.ShouldBind(&listCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("listCategory err: ", err)
	}
	res := listCategory.List(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
