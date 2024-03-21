package v1

import (
	"go_mall/pkg/util"
	"go_mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListFavorites(c *gin.Context) {
	listFavoriteService := service.FavoriteService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&listFavoriteService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("Favorite listService err: ", err)
	}
	res := listFavoriteService.List(c.Request.Context(), claim.ID)
	c.JSON(http.StatusOK, res)
}

func CreateFavorites(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createFavoriteService := service.FavoriteService{}
	err := c.ShouldBind(&createFavoriteService)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("product CreateProduct err: ", err)
	}
	res := createFavoriteService.Create(c.Request.Context(), claim.ID)
	c.JSON(http.StatusOK, res)
}

func DeleteFavorite(c *gin.Context) {
	deleteFavoriteService := service.FavoriteService{}
	err := c.ShouldBind(&deleteFavoriteService)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("Favorite deleteService err: ", err)
	}
	res := deleteFavoriteService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}
