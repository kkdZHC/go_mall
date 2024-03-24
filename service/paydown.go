package service

import (
	"context"
	"errors"
	"fmt"
	"go_mall/dao"
	"go_mall/pkg/util"
	"go_mall/serializer"
	"net/http"
	"strconv"
)

type OrderPay struct {
	OrderId   uint    `json:"order_id" form:"order_id"`
	Money     float64 `json:"money" form:"money"`
	OrderNo   string  `json:"order_no" form:"order_no"`
	ProductId uint    `json:"product_id" form:"product_id"`
	PayTime   string  `json:"pay_time" form:"pay_time"`
	Sign      string  `json:"sign" form:"sign"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	BossName  string  `json:"boss_name" form:"boss_name"`
	Num       int     `json:"num" form:"num"`
	Key       string  `json:"key" form:"key"`
}

func (service *OrderPay) PayDown(ctx context.Context, uId uint) serializer.Response {
	util.Encrypt.SetKey(service.Key)
	orderDao := dao.NewOrderDao(ctx)
	//开启事务
	tx := orderDao.Begin()
	order, err := orderDao.GetOrderByOid(service.OrderId, uId)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "GetOrderByOidError",
			Error:  err.Error(),
		}
	}
	money := order.Money
	num := order.Num
	money = money * float64(num)

	userDao := dao.NewUserDaoByDB(tx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "GetUserError",
			Error:  err.Error(),
		}
	}
	//对金钱解密，减去订单后再加密
	//用户扣钱
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)

	if moneyFloat-money < 0.0 {
		tx.Rollback()
		return serializer.Response{
			Status: 500,
			Msg:    "NotEnougthMoney",
			Error:  errors.New("金额不足").Error(),
		}
	}
	finMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = util.Encrypt.AesEncoding(finMoney)

	userDao = dao.NewUserDaoByDB(userDao.DB)
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		tx.Rollback()
		return serializer.Response{
			Status: 500,
			Msg:    "UpdateUserMoneyError",
			Error:  err.Error(),
		}
	}

	boss, err := userDao.GetUserById(service.BossId)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "GetBossError",
			Error:  err.Error(),
		}
	}
	//商家加钱
	moneyStr = util.Encrypt.AesDecoding(boss.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	finMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(finMoney)
	err = userDao.UpdateUserById(boss.ID, boss)
	if err != nil {
		tx.Rollback()
		return serializer.Response{
			Status: 500,
			Msg:    "UpdateBossMoneyError",
			Error:  err.Error(),
		}
	}

	//商品减库存
	productDao := dao.NewProductDaoByDB(tx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		tx.Rollback()
		return serializer.Response{
			Status: 500,
			Msg:    "GetProductByIdError",
			Error:  err.Error(),
		}
	}
	product.Num -= num
	err = productDao.UpdateProductByPid(product.ID, product)
	if err != nil {
		tx.Rollback()
		return serializer.Response{
			Status: 500,
			Msg:    "UpdateProductByPid",
			Error:  err.Error(),
		}
	}

	//订单删除
	err = orderDao.DeleteOrderByOid(order.ID, uId)
	if err != nil {
		tx.Rollback()
		return serializer.Response{
			Status: 500,
			Msg:    "DeleteOrderByOidError",
			Error:  err.Error(),
		}
	}
	tx.Commit()
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "ok",
	}
}
