package v1

import (
	"go_mall/pkg/util"
	"go_mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	err := c.ShouldBind(&userRegister)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user register api: ", err)
	}
	res := userRegister.Register(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	err := c.ShouldBind(&userLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user login api: ", err)
	}
	res := userLogin.Login(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService

	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&userUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user update api: ", err)
	}
	res := userUpdate.Update(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&uploadAvatar)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user UploadAvatar api: ", err)
	}
	res := uploadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)
	c.JSON(http.StatusOK, res)
}

func SendEmail(c *gin.Context) {
	var sendEmail service.SendEmailService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&sendEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user sendemail api: ", err)
	}
	res := sendEmail.Send(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

func ValidEmail(c *gin.Context) {
	var validEmail service.ValidEmailService
	token := c.GetHeader("Authorization")
	err := c.ShouldBind(&validEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user ValidEmail api: ", err)
	}
	res := validEmail.Valid(c.Request.Context(), token)
	c.JSON(http.StatusOK, res)
}
func ShowMoney(c *gin.Context) {
	var showMoney service.ShowMoneyService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&showMoney)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user showmoney api: ", err)
	}
	res := showMoney.Show(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}
