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

			//收藏夹操作
			authed.GET("favorites", api.ListFavorites)
			authed.POST("favorites", api.CreateFavorites)
			authed.DELETE("favorites/:id", api.DeleteFavorite)

			//地址操作
			authed.POST("addresses", api.CreateAddress)
			authed.GET("addresses/:id", api.ShowAddress)
			authed.GET("addresses", api.ListAddress)
			authed.PUT("addresses/:id", api.UpdateAddress)
			authed.DELETE("addresses/:id", api.DeleteAddress)

			//购物车操作
			authed.POST("carts", api.CreateCart)
			authed.GET("carts/", api.ListCart)
			authed.PUT("carts/:id", api.UpdateCart)
			authed.DELETE("carts/:id", api.DeleteCart)

			//订单操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("orders/", api.ListOrder)
			authed.GET("orders/:id", api.ShowOrder)
			authed.DELETE("orders/:id", api.DeleteOrder)
			// 秒杀专场
			// authed.POST("skill_product/init", api.InitSkillProductHandler())
			// authed.GET("skill_product/list", api.ListSkillProductHandler())
			// authed.GET("skill_product/show", api.GetSkillProductHandler())
			// authed.POST("skill_product/skill", api.SkillProductHandler())
		}
	}
	return r
}
