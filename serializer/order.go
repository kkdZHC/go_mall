package serializer

import (
	"context"
	"go_mall/conf"
	"go_mall/dao"
	"go_mall/model"
)

type OrderVO struct {
	Id           uint    `json:"id"`
	OrderNum     uint64  `json:"order_num"`
	CreatedAt    int64   `json:"created_at"`
	UpdatedAt    int64   `json:"updated_at"`
	UserId       uint    `json:"user_id"`
	ProductId    uint    `json:"product_id"`
	BossId       uint    `json:"boss_id"`
	Num          int     `json:"num"`
	AddressName  string  `json:"address_name"`
	AddressPhone string  `json:"address_phone"`
	Address      string  `json:"address"`
	Type         uint    `json:"type"`
	ProductName  string  `json:"product_name"`
	ImgPath      string  `json:"img_path"`
	Money        float64 `json:"money"`
}

func BuildOrder(order *model.Order, product *model.Product, address *model.Address) *OrderVO {
	return &OrderVO{
		Id:           order.ID,
		OrderNum:     order.OrderNum,
		CreatedAt:    order.CreatedAt.Unix(),
		UpdatedAt:    order.UpdatedAt.Unix(),
		UserId:       order.UserId,
		ProductId:    order.ProductId,
		BossId:       order.BossId,
		Num:          order.Num,
		AddressName:  address.Name,
		AddressPhone: address.Phone,
		Address:      address.Address,
		Type:         order.Type,
		ProductName:  product.Name,
		ImgPath:      conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Money:        order.Money,
	}
}

func BuildOrderList(ctx context.Context, items []*model.Order) (orderListVO []*OrderVO) {
	productDao := dao.NewProductDao(ctx)
	addressDao := dao.NewAddressDao(ctx)
	for _, item := range items {
		product, _ := productDao.GetProductById(item.ProductId)
		address, _ := addressDao.GetAddressByAid(item.AddressId)
		order := BuildOrder(item, product, address)
		orderListVO = append(orderListVO, order)
	}
	return orderListVO
}
