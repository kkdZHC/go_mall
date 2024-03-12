package v1

import (
	"go_mall/pkg/util"
	"go_mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	err := c.ShouldBind(&createProductService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("product CreateProduct err: ", err)
	}
	res := createProductService.Create(c.Request.Context(), claim.ID, files)
	c.JSON(http.StatusOK, res)
}

func ListProduct(c *gin.Context) {
	listProductService := service.ProductService{}
	err := c.ShouldBind(&listProductService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("product ListProduct err: ", err)
	}
	res := listProductService.List(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func ShowProduct(c *gin.Context) {
	showProductService := service.ProductService{}
	err := c.ShouldBind(&showProductService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("product ListProduct err: ", err)
	}
	res := showProductService.Show(c.Request.Context(), c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func SearchProduct(c *gin.Context) {
	searchProductService := service.ProductService{}
	err := c.ShouldBind(&searchProductService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("product ListProduct err: ", err)
	}
	res := searchProductService.Search(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
