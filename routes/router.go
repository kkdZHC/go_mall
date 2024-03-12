package routes

import (
	api "go_mall/api/v1"
	"go_mall/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "success")
		})
		//用户操作
		v1.POST("user/register", api.UserRegister) //注册
		v1.POST("user/login", api.UserLogin)       //登录

		v1.GET("carousels", api.ListCarousel) //轮播图

		v1.GET("products", api.ListProduct)    //获取商品列表
		v1.GET("product/:id", api.ShowProduct) //展示商品详细信息
		v1.GET("imgs/:id", api.ListProductImg) //展示商品详细信息
		v1.GET("categories", api.ListCategory) //展示商品分类

		authed := v1.Group("/") //需要验证登录
		{
			//用户操作
			authed.Use(middleware.JWT())
			authed.PUT("user", api.UserUpdate)               //更新用户
			authed.POST("avatar", api.UploadAvatar)          //上传头像
			authed.POST("user/sending-email", api.SendEmail) //邮件发送
			authed.POST("user/valid-email", api.ValidEmail)  //验证邮箱

			authed.POST("money", api.ShowMoney) //显示金额

			//商品操作
			authed.POST("product", api.CreateProduct)  //创建商品
			authed.POST("products", api.SearchProduct) //搜索商品
		}
	}
	return r
}
