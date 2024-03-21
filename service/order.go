package service

import (
	"context"
	"fmt"
	"go_mall/dao"
	"go_mall/model"
	"go_mall/serializer"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
	ProductId uint    `json:"product_id" form:"product_id"`
	Num       int     `json:"num" form:"num"`
	AddressId uint    `json:"address_id" form:"address_id"`
	Money     float64 `json:"money" form:"money"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	UserId    uint    `json:"user_id" form:"user_id"`
	OrderNum  int     `json:"order_num" form:"order_num"`
	Type      uint    `json:"type" form:"type"`
	model.BasePage
}

func (service *OrderService) Create(ctx context.Context, uId uint) serializer.Response {
	var order *model.Order
	orderDao := dao.NewOrderDao(ctx)
	order = &model.Order{
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       service.Num,
		Type:      1, //1 未支付 2 已支付
		Money:     service.Money,
	}
	//检验地址是否存在
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(service.AddressId)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "GetAddressError||AddressNotExist",
			Error:  err.Error(),
		}
	}
	order.AddressId = address.ID
	//设置订单编号
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000)) // 生成随机number
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(service.UserId))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = orderNum

	err = orderDao.CreateOrder(order)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "CreateOrderError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "ok",
	}
}

func (service *OrderService) Show(ctx context.Context, uId uint, oId string) serializer.Response {
	var order *model.Order
	orderDao := dao.NewOrderDao(ctx)
	orderId, _ := strconv.Atoi(oId)
	order, err := orderDao.GetOrderByOid(uint(orderId), uId)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "GetOrderByUidError",
			Error:  err.Error(),
		}
	}
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(order.AddressId)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "GetAddressError",
			Error:  err.Error(),
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductId)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "GetProductError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "ok",
		Data:   serializer.BuildOrder(order, product, address),
	}
}

func (service *OrderService) List(ctx context.Context, uId uint) serializer.Response {
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	orderDao := dao.NewOrderDao(ctx)

	condition := make(map[string]interface{})
	if service.Type != 0 {
		condition["type"] = service.Type
	}
	condition["user_id"] = uId
	orderes, total, err := orderDao.ListOrderByCondition(condition, service.BasePage)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "ListOrderError",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildOrderList(ctx, orderes), uint(total))
}

func (service *OrderService) Delete(ctx context.Context, uId uint, oId string) serializer.Response {
	orderDao := dao.NewOrderDao(ctx)
	orderId, _ := strconv.Atoi(oId)
	err := orderDao.DeleteOrderByOid(uId, uint(orderId))
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "DeleteOrderByAidError",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "ok",
	}
}
