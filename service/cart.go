package service

import (
	"context"
	"go_mall/dao"
	"go_mall/model"
	"go_mall/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       int  `json:"num" form:"num"`
}

func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	var cart *model.Cart
	//判断有无这个商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "NoThisProduct",
			Error:  err.Error(),
		}
	}

	cartDao := dao.NewCartDao(ctx)
	cart = &model.Cart{
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       uint(service.Num),
	}
	err = cartDao.CreateCart(cart)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "CreateCartError",
			Error:  err.Error(),
		}
	}

	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "GetBossError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "ok",
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

func (service *CartService) List(ctx context.Context, uId uint) serializer.Response {
	cartDao := dao.NewCartDao(ctx)
	carts, err := cartDao.ListCartByUid(uId)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "ListCartError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "ok",
		Data:   serializer.BuildCarts(carts),
	}
}

func (service *CartService) Update(ctx context.Context, uId uint, cId string) serializer.Response {
	cartDao := dao.NewCartDao(ctx)
	cartId, _ := strconv.Atoi(cId)
	err := cartDao.UpdateCartByUid(uId, uint(cartId), uint(service.Num))
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "UpdateCartError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "ok",
	}
}

func (service *CartService) Delete(ctx context.Context, uId uint, cId string) serializer.Response {
	cartDao := dao.NewCartDao(ctx)
	cartId, _ := strconv.Atoi(cId)
	err := cartDao.DeleteCartByAid(uId, uint(cartId))
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "DeleteCartByAidError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "ok",
	}
}
